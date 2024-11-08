package challenge

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
)

type PoWResponse struct {
	Solution string `json:"solution"`
}

// SolveChallenge solves the challenge and sends the solution back to the server.
// The challenge type is used to determine which solver to use.
func SolveChallenge(ctx context.Context, conn net.Conn) error {
	var challengeData Challenge
	if err := json.NewDecoder(conn).Decode(&challengeData); err != nil {
		return err
	}

	solver, ok := solverByType[challengeData.ChallengeType]
	if !ok {
		return fmt.Errorf("unknown challenge type: %s", challengeData.ChallengeType)
	}
	solution, err := solver(ctx, challengeData)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(conn).Encode(PoWResponse{Solution: solution}); err != nil {
		return err
	}
	return nil
}
