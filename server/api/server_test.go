package api

import (
	"context"
	"encoding/json"
	"net"
	"testing"
	"time"

	"PoW-server/challenge"
	"PoW-server/quote"

	"github.com/stretchr/testify/assert"
)

func startTestServer(ctx context.Context) {
	go StartServer(ctx, quote.NewProviderMock("mock quote"), NewChallengeResponseProtocolWrapper(challenge.NewGeneratorMock(challenge.HashInversionChallenge{
		Start:          "278329",
		RequiredPrefix: "1289fe",
	})))
	time.Sleep(time.Second)
}

func startClient(t *testing.T) net.Conn {
	conn, err := net.Dial("tcp", "localhost:8080")
	assert.NoError(t, err)
	return conn
}

func TestServer(t *testing.T) {
	startTestServer(context.Background())

	t.Run("HappyPath", func(t *testing.T) {
		conn := startClient(t)
		defer conn.Close()

		var challengeReq ChallengeRequest
		assert.NoError(t, json.NewDecoder(conn).Decode(&challengeReq))
		assert.Equal(t, "hash_inversion", challengeReq.ChallengeType)

		var ch challenge.HashInversionChallenge
		assert.NoError(t, json.Unmarshal(challengeReq.Payload, &ch))
		assert.Equal(t, "278329", ch.Start)
		assert.Equal(t, "1289fe", ch.RequiredPrefix)

		assert.NoError(t, json.NewEncoder(conn).Encode(ChallengeResponse{Solution: "28391232"}))

		var quoteRes Response
		assert.NoError(t, json.NewDecoder(conn).Decode(&quoteRes))
		assert.Equal(t, "mock quote", quoteRes.Quote)
	})

	t.Run("WrongSolution", func(t *testing.T) {
		conn := startClient(t)
		defer conn.Close()

		var challengeReq ChallengeRequest
		assert.NoError(t, json.NewDecoder(conn).Decode(&challengeReq))
		assert.NoError(t, json.NewEncoder(conn).Encode(ChallengeResponse{Solution: "28391233"}))

		var errResp Response
		assert.NoError(t, json.NewDecoder(conn).Decode(&errResp))
		assert.Equal(t, "Invalid solution", errResp.Error)
	})
}
