package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"

	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
)

func SinginHandler(db *sql.DB) http.HandlerFunc {
	var homeData models.Home
	homeData.Session = false
	homeData.ErrorAuth.EmailError = ""
	homeData.ErrorAuth.GeneralError = ""
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			ok, pageError := middlewares.CheckRequest(r, "/signin", "post")
			if !ok {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: "not found",
				}, pageError)
				return
			}

			datas, err := helper.GetDataTemplate(db, r, false, false, false, true, false)

			if err != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: "user doesn't exist",
				}, http.StatusBadRequest)
				return
			}
			nul := uuid.UUID{}
			if datas.User.ID != nul {
				sess, err := controller.GetSessionIDForUser(db, datas.User.ID)
				if err == nil {
					err := controller.DeleteSession(db, sess)
					if err != nil {
						helper.ErrorPage(w, http.StatusInternalServerError)
						return
					}
				}
				helper.AddSession(w, datas.User.ID, db)
				// Redirect to home
				helper.SendResponse(w, datas, http.StatusOK)
			} else {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: datas.ErrorAuth.GeneralError,
				}, http.StatusBadRequest)
			}
		default:
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "Method not allowed",
			}, http.StatusMethodNotAllowed)
			return
		}
	}
}

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	var homeData models.Home
	homeData.Session = false

	return func(w http.ResponseWriter, r *http.Request) {

		// Handle according to the method
		switch r.Method {
		case http.MethodPost:
			ok, pageError := middlewares.CheckRequest(r, "/register", "post")
			if !ok {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: "not found",
				}, pageError)
				return
			}
			var registerReq models.RegisterRequest
			err := json.NewDecoder(r.Body).Decode(&registerReq)
			if err != nil {
				helper.SendResponse(w,models.ErrorResponse{
					Status: "error",
					Message: "incorrect request",
				},http.StatusBadRequest)
				return
			}
			username := registerReq.UserName
			username = strings.TrimSpace(username)
			email := registerReq.Email
			email = strings.TrimSpace(email)
			password := registerReq.Password
			password = strings.TrimSpace(password)
			confirmPassword := registerReq.Confpassword
			confirmPassword = strings.TrimSpace(confirmPassword)
			// Hasher le mot de passe
			hashedPassword, _ := helper.HashPassword(password)

			ok, ErrAuth := helper.CheckRegisterFormat(username, email, password, confirmPassword, db)

			if !ok {
				//homeData.ErrorAuth = ErrAuth
				helper.SendResponse(w,models.ErrorResponse{
					Status: "error",
					Message: "register format:"+ErrAuth.GeneralError+ErrAuth.EmailError+ErrAuth.PasswordError+ErrAuth.UserNameError,
				},http.StatusBadRequest)
				//homeData.ErrorAuth = models.ErrorAuth{}
				return
			}

			user := models.User{
				Username:  username,
				Email:     email,
				Password:  hashedPassword,
				CreatedAt: time.Now(),
			}

			id, _ := controller.CreateUser(db, user)

			// create a session - TODO
			helper.AddSession(w, id, db)
			helper.SendResponse(w,user,http.StatusOK)
			//helper.RenderTemplate(w, "index", "index", "homedata")
			return

		// case http.MethodGet:
		// 	helper.DeleteSession(w, r)
		// 	ok, pageError := middlewares.CheckRequest(r, "/register", "get")
		// 	if !ok {
		// 		helper.ErrorPage(w, pageError)
		// 		return
		// 	}
		// 	helper.RenderTemplate(w, "register", "auth", homeData)
		default:
			helper.SendResponse(w,models.ErrorResponse{
				Status: "error Method",
				Message: "Method not Allowed",
			},http.StatusMethodNotAllowed)
			return
		}
	}
}
