package itsdangerous

import (
	"bytes"
	"compress/zlib"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	zip "itsdangerous/zlib"
	"strings"
)

// Signature can sign bytes and unsign it and validate the signature
// provided.
//
// Salt can be used to namespace the hash, so that a signed string is only
// valid for a given namespace.  Leaving this at the default value or re-using
// a salt value across different parts of your application where the same
// signed value in one part can mean something different in another part
// is a security risk.
type Signature struct {
	SecretKey     string
	Sep           string
	Salt          string
	KeyDerivation string
	DigestMethod  func() hash.Hash
	Algorithm     SigningAlgorithm
}

// DeriveKey generates a key derivation. Keep in mind that the key derivation in itsdangerous
// is not intended to be used as a security method to make a complex key out of a short password.
// Instead you should use large random secret keys.
func (s *Signature) DeriveKey() ([]byte, error) {
	var key []byte
	var err error

	s.DigestMethod().Reset()

	switch s.KeyDerivation {
	case "concat":
		h := s.DigestMethod()
		h.Write([]byte(s.Salt + s.SecretKey))
		key = h.Sum(nil)
	case "django-concat":
		h := s.DigestMethod()
		h.Write([]byte(s.Salt + "signer" + s.SecretKey))
		key = h.Sum(nil)
	case "hmac":
		h := hmac.New(func() hash.Hash { return s.DigestMethod() }, []byte(s.SecretKey))
		h.Write([]byte(s.Salt))
		key = h.Sum(nil)
	//case "none":
	//	key = s.SecretKey
	default:
		key, err = nil, errors.New("unknown key derivation method")
	}
	return key, err
}

func (s *Signature) Zip(value string) (string, error) {
	deflateString, err := zip.DeflateString(value)
	if err != nil {
		return "", err
	}

	return deflateString, nil

}
func (s *Signature) UnZip(value string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, err
	}
	r, _ := zlib.NewReader(bytes.NewReader(decoded))
	defer func(r io.ReadCloser) {
		err := r.Close()
		if err != nil {

		}
	}(r)
	var out bytes.Buffer
	_, err = io.Copy(&out, r)

	return out.Bytes(), nil
}

// Get returns the signature for the given value.
func (s *Signature) Get(value string) (string, error) {
	key, err := s.DeriveKey()
	if err != nil {
		return "", err
	}

	sig := s.Algorithm.GetSignature(key, []byte(value))
	return Base64Encode(sig), err
}

// Verify verifies the signature for the given value.
func (s *Signature) Verify(value, sig string) (bool, error) {
	key, err := s.DeriveKey()
	if err != nil {
		return false, err
	}

	signed, err := base64Decode(sig)
	if err != nil {
		return false, err
	}
	return s.Algorithm.VerifySignature(key, []byte(value), signed), nil
}

// Sign the given string.
func (s *Signature) Sign(value string) (string, error) {
	encode, err := s.Zip(value)
	if err != nil {
		return "", err
	}
	value = "." + encode
	sig, err := s.Get(value)
	if err != nil {
		return "", err
	}
	return value + s.Sep + sig, nil

}

// Unsign the given string.
func (s *Signature) Unsign(signed string) (string, error) {
	if !strings.Contains(signed, s.Sep) {
		return "", fmt.Errorf("no %s found in value", s.Sep)
	}

	li := strings.LastIndex(signed, s.Sep)
	value, sig := signed[:li], signed[li+len(s.Sep):]

	if ok, _ := s.Verify(value, sig); ok == true {
		return value, nil
	}
	return "", fmt.Errorf("signature %s does not match", sig)
}

func (s *Signature) Dumps(value interface{}) (string, error) {
	str_b, _ := json.Marshal(value)
	str := string(str_b)

	sign, err := s.Sign(str)
	if err != nil {
		return "", err
	}
	return sign, nil
}

func (s *Signature) Loads(value string) (string, error) {
	unsign, err := s.Unsign(value)
	if err != nil {
		return "", err
	}
	li := strings.LastIndex(unsign, s.Sep)
	value = unsign[li+len(s.Sep):]
	unzip, err := s.UnZip(value)
	if err != nil {
		return "", err
	}
	return string(unzip), nil
}

// NewSignature creates a new Signature
func NewSignature(secret, salt, sep, derivation string, digest func() hash.Hash, algo SigningAlgorithm) *Signature {
	if salt == "" {
		salt = "itsdangerous.Signer"
	}
	if sep == "" {
		sep = "."
	}
	if derivation == "" {
		derivation = "django-concat"
	}
	if digest == nil {
		digest = sha1.New
	}
	if algo == nil {
		algo = &HMACAlgorithm{DigestMethod: digest}
	}
	return &Signature{
		SecretKey:     secret,
		Salt:          salt,
		Sep:           sep,
		KeyDerivation: derivation,
		DigestMethod:  digest,
		Algorithm:     algo,
	}
}

// TimestampSignature works like the regular Signature but also records the time
// of the signing and can be used to expire signatures.
type TimestampSignature struct {
	Signature
}

// Sign the given string.
func (s *TimestampSignature) Sign(value string) (string, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, getTimestamp()); err != nil {
		return "", err
	}

	ts := Base64Encode(buf.Bytes())
	val := value + s.Sep + ts

	sig, err := s.Get(val)
	if err != nil {
		return "", err
	}
	return val + s.Sep + sig, nil
}

// Unsign the given string.
func (s *TimestampSignature) Unsign(value string, maxAge uint32) (string, error) {
	var timestamp uint32

	result, err := s.Signature.Unsign(value)
	if err != nil {
		return "", err
	}

	// If there is no timestamp in the result there is something seriously wrong.
	if !strings.Contains(result, s.Sep) {
		return "", errors.New("timestamp missing")
	}

	li := strings.LastIndex(result, s.Sep)
	val, ts := result[:li], result[li+len(s.Sep):]

	sig, err := base64Decode(ts)
	if err != nil {
		return "", err
	}

	buf := bytes.NewReader([]byte(sig))
	if err = binary.Read(buf, binary.BigEndian, &timestamp); err != nil {
		return "", err
	}

	if maxAge > 0 {
		if age := getTimestamp() - timestamp; age > maxAge {
			return "", fmt.Errorf("signature age %d > %d seconds", age, maxAge)
		}
	}
	return val, nil
}

// NewTimestampSignature creates a new TimestampSignature
func NewTimestampSignature(secret, salt, sep, derivation string, digest func() hash.Hash, algo SigningAlgorithm) *TimestampSignature {
	s := NewSignature(secret, salt, sep, derivation, digest, algo)
	return &TimestampSignature{Signature: *s}
}
