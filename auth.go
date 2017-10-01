package auth

import (
	"crypto/subtle"
	"encoding/base64"
)

// Validator defines the interface for all the possible validation processes
type Validator interface {
	IsValid(subject string) bool
}

// NewCredentialsValidator creates a validator for a given credentials pair
func NewCredentialsValidator(credentials Credentials) Validator {
	base := credentials.User + ":" + credentials.Pass
	header := "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
	return authHeader{int32(len(header)), []byte(header)}
}

type authHeader struct {
	lenght  int32
	content []byte
}

// IsValid implements the Validator interface
func (a authHeader) IsValid(subject string) bool {
	if subtle.ConstantTimeEq(int32(len(subject)), a.lenght) == 1 {
		return subtle.ConstantTimeCompare([]byte(subject), a.content) == 1
	}
	// Securely compare actual to itself to keep constant time, but always return false.
	return subtle.ConstantTimeCompare(a.content, a.content) == 1 && false
}
