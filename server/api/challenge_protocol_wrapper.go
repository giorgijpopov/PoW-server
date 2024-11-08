package api

import (
	"encoding/json"
	"net"

	"PoW-server/challenge"
)

type ChallengeResponseProtocolWrapper struct {
	challengeGenerator challenge.ChallengeGenerator
}

func NewChallengeResponseProtocolWrapper(challengeGenerator challenge.ChallengeGenerator) *ChallengeResponseProtocolWrapper {
	return &ChallengeResponseProtocolWrapper{challengeGenerator: challengeGenerator}
}

func (c *ChallengeResponseProtocolWrapper) Wrap(conn net.Conn, next func(net.Conn) error) error {
	if err := powChallengeProtocol(conn, c.challengeGenerator); err != nil {
		return err
	}
	return next(conn)
}

type ChallengeRequest struct {
	ChallengeType string `json:"challengeType"`
	Payload       []byte `json:"payload"`
}

type ChallengeResponse struct {
	Solution string `json:"solution"`
}

func powChallengeProtocol(conn net.Conn, challengeGenerator challenge.ChallengeGenerator) error {
	// generate challenge and send it to the client
	ch := challengeGenerator.GenerateChallenge()
	payload, err := ch.Payload()
	if err != nil {
		return err
	}
	err = json.NewEncoder(conn).Encode(ChallengeRequest{
		ChallengeType: ch.Type(),
		Payload:       payload,
	})
	if err != nil {
		return err
	}

	// receive solution and check it
	var response ChallengeResponse
	if err := json.NewDecoder(conn).Decode(&response); err != nil {
		return err
	}
	if err := ch.CheckSolution(response.Solution); err != nil {
		return err
	}
	return nil
}
