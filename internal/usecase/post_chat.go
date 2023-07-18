package usecase

import "github.com/chatgpt-with-golang/domain/entity"

func (c *ChatInteractor) PostChat(req *entity.ChatRequest) (*entity.ChatResponse, error) {

	res, err := c.OpenAiService.PostChatGpt(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
