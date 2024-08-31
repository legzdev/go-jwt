package jwt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Header struct {
	Type      string `json:"typ"`
	Algorithm string `json:"alg"`
}

type Token struct {
	Header    *Header
	Payload   Claims
	Signature []byte
}

func New(claims Claims) *Token {
	token := &Token{}
	token.Header = &Header{Type: "JWT", Algorithm: "HS256"}
	token.Payload = claims
	token.Signature = make([]byte, 0)
	return token
}

func (token *Token) Encoded() (string, error) {
	var encodedToken string

	headerJson, err := json.Marshal(token.Header)
	if err != nil {
		return encodedToken, err
	}

	payloadJson, err := json.Marshal(token.Payload)
	if err != nil {
		return encodedToken, err
	}

	return fmt.Sprintf("%s.%s", b64encode(headerJson), b64encode(payloadJson)), nil
}

func (token *Token) Signed(secretKey []byte) (string, error) {
	encodedToken, err := token.Encoded()
	if err != nil {
		return "", err
	}

	hasher := hmac.New(sha256.New, secretKey)
	hasher.Write([]byte(encodedToken))

	return fmt.Sprintf("%s.%s", encodedToken, b64encode(hasher.Sum(nil))), nil
}

func (token *Token) Validate(secretKey []byte) error {
	currentTime := time.Now().Unix()

	if token.Payload.GetExpirationTime() <= currentTime {
		return ErrExpiredToken
	}

	encodedToken, err := token.Encoded()
	if err != nil {
		return err
	}

	signature, err := Sign(encodedToken, secretKey)
	if err != nil {
		return err
	}

	if !bytes.Equal(signature, token.Signature) {
		return ErrInvalidSignature
	}

	return nil
}

func parse(tokenString string, claims Claims) (*Token, error) {
	tokenParts := strings.Split(tokenString, ".")
	tokenPartsLenght := len(tokenParts)

	if tokenPartsLenght != 2 && tokenPartsLenght != 3 {
		return nil, ErrInvalidTokenFormat
	}

	tokenHeader, err := b64decode(tokenParts[0])
	if err != nil {
		return nil, err
	}

	tokenPayload, err := b64decode(tokenParts[1])
	if err != nil {
		return nil, err
	}

	token := &Token{}
	token.Header = &Header{}
	token.Payload = claims

	if err := json.Unmarshal(tokenHeader, token.Header); err != nil {
		return nil, err
	}

	claimsPtr := &claims

	if err := json.Unmarshal(tokenPayload, claimsPtr); err != nil {
		return nil, err
	}

	token.Payload = claims

	if tokenPartsLenght == 2 {
		token.Signature = make([]byte, 0)
		return token, nil
	}

	tokenSignature, err := b64decode(tokenParts[2])
	if err != nil {
		return nil, err
	}

	token.Signature = tokenSignature

	return token, nil
}

func Parse(tokenString string) (*Token, error) {
	return parse(tokenString, &CommonClaims{})
}

func ParseWithClaims(tokenString string, claims Claims) (*Token, error) {
	return parse(tokenString, claims)
}

func Sign(encodedToken string, secretKey []byte) ([]byte, error) {
	hasher := hmac.New(sha256.New, secretKey)

	if _, err := hasher.Write([]byte(encodedToken)); err != nil {
		return nil, err
	}

	return hasher.Sum(nil), nil
}
