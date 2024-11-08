package challenge

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	poolCount = 4
)

type HashInversionChallenge struct {
	Start          string `json:"start"`
	RequiredPrefix string `json:"prefix"`
}

func solveHashInversionChallenge(ctx context.Context, ch Challenge) (string, error) {
	var challenge HashInversionChallenge
	if err := json.Unmarshal(ch.Payload, &challenge); err != nil {
		fmt.Println("Error decoding challenge payload:", err)
		return "", err
	}
	return inverseHash(ctx, challenge.Start, challenge.RequiredPrefix)
}

func inverseHash(ctx context.Context, start, requiredPrefix string) (string, error) {
	resultChan := make(chan string)
	errorChan := make(chan error)
	done := make(chan struct{})
	defer func() {
		close(done)
	}()

	search := func(first int) {
		for i := first; ; i += poolCount {
			// check if the context is done or the function is done
			select {
			case <-done:
				return
			case <-ctx.Done():
				errorChan <- ctx.Err()
			default:
			}

			attempt := start + strconv.Itoa(i)
			hash := sha256.Sum256([]byte(attempt))
			hashHex := hex.EncodeToString(hash[:])
			if strings.HasPrefix(hashHex, requiredPrefix) {
				// try to send the result to the result channel
				select {
				case resultChan <- strconv.Itoa(i):
				default:
				}
			}
		}
	}

	for i := 0; i < poolCount; i++ {
		go search(i)
	}

	select {
	case res := <-resultChan:
		return res, nil
	case err := <-errorChan:
		return "", err
	}
}
