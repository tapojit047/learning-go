package db

import "github.com/go-chi/jwtauth/v5"

var TokenAuth *jwtauth.JWTAuth
var CredentialList map[string]string

func InitializeCred() {
	CredentialList = map[string]string{
		"tapojit047": "1234",
	}
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}
