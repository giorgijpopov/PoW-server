package quote

import "context"

type Provider interface {
	GetQuote(ctx context.Context) (string, error)
}
