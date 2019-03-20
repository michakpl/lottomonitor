package notification

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go"
	"log"
	"os"
)

func Send(game string) {
	mgDomain := os.Getenv("MAILGUN_DOMAIN")
	mgApiKey := os.Getenv("MAILGUN_API_KEY")

	mg := mailgun.NewMailgun(mgDomain, mgApiKey)

	sender := os.Getenv("MAILGUN_SENDER")
	recipient := os.Getenv("MAILGUN_RECIPIENT")
	subject := "You won in the lottery!"
	body := fmt.Sprintf("You won in the lottery game: %s", game)

	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID: %s. Response: %s", id, resp)
}