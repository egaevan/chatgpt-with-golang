package main

import (
	"context"
	"github.com/chatgpt-with-golang/internal/config"
	"github.com/chatgpt-with-golang/internal/delivery/rest"
	"github.com/chatgpt-with-golang/internal/service/openai"
	"github.com/chatgpt-with-golang/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
//ctx           = context.TODO()
//cfg           = config.GetConfig()
//log           = config.Logger()
//openAiService = openai.NewOpenAiService(ctx, cfg.OpenAIApiKey, cfg.OpenAIUrl)
)

func main() {
	//apiKey := "YOUR_API_KEY"
	//url := "https://api.openai.com/v1/engines/davinci-codex/completions"

	// Request payload
	//payload := struct {
	//	Prompt string `json:"prompt"`
	//}{
	//	Prompt: "Once upon a time",
	//}

	//payloadBytes, err := json.Marshal(payload)
	//if err != nil {
	//	fmt.Println("Error marshaling request payload:", err)
	//	return
	//}
	//
	//req, err := http.NewRequest("POST", cfg.OpenAIUrl, bytes.NewBuffer(payloadBytes))
	//if err != nil {
	//	fmt.Println("Error creating HTTP request:", err)
	//	return
	//}
	//
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.OpenAIApiKey))
	//req.Header.Set("Content-Type", "application/json")
	//
	//client := http.DefaultClient
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println("Error making HTTP request:", err)
	//	return
	//}
	//
	//defer resp.Body.Close()
	//
	//respBody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error reading HTTP response:", err)
	//	return
	//}
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error response:", resp.Status)
	//	fmt.Println("Response body:", string(respBody))
	//	return
	//}

	//type CompletionResponse struct {
	//	Choices []struct {
	//		Text string `json:"text"`
	//	} `json:"choices"`
	//}
	//
	//var completionResp CompletionResponse
	//err = json.Unmarshal(respBody, &completionResp)
	//if err != nil {
	//	fmt.Println("Error parsing response body:", err)
	//	return
	//}

	//if len(completionResp.Choices) > 0 {
	//	fmt.Println("Generated text:", completionResp.Choices[0].Text)
	//} else {
	//	fmt.Println("No text generated")
	//}

	ctx := context.TODO()
	cfg := config.GetConfig()
	log := config.Logger()

	cfg.OpenAIApiKey = "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"
	//cfg.OpenAIUrl = "https://api.openai.com/v1/engines/davinci-codex/completions"
	cfg.OpenAIUrl = "https://api.openai.com/v1/chat/completions"
	//cfg.OpenAIUrl = "https://api.openai.com/v1/completions"

	openAiService := openai.NewOpenAiService(ctx, cfg.OpenAIApiKey, cfg.OpenAIUrl)

	e := echo.New()

	e.GET("/", HealthCheck)

	chatUseCase := usecase.NewChat(&openAiService)

	chatController := rest.NewChat(log, chatUseCase)

	chat := e.Group("/v1/chat")
	chat.POST("/add", chatController.PostChat)

	e.Logger.Fatal(e.Start(":8000"))

}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
