package usecase

import "github.com/chatgpt-with-golang/domain/service"

type ChatInteractor struct {
	OpenAiService service.OpenAiService
}

func NewChat(openAiService service.OpenAiService) *ChatInteractor {
	return &ChatInteractor{
		OpenAiService: openAiService,
	}
}
