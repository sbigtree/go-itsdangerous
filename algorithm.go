package itsdangerous

import (
	"crypto/hmac"
	"crypto/subtle"
	"hash"
)

// SigningAlgorithm provides interfaces to generate and verify signature
type SigningAlgorithm interface {
	GetSignature(key []byte, value []byte) []byte
	VerifySignature(key []byte, value []byte, sig []byte) bool
}

// HMACAlgorithm provides signature generation using HMACs.
type HMACAlgorithm struct {
	DigestMethod func() hash.Hash
}

// GetSignature returns the signature for the given key and value.
func (a *HMACAlgorithm) GetSignature(key []byte, value []byte) []byte {
	a.DigestMethod().Reset()
	h := hmac.New(func() hash.Hash { return a.DigestMethod() }, key)
	h.Write(value)
	return h.Sum(nil)
}

// VerifySignature verifies the given signature matches the expected signature.
func (a *HMACAlgorithm) VerifySignature(key []byte, value []byte, sig []byte) bool {
	eq := subtle.ConstantTimeCompare(sig, []byte(a.GetSignature(key, value)))
	return eq == 1
}
