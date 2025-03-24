# AI Travel Planner

A smart travel planning application that helps you create personalized vacation plans using AI. The application provides interactive features for planning vacations, getting travel tips, and chatting with an AI travel planner.

## Features

- **Vacation Planning**

  - Personalized vacation plans based on your preferences
  - Multiple travel styles and accommodation options
  - Budget considerations
  - Special needs accommodations
  - Language and currency preferences

- **Travel Tips**

  - Destination-specific advice
  - Duration-based recommendations
  - Interest-based suggestions
  - Safety and practical tips

- **Interactive Chat**
  - Real-time conversation with AI travel planner
  - Quick access to common travel queries
  - Helpful commands for various travel-related topics
  - Context-aware responses

## Prerequisites

- Go 1.16 or higher
- Shivaay API key (for AI capabilities)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/immortal1405/vacay_planner.git
cd vacay_planner
```

2. Install dependencies:

```bash
go mod download
```

3. Set up your environment variables:

```bash
export SHIVAAY_API_KEY=your_api_key_here
export TOP_P=0.7              # Controls diversity of responses (0.0 to 1.0)
export TEMPERATURE=0.8        # Controls randomness in responses (0.0 to 1.0)
```

> **Note**:
>
> - `TOP_P`: Controls the diversity of responses. Lower values (e.g., 0.3) make responses more focused and deterministic, while higher values (e.g., 0.9) allow for more creative and diverse responses.
> - `TEMPERATURE`: Controls the randomness in responses. Lower values (e.g., 0.2) make responses more deterministic and focused, while higher values (e.g., 0.8) make responses more creative and varied.

## Usage

1. Start the server:

```bash
go run cmd/server/main.go
```

2. In a new terminal, start the client:

```bash
go run cmd/client/main.go
```

3. Follow the interactive prompts to:
   - Plan a vacation
   - Get travel tips
   - Chat with the AI planner

## API Endpoints

- `POST /api/v1/plan` - Generate a personalized vacation plan
- `POST /api/v1/tips` - Get travel tips for a destination
- `POST /api/v1/chat` - Chat with the AI travel planner

## Features in Detail

### Vacation Planning

- Destination selection
- Duration planning
- Interest-based activities
- Budget considerations
- Travel style preferences
- Accommodation options
- Transportation choices
- Special needs accommodations
- Language preferences
- Currency selection

### Travel Tips

- Destination-specific advice
- Duration-based recommendations
- Interest-based suggestions
- Safety tips
- Cultural considerations
- Local customs
- Weather information
- Visa requirements

### Interactive Chat

- Quick access to travel information
- Common commands:
  - `help`: Show available commands
  - `clear`: Clear the screen
  - `budget`: Get budget planning advice
  - `packing`: Get packing tips
  - `safety`: Get safety tips
  - `visa`: Get visa information
  - `weather`: Get weather information
  - `timezone`: Get timezone information

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Shivaay AI for providing the language model capabilities
- The Go community for excellent libraries and tools
- Contributors and users of this project

