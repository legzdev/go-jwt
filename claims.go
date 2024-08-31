package jwt

type Claims interface {
	GetIssuer() string
	GetSubject() string
	GetExpirationTime() int64
	GetIssuedAtTime() int64
}

type CommonClaims struct {
	Issuer         string `json:"iss"`
	Subject        string `json:"sub"`
	ExpirationTime int64  `json:"exp"`
	IssuedAtTime   int64  `json:"iat"`
}

func (claims *CommonClaims) GetIssuer() string {
	return claims.Issuer
}

func (claims *CommonClaims) GetSubject() string {
	return claims.Subject
}

func (claims *CommonClaims) GetExpirationTime() int64 {
	return claims.ExpirationTime
}

func (claims *CommonClaims) GetIssuedAtTime() int64 {
	return claims.IssuedAtTime
}
