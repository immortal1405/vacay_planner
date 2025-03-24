package models

type VacationResponse struct {
	Plan string `json:"plan"`
}

type TravelTipsResponse struct {
	Tips string `json:"tips"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

type CompletionResponse struct {
	Answer string `json:"answer"`
}

type ChainOutput struct {
	Response string
}
