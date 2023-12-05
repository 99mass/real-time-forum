package handler

import (
	"forum/controller"
	"forum/helper"
	"forum/models"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID       string
	UserName string
	Conn     *websocket.Conn
}

type Server struct {
	Clients map[string]*Client
}
type Message struct {
	Sender    string `json:"sender"`
	Dest      string `json:"dest"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

func (s *Server) HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		helper.SendResponse(w, models.ErrorResponse{
			Status:  "error",
			Message: "Upgrade error",
		}, http.StatusBadRequest)
	}

	// Create new client
	client := &Client{
		ID:   r.URL.Query().Get("id"),
		Conn: ws,
	}
	var user *models.User
	uid, _ := uuid.FromString(client.ID)
	user, err = controller.GetUserByID(helper.DB, uid)
	if err != nil {
		helper.SendResponse(w, models.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	client.UserName = user.Username
	//fmt.Println(client.Conn)
	// Add new client to the server's clients map
	s.Clients[client.ID] = client

	// Handle incoming messages
	//go s.handleMessages(w,client)
}

func (s *Server) handleMessages(w http.ResponseWriter, client *Client) {
	for {
		// Read in a new message as JSON and map it to a Message object
		var msg Message
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect JSON format",
			}, http.StatusBadRequest)
			continue
		}

		// Send the newly received message to the right client
		if destClient, ok := s.Clients[msg.Dest]; ok {
			err := destClient.Conn.WriteJSON(msg)
			if err != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: "can't send message",
				}, http.StatusBadRequest)
				continue
			}
		}
	}
}

func (s *Server) HandleMessages(client *Client) {
	for {
		// Read in a new message as JSON and map it to a Message object
		var msg Message
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(s.Clients, client.ID)
			break
		}

		// Send the newly received message to the right client
		if destClient, ok := s.Clients[msg.Dest]; ok {
			err := destClient.Conn.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				delete(s.Clients, destClient.ID)
			}
		}
	}
}

// var connection = new WebSocket('ws://localhost:8080/ws?id=yourID');

// connection.onopen = function () {
//   console.log('Connected!');
// };

// connection.onerror = function (error) {
//   console.log('WebSocket Error ' + error);
// };

// connection.onmessage = function (e) {
//   console.log('Server: ' + e.data);
// };

// function sendMessage(msg, dest) {
//   connection.send(JSON.stringify({Dest: dest, Content: msg}));
// }
