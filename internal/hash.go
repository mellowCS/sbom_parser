package internal

type Alg string

const (
	md5        Alg = "MD5"
	sha1       Alg = "SHA-1"
	sha256     Alg = "SHA-256"
	sha384     Alg = "SHA-384"
	sha512     Alg = "SHA-512"
	sha3256    Alg = "SHA3-256"
	sha3384    Alg = "SHA3-384"
	sha3512    Alg = "SHA3-512"
	blake2b256 Alg = "BLAKE2b-256"
	blake2b384 Alg = "BLAKE2b-384"
	blake2b512 Alg = "BLAKE2b-512"
	blake3     Alg = "BLAKE3"
)

type Hash struct {
	Alg Alg
	// Must match regular expression: ^([a-fA-F0-9]{32}|[a-fA-F0-9]{40}|[a-fA-F0-9]{64}|[a-fA-F0-9]{96}|[a-fA-F0-9]{128})$
	Content string
}
