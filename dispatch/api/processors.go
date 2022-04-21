package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quanxiang-cloud/cabin/logger"
	"github.com/quanxiang-cloud/cabin/tailormade/header"
	"github.com/quanxiang-cloud/cabin/tailormade/resp"
	"github.com/quanxiang-cloud/web-processors/dispatch/internal/chain"
)

// Execute is the main function of assign.
func Execute(cmd chain.Command) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := header.MutateContext(c)

		params := &chain.Parameter{}
		if err := c.ShouldBind(params); err != nil {
			logger.Logger.WithName("Bind Parma").Errorw(err.Error(), header.GetRequestIDKV(ctx).Fuzzy()...)
			resp.Format(nil, err).Context(c)
			return
		}

		if err := cmd.Do(ctx, params); err != nil {
			resp.Format(nil, err).Context(c)
			return
		}

		resp.Format(nil, nil).Context(c, http.StatusOK)
	}
}
