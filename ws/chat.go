package ws

import (
	"fmt"
	"forum/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var users map[string]*models.User = make(map[string]*models.User)

func WSHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		//broadcastMessage(fmt.Sprintf("%s has joined the chat", username))
	}
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

type Message struct {
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

func parseMessage(msg Message) (string, string) {
	recipient := msg.Recipient
	messageContent := msg.Message
	return recipient, messageContent
}

func handleMessages(conn *websocket.Conn, username string) {
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		recipient, message := parseMessage(msg)

		if recipient != "" && message != "" {
			sendMessage(recipient, fmt.Sprintf("%s: %s", username, message))
		} else {
			errMessage := map[string]string{
				"error": "Message cannot be empty",
			}
			conn.WriteJSON(errMessage)

		}
	}
}

func sendMessage(recipient string, message string) {
	if user, ok := users[recipient]; ok {
		user.Conn.WriteJSON(message)
	}
}


func CloseConnection(username string) {
    user, ok := users[username]
    if ok {
        err := user.Conn.Close()
        if err != nil {
            log.Printf("Error closing connection for user %s: %v", username, err)
        }
        delete(users, username)
    } else {
        log.Printf("User %s not found", username)
    }
}

func BroadcastMessage(message string) {
	for _, user := range users {
		user.Conn.WriteJSON(message)
	}
}
