package handler

import (
	"database/sql"
	"net/http"

	"github.com/gofrs/uuid"

	"forum/controller"
	"forum/helper"
	"forum/middlewares"
)

func GetPostCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok, pageError := middlewares.CheckRequest(r, "/category", "get")
		if !ok {
			helper.ErrorPage(w, pageError)
			return
		}
		categoryID, err := helper.StringToUuid(r, "categorie")
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}
		if VerifCategory(db, categoryID) {
			category, err := controller.GetCategoryByID(db, categoryID)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			homeData, err := helper.GetDataTemplate(db, r, true, false, false, false, false)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
			}
			posts, err := helper.GetPostForCategory(db, categoryID)
			if err != nil {

				helper.ErrorPage(w, http.StatusInternalServerError)
			}
			homeData.Datas = posts
			datas, errlike := helper.SetLikesAndDislikes(homeData.User, homeData.Datas, db)
			if errlike != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
			}
			homeData.Datas = datas

			homeData.Category = append(homeData.Category, category)

			helper.RenderTemplate(w, "categorie", "categories", homeData)
		} else {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}

	}
}

func VerifCategory(db *sql.DB, Id uuid.UUID) bool {
	category, err := controller.GetCategoryByID(db, Id)
	if err != nil {
		return false
	}
	if &category == nil {
		return false
	}
	return true
}
