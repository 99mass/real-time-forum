package handler

import (
	"database/sql"
	"net/http"

	"forum/helper"
	"forum/middlewares"
)

func Index(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ok, pageError := middlewares.CheckRequest(r, "/", "get")
		if !ok {
			helper.ErrorPage(w, pageError)
			return
		}

		homeData, err := helper.GetDataTemplate(db, r, true, false, true, false, true)

		PostsDetails := homeData.Datas

		homeData.Datas = PostsDetails
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}

		if homeData.Session {
			sessionID, _ := helper.GetSessionRequest(r)
			helper.UpdateCookieSession(w, sessionID, db)
		}
		helper.RenderTemplate(w, "index", "index", homeData)

	}
}
