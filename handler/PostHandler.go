package handler

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"

	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
)

func AddComment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok, pageError := middlewares.CheckRequest(r, "/addcomment", "post")
		if !ok {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "Method not allowed",
			}, pageError)
			return
		}
		var comment models.Comment
		var commentRequest models.AddCommentRequest
		err := json.NewDecoder(r.Body).Decode(&commentRequest)
		if err != nil || commentRequest.PostID == "" || commentRequest.UserID == "" {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect json format",
			}, http.StatusBadRequest)
			return
		}

		postID, errP := uuid.FromString(commentRequest.PostID)
		userID, errU := uuid.FromString(commentRequest.UserID)
		Content := commentRequest.Content
		Content = strings.TrimSpace(Content)
		if errP != nil || errU != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: errP.Error() + errU.Error(),
			}, http.StatusBadRequest)
			return
		}

		homeDataSess, err := helper.GetDataTemplate(commentRequest.PostID, db, r, true, false, false, false, false)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: err.Error(),
			}, http.StatusBadRequest)
			return
		}

		if !homeDataSess.Session {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "session invalid or expired",
			}, http.StatusBadRequest)
			return
		}
		if Content == "" {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "a comment can't be empty",
			}, http.StatusBadRequest)
			return
		}
		if len(Content) > 500 {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "the number of characters must not exceed 1000",
			}, http.StatusBadRequest)
			return
		}
		comment.PostID = postID
		comment.UserID = userID
		comment.Content = Content

		_, erro := controller.CreateComment(db, comment)
		if erro != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: erro.Error(),
			}, http.StatusBadRequest)
			return
		}
		homeData, err := helper.GetDataTemplate(commentRequest.PostID, db, r, true, true, false, false, false)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "hello" + err.Error(),
			}, http.StatusBadRequest)
			return
		}
		posts, err := helper.GetPostsForOneUser(db, homeData.User.ID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		//Set likes and dislikes
		postsliked, err := helper.SetLikesAndDislikes(homeData.User, posts, db)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: err.Error(),
			}, http.StatusBadRequest)
			return
		}

		posts = postsliked

		category, err := controller.GetCategoriesByPost(db, homeData.PostData.Posts.ID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		homeData.Category = category
		homeData.Datas = posts
		helper.SendResponse(w, homeData.PostData, http.StatusOK)
	}
}

func GetOnePost(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ok, pageError := middlewares.CheckRequest(r, "/post", "post")
		if !ok {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "Method not allowed",
			}, pageError)
			return
		}

		var post models.OnePostRequest
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect json format",
			}, http.StatusBadRequest)
			return
		}
		homeData, err := helper.GetDataTemplate(post.PostID, db, r, true, true, false, false, false)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		postid := post.PostID
		posts, err := helper.GetPostsForOneUser(db, homeData.PostData.User.ID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		postsliked, err := helper.SetLikesAndDislikes(homeData.User, posts, db)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: err.Error(),
			}, http.StatusBadRequest)
			return
		}

		posts = postsliked
		for i := range posts {
			posts[i].Route = "post?post_id=" + postid
			for j := range posts[i].Comment {
				posts[i].Comment[j].Route = "post?post_id=" + postid
			}
		}
		category, err := controller.GetCategoriesByPost(db, homeData.PostData.Posts.ID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "can't get categories for this post" + err.Error(),
			}, http.StatusBadRequest)
			return
		}
		homeData.Category = category
		homeData.Datas = posts
		homeData.PostData.Route = "post?post_id=" + postid
		helper.SendResponse(w, homeData.PostData.Comment, http.StatusOK)

	}
}

func AddPostHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ok, errorPage := middlewares.CheckRequest(r, "/addpost", "post")
		if !ok {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "Method is not allowed",
			}, errorPage)
			return
		}
		session, err := helper.GetSessionRequest(r)
		if err != nil {
			return
		}

		if helper.VerifySession(db, session) {

			var newPost models.AddPostRequest

			err := json.NewDecoder(r.Body).Decode(&newPost)
			if err != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: "incorrect json format : " + err.Error(),
				}, http.StatusBadRequest)
				return
			}

			errForm := helper.CheckFormAddPost(newPost, db)
			if errForm != nil {
				homeData, err := helper.GetDataTemplate("", db, r, true, false, true, false, true)

				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}

				if homeData.Session {
					sessionID, _ := helper.GetSessionRequest(r)
					helper.UpdateCookieSession(w, sessionID, db)
				}
				homeData.Error = errForm.Error()
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: homeData.Error,
				}, http.StatusBadRequest)
				//helper.RenderTemplate(w, "index", "index", homeData)
				return
			}
			// Analyser le formulaire avec le champ de fichier "image_post"
			if r.Header.Get("Content-Type") == "multipart/form-data" {
				// Parse multipart form data
				err := r.ParseMultipartForm(20000)
				if err != nil {
					helper.SendResponse(w, models.ErrorResponse{
						Status:  "error",
						Message: "size of file is greater than 20KB",
					}, http.StatusBadRequest)
					return
				}
			}
			postTitle := newPost.Title
			postContent := newPost.Content
			_postCategorystring := newPost.Category
			// var _postCategoryuuid []uuid.UUID
			var _postCategories []models.Category
			// for _, v := range _postCategorystring {
			// 	catuuid, _ := uuid.FromString(v)
			// 	_postCategoryuuid = append(_postCategoryuuid, catuuid)
			// }
			for _, v := range _postCategorystring {
				var cat models.Category
				catuuid, _ := uuid.FromString(v)
				cat.ID = catuuid
				_postCategories = append(_postCategories, cat)
			}

			user, err := controller.GetUserBySessionId(session, db)
			if err != nil {
				controller.DeleteSession(db, session)
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: "this session is not valid",
				}, http.StatusBadRequest)
				//http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			dataUrl := newPost.Image
			if newPost.Image != "" {
				mimeType := strings.Split(dataUrl, ";")[0]
				mimeType = strings.TrimPrefix(mimeType, "data:")
				if mimeType != "image/jpeg" && mimeType != "image/png" {
					fmt.Println(mimeType)
					helper.SendResponse(w, models.ErrorResponse{
						Status:  "error",
						Message: "File format is not valid : " + err.Error(),
					}, http.StatusBadRequest)
					return
				}
			}

			base64Data := dataUrl[strings.IndexByte(dataUrl, ',')+1:]
			img, err := base64.StdEncoding.DecodeString(base64Data)
			if err != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: "invalid base64 string of image : " + err.Error(),
				}, http.StatusBadRequest)
				return
			}
			var post models.Post
			//image, img_header, err := r.FormFile("image_post")
			if newPost.Image == "" {
				post = models.Post{
					UserID:     user.ID,
					Title:      postTitle,
					Content:    postContent,
					Categories: _postCategories,
				}
			} else {

				imgSize := (float64(len(img)) / 1024.0) / 1024.0
				if imgSize > 20 {
					helper.SendResponse(w, models.ErrorResponse{
						Status:  "error",
						Message: "the size of image is bigger than 20ko",
					}, http.StatusBadRequest)
					return
				}
				name := helper.NewName()
				err = ioutil.WriteFile("./static/image/"+name+".png", img, 0644)
				if err != nil {
					helper.SendResponse(w, models.ErrorResponse{
						Status:  "error",
						Message: err.Error(),
					}, http.StatusBadRequest)
					return
				}

				post = models.Post{
					UserID:     user.ID,
					Title:      postTitle,
					Content:    postContent,
					Categories: _postCategories,
					Image:      name + ".png",
				}
			}

			_, err = controller.CreatePost(db, post)
			if err != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: "the server is unable to create this post",
				}, http.StatusInternalServerError)
				//helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			helper.SendResponse(w, post, http.StatusOK)
			//http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}
}

// Like post
func LikePoste(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		like := models.PostLike{}

		ok, errorPage := middlewares.CheckRequest(r, "/likepost", "post")
		if !ok {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "Method not allowed",
			}, errorPage)
			return
		}

		//check the session and get the user
		sessionID, errsess := helper.GetSessionRequest(r)
		if errsess != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: errsess.Error(),
			}, http.StatusBadRequest)
			return
		} else {

			session, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil || &session == nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: errgets.Error(),
				}, http.StatusBadRequest)
				return
			}
			User, errgetu := controller.GetUserBySessionId(sessionID, db)
			if errgetu != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: errgetu.Error(),
				}, http.StatusBadRequest)
				return
			}
			like.UserID = User.ID
		}
		var PostToLike models.OnePostRequest
		err := json.NewDecoder(r.Body).Decode(&PostToLike)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect json request",
			}, http.StatusBadRequest)
			return
		}
		postID, err := uuid.FromString(PostToLike.PostID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect post ID",
			}, http.StatusBadRequest)
			return
		}

		_, err = controller.GetPostByID(db, postID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "The post does not exist",
			}, http.StatusBadRequest)
			return
		}

		like.PostID = postID
		_, err = controller.CreatePostLike(db, like)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "the server can't create the like",
			}, http.StatusInternalServerError)
			return
		}
		//route := r.FormValue("route")
		//fmt.Println(route)
		helper.SendResponse(w, like, http.StatusOK)

	}
}

