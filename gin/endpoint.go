package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	auth "github.com/kpacha/krakend-http-auth"
	"github.com/luraproject/lura/config"
	"github.com/luraproject/lura/proxy"
	luragin "github.com/luraproject/lura/router/gin"
)

// HandlerFactory decorates a krakendgin.HandlerFactory with the auth layer
func HandlerFactory(hf luragin.HandlerFactory) luragin.HandlerFactory {
	return func(configuration *config.EndpointConfig, proxy proxy.Proxy) gin.HandlerFunc {
		next := hf(configuration, proxy)

		v := auth.ConfigGetter(configuration.ExtraConfig)
		if v == nil {
			return next
		}
		credentials, ok := v.(auth.Credentials)
		if !ok {
			return next
		}

		validator := auth.NewCredentialsValidator(credentials)

		return func(c *gin.Context) {
			if !validator.IsValid(c.GetHeader("Authorization")) {
				c.String(http.StatusForbidden, "wrong auth header")
				return
			}
			next(c)
		}
	}
}
