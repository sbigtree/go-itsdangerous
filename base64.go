package itsdangerous

import (
	"encoding/base64"
	"strings"
	"unicode"
)

func isValidBase64(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '+' || r == '/' || r == '=') {
			return false
		}
	}
	if len(s)%4 != 0 {
		return false
	}
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return false
	}
	// 严格：解码再编码应一致
	return base64.StdEncoding.EncodeToString(b) == s
}

func isValidBase64URL(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' || r == '=') {
			return false
		}
	}
	// 兼容有/无 padding 的 base64url
	var (
		b   []byte
		err error
	)
	if strings.Contains(s, "=") {
		b, err = base64.URLEncoding.DecodeString(s)
		if err != nil {
			return false
		}
		// 允许带 padding 的 base64url
		return base64.URLEncoding.EncodeToString(b) == s
	}
	b, err = base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return false
	}
	// 无 padding 的 canonical 校验
	return base64.RawURLEncoding.EncodeToString(b) == s
}

func base64ToBase64URL(b64 string) string {
	// 输出无 padding 的 base64url
	return strings.TrimRight(strings.NewReplacer("+", "-", "/", "_").Replace(b64), "=")
}

func base64URLToBase64(b64url string) string {
	b64 := strings.NewReplacer("-", "+", "_", "/").Replace(b64url)
	pad := (4 - len(b64)%4) % 4
	return b64 + strings.Repeat("=", pad)
}

// 单段：转 base64url（不是 base64/base64url 就原样返回）
func ToBase64URLIfPossible(input string) string {
	if isValidBase64(input) {
		return base64ToBase64URL(input)
	}
	if isValidBase64URL(input) {
		// 统一成无 padding 的 base64url
		if strings.Contains(input, "=") {
			// 先解再编，得到 canonical raw url 形式
			b, err := base64.URLEncoding.DecodeString(input)
			if err == nil {
				return base64.RawURLEncoding.EncodeToString(b)
			}
		}
		return input
	}
	return input
}

// 单段：转 base64（不是 base64/base64url 就原样返回）
func ToBase64IfPossible(input string) string {
	if isValidBase64(input) {
		return input
	}
	if isValidBase64URL(input) {
		return base64URLToBase64(strings.TrimRight(input, "="))
	}
	return input
}

// 多段（按 . 分段）转 base64url
func ToBase64URLByDotIfPossible(token string) string {
	parts := strings.Split(token, ".")
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = ToBase64URLIfPossible(p)
	}
	return strings.Join(parts, ".")
}

// 多段（按 . 分段）转 base64
func ToBase64ByDotIfPossible(token string) string {
	parts := strings.Split(token, ".")
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = ToBase64IfPossible(p)
	}
	return strings.Join(parts, ".")
}
