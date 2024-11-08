package challenge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHashInversionChallenge(t *testing.T) {
	ch := HashInversionChallenge{
		Start:          "583213",
		RequiredPrefix: "eb000dca",
	}

	assert.Equal(t, "hash_inversion", ch.Type())
	assert.NoError(t, ch.CheckSolution("1234565"))
	assert.Error(t, ch.CheckSolution("1234566"))
	assert.Error(t, ch.CheckSolution(""))
}
