package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/immortal1405/vacay_planner/internal/models"
	"github.com/manifoldco/promptui"
)

const (
	defaultServerURL = "http://localhost:8080"
	INTERESTS        = `Adventure
Culture
Nature
Food
Shopping
Relaxation
History
Nightlife
Sports
Other`

	BUDGET_OPTIONS = `Budget
Moderate
Luxury
Ultra-Luxury`

	TRAVEL_STYLES = `Backpacking
Cultural
Luxury
Adventure
Relaxation
Family
Solo
Group`

	ACCOMMODATION_TYPES = `Hostel
Budget Hotel
Mid-Range Hotel
Luxury Hotel
Resort
Vacation Rental
Camping
Other`

	TRANSPORTATION_TYPES = `Public Transport
Private Transport
Rental Car
Walking
Cycling
Other`

	SPECIAL_NEEDS = `None
Wheelchair Accessible
Dietary Restrictions
Medical Requirements
Language Assistance
Other`

	LANGUAGES = `English
Spanish
French
German
Chinese
Japanese
Korean
Other`

	CURRENCIES = `USD (US Dollar)
EUR (Euro)
GBP (British Pound)
JPY (Japanese Yen)
INR (Indian Rupee)
CNY (Chinese Yuan)
AUD (Australian Dollar)
CAD (Canadian Dollar)
CHF (Swiss Franc)
NZD (New Zealand Dollar)
SGD (Singapore Dollar)
HKD (Hong Kong Dollar)
KRW (South Korean Won)
RUB (Russian Ruble)
BRL (Brazilian Real)
MXN (Mexican Peso)
ZAR (South African Rand)
AED (UAE Dirham)
SAR (Saudi Riyal)
TRY (Turkish Lira)`
)

func printHeader(title string) {
	fmt.Print("\n" + strings.Repeat("=", 50) + "\n")
	fmt.Print(title + "\n")
	fmt.Print(strings.Repeat("=", 50) + "\n")
}

func main() {
	fmt.Print("Welcome to the AI Travel Planner!\n")
	fmt.Print("Let's plan your perfect vacation...\n\n")

	for {
		menuPrompt := promptui.Select{
			Label: "What would you like to do?",
			Items: []string{
				"Plan a Vacation",
				"Get Travel Tips",
				"Chat with Planner",
				"Exit",
			},
		}
		_, result, err := menuPrompt.Run()
		if err != nil {
			log.Fatal(err)
		}

		switch result {
		case "Plan a Vacation":
			planVacation()
		case "Get Travel Tips":
			getTravelTips(defaultServerURL)
		case "Chat with Planner":
			chatWithPlanner(defaultServerURL)
		case "Exit":
			fmt.Print("\nThank you for using AI Travel Planner!\n")
			return
		}
	}
}

func planVacation() {
	printHeader("Plan Your Vacation")

	destinationPrompt := promptui.Prompt{
		Label: "Where would you like to go?",
	}
	destination, err := destinationPrompt.Run()
	if err != nil {
		fmt.Printf("Error getting destination: %v\n", err)
		return
	}

	durationPrompt := promptui.Prompt{
		Label: "How long would you like to stay? (e.g., 7 days)",
	}
	duration, err := durationPrompt.Run()
	if err != nil {
		fmt.Printf("Error getting duration: %v\n", err)
		return
	}

	interestsPrompt := promptui.Select{
		Label: "What are your interests?",
		Items: strings.Split(INTERESTS, "\n"),
	}
	_, interests, err := interestsPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	budgetPrompt := promptui.Select{
		Label: "What's your budget level?",
		Items: strings.Split(BUDGET_OPTIONS, "\n"),
	}
	_, budget, err := budgetPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	stylePrompt := promptui.Select{
		Label: "What's your preferred travel style?",
		Items: strings.Split(TRAVEL_STYLES, "\n"),
	}
	_, travelStyle, err := stylePrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	accommodationPrompt := promptui.Select{
		Label: "What type of accommodation do you prefer?",
		Items: strings.Split(ACCOMMODATION_TYPES, "\n"),
	}
	_, accommodation, err := accommodationPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	transportationPrompt := promptui.Select{
		Label: "What type of transportation do you prefer?",
		Items: strings.Split(TRANSPORTATION_TYPES, "\n"),
	}
	_, transportation, err := transportationPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	specialNeedsPrompt := promptui.Select{
		Label: "Do you have any special needs?",
		Items: strings.Split(SPECIAL_NEEDS, "\n"),
	}
	_, specialNeeds, err := specialNeedsPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	languagePrompt := promptui.Select{
		Label: "What's your preferred language?",
		Items: strings.Split(LANGUAGES, "\n"),
	}
	_, language, err := languagePrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	currencyPrompt := promptui.Select{
		Label: "What's your preferred currency?",
		Items: strings.Split(CURRENCIES, "\n"),
	}
	_, currency, err := currencyPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	req := &models.VacationRequest{
		Destination:    destination,
		Duration:       duration,
		Interests:      []string{interests},
		Budget:         budget,
		TravelStyle:    travelStyle,
		Accommodation:  accommodation,
		Transportation: transportation,
		SpecialNeeds:   []string{specialNeeds},
		Language:       language,
		Currency:       currency,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("Error marshaling request: %v\n", err)
		return
	}

	resp, err := http.Post(defaultServerURL+"/api/v1/plan", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	var vacationResp models.VacationResponse
	if err := json.Unmarshal(body, &vacationResp); err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		return
	}

	fmt.Println("\nYour Personalized Vacation Plan:")
	fmt.Println("=================================")
	fmt.Print(vacationResp.Plan)

	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("vacation_plan_%s.txt", timestamp)
	if err := os.WriteFile(filename, []byte(vacationResp.Plan), 0644); err != nil {
		fmt.Printf("Error saving plan to file: %v\n", err)
		return
	}
	fmt.Printf("\nYour vacation plan has been saved to %s\n", filename)
}

