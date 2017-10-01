package auth

import (
	"github.com/devopsfaith/krakend/config"
)

const Namespace = "github.com/kpacha/krakend-http-auth"

type Credentials struct {
	User string
	Pass string
}

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
