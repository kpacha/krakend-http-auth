package auth

import (
	"testing"

	"github.com/devopsfaith/krakend/config"
)

func TestConfigGetter(t *testing.T) {
	v := ConfigGetter(config.ExtraConfig(map[string]interface{}{Namespace: map[string]interface{}{"user": "a", "pass": "b"}}))
	if v == nil {
		t.Fail()
	}
	credentials, ok := v.(Credentials)
	if !ok {
		t.Fail()
	}
	if credentials.User != "a" || credentials.Pass != "b" {
		t.Fail()
	}

	if v = ConfigGetter(config.ExtraConfig(map[string]interface{}{Namespace: map[string]interface{}{"user": "a"}})); v != nil {
		t.Fail()
	}

	if v = ConfigGetter(config.ExtraConfig(map[string]interface{}{Namespace: map[string]interface{}{"user": "a", "pass": true}})); v != nil {
		t.Fail()
	}

	if v = ConfigGetter(config.ExtraConfig(map[string]interface{}{Namespace: map[string]interface{}{"pass": "b"}})); v != nil {
		t.Fail()
	}

	if v = ConfigGetter(config.ExtraConfig(map[string]interface{}{Namespace: map[string]interface{}{"user": 42, "pass": "b"}})); v != nil {
		t.Fail()
	}

	if v = ConfigGetter(config.ExtraConfig(map[string]interface{}{Namespace: true})); v != nil {
		t.Fail()
	}

	if v = ConfigGetter(config.ExtraConfig(map[string]interface{}{"user": "a", "pass": "b"})); v != nil {
		t.Fail()
	}
}