// dDislike posts
func DislikePoste(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dislike := models.PostDislike{}

		ok, errorPage := middlewares.CheckRequest(r, "/dislikepost", "post")
		if !ok {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "Method not allowed",
			}, errorPage)
			return
		}

		//check the session and get the user
		sessionID, errsess := helper.GetSessionRequest(r)
		if errsess != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: errsess.Error(),
			}, http.StatusBadRequest)
			return
		} else {

			session, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil || &session == nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: errgets.Error(),
				}, http.StatusBadRequest)
				return
			}
			User, errgetu := controller.GetUserBySessionId(sessionID, db)
			if errgetu != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: errgetu.Error(),
				}, http.StatusBadRequest)
				return
			}
			dislike.UserID = User.ID
		}

		var PostToLike models.OnePostRequest
		err := json.NewDecoder(r.Body).Decode(&PostToLike)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect json request",
			}, http.StatusBadRequest)
			return
		}
		postID, err := uuid.FromString(PostToLike.PostID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect post ID",
			}, http.StatusBadRequest)
			return
		}

		_, err = controller.GetPostByID(db, postID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "The post does not exist",
			}, http.StatusBadRequest)
			return
		}

		dislike.PostID = postID
		_, err = controller.CreatePostDislike(db, dislike)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "the server can't create the dislike",
			}, http.StatusInternalServerError)
			return
		}
		//route := r.FormValue("route")
		helper.SendResponse(w, dislike, http.StatusOK)
	}
}

// Like Comments
func LikeComment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		like := models.CommentLike{}

		ok, errorPage := middlewares.CheckRequest(r, "/likecomment", "post")
		if !ok {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "Method not allowed",
			}, errorPage)
			return
		}

		//check the session and get the user
		sessionID, errsess := helper.GetSessionRequest(r)
		if errsess != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: errsess.Error(),
			}, http.StatusBadRequest)
			return
		} else {

			session, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil || &session == nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: errgets.Error(),
				}, http.StatusBadRequest)
				return
			}
			User, errgetu := controller.GetUserBySessionId(sessionID, db)
			if errgetu != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: errgetu.Error(),
				}, http.StatusBadRequest)
				return
			}
			like.UserID = User.ID
		}
		var CommentTolike models.OneCommentRequest
		err := json.NewDecoder(r.Body).Decode(&CommentTolike)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect json request",
			}, http.StatusBadRequest)
			return
		}
		commentID, err := uuid.FromString(CommentTolike.CommentID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect comment ID",
			}, http.StatusBadRequest)
			return
		}

		_, err = controller.GetCommentByID(db, commentID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "The comment does not exist",
			}, http.StatusBadRequest)
			return
		}

		like.CommentID = commentID
		_, err = controller.CreateCommentLike(db, like)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "the server can't create the like",
			}, http.StatusInternalServerError)
			return
		}
		//route := r.FormValue("route")

		helper.SendResponse(w, like, http.StatusOK)

	}
}

// Dislike comments
func DislikeComment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dislike := models.CommentDislike{}

		ok, errorPage := middlewares.CheckRequest(r, "/dislikecomment", "post")
		if !ok {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "Method not allowed",
			}, errorPage)
			return
		}

		//check the session and get the user
		sessionID, errsess := helper.GetSessionRequest(r)
		if errsess != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: errsess.Error(),
			}, http.StatusBadRequest)
			return
		} else {

			session, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil || &session == nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: errgets.Error(),
				}, http.StatusBadRequest)
				return
			}
			User, errgetu := controller.GetUserBySessionId(sessionID, db)
			if errgetu != nil {
				helper.SendResponse(w, models.ErrorResponse{
					Status:  "error",
					Message: errgetu.Error(),
				}, http.StatusBadRequest)
				return
			}
			dislike.UserID = User.ID
		}

		var CommentTolike models.OneCommentRequest
		err := json.NewDecoder(r.Body).Decode(&CommentTolike)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect json request",
			}, http.StatusBadRequest)
			return
		}
		commentID, err := uuid.FromString(CommentTolike.CommentID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "incorrect comment ID",
			}, http.StatusBadRequest)
			return
		}

		_, err = controller.GetCommentByID(db, commentID)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "The comment does not exist",
			}, http.StatusBadRequest)
			return
		}

		dislike.CommentID = commentID
		_, err = controller.CreateCommentDislike(db, dislike)
		if err != nil {
			helper.SendResponse(w, models.ErrorResponse{
				Status:  "error",
				Message: "the server can't create the dislike",
			}, http.StatusInternalServerError)
			return
		}
		//	route := r.FormValue("route")

		helper.SendResponse(w, dislike, http.StatusOK)

	}
}
