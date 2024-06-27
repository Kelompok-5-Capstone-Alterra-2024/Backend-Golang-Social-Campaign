package service

import (
	"capstone/dto"
	"capstone/repositories"
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

type ChatbotService interface {
	Create(chatbot dto.Chatbot) (dto.Chatbot, error)
	GetByID(id string) ([]dto.Chatbot, error)
}

type chatbotService struct {
	repo repositories.ChatbotRepository
}

func NewChatbotService(repo repositories.ChatbotRepository) ChatbotService {
	return &chatbotService{repo: repo}
}

func (s *chatbotService) Create(chatbot dto.Chatbot) (dto.Chatbot, error) {
	var sysChatbot dto.Message
	sysChatbot.Role = "system"
	sysChatbot.Message = "Anda adalah seorang ahli dalam bidang sosial, anda akan diberikan pertanyaan tentang donasi di bidang sosial, pendidikan, alam, dan bencana,anda harus menjawab pertanyaan tersebut dengan format narasi"

	var payloadChatbot []dto.Message
	var ID string

	payloadChatbot = append(payloadChatbot, sysChatbot)

	if chatbot.ChatID == "" {
		ID = uuid.New().String()
	} else {
		ID = chatbot.ChatID
		res, err := s.repo.FindByID(chatbot.ChatID)
		if err != nil {
			return dto.Chatbot{}, err
		}
		for _, v := range res {
			payloadChatbot = append(payloadChatbot, dto.Message{
				Role:    v.Role,
				Message: v.Message,
			})
		}
	}

	chatbot.ID = uuid.New().String()
	chatbot.ChatID = ID
	chatbot.Role = "user"

	_, err := s.repo.Create(chatbot)
	if err != nil {
		return dto.Chatbot{}, err
	}

	payloadChatbot = append(payloadChatbot, dto.Message{
		Role:    chatbot.Role,
		Message: chatbot.Message,
	})

	// var openAIPayload []map[string]string
	// for _, v := range payloadChatbot {
	// 	openAIPayload = append(openAIPayload, map[string]string{
	// 		"role":    v.Role,
	// 		"content": v.Message,
	// 	})
	// }

	var openAIPayload []openai.ChatCompletionMessage
	for _, v := range payloadChatbot {
		openAIPayload = append(openAIPayload, openai.ChatCompletionMessage{
			Role:    v.Role,
			Content: v.Message,
		})
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: openAIPayload,
		},
	)

	if err != nil {
		return dto.Chatbot{}, err
	}

	// openaiURL := "https://api.openai.com/v1/chat/completions"
	// openaiRequest := map[string]interface{}{
	// 	"model":    "gpt-4",
	// 	"messages": openAIPayload,
	// }

	// requestBody, err := json.Marshal(openaiRequest)
	// if err != nil {
	// 	return dto.Chatbot{}, err
	// }

	// req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestBody))
	// if err != nil {
	// 	return dto.Chatbot{}, err
	// }

	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	return dto.Chatbot{}, err
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return dto.Chatbot{}, err
	// }

	// var openaiResp dto.OpenaiResponse
	// if err := json.Unmarshal(body, &openaiResp); err != nil {
	// 	return dto.Chatbot{}, err
	// }

	var chatBotAssistant dto.Chatbot
	chatBotAssistant.ID = uuid.New().String()
	chatBotAssistant.ChatID = ID
	chatBotAssistant.Role = "assistant"
	chatBotAssistant.Message = resp.Choices[0].Message.Content
	res, err := s.repo.Create(chatBotAssistant)
	if err != nil {
		return dto.Chatbot{}, err
	}

	return res, nil
}

func (s *chatbotService) GetByID(id string) ([]dto.Chatbot, error) {
	return s.repo.FindByID(id)
}
