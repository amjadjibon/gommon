package pw

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// BCryptHasher implements Bcrypt password hasher.
type BCryptHasher struct {
	// Algorithm identifier.
	Algorithm string
	// Defines the hash function used to avoid bcrypt's 72 bytes password truncation.
	Digest func() hash.Hash
	// Defines the number of rounds used to encode the password.
	Cost int
}

// HashPassword returns a bcrypt hashed password.
func (h *BCryptHasher) HashPassword(password string) (string, error) {
	if h.Digest != nil {
		d := h.Digest()
		d.Write([]byte(password))
		password = hex.EncodeToString(d.Sum(nil))
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.Cost)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s$%s", h.Algorithm, string(bytes)), nil
}

// VerifyPassword compares a bcrypt hashed password with its possible cleartext
// equivalent. Returns true if the password and hash match, otherwise false.
func (h *BCryptHasher) VerifyPassword(password, encoded string) (bool, error) {
	s := strings.SplitN(encoded, "$", 2)

	if len(s) != 2 {
		return false, errors.New("invalid hash format")
	}

	algorithm, hashedPw := s[0], s[1]

	if algorithm != h.Algorithm {
		return false, errors.New("invalid algorithm")
	}

	if h.Digest != nil {
		d := h.Digest()
		d.Write([]byte(password))
		password = hex.EncodeToString(d.Sum(nil))
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

// NewBCryptHasher returns a new BCryptHasher.
func NewBCryptHasher(cost int) Password {
	return &BCryptHasher{
		Algorithm: BCrypt,
		Cost:      cost,
	}
}

// NewBCryptSHA256Hasher returns a new BCryptHasher with SHA256 digest.
func NewBCryptSHA256Hasher(cost int) Password {
	return &BCryptHasher{
		Algorithm: BCryptSHA256,
		Digest:    sha256.New,
		Cost:      cost,
	}
}
