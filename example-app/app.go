package main

import "net/http"

func sqsHandler(w http.ResponseWriter, r *http.Request) {
}

func snsHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/publish-sqs", sqsHandler)
	http.HandleFunc("/publish-sns", snsHandler)

	http.ListenAndServe(":8888", nil)
}
