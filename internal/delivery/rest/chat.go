package rest

import (
	"github.com/chatgpt-with-golang/domain/usecase"
	"github.com/chatgpt-with-golang/internal/delivery/request"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type ChatCapsule struct {
	Super       InitHandler
	ChatUseCase usecase.ChatUseCase
}

type Chat interface {
	PostChat(c echo.Context) error
}

func NewChat(log *zap.Logger, chatUseCase usecase.ChatUseCase) Chat {
	return &ChatCapsule{
		Super: InitHandler{
			Logger: log,
		},
		ChatUseCase: chatUseCase,
	}
}

type responseError struct {
	Message string
}

func (c *ChatCapsule) PostChat(ctx echo.Context) error {

	dataReq := request.ChatRequest{}
	err := ctx.Bind(&dataReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responseError{
			Message: "invalid data request",
		})
		return echo.ErrBadRequest
	}

	req := request.ChatPostRequest(dataReq)

	res, err := c.ChatUseCase.PostChat(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responseError{
			Message: err.Error(),
		})
		return echo.ErrBadRequest
	}

	return ctx.JSON(http.StatusOK, res)
}
