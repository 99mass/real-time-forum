package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"

	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
)

func GetPostCategory(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok, pageError := middlewares.CheckRequest(r, "/category", "post")
		if !ok {
			helper.SendResponse(w,models.ErrorResponse{
				Status:  "error",
                Message: "Method not allowed",
			},pageError)
			return
		}
		var category models.GetPostsByCategoryRequest
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect json format",
			}, http.StatusBadRequest)
			return
		}
		categoryID, err := uuid.FromString(category.CategoryID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect category ID",
			}, http.StatusBadRequest)
			return
		}
		if VerifCategory(db, categoryID) {
			category, err := controller.GetCategoryByID(db, categoryID)
			if err != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: err.Error(),
				}, http.StatusInternalServerError)
			}
			homeData, err := helper.GetDataTemplate("", db, r, true, false, false, false, false)
			if err != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: err.Error(),
				}, http.StatusBadRequest)
				return
			}
			posts, err := helper.GetPostForCategory(db, categoryID)
			if err != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: err.Error(),
				}, http.StatusBadRequest)
				return
			}
			homeData.Datas = posts
			datas, errlike := helper.SetLikesAndDislikes(homeData.User, homeData.Datas, db)
			if errlike != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: err.Error(),
				}, http.StatusBadRequest)
				return
			}
			homeData.Datas = datas

			homeData.Category = append(homeData.Category, category)

			helper.SendResponse(w,homeData.Datas, http.StatusOK)
		} else {
			helper.SendResponse(w,models.ErrorResponse{
				Status:  "error",
				Message: "invalid category",
			},http.StatusBadRequest)
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
