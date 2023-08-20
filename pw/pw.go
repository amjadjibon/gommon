package pw

type Password interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) (bool, error)
}

const DefaultIteration = 1
const DefaultSaltSize = 12

const (
	Argon2       = "argon2"
	Argon2i      = "argon2i"
	Argon2id     = "argon2id"
	BCrypt       = "bcrypt"
	BCryptSHA256 = "bcrypt_sha256"
	Crypt        = "crypt"
	MD5          = "md5"
	PBKDF2SHA1   = "pbkdf2_sha1"
	PBKDF2SHA256 = "pbkdf2_sha256"
	PBKDF2SHA512 = "pbkdf2_sha512"
	SHA1         = "sha1"
	UnsaltedMD5  = "unsalted_md5"
	UnsaltedSHA1 = "unsalted_sha1"
)
