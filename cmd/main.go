package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chatgpt-with-golang/config"
	"io/ioutil"
	"net/http"
)

var (
	cfg = config.GetConfig()
)

func main() {
	//apiKey := "YOUR_API_KEY"
	//url := "https://api.openai.com/v1/engines/davinci-codex/completions"

	// Request payload
	payload := struct {
		Prompt string `json:"prompt"`
	}{
		Prompt: "Once upon a time",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling request payload:", err)
		return
	}

	req, err := http.NewRequest("POST", cfg.OpenAIUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.OpenAIApiKey))
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response:", resp.Status)
		fmt.Println("Response body:", string(respBody))
		return
	}

	type CompletionResponse struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	var completionResp CompletionResponse
	err = json.Unmarshal(respBody, &completionResp)
	if err != nil {
		fmt.Println("Error parsing response body:", err)
		return
	}

	if len(completionResp.Choices) > 0 {
		fmt.Println("Generated text:", completionResp.Choices[0].Text)
	} else {
		fmt.Println("No text generated")
	}
}
