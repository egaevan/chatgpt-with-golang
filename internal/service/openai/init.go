package openai

import "context"

type OpenAiService struct {
	ctx       context.Context
	openAiApi string
	openAiUrl string
}

func NewOpenAiService(ctx context.Context, openAiApi, openAiUrl string) OpenAiService {
	return OpenAiService{
		ctx:       ctx,
		openAiApi: openAiApi,
		openAiUrl: openAiUrl,
	}
}
