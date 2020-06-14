package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/cyberninja89/website/internal/personal"
	"github.com/cyberninja89/website/internal/server/api"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/go-nacelle/nacelle"
)

type Process struct {
	Logger nacelle.Logger `service:"logger"`
	Health nacelle.Health `service:"health"`
	Config *Config
	CORS   cors.Config
	API    *http.Server

	stopChannel chan struct{}
	stopOnce    sync.Once
	killChannel chan os.Signal
}

type Config struct {
	Port string `env:"PORT" default:"80"`
	Root string `env:"ROOT_PATH" default:"/"`

	TLSCert []byte `env:"TLS_CERT"`
	TLSKey  []byte `env:"TLS_KEY"`
}

type CORS struct {
	AllowOrigins     []string      `env:"ALLOW_ORIGINS" default:"[\"https://localhost\", \"https://localhost\"]"`
	AllowMethods     []string      `env:"ALLOW_METHODS" default:"[\"*\"]"`
	AllowHeaders     []string      `env:"ALLOW_HEADERS" default:"[\"*\"]"`
	ExposeHeaders    []string      `env:"EXPOSE_HEADERS"`
	AllowCredentials bool          `env:"ALLOW_CREDENTIALS"`
	MaxAge           time.Duration `env:"MAX_AGE"`
}

func NewProcess() *Process {
	return &Process{
		Logger:      nacelle.NewNilLogger(),
		API:         &http.Server{},
		stopChannel: make(chan struct{}),
		killChannel: make(chan os.Signal, 1),
	}
}

func (p *Process) Init(config nacelle.Config) error {
	// Reading the API related environment config
	apiConfig := &Config{}
	if err := config.Load(apiConfig); err != nil {
		return err
	}
	p.Config = apiConfig

	// HTTPS server config currently set to skip over if not provided
	var tlsCfg *tls.Config
	if len(p.Config.TLSCert) > 0 && len(p.Config.TLSKey) > 0 {
		cert, err := tls.X509KeyPair(p.Config.TLSCert, p.Config.TLSKey)
		if err != nil {
			return fmt.Errorf("could not load TLS pair: %s", err)
		}
		tlsCfg = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
		p.API.TLSConfig = tlsCfg
	}
	p.API.Addr = fmt.Sprintf(":%s", p.Config.Port)

	// Reading the CORS related environment config
	corsConfig := &CORS{}
	if err := config.Load(corsConfig); err != nil {
		return err
	}
	p.CORS = cors.Config{
		AllowOrigins:     corsConfig.AllowOrigins,
		AllowMethods:     corsConfig.AllowMethods,
		AllowHeaders:     corsConfig.AllowHeaders,
		ExposeHeaders:    corsConfig.ExposeHeaders,
		AllowCredentials: corsConfig.AllowCredentials,
		MaxAge:           corsConfig.MaxAge,
	}

	signal.Notify(p.killChannel, syscall.SIGINT, syscall.SIGTERM)

	return nil
}

func (p *Process) InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.RedirectTrailingSlash = true
	gin.DefaultWriter = ioutil.Discard

	// Use CORS
	r.Use(cors.New(p.CORS))

	// Use the Gin's builtin logger framework
	r.Use(GinLogger(p.Logger))

	// Jaeger Tracing
	// Tracer := opentracing.GlobalTracer()
	// r.Use(ginhttp.Middleware(Tracer))

	// PPROF debugging
	pprof.Register(r, p.Config.Root+"/debug/pprof")

	// Prometheus
	prom := ginprometheus.NewPrometheus("")
	prom.MetricsPath = p.Config.Root + "/metrics"
	prom.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			// Range over url metric replacers
			for k, v := range personal.URLReplacers {
				if p.Key == v {
					url = strings.Replace(url, p.Value, k, 1)
					break
				}
			}
		}
		return url
	}
	prom.Use(r)

	// Configure routes for the API implementation
	apiImpl := api.New(
		p.Logger,
		api.WithHealth(p.Health),
	)

	// Add the versioned business logic routes
	apiImpl.ConfigureRouters(r.Group(p.Config.Root))

	return r
}

func (p *Process) Start() error {
	p.API.Handler = p.InitRouter()

	go func() {
		if err := p.API.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			p.Logger.Error("http server listen error:", err)
		}
	}()

	select {
	case <-p.killChannel:
		p.Logger.Info("http server received kill signal...")
		if err := p.Shutdown(); err != nil {
			p.Logger.Error("http server error shutting down:", err)
		}
	case <-p.stopChannel:
		p.Logger.Info("http server received stop signal...")
		if err := p.Shutdown(); err != nil {
			p.Logger.Error("http server error shutting down:", err)
		}
	}

	return nil
}

func (p *Process) Shutdown() error {
	p.Logger.Info("http server shutdown sequence triggered...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := p.API.Shutdown(ctx); err != nil {
		return err
	}
	p.Logger.Info("...http server exit complete.")

	return nil
}

func (p *Process) Stop() error {
	p.stopOnce.Do(func() {
		close(p.stopChannel)
	})
	return nil
}
