package main

import (
	"bongserver/db"
	"bongserver/utils"
	"bongserver/websocket"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
)

var (
	number         string
	sendSmsEnabled = false

	conn *sql.DB
	hub *websocket.Hub
)

func init() {
	c, err := db.Connect("bongserver.db")
	if err != nil {
		switch err.(type) {
		case *net.OpError:
			log.Fatalf("Unable to connect to database, check if the database is running\nError: %s\n", err)
		default:
			log.Fatalf("Unknown error occured: %s", err)
		}
	} else {
		log.Print("Connected to database")
	}

	conn = c
	hub = websocket.NewHub()
}

// TODO: add websocket
func main() {
	defer conn.Close()

	utils.LogDebug("Creating table")
	db.CreateTable(conn)

	err := db.PrintIssues(conn)
	if err != nil {
		utils.LogDebug("Failed to print issues: %s", err)
	}

	if len(os.Args) < 2 {
		log.Printf("Number is not passed, disabling sending a message. Usage: %s [number]\n", os.Args[0])
	} else {
		number = os.Args[1]
	}

	go hub.Run()

	http.HandleFunc("/reportIssue", reportIssue)
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		websocket.ServeWs(hub, writer, request)
	})

	localip, err := utils.GetLocalIP()
	if err != nil {
		localip = "localhost"
		log.Println("Unable to get localip")
	}

	port := 3333

	log.Printf("Server is listening at http://%s:%d\n", localip, port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Printf("error occured with server: %s\n", err)
	}
}

func sendSMS(writer http.ResponseWriter, request *http.Request) {
	if !sendSmsEnabled {
		utils.LogDebug("Sending SMS is disabled")
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

// TODO: Add validation for the fields
func issueFromRequest(request *http.Request) *db.Issue {
	decoder := json.NewDecoder(request.Body)

	var issue db.Issue
	err := decoder.Decode(&issue)
	if err != nil {
		log.Fatalf("Unable to decode request: %s\n", err)
	}

	return &issue
}

func reportIssue(writer http.ResponseWriter, request *http.Request) {
	// issue := issueFromRequest(request)

	// err := db.InsertIssue(conn, issue)
	// if err != nil {
	// 	log.Fatalf("Unable to insert issue: %s\n", err)
	// }

	hub.Broadcast <- []byte("New issue reported")

	db.PrintIssues(conn)
}
