package main

import (
	"bongserver/db"
	"bongserver/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var number string
var sendSmsEnabled = false

var GlobalState db.Library

func init() {
	GlobalState = db.Library{
		Labs: []db.Lab{
			{
				Name: "1",
				Computers: []db.Computer{
					{
						Number: "1",
						Issues: []db.Issue{
							{
								Timestamp:   utils.CurrentTimeUnixString(),
								Description: "Broken screen",
							},
						},
					},
					{
						Number: "2",
						Issues: []db.Issue{
							{
								Timestamp:   utils.CurrentTimeUnixString(),
								Description: "Broken keyboard",
							},
						},
					},
				},
			},
			{
				Name: "2",
				Computers: []db.Computer{
					{
						Number: "1",
						Issues: []db.Issue{
							{
								Timestamp:   utils.CurrentTimeUnixString(),
								Description: "Broken mouse",
							},
						},
					},
					{
						Number: "2",
						Issues: []db.Issue{
							{
								Timestamp:   utils.CurrentTimeUnixString(),
								Description: "Python not installed",
							},
						},
					},
				},
			},
		},
	}
}

func sendSMS(writer http.ResponseWriter, request *http.Request) {
	if !sendSmsEnabled {
		return
	}

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

func main() {
	_, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Connected to database")
	}

	http.HandleFunc("/sendSMS", sendSMS)	

	if len(os.Args) < 2 {
		log.Printf("Number is not passed, disabling sending a message. Usage: %s [number]\n", os.Args[0])
	} else {
		number = os.Args[1]
	}

	localip, err := utils.GetLocalIP()
	if err != nil {
		localip = "localhost"
		log.Println("Unable to get localip")
	}

	port := 3333

	fmt.Printf("Server is listening at http://%s:%d", localip, port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Printf("error occured with server: %s\n", err)
	}
}
