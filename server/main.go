package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func sendSMS(writer http.ResponseWriter, request *http.Request) {
	message := request.PostFormValue("message")
	if message == "" {
		log.Println("message was empty")
	}

	log.Printf("received message: %s\n", message)

	cmd := exec.Command("echo", "doing echo:", message)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Failed to senf message")
	}

	log.Println(string(out))
	// TODO: send sms using termux-api-sms
}

func main() {
	http.HandleFunc("/sendSMS", sendSMS)

	fmt.Println("Starting server...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Printf("error occured with server: %s\n", err)
	}
}
