package ws

import (
	"fmt"
	"forum/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var usersConnected []string
var users map[string]*models.User = make(map[string]*models.User)

type ConnectedUser struct {
	Users []string
}

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

		if user, ok := users[username]; ok {
			// Si l'utilisateur existe déjà, mettez à jour la connexion
			user.Conn = conn
		} else {
			// Sinon, créez un nouvel utilisateur
			users[username] = &models.User{Conn: conn, Username: username}
			usersConnected = append(usersConnected, username)
		}

		fmt.Println(users)
		// users[username] = &models.User{Conn: conn, Username: username}
		// usersConnected = append(usersConnected, username)

		fmt.Println(usersConnected)

		if len(usersConnected) != 1 {
			BroadcastMessage(usersConnected)
		} else {
			noUser := map[string]string{
				"message": "there's no user online",
			}
			conn.WriteJSON(noUser)
		}

		go handleMessages(conn, username)

		//broadcastMessage(fmt.Sprintf("%s has joined the chat", username))
	}
}

type UsernameMessage struct {
	Username string `json:"Username"`
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
		usersConnected = removeUser(usersConnected, username)
		fmt.Println(usersConnected)
	} else {
		log.Printf("User %s not found", username)
	}
}

func removeUser(slice []string, user string) []string {
	var tab []string
	for _, val := range slice {
		if val != user {
			tab = append(tab, val)
		}
	}
	return tab
}

func BroadcastMessage(message []string) {
	for _, user := range users {
		usersConn := removeUser(message, user.Username)

		user.Conn.WriteJSON(usersConn)

	}
}

// _, ok := users[username]
// if ok {
// 	log.Println("user is already connected")
// 	conn.Close()
// 	return
// }

// func EndPointConnectedUser(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		ok, pageError := middlewares.CheckRequest(r, "/connectedUsers", "get")
// 		if !ok {
// 			helper.SendResponse(w, models.ErrorResponse{
// 				Status:  "error",
// 				Message: "Method not Allowed",
// 			}, pageError)
// 			return
// 		}

// 		sessionID, err := helper.GetSessionRequest(r)
// 		if err != nil {
// 			helper.SendResponse(w, models.ErrorResponse{
// 				Status:  "error",
// 				Message: "there's no session",
// 			}, http.StatusBadRequest)
// 			return
// 		}
// 		user, err := controller.GetUserBySessionId(sessionID, db)
// 		if err != nil {
// 			helper.SendResponse(w, models.ErrorResponse{
// 				Status:  "error",
// 				Message: "session is not valid",
// 			}, http.StatusBadRequest)
// 			return
// 		}

// 		username := user.Username

// 		usersConn := removeUser(usersConnected, username)
// 		if len(usersConn) != 0 {
// 			var connected ConnectedUser
// 			connected.Users = usersConn

// 			helper.SendResponse(w, connected, http.StatusOK)
// 		} else {
// 			noUser := map[string]string{
// 				"message": "there's no user online",
// 			}
// 			helper.SendResponse(w, noUser, http.StatusOK)
// 		}

// 	}
// }
