package api

import (
	"github.com/cyberninja89/website/internal/server/api/auxiliary"

	"github.com/gin-gonic/gin"
	"github.com/go-nacelle/nacelle"
)

type Config struct {
	logger nacelle.Logger
	health nacelle.Health
}

func New(logger nacelle.Logger, opts ...APIOption) *Config {
	napi := Config{}
	napi.logger = logger

	for _, opt := range opts {
		opt(&napi)
	}

	return &napi
}

type APIOption func(api *Config)

func WithHealth(health nacelle.Health) APIOption {
	return func(api *Config) {
		api.health = health
	}
}

func (c *Config) ConfigureRouters(ge gin.IRouter) {
	// V1 Routes
	// v1Router := v1.NewRouter(
	// 	v1.WithHealth(a.health),
	// 	v1.WithLogger(a.logger),
	// 	v1.WithService(
	// 		entity.NewService(
	// 			a.logger,
	// 		)),
	// )
	// v1Router.Configure(ge.Group("/v1"))

	// Aux routes
	auxRouter := auxiliary.NewRouter(
		auxiliary.WithLogger(c.logger),
		auxiliary.WithHealth(c.health),
	)
	auxRouter.Configure(ge.Group("/"))
}
