package challenge

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	prefixLength  = 6
	maxChallenge  = 1000000
	hexChars      = "0123456789abcdef"
	challengeType = "hash_inversion"
)

// HashInversionChallenge is a challenge where the client has to find a string,
// that after being concatenated with a given 'Start' string and hashed with SHA-256,
// gives a hash, that in hexadecimal representation starts with a given prefix.
type HashInversionChallenge struct {
	Start          string `json:"start"`
	RequiredPrefix string `json:"prefix"`
}

var _ Challenge = (*HashInversionChallenge)(nil)

func NewHashInversionChallenge() Challenge {
	challenge := strconv.Itoa(rand.Intn(maxChallenge))
	prefix := generateRandomPrefix(prefixLength)

	return HashInversionChallenge{
		Start:          challenge,
		RequiredPrefix: prefix,
	}
}

func (ch HashInversionChallenge) Type() string {
	return challengeType
}

func (ch HashInversionChallenge) Payload() ([]byte, error) {
	return json.Marshal(ch)
}

func (ch HashInversionChallenge) CheckSolution(solution string) error {
	hash := sha256.Sum256([]byte(ch.Start + solution))
	hashHex := hex.EncodeToString(hash[:])

	if len(hashHex) >= len(ch.RequiredPrefix) && hashHex[:len(ch.RequiredPrefix)] == ch.RequiredPrefix {
		return nil
	}
	return fmt.Errorf("Invalid solution")
}

func generateRandomPrefix(length int) string {
	prefix := make([]byte, length)
	for i := range prefix {
		prefix[i] = hexChars[rand.Intn(len(hexChars))]
	}
	return string(prefix)
}
