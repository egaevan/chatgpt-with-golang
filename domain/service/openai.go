package service

import "github.com/chatgpt-with-golang/domain/entity"

type OpenAiService interface {
	PostChatGpt(*entity.ChatRequest) (*entity.ChatResponse, error)
}
