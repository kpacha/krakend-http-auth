package auth

import (
	"encoding/base64"
	"testing"
)

func TestIsValid(t *testing.T) {
	v := NewCredentialsValidator(Credentials{User: "a", Pass: "b"})

	for k, pair := range [][]string{
		[]string{"a", "be"},
		[]string{"aaa", "be"},
	} {
		h := pair[0] + ":" + pair[1]
		if v.IsValid("Basic " + base64.StdEncoding.EncodeToString([]byte(h))) {
			t.Error("Unexpected ok for the subject #d", k)
		}
	}

	if !v.IsValid("Basic " + base64.StdEncoding.EncodeToString([]byte("a:b"))) {
		t.Error("Unexpected ko for the pair a:b")
	}
}
