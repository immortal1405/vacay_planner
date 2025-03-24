package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/immortal1405/vacay_planner/internal/models"
)

const (
	baseURL = "https://api_v2.futurixai.com/api/lara/v1"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

func (c *Client) CreateCompletion(messages []models.Message, temperature, topP float64) (*models.CompletionResponse, error) {
	reqBody := models.CompletionRequest{
		Messages:    messages,
		Temperature: temperature,
		TopP:        topP,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	log.Printf("Sending request to Shivaay API: %s", string(jsonData))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/completion", baseURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-subscription-key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	log.Printf("Received response from Shivaay API: %s", string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var completionResp models.CompletionResponse
	if err := json.Unmarshal(body, &completionResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &completionResp, nil
}
