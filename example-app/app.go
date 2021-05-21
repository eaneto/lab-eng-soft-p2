package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eaneto/notify"
	sqs "github.com/eaneto/notify/service/amazonsqs"
)

func sqsHandler(w http.ResponseWriter, r *http.Request) {
	notifier := notify.New()
	fmt.Println(notifier)

	// Create a telegram service. Ignoring error for demo simplicity.
	sqsService, err := sqs.New("test", "test", "us-east-1")

	if err != nil {
		fmt.Println("Yes, error")
	}

	// Passing a telegram chat id as receiver for our messages.
	// Basically where should our message be sent?
	sqsService.AddReceivers("http://localhost:4566/000000000000/notification-queue")

	// Tell our notifier to use the telegram service. You can repeat the above process
	// for as many services as you like and just tell the notifier to use them.
	// Inspired by http middlewares used in higher level libraries.
	notifier.UseServices(sqsService)

	// Send a test message.
	error := notifier.Send(
		context.Background(),
		"Subject",
		"Example message",
	)

	fmt.Println(error)
}

func snsHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/publish-sqs", sqsHandler)
	http.HandleFunc("/publish-sns", snsHandler)

	http.ListenAndServe(":8888", nil)
}
