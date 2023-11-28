package handler

import (
	"database/sql"
	"net/http"

	"github.com/gofrs/uuid"

	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
)

func GetMypage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Check if the user is connected
		var sessiondata bool

		ok, pageError := middlewares.CheckRequest(r, "/mypage", "get")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}

		sessionID, errsess := helper.GetSessionRequest(r)
		if errsess != nil {
			sessiondata = false
		} else {

			sessiondata = true

			_, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil {
				sessiondata = false
			}

		}
		if sessiondata {
			user, err := controller.GetUserBySessionId(sessionID, db)
			if err != nil {
				controller.DeleteSession(db, sessionID)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			category, err := controller.GetAllCategories(db)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			CatId := r.FormValue("categorie")
			if CatId != "" {
				CategID, err := uuid.FromString(CatId)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
					return
				}
				PostsDetails, err := helper.GetPostsForOneUserAndCategory(db, user.ID, CategID)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
				}
				for i := range PostsDetails {
					PostsDetails[i].Route = "mypage"
					for j := range PostsDetails[i].Comment {
						PostsDetails[i].Comment[j].Route = "mypage"
					}
				}

				//Set likes and dislikes
				Datasliked, err := helper.SetLikesAndDislikes(user, PostsDetails, db)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
				}

				PostsDetails = Datasliked

				datas := new(models.DataMypage)
				datas.Datas = PostsDetails
				datas.Session = sessiondata
				datas.User = user
				datas.CategoryID = CategID
				datas.Category = category
				helper.RenderTemplate(w, "mypage", "mypages", datas)
			} else {
				PostsDetails, err := helper.GetPostsForOneUser(db, user.ID)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}
				//fmt.Println(PostsDetails)
				for i := range PostsDetails {
					PostsDetails[i].Route = "mypage"
					//fmt.Println(PostsDetails[i].Route)
					for j := range PostsDetails[i].Comment {
						PostsDetails[i].Comment[j].Route = "mypage"
					}
				}
				//Set likes and dislikes
				Datasliked, err := helper.SetLikesAndDislikes(user, PostsDetails, db)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
				}

				PostsDetails = Datasliked

				datas := new(models.DataMypage)
				datas.Datas = PostsDetails

				datas.Session = sessiondata
				datas.User = user
				datas.Category = category
				helper.RenderTemplate(w, "mypage", "mypages", datas)
			}

		} else {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

	}
}

// Like post
func LikePosteByMyPage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		like := models.PostLike{}

		ok, errorPage := middlewares.CheckRequest(r, "/likepostmypage", "mypage")
		if !ok {
			helper.ErrorPage(w, errorPage)
			return
		}

		//check the session and get the user
		sessionID, errsess := helper.GetSessionRequest(r)
		if errsess != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {

			session, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil || &session == nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			User, errgetu := controller.GetUserBySessionId(sessionID, db)
			if errgetu != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			like.UserID = User.ID
		}

		postID, _ := helper.StringToUuid(r, "post_id")

		like.PostID = postID
		_, err := controller.CreatePostLike(db, like)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/mypage", http.StatusSeeOther)

	}
}