func getTravelTips(serverURL string) {
	printHeader("Get Travel Tips")

	destinationPrompt := promptui.Prompt{
		Label: "Where are you traveling to?",
	}
	destination, err := destinationPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	durationPrompt := promptui.Prompt{
		Label: "How many days will you be staying?",
	}
	durationStr, err := durationPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil || duration < 1 {
		fmt.Println("Invalid duration. Please enter a positive number.")
		return
	}

	interestsPrompt := promptui.Select{
		Label: "What are your interests?",
		Items: strings.Split(INTERESTS, "\n"),
	}
	_, interests, err := interestsPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	request := models.TravelTipsRequest{
		Destination: destination,
		Duration:    duration,
		Interests:   []string{interests},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	resp, err := http.Post(serverURL+"/api/v1/tips", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Server returned error: %s\n", body)
		return
	}

	var response models.TravelTipsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		return
	}

	printHeader(fmt.Sprintf("Travel Tips for %s", destination))
	fmt.Print(response.Tips)
}

func chatWithPlanner(serverURL string) {
	printHeader("Chat with Vacation Planner")
	fmt.Println("Ask any travel-related questions or get personalized advice.")
	fmt.Println("Type 'exit' to return to the main menu")
	fmt.Println("Type 'help' to see available commands")
	fmt.Println("Type 'clear' to clear the screen")

	for {
		messagePrompt := promptui.Prompt{
			Label: "You",
		}
		message, err := messagePrompt.Run()
		if err != nil {
			log.Fatal(err)
		}
		message = strings.TrimSpace(message)

		switch strings.ToLower(message) {
		case "exit":
			return
		case "help":
			fmt.Println("\nAvailable commands:")
			fmt.Println("- 'exit': Return to main menu")
			fmt.Println("- 'help': Show this help message")
			fmt.Println("- 'clear': Clear screen")
			fmt.Println("- 'budget': Get budget planning advice")
			fmt.Println("- 'packing': Get packing tips")
			fmt.Println("- 'safety': Get safety tips")
			fmt.Println("- 'visa': Get visa information")
			fmt.Println("- 'weather': Get weather information")
			fmt.Println("- 'timezone': Get timezone information")
			continue
		case "clear":
			fmt.Print("\033[H\033[2J") // ANSI escape code to clear screen
			continue
		}

		request := models.ChatRequest{
			Message: message,
		}

		jsonData, err := json.Marshal(request)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			continue
		}

		resp, err := http.Post(serverURL+"/api/v1/chat", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Server returned error: %s\n", body)
			continue
		}

		var response models.ChatResponse
		if err := json.Unmarshal(body, &response); err != nil {
			fmt.Printf("Error parsing response: %v\n", err)
			continue
		}

		fmt.Printf("\nPlanner: %s\n", response.Response)
	}
}
