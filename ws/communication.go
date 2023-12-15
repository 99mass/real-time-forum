package ws

import (
	"database/sql"
	"fmt"
	"net/http"

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
		fmt.Println(request.User1)
		fmt.Println(request.User2)
		discuss := GetCommunication(db, request.User1, request.User2)
		goodDiscuss := GoodToSend(db, discuss)
		fmt.Println(goodDiscuss)
		conn.WriteJSON(goodDiscuss)

	}
}

type messageToSend struct {
	Sender    string
	Recipient string
	Message   string
	Created   string
}

func GoodToSend(db *sql.DB, discuss []Message) []messageToSend {

	var messToSend []messageToSend
	for _, m := range discuss {
		//fmt.Println(m.Message)
		var mes messageToSend
		mes.Sender = GetUsername(db, m.Sender)
		mes.Recipient = GetUsername(db, m.Recipient)
		mes.Message = m.Message
		mes.Created = m.Created.Format("2006-01-02 15:04:05")
		messToSend = append(messToSend, mes)
	}
	//fmt.Println(messToSend)
	return messToSend
}
