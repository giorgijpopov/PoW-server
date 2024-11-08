package api

import (
	"encoding/json"
	"fmt"
	"net"
	"testing"

	"PoW-server/challenge"

	"github.com/stretchr/testify/assert"
)

type DummyChallenge struct {
	Value1 string `json:"value_1"`
	Value2 string `json:"value_2"`
}

var _ challenge.Challenge = (*DummyChallenge)(nil)

func (d DummyChallenge) Type() string {
	return "dummy"
}

func (d DummyChallenge) Payload() ([]byte, error) {
	return json.Marshal(d)
}

func (d DummyChallenge) CheckSolution(solution string) error {
	if solution == d.Value1+d.Value2 {
		return nil
	}
	return fmt.Errorf("Invalid solution")
}

func TestChallengeResponseProtocolWrapper(t *testing.T) {
	wrapper := NewChallengeResponseProtocolWrapper(challenge.NewGeneratorMock(DummyChallenge{
		Value1: "a",
		Value2: "b",
	}))

	createConnection := func() net.Conn {
		serverConn, clientConn := net.Pipe()

		go func(conn net.Conn) {
			defer conn.Close()
			wrapper.Wrap(conn, func(conn net.Conn) error {
				_, err := conn.Write([]byte("success"))
				return err
			})
		}(serverConn)
		return clientConn
	}

	t.Run("ChallengeFailure", func(t *testing.T) {
		clientConn := createConnection()
		defer clientConn.Close()

		var challengeReq ChallengeRequest
		assert.NoError(t, json.NewDecoder(clientConn).Decode(&challengeReq))
		assert.Equal(t, "dummy", challengeReq.ChallengeType)
		var ch DummyChallenge
		assert.NoError(t, json.Unmarshal(challengeReq.Payload, &ch))
		assert.Equal(t, "a", ch.Value1)
		assert.Equal(t, "b", ch.Value2)

		assert.NoError(t, json.NewEncoder(clientConn).Encode(ChallengeResponse{Solution: "c"}))

		n, err := clientConn.Read(make([]byte, 7))
		assert.Equal(t, 0, n)
		assert.ErrorContains(t, err, "EOF")
	})

	t.Run("ChallengeSuccess", func(t *testing.T) {
		clientConn := createConnection()
		defer clientConn.Close()

		var challengeReq ChallengeRequest
		assert.NoError(t, json.NewDecoder(clientConn).Decode(&challengeReq))
		assert.Equal(t, "dummy", challengeReq.ChallengeType)
		var ch DummyChallenge
		assert.NoError(t, json.Unmarshal(challengeReq.Payload, &ch))
		assert.Equal(t, "a", ch.Value1)
		assert.Equal(t, "b", ch.Value2)

		assert.NoError(t, json.NewEncoder(clientConn).Encode(ChallengeResponse{Solution: "ab"}))
		buf := make([]byte, 7)
		_, err := clientConn.Read(buf)
		assert.NoError(t, err)
		assert.Equal(t, "success", string(buf))
	})
}
