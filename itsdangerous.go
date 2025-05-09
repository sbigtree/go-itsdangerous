/*
Package itsdangerous implements various functions to deal with untrusted sources.
Mainly useful for web applications.

This package exists purely as a port of https://github.com/mitsuhiko/itsdangerous,
where the original version is written in Python.
*/
package itsdangerous

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

// 2011/01/01 in UTC
const EPOCH = 1293840000

// Encodes a single string. The resulting string is safe for putting into URLs.
func Base64Encode(src []byte) string {
	s := base64.URLEncoding.EncodeToString(src)
	return strings.Trim(s, "=")
}

// Decodes a single string.
func base64Decode(s string) ([]byte, error) {
	var padLen int
	if l := len(s) % 4; l > 0 {
		padLen = 4 - l
		s = s + strings.Repeat("=", padLen)
	}
	// 将 url_base64 转成  标准的base64
	s = strings.ReplaceAll(s, "-", "+")
	s = strings.ReplaceAll(s, "_", "/")

	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println(s)
		return []byte(""), err
	}
	return b, nil
}

// Returns the current timestamp.  This implementation returns the
// seconds since 1/1/2011.
func getTimestamp() uint32 {
	return uint32(time.Now().Unix() - EPOCH)
}
