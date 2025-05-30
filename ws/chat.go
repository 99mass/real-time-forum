package ws

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/controller"
	"forum/models"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

var usersConnected []string
var users map[string]*models.User = make(map[string]*models.User)
var userList []UserToShow
var usersMessage map[string]*models.User = make(map[string]*models.User)

type UserToShow struct {
	Username string `json:"Username"`
	Status   string `json:"Status"`
}
type ConnectedUser struct {
	Users []string
}

type MessageUnread struct {
	Sender        string
	NumberMessage int
}

func GetNumberMessage(db *sql.DB, user []UserToShow, receiver string) []MessageUnread {
	var messagesUnread []MessageUnread
	receiverID, _ := GetUserIDByUserName(db, receiver)

	for _, us := range user {
		if us.Username != receiver {
			usID, _ := GetUserIDByUserName(db, us.Username)
			message, _ := GetMessageSentByOneUserToAnotherOne(db, usID, receiverID)
			var n int
			for _, m := range message {
				if !m.Read {
					n++
				}
			}
			var mess MessageUnread
			mess.Sender = us.Username
			mess.NumberMessage = n
			messagesUnread = append(messagesUnread, mess)
		}
	}

	return messagesUnread

}
func GetUserOrder(db *sql.DB, receiver string, users []UserToShow) []UserToShow {
	var userList []UserToShow
	var messages []Message
	receiverID, _ := GetUserIDByUserName(db, receiver)
	for _, us := range users {
		usID, _ := GetUserIDByUserName(db, us.Username)
		message, _ := GetMessageSentByOneUserToAnotherOne(db, usID, receiverID)
		if message == nil {
			user, _ := controller.GetUserByUsername(db, us.Username)
			if us.Username != receiver {
				var mess Message
				mess.Sender = user.ID
				message = append(message, mess)
			}
		}

		messages = append(messages, message...)
	}
	for _, us := range users {
		usID, _ := GetUserIDByUserName(db, us.Username)
		message, _ := GetMessageSentByOneUserToAnotherOne(db, receiverID, usID)
		if message != nil {
			messages = append(messages, message...)
		}

	}

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Created.After(messages[j].Created)
	})

	for _, m := range messages {
		var us UserToShow

		name, _ := GetUsername(db, m.Sender)
		if name != receiver {
			us.Username = name
		} else {
			na, _ := GetUsername(db, m.Recipient)
			us.Username = na
		}

		us.Status = "offline"

		// Check if user already exists in userList
		exists := false
		for _, user := range userList {
			if user.Username == us.Username {
				exists = true
				break
			}
		}

		// If user does not exist in userList, append
		if !exists {
			userList = append(userList, us)
		}
	}

	return userList
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

		// correctUserOrder := GetUserOrder(db, username, userList)
		// changeElementStatus(correctUserOrder, usersConnected)
		// fmt.Println("this is the correct user order :", correctUserOrder)

		BroadcastUsers(db, userList)

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
			log.Println(err)
			return
		}

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
			sendMessage(db, recipient, msg)
			conn.WriteJSON(msg)
		}
		//update order of the users list
		BroadcastUsers(db, userList)
	}
}

func SaveMessage(db *sql.DB, sender string, recipient string, message string) error {
	senderID, err := GetUserIDByUserName(db, sender)
	if err != nil {
		return err
	}
	recipientID, err := GetUserIDByUserName(db, recipient)
	if err != nil {
		return err
	}
	var Mes Message
	Mes.Sender = senderID
	Mes.Recipient = recipientID
	Mes.Message = message
	_, err = CreateMessage(db, Mes)
	if err != nil {
		return errors.New("can't create message : " + err.Error())
	}
	log.Println("message saved successfully")
	return nil
}

func sendMessage(db *sql.DB, recipient string, message GetMessage) {
	if user, ok := usersMessage[recipient]; ok {
		user.Conn.WriteJSON(message)
	}
	fmt.Println(users)
	if _, ok := users[recipient]; ok {
		fmt.Println(ok)
		sender, recipient, _, _ := parseMessage(message)
		sendr, _ := GetUserIDByUserName(db, sender)
		receprnt, _ := GetUserIDByUserName(db, recipient)
		message1, _ := GetMessageSentByOneUserToAnotherOne(db, sendr, receprnt)
		for _, m := range message1 {
			if !m.Read {
				MarkMessageAsRead(db, m.ID)
			}
		}
	}
	log.Println("message sent successfully to : " + recipient)
}

func GetUserIDByUserName(db *sql.DB, userName string) (uuid.UUID, error) {
	user, err := controller.GetUserByUsername(db, userName)
	if err != nil {
		return uuid.Nil, err
	}
	return user.ID, nil
}

func removeElement(slice []string, el string) []string {
	for i, a := range slice {
		if a == el {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func CloseConnection(db *sql.DB, username string) {
	user, ok := users[username]
	if ok {
		err := user.Conn.Close()
		if err != nil {
			log.Printf("Error closing connection for user %s: %v", username, err)
		}
		delete(users, username)

		usersConnected = removeElement(usersConnected, username)
		for i, us := range userList {
			if us.Username == username {
				userList[i].Status = "offline"
			}
		}
		//fmt.Println(userList)
		BroadcastUsers(db, userList)

	} else {
		log.Printf("User %s not found", username)
	}

	userMessage, okMessage := usersMessage[username]
	if okMessage {
		err := userMessage.Conn.Close()
		if err != nil {
			log.Printf("Error closing connection for user %s: %v", username, err)
		}
		delete(usersMessage, username)
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

func BroadcastUsers(db *sql.DB, userList []UserToShow) {
	for _, user := range users {
		cur := GetUserOrder(db, user.Username, userList)
		changeElementStatus(cur, usersConnected)

		usersConn := removeUser(cur, user.Username)
		mess := GetNumberMessage(db, cur, user.Username)
		//fmt.Println(mess)
		user.Conn.WriteJSON(mess)
		user.Conn.WriteJSON(usersConn)

	}
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
	sort.Slice(userList, func(i, j int) bool {
		return userList[i].Username < userList[j].Username
	})
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
