package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// Configurações de envio de SMS
type SMSConfig struct {
	PhoneNumber string
	Message     string
}

// Cria uma nova sessão AWS
func createAWSSession() (*session.Session, error) {
	// Cria a sessão usando as credenciais da AWS armazenadas localmente em ~/.aws/credentials
	return session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
}

// Envia uma mensagem SMS utilizando o Amazon SNS
func sendSMS(session *session.Session, config SMSConfig) (*sns.PublishOutput, error) {
	// Cria um cliente SNS
	snsClient := sns.New(session)

	// Define os parâmetros para a mensagem
	params := &sns.PublishInput{
		PhoneNumber: aws.String(config.PhoneNumber),
		Message:     aws.String(config.Message),
	}

	// Envia a mensagem SMS
	response, err := snsClient.Publish(params)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar SMS: %w", err)
	}

	return response, nil
}

func main() {
	// Configuração para envio de SMS
	config := SMSConfig{
		PhoneNumber: "+5511999999999",
		Message:     "Hello, this is a test message from Amazon SNS.",
	}

	// Cria uma sessão AWS
	session, err := createAWSSession()
	if err != nil {
		log.Fatalf("Erro ao criar sessão AWS: %v", err)
	}

	// Envia o SMS
	response, err := sendSMS(session, config)
	if err != nil {
		log.Fatal(err)
	}

	// Exibe a resposta do SNS
	fmt.Println("Resposta do SNS:", response)
}
