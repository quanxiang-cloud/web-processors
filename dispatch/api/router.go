package api

import (
	"github.com/gin-gonic/gin"

	gin2 "github.com/quanxiang-cloud/cabin/tailormade/gin"
	"github.com/quanxiang-cloud/web-processors/dispatch/internal/chain"
	"github.com/quanxiang-cloud/web-processors/dispatch/pkg/config"
)

const (
	processors = "processors"
)

// Router is the router.
type Router struct {
	conf   *config.Config
	engine *gin.Engine
}

type router func(c *config.Config, r map[string]*gin.RouterGroup) error

var routers = []router{
	processorsRouter,
}

// NewRouter returns a new router.
func NewRouter(c *config.Config) (*Router, error) {
	engine := gin.New()
	engine.Use(gin2.LoggerFunc(), gin2.RecoveryFunc())

	r := map[string]*gin.RouterGroup{
		processors: engine.Group("/api/v1/web-processors"),
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

// Run serve starts the server.
func (r *Router) Run() {
	r.engine.Run(r.conf.Port)
}

func processorsRouter(c *config.Config, r map[string]*gin.RouterGroup) error {
	r[processors].POST("/execute", Execute(chain.NewEvolution(c)))

	return nil
}
