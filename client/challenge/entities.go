package challenge

import "context"

type Challenge struct {
	ChallengeType ChallengeType `json:"challengeType"`
	Payload       []byte        `json:"payload"`
}

type ChallengeType string
type ChallengeSolver func(ctx context.Context, ch Challenge) (string, error)

var (
	HashInversionChallengeType = registerChallengeType("hash_inversion", solveHashInversionChallenge)

	solverByType = map[ChallengeType]ChallengeSolver{}
)

func registerChallengeType(challengeType ChallengeType, s ChallengeSolver) ChallengeType {
	solverByType[challengeType] = s
	return challengeType
}
