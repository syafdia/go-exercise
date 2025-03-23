package communication

import "context"

type EmailMessage = string

type EmailService interface {
	SendEmail(ctx context.Context, message EmailMessage)
}
