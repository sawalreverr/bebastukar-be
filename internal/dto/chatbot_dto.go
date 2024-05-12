package dto

type ChatBotInput struct {
	Question string `form:"question" validate:"required,min=5"`
}

type ChatBotResponse struct {
	Question string `json:"question"`
	AnswerAI string `json:"answer_ai"`
}
