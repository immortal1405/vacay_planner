package llm

import (
	"context"
	"fmt"

	"github.com/immortal1405/vacay_planner/pkg/shivaay"
)

type LangChainService struct {
	client *shivaay.Client
}

func NewLangChainService(client *shivaay.Client) *LangChainService {
	return &LangChainService{
		client: client,
	}
}

func (s *LangChainService) GetResponse(ctx context.Context, prompt string) (string, error) {
	systemMessage := shivaay.Message{
		Role:    "system",
		Content: "You are an expert travel planner with deep knowledge of destinations worldwide. Your responses are detailed, practical, and tailored to the traveler's specific needs and preferences.",
	}

	userMessage := shivaay.Message{
		Role:    "user",
		Content: prompt,
	}

	messages := []shivaay.Message{systemMessage, userMessage}

	response, err := s.client.CreateCompletion(messages, 0.7, 1.0)
	if err != nil {
		return "", fmt.Errorf("failed to get completion: %w", err)
	}

	return response.Answer, nil
}

func (s *LangChainService) GetChatResponse(ctx context.Context, messages []shivaay.Message) (string, error) {
	response, err := s.client.CreateCompletion(messages, 0.7, 1.0)
	if err != nil {
		return "", fmt.Errorf("failed to get completion: %w", err)
	}

	return response.Answer, nil
}
