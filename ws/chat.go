package ws

import (
	"encoding/json"
	"fmt"
	"forum/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var users map[string]*models.User = make(map[string]*models.User)

func WSHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer conn.Close()

	username, err := readUsername(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	users[username] = &models.User{Conn: conn, Username: username}
	fmt.Println(users)
	go handleMessages(conn, username)

	broadcastMessage(fmt.Sprintf("%s has joined the chat", username))
}

type UsernameMessage struct {
	Username string `json:"username"`
}

func readUsername(conn *websocket.Conn) (string, error) {
	var msg UsernameMessage
	err := conn.ReadJSON(&msg)
	if err != nil {
		return "", err
	}
	return msg.Username, nil
}

func parseMessage(message string) (string, string) {
	var msg map[string]interface{}
	err := json.Unmarshal([]byte(message), &msg)
	if err != nil {
		// Handle error
		return "", ""
	}

	recipient, ok := msg["recipient"].(string)
	if !ok {
		recipient = ""
	}

	messageContent, ok := msg["message"].(string)
	if !ok {
		messageContent = ""
	}

	return recipient, messageContent
}

func handleMessages(conn *websocket.Conn, username string) {
	for {
		var message string
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		recipient, message := parseMessage(message)

		if recipient != "" {
			sendMessage(recipient, fmt.Sprintf("%s: %s", username, message))
		} else {
			broadcastMessage(fmt.Sprintf("%s: %s", username, message))
		}
	}
}

func sendMessage(recipient string, message string) {
	if user, ok := users[recipient]; ok {
		user.Conn.WriteJSON(message)
	}
}

func broadcastMessage(message string) {
	for _, user := range users {
		user.Conn.WriteJSON(message)
	}
}
