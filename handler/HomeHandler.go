package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/helper"
	"forum/middlewares"
	"forum/models"

	"github.com/gofrs/uuid"
)

func Index(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ok, pageError := middlewares.CheckRequest(r, "/verifySession", "post")
		if !ok {
			helper.ErrorPage(w, pageError)
			return
		}
		var sessionReq models.SessionRequest
		err := json.NewDecoder(r.Body).Decode(&sessionReq)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect request",
			}, http.StatusBadRequest)
			return
		}
		sess, err := uuid.FromString(sessionReq.Session)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "Session ID incorrect",
			}, http.StatusBadRequest)
			return
		}
		if helper.VerifySession(db, sess) {
			homeData, err := helper.GetDataTemplate(db, r, true, false, true, false, true)

			PostsDetails := homeData.Datas

			homeData.Datas = PostsDetails
			if err != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: "something goes wrong",
				}, http.StatusInternalServerError)
				return
			}

			if homeData.Session {
				sessionID, _ := helper.GetSessionRequest(r)
				helper.UpdateCookieSession(w, sessionID, db)
			}

			helper.SendResponse(w, homeData, http.StatusOK)

		} else {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "Session invalid or expired",
			}, http.StatusBadRequest)
		}

	}
}
