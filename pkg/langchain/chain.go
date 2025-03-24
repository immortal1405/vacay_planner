package langchain

import (
	"context"
	"fmt"

	"github.com/immortal1405/vacay_planner/pkg/shivaay"
)

type Chain struct {
	client       *shivaay.Client
	systemPrompt string
}

func NewChain(client *shivaay.Client, systemPrompt string) *Chain {
	return &Chain{
		client:       client,
		systemPrompt: systemPrompt,
	}
}

type ChainInput struct {
	Messages []shivaay.Message
}

type ChainOutput struct {
	Response string
}

func (c *Chain) Run(ctx context.Context, input ChainInput) (*ChainOutput, error) {
	messages := append([]shivaay.Message{
		{
			Role:    "system",
			Content: c.systemPrompt,
		},
	}, input.Messages...)

	resp, err := c.client.CreateCompletion(messages, 0.7, 1.0)
	if err != nil {
		return nil, fmt.Errorf("error creating completion: %w", err)
	}

	if resp.Answer == "" {
		return nil, fmt.Errorf("empty response from API")
	}

	return &ChainOutput{
		Response: resp.Answer,
	}, nil
}
