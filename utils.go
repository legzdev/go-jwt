package jwt

import "encoding/base64"

func b64encode(data []byte) []byte {
	dst := make([]byte, base64.RawURLEncoding.EncodedLen(len(data)))
	base64.RawURLEncoding.Encode(dst, data)
	return dst
}

func b64decode(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}
