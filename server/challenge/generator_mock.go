package challenge

const (
	ChallengeTypeMock = "mock"
)

type GeneratorMock struct {
	ch Challenge
}

var _ ChallengeGenerator = (*GeneratorMock)(nil)

func NewGeneratorMock(ch Challenge) ChallengeGenerator {
	return &GeneratorMock{ch: ch}
}

func (g *GeneratorMock) GenerateChallenge() Challenge {
	return g.ch
}

func (g *GeneratorMock) ChallengeType() string {
	return ChallengeTypeMock
}
