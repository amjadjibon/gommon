package pw

import (
	"fmt"

	"golang.org/x/crypto/argon2"
)

// Argon2Hasher implements Argon2i password hasher.
type Argon2Hasher struct {
	// Algorithm identifier.
	Algorithm string
	// Defines the amount of computation time, given in number of iterations.
	Time uint32
	// Defines the memory usage (KiB).
	Memory uint32
	// Defines the number of parallel threads.
	Threads uint8
	// Defines the length of the hash in bytes.
	Length uint32

	// Mode defines the Argon2 mode.
	// Possible values are:
	// - argon2i
	// - argon2id
	Mode string
}

func (a Argon2Hasher) Hash(password, salt string) (string, error) {
	var b64Hash []byte
	switch a.Algorithm {
	case Argon2i:
		b64Hash = argon2.Key([]byte(password), []byte(salt), a.Time, a.Memory, a.Threads, a.Length)
	case Argon2id:
		b64Hash = argon2.IDKey([]byte(password), []byte(salt), a.Time, a.Memory, a.Threads, a.Length)
	}

	return fmt.Sprintf("%s$%s$v=%d$m=%d,t=%d,p=%d$%s$%s",
		a.Algorithm,
		a.Mode,
		argon2.Version,
		a.Memory,
		a.Time,
		a.Threads,
		salt,
		b64Hash,
	), nil
}

func (a Argon2Hasher) HashPassword(password string) (string, error) {
	return a.Hash(password, GetRandomString(DefaultSaltSize))
}

func (a Argon2Hasher) VerifyPassword(password, hash string) (bool, error) {
	// TODO implement me
	panic("implement me")
}

func DefaultArgon2Hasher() *Argon2Hasher {
	return &Argon2Hasher{
		Algorithm: Argon2,
		Time:      1,
		Memory:    512,
		Threads:   2,
		Length:    16,
	}
}

// Hash generates a hash from a password.
