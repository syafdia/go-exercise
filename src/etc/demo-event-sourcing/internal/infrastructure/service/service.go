package service

import (
	"context"
	"fmt"

	"github.com/syafdia/demo-es/internal/domain/communication"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (e *EmailService) SendEmail(ctx context.Context, message communication.EmailMessage) {
	fmt.Println("Send email:", message)
}
