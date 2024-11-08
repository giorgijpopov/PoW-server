package quote

import "context"

type providerMock struct {
	quote string
}

var _ Provider = (*providerMock)(nil)

func NewProviderMock(q string) Provider {
	return &providerMock{quote: q}
}

func (p providerMock) GetQuote(ctx context.Context) (string, error) {
	return p.quote, nil
}
