package secp256k1

import (
	"github.com/dchest/blake256"
)

// HashBlake256 calculates hash(b) and returns the resulting bytes.
func HashBlake256(b []byte) []byte {
	hash := blake256.Sum256(b)
	return hash[:]
}
