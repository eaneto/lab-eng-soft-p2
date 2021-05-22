package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/eaneto/notify"
	sqs "github.com/eaneto/notify/service/amazonsqs"
	"github.com/sirupsen/logrus"
)

func sqsHandler(w http.ResponseWriter, r *http.Request) {
	notifier := notify.New()

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_DEFAULT_REGION")
	sqsService, err := sqs.New(accessKey, secretKey, region)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	account := os.Getenv("AWS_ACCOUNT")
	baseUrl := "https://sqs.%s.amazonaws.com/%s/notification-queue"
	sqsService.AddReceivers(fmt.Sprintf(baseUrl, region, account))

	notifier.UseServices(sqsService)

	err = notifier.Send(
		context.Background(),
		"Subject",
		"Example message",
	)

	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	logrus.Info("Message published succesfully")
}

func snsHandler(w http.ResponseWriter, r *http.Request) {
}

// healthHandler returns ok always, used by nagios.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/publish-sqs", sqsHandler)
	http.HandleFunc("/publish-sns", snsHandler)
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8888", nil)
}
