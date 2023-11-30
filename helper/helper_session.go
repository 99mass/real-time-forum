package helper

import (
	"database/sql"
	"forum/controller"
	"forum/models"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

// Example function to create and send a login session cookie
func AddSession(w http.ResponseWriter, userID uuid.UUID, db *sql.DB) {

	expiration := time.Now().Add(24 * time.Hour)
	if userID != uuid.Nil {
		session := models.Session{
			UserID:    userID,
			ExpiresAt: expiration,
			CreatedAt: time.Now(),
		}
		sessionID, err := controller.CreateSession(db, session) // You'll need to implement this function
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Create the session cookie
		cookie := http.Cookie{
			Name:     "sessionID",
			Value:    sessionID.String(),
			Expires:  expiration,
			HttpOnly: false,
			Path:     "/",
		}

		http.SetCookie(w, &cookie)
	}

}

func UpdateSession(db *sql.DB, sessionID uuid.UUID, newExpiration time.Time) error {
	query := `
		UPDATE sessions
		SET expires_at = ?
		WHERE id = ?;
	`

	_, err := db.Exec(query, newExpiration, sessionID)
	if err != nil {
		return err
	}

	return nil
}

func IsEmptySession(s models.Session) bool {
	return s == models.Session{}
}

func GetSessionRequest(r *http.Request) (uuid.UUID, error) {
	// Retrieve the session cookie named "sessionID"
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		// No session cookie foun
		return uuid.Nil, err
	}

	// Extract the value of the session cookie
	sessionid := cookie.Value

	sessionID, err := uuid.FromString(sessionid)
	if err != nil {
		return uuid.Nil, err
	}

	return sessionID, nil

}

func VerifySession(db *sql.DB, sessionID uuid.UUID) bool {
	session, err := controller.GetSessionByID(db, sessionID)
	if err != nil {
		return false
	}
	if &session == nil {
		return false
	}
	return true
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	// Create a new cookie with the same name as the session cookie
	cookie := http.Cookie{
		Name:     "sessionID",
		Value:    "",         // Empty value
		Expires:  time.Now(), // Set to a time in the past
		HttpOnly: false,
		Path:     "/",
	}

	// Set the cookie in the response, effectively deleting it
	http.SetCookie(w, &cookie)
}

// Example function to create and send a login session cookie
func UpdateCookieSession(w http.ResponseWriter, sessionID uuid.UUID, db *sql.DB) {

	expiration := time.Now().Add(24 * time.Hour)

	// Create the session cookie
	cookie := http.Cookie{
		Name:     "sessionID",
		Value:    sessionID.String(),
		Expires:  expiration,
		HttpOnly: false,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)
	UpdateSession(db, sessionID, expiration)
}
