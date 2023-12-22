package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
	"forum/ws"

	"github.com/gofrs/uuid"
)

func SignOutHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ok, pageError := middlewares.CheckRequest(r, "/signout", "post")
		if !ok {
			helper.ErrorPage(w, pageError)
			return
		}

		var SessionReq models.SessionRequest
		err := json.NewDecoder(r.Body).Decode(&SessionReq)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect request",
			}, http.StatusBadRequest)
			return
		}
		sessionID := SessionReq.Session
		IDsess, _ := uuid.FromString(sessionID)
		user, _ := controller.GetUserBySessionId(IDsess, db)
		ws.CloseConnection(db, user.Username)
		helper.DeleteSession(db, sessionID, w, r)

		helper.SendResponse(w, models.ErrorResponse{}, http.StatusOK)
	}
}
