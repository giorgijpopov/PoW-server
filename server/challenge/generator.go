package challenge

type generator struct{}

var _ ChallengeGenerator = (*generator)(nil)

func NewGenerator() ChallengeGenerator {
	return &generator{}
}

func (g *generator) GenerateChallenge() Challenge {
	return NewHashInversionChallenge()
}
