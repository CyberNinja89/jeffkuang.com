package auxiliary

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nacelle/nacelle"
)

type Router struct {
	Logger nacelle.Logger
	Health nacelle.Health
}

type RouterOption func(ar *Router)

func WithLogger(logger nacelle.Logger) RouterOption {
	return func(ar *Router) {
		ar.Logger = logger
	}
}

func WithHealth(health nacelle.Health) RouterOption {
	return func(ar *Router) {
		ar.Health = health
	}
}

func NewRouter(opts ...RouterOption) *Router {
	nar := Router{}

	for _, opt := range opts {
		opt(&nar)
	}

	return &nar
}

func (r *Router) Configure(ge gin.IRouter) {
	ge.GET("/status", r.Status)
	ge.GET("/version", r.Version)
}
