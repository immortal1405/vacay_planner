package planner

import (
	"context"
	"fmt"
	"strings"

	"github.com/immortal1405/vacay_planner/internal/models"
	"github.com/immortal1405/vacay_planner/internal/services/llm"
)

type VacationPlanner struct {
	llmService *llm.LangChainService
}

func NewVacationPlanner(llmService *llm.LangChainService) *VacationPlanner {
	return &VacationPlanner{
		llmService: llmService,
	}
}

func (p *VacationPlanner) GeneratePlan(ctx context.Context, req *models.VacationRequest) (*models.VacationResponse, error) {
	prompt := fmt.Sprintf(`You are an expert travel planner with deep knowledge of destinations worldwide. 
Create a detailed, personalized vacation plan for the following requirements:

Destination: %s
Duration: %s
Interests: %s
Budget Level: %s
Travel Style: %s
Accommodation Type: %s
Transportation Type: %s
Special Needs: %s
Preferred Language: %s
Currency: %s

Please provide a comprehensive vacation plan that includes:

1. Destination Overview:
   - Brief introduction to the destination
   - Best time to visit
   - Local customs and etiquette
   - Safety considerations

2. Daily Itinerary:
   - Day-by-day breakdown of activities
   - Time estimates for each activity
   - Travel time between locations
   - Meal suggestions with local specialties

3. Accommodation Details:
   - Recommended areas to stay
   - Specific hotel/resort suggestions based on budget
   - Amenities and facilities
   - Booking tips

4. Transportation Guide:
   - Airport transfers
   - Local transportation options
   - Cost estimates
   - Travel tips and hacks

5. Activities and Attractions:
   - Must-see attractions
   - Hidden gems
   - Activity difficulty levels
   - Booking recommendations

6. Dining Guide:
   - Local cuisine highlights
   - Restaurant recommendations by budget
   - Food safety tips
   - Dietary considerations

7. Practical Information:
   - Visa requirements
   - Currency exchange tips
   - Emergency contacts
   - Local customs and etiquette

8. Budget Breakdown:
   - Estimated costs for major expenses
   - Money-saving tips
   - Value-for-money recommendations
   - Hidden costs to consider

9. Health and Safety:
   - Required vaccinations
   - Health precautions
   - Emergency services
   - Travel insurance recommendations

10. Additional Tips:
    - Packing suggestions
    - Weather considerations
    - Cultural sensitivities
    - Local scams to avoid

Please format the response in a clear, easy-to-read structure with proper sections and bullet points. 
Include specific recommendations based on the traveler's interests, budget, and special needs.
Make the plan practical and actionable, with real-world tips and considerations.`,
		req.Destination,
		req.Duration,
		strings.Join(req.Interests, ", "),
		req.Budget,
		req.TravelStyle,
		req.Accommodation,
		req.Transportation,
		strings.Join(req.SpecialNeeds, ", "),
		req.Language,
		req.Currency)

	response, err := p.llmService.GetResponse(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("error generating plan: %w", err)
	}

	formattedResponse := formatResponse(response)

	return &models.VacationResponse{
		Plan: formattedResponse,
	}, nil
}

func formatResponse(response string) string {
	sections := strings.Split(response, "\n\n")

	var formattedSections []string
	for _, section := range sections {
		// Add section headers with emojis
		if strings.Contains(strings.ToLower(section), "overview") {
			section = "🌍 " + section
		} else if strings.Contains(strings.ToLower(section), "itinerary") {
			section = "📅 " + section
		} else if strings.Contains(strings.ToLower(section), "accommodation") {
			section = "🏨 " + section
		} else if strings.Contains(strings.ToLower(section), "transportation") {
			section = "🚗 " + section
		} else if strings.Contains(strings.ToLower(section), "activities") {
			section = "🎯 " + section
		} else if strings.Contains(strings.ToLower(section), "dining") {
			section = "🍽️ " + section
		} else if strings.Contains(strings.ToLower(section), "practical") {
			section = "ℹ️ " + section
		} else if strings.Contains(strings.ToLower(section), "budget") {
			section = "💰 " + section
		} else if strings.Contains(strings.ToLower(section), "health") {
			section = "🏥 " + section
		} else if strings.Contains(strings.ToLower(section), "tips") {
			section = "💡 " + section
		}

		lines := strings.Split(section, "\n")
		for i, line := range lines {
			if strings.TrimSpace(line) != "" && !strings.HasPrefix(line, "•") {
				lines[i] = "• " + line
			}
		}
		section = strings.Join(lines, "\n")
		formattedSections = append(formattedSections, section)
	}

	return strings.Join(formattedSections, "\n\n")
}

func (p *VacationPlanner) GetTravelTips(ctx context.Context, req *models.TravelTipsRequest) (*models.TravelTipsResponse, error) {
	prompt := fmt.Sprintf(`Provide essential travel tips and advice for visiting %s. Include information about:
1. Local customs and etiquette
2. Safety considerations
3. Transportation options
4. Must-know information for tourists
5. Cultural sensitivities
6. Health and medical considerations
7. Emergency contacts
8. Common scams to avoid

Please format the response in a clear, easy-to-read structure with proper sections and bullet points.`,
		req.Destination)

	response, err := p.llmService.GetResponse(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("error getting travel tips: %w", err)
	}

	formattedResponse := formatResponse(response)

	return &models.TravelTipsResponse{
		Tips: formattedResponse,
	}, nil
}

func (p *VacationPlanner) ChatWithPlanner(ctx context.Context, req *models.ChatRequest) (*models.ChatResponse, error) {
	response, err := p.llmService.GetResponse(ctx, req.Message)
	if err != nil {
		return nil, fmt.Errorf("error in chat: %w", err)
	}

	formattedResponse := formatResponse(response)

	return &models.ChatResponse{
		Response: formattedResponse,
	}, nil
}
