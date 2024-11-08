package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"PoW-server/challenge"
)

func main() {
	ctx := context.Background()
	for {
		getQuote(ctx)
	}
}

type QuoteResponse struct {
	Quote        string `json:"quote"`
	ErrorMessage string `json:"error"`
}

func getQuote(ctx context.Context) {
	conn, err := net.Dial("tcp", "wisdom-server:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	err = challenge.SolveChallenge(ctx, conn)
	if err != nil {
		fmt.Println("Error solving challenge:", err)
		return
	}

	var quoteResponse QuoteResponse
	if err := json.NewDecoder(conn).Decode(&quoteResponse); err != nil {
		fmt.Println("Error decoding server response:", err)
		return
	}
	if len(quoteResponse.ErrorMessage) != 0 {
		fmt.Println("Received error from server:", quoteResponse.ErrorMessage)
		return
	}
	fmt.Println("Received quote from server:", quoteResponse.Quote)
}
