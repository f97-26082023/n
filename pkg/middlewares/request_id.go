package middlewares

import (
	"github.com/f97/n/pkg/core"
	"github.com/f97/n/pkg/requestid"
	"github.com/f97/n/pkg/settings"
)

const requestIdHeader = "X-Request-ID"

// RequestId generates a new request id and add it to context and response header
func RequestId(config *settings.Config) core.MiddlewareHandlerFunc {
	return func(c *core.Context) {
		if requestid.Container.Current == nil {
			c.Next()
			return
		}

		requestId := requestid.Container.Current.GenerateRequestId(c.ClientIP())
		c.SetRequestId(requestId)

		if config.EnableRequestIdHeader {
			c.Header(requestIdHeader, requestId)
		}

		c.Next()
	}
}
