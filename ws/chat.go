package ws

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/controller"
	"forum/models"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

var usersConnected []string
var users map[string]*models.User = make(map[string]*models.User)
var userList []UserToShow
var usersMessage map[string]*models.User = make(map[string]*models.User)

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

		BroadcastUsers(userList)

		//broadcastMessage(fmt.Sprintf("%s has joined the chat", username))
	}
}

func HandlerMessages(db *sql.DB) http.HandlerFunc {
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
		//fmt.Println("user : ", username)

		if user, ok := usersMessage[username]; ok {
			// Si l'utilisateur existe déjà, mettez à jour la connexion
			user.Conn = conn
		} else {
			// Sinon, créez un nouvel utilisateur
			usersMessage[username] = &models.User{Conn: conn, Username: username}
		}
		//fmt.Println("list : ", usersMessage)

		go handleMessages(db, conn, username)
	}
}

func CommunicationHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		type requestUSer struct {
			User1    string	`json:"User1"`
			User2 string	`json:"User2"`
		}
		var request requestUSer
		err = conn.ReadJSON(&request)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(request.User1)
		fmt.Println(request.User2)
		// val1,_ := uuid.FromString("83a723ca-4d38-4af1-bd0d-1404adff7a7f")
		// val2,_ := uuid.FromString("b595fade-d0d2-463b-8396-f4cf6ea8661e")
		// dis := GetDiscussion(db,val1,val2)
		// fmt.Println(dis)

		discuss := GetCommunication(db, request.User1, request.User2)
		fmt.Println(discuss)
		conn.WriteJSON(discuss)

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

type GetMessage struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
	Created   string `json:"created"`
}

func parseMessage(msg GetMessage) (string, string, string, error) {
	sender := msg.Sender
	recipient := msg.Recipient
	messageContent := msg.Message
	if sender == "" || recipient == "" || messageContent == "" {
		return "", "", "", errors.New("invalid type for message")
	}
	log.Println("message parsed successfully")
	return sender, recipient, messageContent, nil
}

func handleMessages(db *sql.DB, conn *websocket.Conn, username string) {
	for {
		if user, ok := usersMessage[username]; ok {
			// Si l'utilisateur existe déjà, mettez à jour la connexion
			user.Conn = conn
		} else {
			// Sinon, créez un nouvel utilisateur
			usersMessage[username] = &models.User{Conn: conn, Username: username}
		}

		//fmt.Println("list 2: ", usersMessage)
		var msg GetMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		//fmt.Println("modele : ", msg)

		sender, recipient, message, err := parseMessage(msg)
		if err == nil {
			if len(message) > 100 {
				conn.WriteJSON("message is too long")
				continue
			}
			err = SaveMessage(db, sender, recipient, message)
			if err != nil {
				fmt.Println(err)
			}
			msg.Created = time.Now().Format("2006-01-02 15:04:05")
			sendMessage(recipient, msg)
			sendMessage(sender, msg)
		}

	}
}

func SaveMessage(db *sql.DB, sender string, recipient string, message string) error {
	senderID := GetUserIDByUserName(db, sender)
	recipientID := GetUserIDByUserName(db, recipient)
	var Mes Message
	Mes.Sender = senderID
	Mes.Recipient = recipientID
	Mes.Message = message
	_, err := CreateMessage(db, Mes)
	if err != nil {
		return errors.New("can't create message : " + err.Error())
	}
	log.Println("message saved successfully")
	return nil
}

func sendMessage(recipient string, message GetMessage) {
	if user, ok := usersMessage[recipient]; ok {
		user.Conn.WriteJSON(message)
	}
	log.Println("message sent successfully to : " + recipient)
}

func GetUserIDByUserName(db *sql.DB, userName string) uuid.UUID {
	user, err := controller.GetUserByUsername(db, userName)
	if err != nil {
		return uuid.Nil
	}
	return user.ID
}

func removeElement(slice []string, el string) []string {
	for i, a := range slice {
		if a == el {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func CloseConnection(username string) {
	user, ok := users[username]
	if ok {
		err := user.Conn.Close()
		if err != nil {
			log.Printf("Error closing connection for user %s: %v", username, err)
		}
		delete(users, username)
		delete(usersMessage, username)
		usersConnected = removeElement(usersConnected, username)
		for i, us := range userList {
			if us.Username == username {
				userList[i].Status = "offline"
			}
		}
		fmt.Println(userList)
		BroadcastUsers(userList)

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
