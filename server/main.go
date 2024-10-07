package main

import (
	"fmt"
	"log"
	"net"
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

	cmd := exec.Command("termux-sms-send", "-n", number, message)
	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(writer, fmt.Sprintf("Failed sending sms: %s", err), http.StatusInternalServerError)
		log.Printf("Failed sending sms: %s\n", err)
	} else {
		log.Println("Message sent.")
	}

	log.Printf("Command run with output: %s\n", out)
}

func getBoundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	http.HandleFunc("/sendSMS", sendSMS)
		
	if len(os.Args) < 2 {
		log.Fatalf("Number is not passed. Usage: %s [number]", os.Args[0])
	}
	number = os.Args[1]

	localip := getBoundIP()
	port := 3333

	fmt.Printf("Server is listening at http://%s:%d", localip, port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Printf("error occured with server: %s\n", err)
	}
}
