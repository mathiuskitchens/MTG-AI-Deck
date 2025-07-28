# MTG Deck Guide AI

An AI-powered Magic: The Gathering deck helper and analysis tool.

The goal of this project was to make a simple application that utilizes a frontend and backend. 

On the frontend, you can search for a Magic: The Gathering card of your choice. This will retrieve card info from the Scryfall public API. 

From there, if you like the card, you can select "Generate Deck Idea", which will send your card of choice to the OpenAI API and generate ideas for other cards in the deck, suggest synergies and strategies for how to use the deck!

This process is handled by the Go backend server. The server simply listens for requests, extracts the card name entered, and sends this info in a pre-generated prompt to OpenAI. 

## Project Structure

- `frontend/` - Astro-based web frontend
- `backend/` - Go backend server

## Getting Started

### Frontend (Astro)

```bash
cd frontend
npm install
npm run dev
```

### Backend (Go)

```bash
cd backend
go run main.go
```

## Development

This project consists of a modern web frontend built with Astro and a Go backend server.

### Prerequisites

- Node.js (for frontend)
- Go 1.19+ (for backend)
- OpenAI API key (required for AI functionality)

### OpenAI API Key Setup

**Important:** This application requires an OpenAI API key to function. You must provide your own API key.

1. Get your API key from [OpenAI's platform](https://platform.openai.com/api-keys)
2. Set it as an environment variable called `OPENAI_API_KEY`

#### On macOS/Linux:
```bash
export OPENAI_API_KEY="your-api-key-here"
```

To make it permanent, add the above line to your `~/.zshrc`, `~/.bashrc`, or `~/.bash_profile`:
```bash
echo 'export OPENAI_API_KEY="your-api-key-here"' >> ~/.zshrc
source ~/.zshrc
```

#### On Windows:
```cmd
set OPENAI_API_KEY=your-api-key-here
```

Or for PowerShell:
```powershell
$env:OPENAI_API_KEY="your-api-key-here"
```

To make it permanent on Windows, set it through System Properties > Environment Variables.

#### Verify Setup:
```bash
echo $OPENAI_API_KEY  # Should display your API key
```

### Installation

1. Clone the repository
2. Install frontend dependencies: `cd frontend && npm install`
3. Run the development servers as described above

## License

[Add your license here]
