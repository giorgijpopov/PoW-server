package challenge

type Challenge interface {
	Type() string
	Payload() ([]byte, error)
	CheckSolution(solution string) error
}

type ChallengeGenerator interface {
	GenerateChallenge() Challenge
}
