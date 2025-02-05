package auth

import (
	"github.com/go-webauthn/webauthn/webauthn"
)

var webAuthn *webauthn.WebAuthn

func InitWebAuthn() error {
	var err error
	webAuthn, err = webauthn.New(&webauthn.Config{
		RPDisplayName: "HD Wallet",
		RPID:          "localhost",
		RPOrigin:      "http://localhost:8080",
	})
	return err
}

func GetWebAuthn() *webauthn.WebAuthn {
	return webAuthn
}
