package planner

import (
	"context"
	"fmt"

	"github.com/immortal1405/vacay_planner/internal/models"
	"github.com/immortal1405/vacay_planner/internal/services/llm"
	"github.com/immortal1405/vacay_planner/pkg/shivaay"
)

type Service struct {
	langChainService *llm.LangChainService
}

func NewService(langChainService *llm.LangChainService) *Service {
	return &Service{
		langChainService: langChainService,
	}
}

func (s *Service) GeneratePlan(ctx context.Context, req *models.VacationRequest) (*models.VacationResponse, error) {
	prompt := fmt.Sprintf(`Create a detailed vacation plan for %s for %s. 
Interests: %s
Budget: %s
Travel Style: %s
Accommodation: %s
Transportation: %s
Special Needs: %s
Language: %s
Currency: %s

Please provide a comprehensive plan including:
1. Daily itinerary
2. Recommended accommodations
3. Transportation options
4. Activities and attractions
5. Local tips and recommendations
6. Cultural considerations
7. Safety advice
8. Budget breakdown`,
		req.Destination, req.Duration, req.Interests, req.Budget, req.TravelStyle,
		req.Accommodation, req.Transportation, req.SpecialNeeds, req.Language, req.Currency)

	response, err := s.langChainService.GetResponse(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("error generating plan: %w", err)
	}

	return &models.VacationResponse{
		Plan: response,
	}, nil
}

func (s *Service) GetTravelTips(ctx context.Context, req *models.TravelTipsRequest) (*models.TravelTipsResponse, error) {
	prompt := fmt.Sprintf("Provide essential travel tips and advice for visiting %s. Include information about local customs, safety, transportation, and must-know information for tourists.",
		req.Destination)

	response, err := s.langChainService.GetResponse(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("error getting travel tips: %w", err)
	}

	return &models.TravelTipsResponse{
		Tips: response,
	}, nil
}

func (s *Service) ChatWithPlanner(ctx context.Context, req *models.ChatRequest) (*models.ChatResponse, error) {
	messages := []shivaay.Message{
		{
			Role:    "user",
			Content: req.Message,
		},
	}

	response, err := s.langChainService.GetChatResponse(ctx, messages)
	if err != nil {
		return nil, fmt.Errorf("error in chat: %w", err)
	}

	return &models.ChatResponse{
		Response: response,
	}, nil
}
