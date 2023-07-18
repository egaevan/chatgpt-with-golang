package openai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chatgpt-with-golang/domain/entity"
	"io/ioutil"
	"net/http"
)

var (
	responseGenerate string
)

type payload struct {
	Model    string `json:"model"`
	Prompt   string `json:"prompt"`
	MaxTurns int    `json:"max_tokens"`
}

func (o *OpenAiService) PostChatGpt(reqData *entity.ChatRequest) (*entity.ChatResponse, error) {

	sendPayload := payload{
		Model:    "davinci-codex",
		Prompt:   reqData.Prompt,
		MaxTurns: 5,
	}

	payloadBytes, err := json.Marshal(sendPayload)
	if err != nil {
		//fmt.Println("Error marshaling request payload:", err)
		return nil, errors.New("error marshaling request payload")
	}

	req, err := http.NewRequest("POST", o.openAiUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		//fmt.Println("Error creating HTTP request:", err)
		return nil, errors.New("error creating HTTP request")
	}

	//fmt.Println("CEK API : ", o.openAiApi)
	//fmt.Println("CEK URL : ", o.openAiUrl)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", o.openAiApi))
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println("Error making HTTP request:", err)
		return nil, errors.New("error making HTTP request")
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("Error reading HTTP response:", err)
		return nil, errors.New("error reading HTTP response")
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response:", resp.Status)
		fmt.Println("Response body:", string(respBody))
		return nil, errors.New("error response body")
	}

	type CompletionResponse struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	var completionResp CompletionResponse
	err = json.Unmarshal(respBody, &completionResp)
	if err != nil {
		//fmt.Println("Error parsing response body:", err)
		return nil, errors.New("error parsing response body")
	}

	if len(completionResp.Choices) > 0 {
		responseGenerate = completionResp.Choices[0].Text
	}

	return &entity.ChatResponse{
		Prompt: responseGenerate,
	}, nil
}
