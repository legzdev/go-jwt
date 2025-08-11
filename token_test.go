package jwt_test

import (
	"testing"
	"time"

	"github.com/legzdev/go-jwt"
)

var secretKey = []byte("mysecretkey")

func TestMain(t *testing.T) {
	now := time.Now()

	claims := &jwt.CommonClaims{
		Subject:        "test",
		ExpirationTime: now.AddDate(0, 1, 0).Unix(),
		IssuedAtTime:   now.Unix(),
	}

	validToken := jwt.New(claims)

	validTokenSigned, err := validToken.Signed(secretKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("signed token:", validTokenSigned)

	validTokenParsed, err := jwt.Parse(validTokenSigned)
	if err != nil {
		t.Fatal("parsing token:", err)
	}

	if err := validTokenParsed.Validate(secretKey); err != nil {
		t.Fatal(err)
	}

	validPayload := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJhdXRoLXRlc3QiLCJzdWIiOiJ0ZXN0LXVzZXIiLCJleHAiOjE3MjE3NTE1MjEsImlhdCI6MTcxOTE1OTUyMX0"
	invalidSignature := "Z6INFs_Ajv87ikNTGnVk05NPGRmUmHYK-hgGMiM5CcQ"

	invalidToken, err := jwt.Parse(validPayload + "." + invalidSignature)
	if err != nil {
		t.Fatal(err)
	}

	if err := invalidToken.Validate(secretKey); err == nil {
		t.Fatal("invalid token validation")
	}
}
