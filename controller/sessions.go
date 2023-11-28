package controller

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"

	"forum/models"
)

func CreateSession(db *sql.DB, session models.Session) (uuid.UUID, error) {
	query := `
        INSERT INTO sessions (id, user_id, expires_at)
        VALUES (?, ?, ?);
    `

	newUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = db.Exec(query, newUUID.String(), session.UserID, session.ExpiresAt)
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUUID, nil
}

// GetSessionByID retrieves a session by ID from the database.
func GetSessionByID(db *sql.DB, sessionID uuid.UUID) (models.Session, error) {
	var session models.Session
	query := `
        SELECT id, user_id, expires_at
        FROM sessions
        WHERE id = ?;
    `

	err := db.QueryRow(query, sessionID).Scan(&session.ID, &session.UserID, &session.ExpiresAt)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}

// GetSessionUserID retrieves the user ID associated with a session from the database.
func GetSessionUserID(db *sql.DB, sessionID uuid.UUID) (uuid.UUID, error) {
	query := `
        SELECT user_id
        FROM sessions
        WHERE id = ? AND expires_at > ?;
    `

	var userID uuid.UUID
	err := db.QueryRow(query, sessionID, time.Now()).Scan(&userID)
	if err != nil {
		return uuid.UUID{}, err
	}

	return userID, nil
}

// DeleteSession deletes a session by ID from the database.
func DeleteSession(db *sql.DB, sessionID uuid.UUID) error {
	query := `
        DELETE FROM sessions
        WHERE id = ?;
    `
	
	_, err := db.Exec(query, sessionID)
	return err
}

// ValidateSession checks if a session is valid based on its expiration time.
func ValidateSession(session models.Session) bool {
	return session.ExpiresAt.After(time.Now())
}

// GetSessionIDForUser retrieves the session ID for a given user ID from the database.
func GetSessionIDForUser(db *sql.DB, userID uuid.UUID) (uuid.UUID, error) {
	var sessionID uuid.UUID
	query := `
		SELECT id FROM sessions
		WHERE user_id = ?
		LIMIT 1;
	`

	row := db.QueryRow(query, userID)
	err := row.Scan(&sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			// No session found for the user
			return uuid.Nil, nil
		}
		return uuid.Nil, err
	}

	return sessionID, nil
}
