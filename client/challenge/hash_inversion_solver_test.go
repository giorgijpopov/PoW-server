package challenge

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveHashInversionChallenge(t *testing.T) {
	challenge := HashInversionChallenge{
		Start:          "1234567",
		RequiredPrefix: "afc4",
	}

	payload, err := json.Marshal(challenge)
	assert.NoError(t, err)

	t.Run("TestSolveHashInversionChallenge", func(t *testing.T) {
		ctx := context.Background()
		res, err := solveHashInversionChallenge(ctx, Challenge{Payload: payload})
		assert.NoError(t, err)

		assert.True(t, checkSolution(challenge.Start, challenge.RequiredPrefix, res))
	})

	t.Run("Timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 0)
		defer cancel()

		_, err := solveHashInversionChallenge(ctx, Challenge{Payload: payload})
		assert.ErrorContains(t, err, "context deadline exceeded")
	})
}

func checkSolution(start, prefix, solution string) bool {
	hash := sha256.Sum256([]byte(start + solution))
	hashHex := hex.EncodeToString(hash[:])

	return len(hashHex) >= len(prefix) && hashHex[:len(prefix)] == prefix
}
