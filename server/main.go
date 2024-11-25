package main

import (
	"bongserver/db"
	"bongserver/protobuf"
	"bongserver/utils"
	"bongserver/websocket"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"

	"google.golang.org/protobuf/proto"
)

var (
	number         string
	sendSmsEnabled = false

	conn *sql.DB
	hub  *websocket.Hub
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

func main() {
	defer conn.Close()

	utils.LogDebug("Creating table")
	err := db.CreateTable(conn)
	if err != nil {
		log.Printf("Unable to create table: %s\n", err)
	}

	if len(os.Args) < 2 {
		log.Printf("Number is not passed, disabling sending a message. Usage: %s [number]\n", os.Args[0])
	} else {
		number = os.Args[1]
	}

	go hub.Run()

	http.HandleFunc("/reportIssue", reportIssue)
	http.HandleFunc("/updateStatus", updateStatus)
	http.HandleFunc("/getIssues", getIssues)
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

func serializeMessage(issue *protobuf.Issue) ([]byte, error) {
	out, err := proto.Marshal(issue)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func getIssues(writer http.ResponseWriter, request *http.Request) {
	issues, err := db.GetIssues(conn)
	if err != nil {
		utils.LogDebug("Failed to get issues %s", err)
		http.Error(writer, "Unable to marshal issue", http.StatusInternalServerError)
	}

	byte, err := proto.Marshal(issues)
	if err != nil {
		utils.LogDebug("Unable to marshal issue: %s", err)
		http.Error(writer, "Unable to marshal issue", http.StatusInternalServerError)
	}

	utils.LogDebug("Sending issues: %s", issues.Issues)
	utils.LogDebug("Origin: %s", request.Header.Get("Origin"))

	writer.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	writer.Header().Set("Content-Type", "application/x-protobuf")
	writer.Write(byte)
}

func httpError(writer http.ResponseWriter, code int, message string, args ...interface{}) {
	message = fmt.Sprintf(message, args...)
	utils.LogDebug(message)
	http.Error(writer, message, code)
}

func updateStatus(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		httpError(writer, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	err := request.ParseForm()
	if err != nil {
		httpError(writer, http.StatusInternalServerError, "Unable to parse form: %s", err)
		return
	}

	id := request.Form.Get("id")
	status := request.Form.Get("status")

	utils.LogDebug("id: %s", id)
	utils.LogDebug("id: %s", status)

	err = db.UpdateComponentStatus(conn, id, status)
	if err != nil {
		httpError(writer, http.StatusInternalServerError, "Failed to update component status: %s", err)
		return
	}
}

func reportIssue(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		httpError(writer, http.StatusInternalServerError, "Unable to read body %s", err)
		return
	}
	defer request.Body.Close()

	var issue protobuf.Issue
	err = proto.Unmarshal(body, &issue)
	if err != nil {
		log.Fatalf("Unable to insert issue: %s\n", err)
	}

	utils.LogDebug("Received issue: %v", &issue)

	id, err := db.InsertIssue(conn, &issue)
	if err != nil {
		log.Fatalf("Unable to insert issue: %s\n", err)
	}
	issue.Id = id

	body, err = proto.Marshal(&issue)
	if err != nil {
		httpError(writer, http.StatusInternalServerError, "Unable to marshal issue: %s", err)
	}

	utils.LogDebug("Broadcasting issue: %s", body)

	hub.Broadcast <- body
}
