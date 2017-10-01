package mux

import (
	"net/http"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/proxy"
	krakendmux "github.com/devopsfaith/krakend/router/mux"
	auth "github.com/kpacha/krakend-http-auth"
)

// HandlerFactory decorates a krakendmux.HandlerFactory with the auth layer
func HandlerFactory(hf krakendmux.HandlerFactory) krakendmux.HandlerFactory {
	return func(configuration *config.EndpointConfig, proxy proxy.Proxy) http.HandlerFunc {
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

		return func(w http.ResponseWriter, r *http.Request) {
			if !validator.IsValid(r.Header.Get("Authorization")) {
				http.Error(w, "wrong auth header", http.StatusForbidden)
				return
			}
			next(w, r)
		}
	}
}
