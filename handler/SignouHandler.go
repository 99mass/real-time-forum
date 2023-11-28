package handler

import (
	"net/http"

	"forum/helper"
	"forum/middlewares"
)

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	ok, pageError := middlewares.CheckRequest(r, "/signout", "get")
	if !ok {
		helper.ErrorPage(w, pageError)
		return
	}
	helper.DeleteSession(w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
