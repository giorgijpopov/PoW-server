package challenge

import (
	"context"
	"encoding/json"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {

	createConnection := func(ctx context.Context) (net.Conn, chan error) {
		serverConn, clientConn := net.Pipe()
		errChan := make(chan error)
		go func(conn net.Conn) {
			defer conn.Close()
			errChan <- SolveChallenge(ctx, conn)
		}(clientConn)
		return serverConn, errChan
	}

	t.Run("Success", func(t *testing.T) {
		serverConn, errChan := createConnection(context.Background())
		defer serverConn.Close()

		ch := Challenge{
			ChallengeType: HashInversionChallengeType,
			Payload:       []byte(`{"start":"1234567","prefix":"afc4"}`),
		}
		assert.NoError(t, json.NewEncoder(serverConn).Encode(ch))

		var res PoWResponse
		assert.NoError(t, json.NewDecoder(serverConn).Decode(&res))
		assert.True(t, checkSolution("1234567", "afc4", res.Solution))

		assert.NoError(t, <-errChan)
	})

	t.Run("UnknownChallengeType", func(t *testing.T) {
		serverConn, errChan := createConnection(context.Background())
		defer serverConn.Close()

		ch := Challenge{
			ChallengeType: "puzzles",
		}
		assert.NoError(t, json.NewEncoder(serverConn).Encode(ch))
		assert.ErrorContains(t, <-errChan, "unknown challenge type: puzzles")
	})

	t.Run("Timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 0)
		defer cancel()

		serverConn, errChan := createConnection(ctx)
		defer serverConn.Close()

		ch := Challenge{
			ChallengeType: HashInversionChallengeType,
			Payload:       []byte(`{"start":"1234567","prefix":"afc4"}`),
		}
		assert.NoError(t, json.NewEncoder(serverConn).Encode(ch))

		assert.ErrorContains(t, <-errChan, "context deadline exceeded")
	})
}
