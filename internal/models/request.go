package models

type VacationRequest struct {
	Destination    string   `json:"destination" binding:"required"`
	Duration       string   `json:"duration" binding:"required"`
	Interests      []string `json:"interests" binding:"required"`
	Budget         string   `json:"budget" binding:"required"`
	TravelStyle    string   `json:"travel_style" binding:"required"`
	Accommodation  string   `json:"accommodation" binding:"required"`
	Transportation string   `json:"transportation" binding:"required"`
	SpecialNeeds   []string `json:"special_needs" binding:"required"`
	Language       string   `json:"language" binding:"required"`
	Currency       string   `json:"currency" binding:"required"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionRequest struct {
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	TopP        float64   `json:"top_p"`
}

type TravelTipsRequest struct {
	Destination string   `json:"destination" binding:"required"`
	Duration    int      `json:"duration" binding:"required"`
	Interests   []string `json:"interests" binding:"required"`
}

type ChatRequest struct {
	Message string `json:"message"`
}
