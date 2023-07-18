package usecase

import "github.com/chatgpt-with-golang/domain/entity"

type ChatUseCase interface {
	PostChat(req *entity.ChatRequest) (*entity.ChatResponse, error)
}
