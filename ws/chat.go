package ws

import (
	"database/sql"
	"fmt"
	"forum/controller"
	"forum/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var usersConnected []string
var users map[string]*models.User = make(map[string]*models.User)
var userList []UserToShow
type ConnectedUser struct {
	Users []string
}

func WSHandler(db *sql.DB) http.HandlerFunc {
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

		userList, err = GetAllUserNames(db)
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

		changeElementStatus(userList, usersConnected)

		fmt.Println(userList)

		fmt.Println(users)

		fmt.Println(usersConnected)

		if len(usersConnected) != 1 {
			BroadcastUsers(userList)
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
		for i,us := range userList{
			if us.Username == username {
				userList[i].Status = "offline"
			}
		}
		fmt.Println(userList)
		if len(usersConnected) != 1 {
			BroadcastUsers(userList)
		} else {
			noUser := map[string]string{
				"message": "there's no user online",
			}
			for _, us := range users {
				us.Conn.WriteJSON(noUser)
			}
		}
	} else {
		log.Printf("User %s not found", username)
	}
}

func removeUser(slice []UserToShow, user string) []UserToShow {
	var tab []UserToShow
	for _, val := range slice {
		if val.Username != user {
			tab = append(tab, val)
		}
	}
	return tab
}

func BroadcastUsers(userList []UserToShow) {
	for _, user := range users {
		usersConn := removeUser(userList, user.Username)

		user.Conn.WriteJSON(usersConn)

	}
}

type UserToShow struct {
	Username string `json:"Username"`
	Status   string `json:"Status"`
}

func GetAllUserNames(db *sql.DB) ([]UserToShow, error) {
	var userList []UserToShow
	users, err := controller.GetAllUsers(db)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		var us UserToShow
		us.Username = user.Username
		us.Status = "offline"
		userList = append(userList, us)
	}
	return userList, nil
}

func changeElementStatus(slice1 []UserToShow, slice2 []string) {

	for i, el := range slice1 {
		if contains(slice2, el.Username) {
			slice1[i].Status = "online"
		}
	}

}

func contains(slice []string, el string) bool {
	for _, a := range slice {
		if a == el {
			return true
		}
	}
	return false
}
