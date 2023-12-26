package ws

import (
	"database/sql"
	"fmt"
	"forum/controller"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

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
			User1 string `json:"User1"`
			User2 string `json:"User2"`
		}
		var request requestUSer
		err = conn.ReadJSON(&request)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(request.User1)
		// fmt.Println(request.User2)
		discuss, err := GetCommunication(db, request.User1, request.User2)
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}

		goodDiscuss, err := GoodToSend(db, discuss)
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		//fmt.Println(goodDiscuss)
		conn.WriteJSON(goodDiscuss)
		conn.Close()

	}
}

type messageToSend struct {
	Sender    string
	Recipient string
	Message   string
	Created   string
}

func GoodToSend(db *sql.DB, discuss []Message) ([]messageToSend, error) {

	var messToSend []messageToSend
	for _, m := range discuss {
		//fmt.Println(m.Message)
		var mes messageToSend
		send, err := GetUsername(db, m.Sender)
		if err != nil {
			return nil, err
		}
		mes.Sender = send
		recep, err := GetUsername(db, m.Recipient)
		if err != nil {
			return nil, err
		}
		mes.Recipient = recep
		mes.Message = m.Message
		mes.Created = m.Created.Format("2006-01-02 15:04:05")
		messToSend = append(messToSend, mes)
	}
	//fmt.Println(messToSend)
	return messToSend, nil
}

func GetUsername(db *sql.DB, user uuid.UUID) (string, error) {
	us, err := controller.GetUserByID(db, user)
	if err != nil {
		return "", err
	}
	return us.Username, nil
}

func GetCommunication(db *sql.DB, user1 string, user2 string) ([]Message, error) {
	us1, err := GetUserIDByUserName(db, user1)
	if err != nil {
		return nil, err
	}
	us2, err := GetUserIDByUserName(db, user2)
	if err != nil {
		return nil, err
	}
	message1, err := GetMessageSentByOneUserToAnotherOne(db, us2, us1)
	if err != nil {
		return nil, err
	}
	message2, err := GetMessageSentByOneUserToAnotherOne(db, us1, us2)
	if err != nil {
		return nil, err
	}

	discussion := GetDiscussion(db, us1, us2)

	for _, m := range message1 {
		if !m.Read {
			MarkMessageAsRead(db, m.ID)
		}
	}
	for _, m := range message2 {
		if !m.Read {
			MarkMessageAsRead(db, m.ID)
		}
	}

	return discussion, nil
}
