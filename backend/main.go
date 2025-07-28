package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type DeckRequest struct {
	Card string `json: "card"`
}

type DeckResponse struct {
	Idea string `json:"idea"`
}

type ChatMessage struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type OpenAIRequest struct {
	Model string 	`json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/deck-idea", deckIdeaHandler)	

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	type HelloResponse struct {
		Message string `json:"message"`
	}


	res := HelloResponse{Message: "Hello from Go!"}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(res)

}

func deckIdeaHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == http.MethodOptions {
        // Preflight request, just respond OK with headers
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}



	var req DeckRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Card == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		http.Error(w, "OpenAI API key not set", http.StatusInternalServerError)
		return
	}

	prompt := "Build a Standard-legal Magic: The Gathering deck around this card: " + req.Card + ". Suggest archetype, synergy, and card suggestions."

	openAIReq := OpenAIRequest{
		Model: "gpt-4.1-nano",
		Messages: []ChatMessage{
			{Role: "user", Content: prompt},
		},
	}

	body, err := json.Marshal(openAIReq)
	if err != nil {
		http.Error(w, "Failed to create reuqest", http.StatusInternalServerError)
		return
	}

	request, err := http.NewRequest("POST","https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Failed ot build OpenAI request", http.StatusInternalServerError)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		http.Error(w, "Failed to reach OpenAI", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var openAIResp OpenAIResponse
	respBody, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, &openAIResp)
	if err != nil {
		http.Error(w, "Failed to parse OpenAI response", http.StatusInternalServerError)
		return
	}

	if len(openAIResp.Choices) == 0 {
		http.Error(w, "OpenAI returned no results", http.StatusInternalServerError)
		return
	}

	idea := openAIResp.Choices[0].Message.Content

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(DeckResponse{
		Idea: idea,
	})
}