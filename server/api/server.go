package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"PoW-server/quote"
)

func StartServer(ctx context.Context, quoteProvider quote.Provider, protocolWrapper *ChallengeResponseProtocolWrapper) error {
	var lc net.ListenConfig
	listener, err := lc.Listen(ctx, "tcp", ":8080")
	if err != nil {
		return fmt.Errorf("Error starting server: %w", err)
	}
	defer listener.Close()
	fmt.Println("Server started on port 8080")

	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn, quoteProvider, protocolWrapper)
	}
}

func handleConnection(conn net.Conn, quoteProvider quote.Provider, protocolWrapper *ChallengeResponseProtocolWrapper) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	defer conn.Close()

	err := protocolWrapper.Wrap(conn, func(conn net.Conn) error {
		q, err := quoteProvider.GetQuote(ctx)
		if err != nil {
			return err
		}

		return json.NewEncoder(conn).Encode(newSuccessResponse(q))
	})
	if err != nil {
		json.NewEncoder(conn).Encode(newErrorResponse(err))
	}
}
