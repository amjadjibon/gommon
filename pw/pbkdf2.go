package pw

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"hash"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

type PBKDF2Hasher struct {
	// Algorithm identifier.
	Algorithm string
	// Defines the number of rounds used to encode the password.
	Iterations int
	// Defines the length of the hash in bytes.
	Size int
	// Defines the hash function used to encode the password.
	Digest func() hash.Hash
	// Salt length in bytes.
	SaltLength int
}

// Encode returns a hashed password using PBKDF2.
func (h *PBKDF2Hasher) Encode(password string, salt string, iterations int) (string, error) {
	hashedPw := pbkdf2.Key([]byte(password), []byte(salt), iterations, h.Size, h.Digest)
	b64Hash := base64.StdEncoding.EncodeToString(hashedPw)
	return fmt.Sprintf("%s$%d$%s$%s", h.Algorithm, iterations, salt, b64Hash), nil
}

// HashPassword returns a hashed password using PBKDF2.
func (h *PBKDF2Hasher) HashPassword(password string) (string, error) {
	salt := GetRandomString(h.SaltLength)
	iterations := h.Iterations
	if iterations <= 0 {
		iterations = DefaultIteration
	}
	return h.Encode(password, salt, iterations)
}

// VerifyPassword verifies a hashed password using PBKDF2.
func (h *PBKDF2Hasher) VerifyPassword(password, hashedPw string) (bool, error) {
	splits := strings.Split(hashedPw, "$")
	if len(splits) != 4 {
		return false, fmt.Errorf("invalid hash format")
	}

	algorithm := splits[0]
	iterations := splits[1]
	salt := splits[2]
	b64Hash := splits[3]

	if algorithm != h.Algorithm {
		return false, fmt.Errorf("invalid algorithm")
	}

	if len(iterations) == 0 {
		return false, fmt.Errorf("invalid iterations")
	}

	if len(salt) == 0 {
		return false, fmt.Errorf("invalid salt")
	}

	if len(b64Hash) == 0 {
		return false, fmt.Errorf("invalid hash")
	}

	iterationsInt, err := strconv.Atoi(iterations)
	if err != nil {
		return false, err
	}

	hashedPw2, err := h.Encode(password, salt, iterationsInt)
	if err != nil {
		return false, err
	}

	return hmac.Equal([]byte(hashedPw), []byte(hashedPw2)), nil
}

// NewPBKDF2SHA1Hasher returns a new PBKDF2 SHA1 Hasher
func NewPBKDF2SHA1Hasher(
	iterations int,
) *PBKDF2Hasher {
	return &PBKDF2Hasher{
		Algorithm:  PBKDF2SHA1,
		Iterations: iterations,
		Size:       sha1.Size,
		Digest:     sha1.New,
	}
}

// NewPBKDF2SHA256Hasher returns a new PBKDF2 SHA256 Hasher
func NewPBKDF2SHA256Hasher(
	iterations int,
) *PBKDF2Hasher {
	return &PBKDF2Hasher{
		Algorithm:  PBKDF2SHA256,
		Iterations: iterations,
		Size:       sha256.Size,
		Digest:     sha256.New,
	}
}

// NewPBKDF2SHA512Hasher returns a new PBKDF2 Hasher
func NewPBKDF2SHA512Hasher(
	iterations int,
) *PBKDF2Hasher {
	return &PBKDF2Hasher{
		Algorithm:  PBKDF2SHA512,
		Iterations: iterations,
		Size:       sha512.Size,
		Digest:     sha512.New,
	}
}
