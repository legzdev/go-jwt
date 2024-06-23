package jwt

import "encoding/json"

type Claims interface {
	Marshal() ([]byte, error)
	GetExpirationTime() int64
}

type CommonClaims struct {
	Issuer         string `json:"iss"`
	Subject        string `json:"sub"`
	ExpirationTime int64  `json:"exp"`
	IssuedAtTime   int64  `json:"iat"`
}

func (claims *CommonClaims) Marshal() ([]byte, error) {
	return json.Marshal(claims)
}

func (claims *CommonClaims) GetExpirationTime() int64 {
	return claims.ExpirationTime
}
