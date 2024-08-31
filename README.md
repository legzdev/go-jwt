<div align="center">
<h1>go-jwt</h1>

Simple, minimal and zero dependency Go JWT library.

[![Go Version](https://img.shields.io/badge/go-1.22.3-007d9c?logo=go)](https://go.dev)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghostsama2503/go-jwt.svg)](https://pkg.go.dev/github.com/ghostsama2503/go-jwt)
[![JWT Logo](https://img.shields.io/badge/JWT-gray?logo=JSON%20web%20tokens&logoColor=white)](https://jwt.io)

</div>

<br>

## Installation
```shell
go get -u github.com/ghostsama2503/go-jwt
```

## Building and signing a token
```go
package main

import (
	"fmt"
	"time"

	"github.com/ghostsama2503/go-jwt"
)

func main() {
	secretKey := []byte("secret")
	currentTime := time.Now().UTC()

	claims := &jwt.CommonClaims{
		Subject:        "test",
		ExpirationTime: currentTime.AddDate(0, 1, 0).Unix(),
		IssuedAtTime:   currentTime.Unix(),
	}

	token := jwt.New(claims)

	signedToken, err := token.Signed(secretKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(signedToken)
}
```

## Parsing and validating a token
```go
package main

import (
	"fmt"
	"os"

	"github.com/ghostsama2503/go-jwt"
)

func main() {
	secretKey := []byte("secret")
	tokenString := os.Getenv("TOKEN")

	token, err := jwt.Parse(tokenString)
	if err != nil {
		panic(err)
	}

	if err := token.Validate(secretKey); err != nil {
		panic(err)
	}

	fmt.Println("valid token")
}
```

## Disclaimer

It's provided as-is, with no guarantees or warranties. It's a lightweight and easy-to-use solution for basic needs. For more robust solutions, please consider alternatives like [golang-jwt/jwt](https://github.com/golang-jwt/jwt).
