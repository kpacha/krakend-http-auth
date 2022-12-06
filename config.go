package auth

import (
	"github.com/luraproject/lura/config"
)

// Namespace is the key to look for extra configuration details
const Namespace = "github.com/kpacha/krakend-http-auth"

// Credentials contains the pair user:pass
type Credentials struct {
	User string
	Pass string
}

// ConfigGetter extracts the credentials from the extra config details
func ConfigGetter(e config.ExtraConfig) interface{} {
	cfg, ok := e[Namespace]
	if !ok {
		return nil
	}
	data, ok := cfg.(map[string]interface{})
	if !ok {
		return nil
	}

	v, ok := data["user"]
	if !ok {
		return nil
	}

	user, ok := v.(string)
	if !ok {
		return nil
	}

	v, ok = data["pass"]
	if !ok {
		return nil
	}

	pass, ok := v.(string)
	if !ok {
		return nil
	}

	return Credentials{user, pass}
}
