package api

import (
	"github.com/gin-gonic/gin"

	gin2 "github.com/quanxiang-cloud/cabin/tailormade/gin"
	"github.com/quanxiang-cloud/web-processors/dispatch/internal/chain"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/config"
)

const (
	assign = "assign"
)

type Router struct {
	conf   *config.Config
	engine *gin.Engine
}

type router func(c *config.Config, r map[string]*gin.RouterGroup) error

var routers = []router{
	assignRouter,
}

func NewRouter(c *config.Config) (*Router, error) {
	engine := gin.New()
	engine.Use(gin2.LoggerFunc(), gin2.RecoveryFunc())

	r := map[string]*gin.RouterGroup{
		assign: engine.Group("/api/v1/assign"),
	}

	for _, f := range routers {
		if err := f(c, r); err != nil {
			return nil, err
		}
	}

	return &Router{
		conf:   c,
		engine: engine,
	}, nil
}

func (r *Router) Run() {
	r.engine.Run(r.conf.Port)
}

func assignRouter(c *config.Config, r map[string]*gin.RouterGroup) error {
	r[assign].POST("/exec", Execute(chain.NewEvolution(c)))

	return nil
}
