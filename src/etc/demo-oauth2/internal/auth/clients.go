package auth

import (
	"context"
)

type Client struct {
	HydraAdmin *HydraAdmin
}

func NewClient() *Client {
	return &Client{}
}

type HydraAdmin struct {
}

type AcceptLoginChallengeRequest struct {
	LoginChallenge string
}

func (h *HydraAdmin) AcceptLoginChallenge(ctx context.Context, in AcceptLoginChallengeRequest) error {
	return nil
}
