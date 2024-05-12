package chatbot

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/bebastukar-be/config"
	"github.com/sawalreverr/bebastukar-be/internal/dto"
	"github.com/sawalreverr/bebastukar-be/internal/helper"
	"google.golang.org/api/option"
)

var (
	apiKey            = config.GetConfig().GenerativeAI.ApiKey
	modelGenerative   = "gemini-1.5-pro-latest"
	instructionSystem = "Anda adalah seorang yang sangat ahli dalam masalah pengelolaan dan pertukaran barang bekas, saya ingin anda menjawab pertanyaan apapun terkait barang bekas yang akan diberikan oleh user dengan jawaban yang mudah dipahami. Selain pertanyaan seputar barang bekas anda hanya akan menjawab 'Tanyakan seputar barang bekas, Selain itu saya tidak tahu apapun'"
)

type ChatBotHandler interface {
	newModel() (*genai.Client, *genai.GenerativeModel)
	QuestionHandler(c echo.Context) error
}

type chatBotHandler struct{}

func NewDiscussionHandler() ChatBotHandler {
	return &chatBotHandler{}
}

func (h *chatBotHandler) newModel() (*genai.Client, *genai.GenerativeModel) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel(modelGenerative)
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(instructionSystem)},
	}

	return client, model
}

func (h *chatBotHandler) QuestionHandler(c echo.Context) error {
	var question dto.ChatBotInput

	if err := c.Bind(&question); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "bind error")
	}

	if err := c.Validate(&question); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "question must not be empty and at least 5 letter")
	}

	ctx := context.Background()
	client, model := h.newModel()
	defer client.Close()

	resp, err := model.GenerateContent(ctx, genai.Text(question.Question))
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "chatbot server error")
	}

	answerAI := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])

	response := dto.ChatBotResponse{
		Question: question.Question,
		AnswerAI: answerAI,
	}

	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "ok", response))
}
