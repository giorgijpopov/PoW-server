package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"PoW-server/api"
	"PoW-server/challenge"
	"PoW-server/quote"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := api.StartServer(ctx, quote.NewProvider(), api.NewChallengeResponseProtocolWrapper(challenge.NewGenerator()))
	if err != nil {
		fmt.Println("Server stopped with error:", err)
	}
}
