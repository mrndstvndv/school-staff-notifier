package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var number string

func sendSMS(writer http.ResponseWriter, request *http.Request) {
	message := request.PostFormValue("message")

	if message == "" {
		log.Println("message was empty")
	}

	log.Printf("Trying to send message: %s\n", message)

	cmd := exec.Command("termux-api-sms", "-n", number, message)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Failed to senf message")
	}

	log.Println(string(out))
}

func main() {
	http.HandleFunc("/sendSMS", sendSMS)
		
	if len(os.Args) < 2 {
		log.Fatalf("Number is not passed. Usage: %s [number]", os.Args[0])
	}
	number = os.Args[1]

	fmt.Println("Server is listening at port 3333")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Printf("error occured with server: %s\n", err)
	}
}
