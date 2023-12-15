package ws

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

type Message struct {
	ID        uuid.UUID
	Sender    uuid.UUID
	Recipient uuid.UUID
	Message   string
	Created   string
}

func CreateMessage(db *sql.DB, message Message) (uuid.UUID, error) {
	query := `
        INSERT INTO messages (id, sender, recipient, message, created_at)
        VALUES (?, ?, ?, ?, ?);
    `

	newUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = db.Exec(query, newUUID.String(), message.Sender, message.Recipient, message.Message, time.Now())
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUUID, nil
}

// func GetDiscussion(db *sql.DB, user1 uuid.UUID, user2 uuid.UUID) []Message {
// 	var messages []Message

// 	// Get all messages sent by user 1
// 	messages1 := getMessagesForUser(db, user1)

// 	// Get all messages sent by user 2
// 	messages2 := getMessagesForUser(db, user2)

// 	// Find messages sent to both users
// 	for _, m1 := range messages1 {
// 		for _, m2 := range messages2 {
// 			if m1.Recipient == user2.String() && m2.Recipient == user1.String() {
// 				messages = append(messages, m1, m2)
// 			}
// 		}
// 	}

// 	// Sort messages by creation date
// 	sort.Slice(messages, func(i, j int) bool {
// 		return messages[i].Created.Before(messages[j].Created)
// 	})

// 	return messages
// }

func getMessagesForUser(db *sql.DB, userID uuid.UUID) []Message {
	query := `
        SELECT id, sender, recipient, message, created_at
        FROM messages
        WHERE recipient = ?
        OR sender = ?;
    `

	rows, err := db.Query(query, userID.String(), userID.String())
	if err != nil {
		return []Message{}
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var message Message
		err := rows.Scan(&message.ID, &message.Sender, &message.Recipient, &message.Message, &message.Created)
		if err != nil {
			return []Message{}
		}
		messages = append(messages, message)
	}

	return messages
}
