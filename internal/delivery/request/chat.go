package request

import "github.com/chatgpt-with-golang/domain/entity"

type ChatRequest struct {
	Prompt string `json:"prompt"`
}

func ChatPostRequest(req ChatRequest) *entity.ChatRequest {
	return &entity.ChatRequest{
		Prompt: req.Prompt,
	}
}
