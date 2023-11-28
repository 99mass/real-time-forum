package handler

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"

	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
)

func GetOnePost(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			ok, pageError := middlewares.CheckRequest(r, "/post", "get")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}
			postid := r.FormValue("post_id")
			homeData, err := helper.GetDataTemplate(db, r, true, true, false, false, false)
			if err != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
				return
			}
			posts, err := helper.GetPostsForOneUser(db, homeData.PostData.User.ID)
			if err != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
				return
			}
			postsliked, err := helper.SetLikesAndDislikes(homeData.User, posts, db)
			if err != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
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
				helper.ErrorPage(w, http.StatusBadRequest)
				return
			}
			homeData.Category = category
			homeData.Datas = posts
			homeData.PostData.Route = "post?post_id=" + postid
			helper.RenderTemplate(w, "post", "posts", homeData)
		case http.MethodPost:
			ok, pageError := middlewares.CheckRequest(r, "/post", "post")
			if !ok {
				helper.ErrorPage(w, pageError)
				return
			}
			var comment models.Comment

			postID, errP := helper.StringToUuid(r, "post_id")
			userID, errU := helper.StringToUuid(r, "user_id")
			Content := r.FormValue("content")
			Content = strings.TrimSpace(Content)

			if errP != nil || errU != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
				return
			}
			homeDataSess, err := helper.GetDataTemplate(db, r, true, false, false, false, false)
			if err != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
			}
			if !homeDataSess.Session {
				homeData, err := helper.GetDataTemplate(db, r, true, true, false, false, false)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
				}
				posts, err := helper.GetPostsForOneUser(db, homeData.PostData.User.ID)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
					return
				}
				postsliked, err := helper.SetLikesAndDislikes(homeData.User, posts, db)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
				}

				posts = postsliked
				for i := range posts {
					posts[i].Route = "post"
					for j := range posts[i].Comment {
						posts[i].Comment[j].Route = "post"
					}
				}
				category, err := controller.GetCategoriesByPost(db, homeData.PostData.Posts.ID)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
					return
				}

				homeData.Category = category
				homeData.Datas = posts
				helper.RenderTemplate(w, "post", "posts", homeData)
				return
			}
			if Content == "" {
				homeData, err := helper.GetDataTemplate(db, r, true, true, false, false, false)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
				}
				posts, err := helper.GetPostsForOneUser(db, homeData.PostData.User.ID)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
					return
				}
				//Set likes and dislikes
				postsliked, err := helper.SetLikesAndDislikes(homeData.User, posts, db)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
				}

				posts = postsliked

				category, err := controller.GetCategoriesByPost(db, homeData.PostData.Posts.ID)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
					return
				}
				homeData.Category = category
				homeData.Datas = posts
				homeData.Error = "comments cannot be empty"
				helper.RenderTemplate(w, "post", "posts", homeData)
				return
			}
			if Content == "@$" {
				homeData, err := helper.GetDataTemplate(db, r, true, true, false, false, false)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
				}
				posts, err := helper.GetPostsForOneUser(db, homeData.PostData.User.ID)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
					return
				}
				//Set likes and dislikes
				postsliked, err := helper.SetLikesAndDislikes(homeData.User, posts, db)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
				}

				posts = postsliked

				category, err := controller.GetCategoriesByPost(db, homeData.PostData.Posts.ID)
				if err != nil {
					helper.ErrorPage(w, http.StatusBadRequest)
					return
				}
				homeData.Category = category
				homeData.Datas = posts
				homeData.Error = "the number of characters must not exceed 1000"
				helper.RenderTemplate(w, "post", "posts", homeData)
				return
			}
			comment.PostID = postID
			comment.UserID = userID
			comment.Content = Content

			_, erro := controller.CreateComment(db, comment)
			if erro != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
				return
			}
			homeData, err := helper.GetDataTemplate(db, r, true, true, true, false, false)
			if err != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
			}
			posts, err := helper.GetPostsForOneUser(db, homeData.User.ID)
			if err != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
				return
			}
			//Set likes and dislikes
			postsliked, err := helper.SetLikesAndDislikes(homeData.User, posts, db)
			if err != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
			}

			posts = postsliked

			category, err := controller.GetCategoriesByPost(db, homeData.PostData.Posts.ID)
			if err != nil {
				helper.ErrorPage(w, http.StatusBadRequest)
				return
			}
			homeData.Category = category
			homeData.Datas = posts
			helper.RenderTemplate(w, "post", "posts", homeData)
		default:
			helper.ErrorPage(w, http.StatusMethodNotAllowed)
			return
		}

	}
}

func AddPostHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ok, errorPage := middlewares.CheckRequest(r, "/addpost", "post")
		if !ok {
			helper.ErrorPage(w, errorPage)
			return
		}
		session, err := helper.GetSessionRequest(r)
		if err != nil {
			return
		}

		if helper.VerifySession(db, session) {

			errForm := helper.CheckFormAddPost(r, db)
			if errForm != nil {
				homeData, err := helper.GetDataTemplate(db, r, true, false, true, false, true)

				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}

				if homeData.Session {
					sessionID, _ := helper.GetSessionRequest(r)
					helper.UpdateCookieSession(w, sessionID, db)
				}
				homeData.Error = errForm.Error()

				helper.RenderTemplate(w, "index", "index", homeData)
				return
			}
			// Analyser le formulaire avec le champ de fichier "image_post"
			err := r.ParseMultipartForm(20000) // Limite de taille du fichier de 10 Mo
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			postTitle := r.FormValue("title")
			postContent := r.FormValue("content")
			_postCategorystring := r.Form["category"]
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
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			var post models.Post
			image, img_header, err := r.FormFile("image_post")
			if err != nil {
				post = models.Post{
					UserID:     user.ID,
					Title:      postTitle,
					Content:    postContent,
					Categories: _postCategories,
				}
			} else {
				if !helper.VerifImage(img_header.Filename) {
					homeData, err := helper.GetDataTemplate(db, r, true, false, true, false, true)

					if err != nil {
						helper.ErrorPage(w, http.StatusInternalServerError)
						return
					}

					if homeData.Session {
						sessionID, _ := helper.GetSessionRequest(r)
						helper.UpdateCookieSession(w, sessionID, db)
					}
					homeData.Error = "this type of file is not allowed"

					helper.RenderTemplate(w, "index", "index", homeData)
					return
				}
				if img_header.Size > 20971520 {
					homeData, err := helper.GetDataTemplate(db, r, true, false, true, false, true)

					if err != nil {
						helper.ErrorPage(w, http.StatusInternalServerError)
						return
					}

					if homeData.Session {
						sessionID, _ := helper.GetSessionRequest(r)
						helper.UpdateCookieSession(w, sessionID, db)
					}
					homeData.Error = "the file size can't be over than 20Mo"

					helper.RenderTemplate(w, "index", "index", homeData)
					return
				}
				imageData, err := ioutil.ReadAll(image)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}
				name := helper.NewName()
				err = ioutil.WriteFile("./static/image/"+name+img_header.Filename, imageData, 0644)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}

				post = models.Post{
					UserID:     user.ID,
					Title:      postTitle,
					Content:    postContent,
					Categories: _postCategories,
					Image:      name + img_header.Filename,
				}
			}

			_, err = controller.CreatePost(db, post)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}
}

func AddPostHandlerForMyPage(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ok, errorPage := middlewares.CheckRequest(r, "/addpostmypage", "post")
		if !ok {
			helper.ErrorPage(w, errorPage)
			return
		}
		session, err := helper.GetSessionRequest(r)
		if err != nil {
			return
		}

		if helper.VerifySession(db, session) {
			sessiondata := true

			errForm := helper.CheckFormAddPost(r, db)
			if errForm != nil {
				user, err := controller.GetUserBySessionId(session, db)
				if err != nil {
					controller.DeleteSession(db, session)
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
						return
					}

					datas := new(models.DataMypage)
					datas.Datas = PostsDetails
					datas.Session = sessiondata
					datas.User = user
					datas.CategoryID = CategID
					datas.Category = category
					datas.Error = errForm.Error()
					helper.RenderTemplate(w, "mypage", "mypages", datas)
				} else {
					PostsDetails, err := helper.GetPostsForOneUser(db, user.ID)
					if err != nil {
						helper.ErrorPage(w, http.StatusInternalServerError)
						return
					}
					datas := new(models.DataMypage)
					datas.Datas = PostsDetails

					datas.Session = sessiondata
					datas.User = user
					datas.Category = category
					datas.Error = errForm.Error()
					helper.RenderTemplate(w, "mypage", "mypages", datas)
				}
				return
			}
			postTitle := r.FormValue("title")
			postContent := r.FormValue("content")
			_postCategorystring := r.Form["category"]
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
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			var post models.Post
			image, img_header, err := r.FormFile("image_post")
			if err != nil {
				post = models.Post{
					UserID:     user.ID,
					Title:      postTitle,
					Content:    postContent,
					Categories: _postCategories,
				}
			} else {
				if !helper.VerifImage(img_header.Filename) {
					PostsDetails, err := helper.GetPostsForOneUser(db, user.ID)
					if err != nil {
						helper.ErrorPage(w, http.StatusInternalServerError)
						return
					}
					category, err := controller.GetAllCategories(db)
					if err != nil {
						helper.ErrorPage(w, http.StatusInternalServerError)
						return
					}
					datas := new(models.DataMypage)
					datas.Datas = PostsDetails

					datas.Session = sessiondata
					datas.User = user
					datas.Category = category
					datas.Error = "this type of file is not allowed"
					helper.RenderTemplate(w, "mypage", "mypages", datas)
					return
				}
				if img_header.Size > 20971520 {
					PostsDetails, err := helper.GetPostsForOneUser(db, user.ID)
					if err != nil {
						helper.ErrorPage(w, http.StatusInternalServerError)
						return
					}
					category, err := controller.GetAllCategories(db)
					if err != nil {
						helper.ErrorPage(w, http.StatusInternalServerError)
						return
					}
					datas := new(models.DataMypage)
					datas.Datas = PostsDetails

					datas.Session = sessiondata
					datas.User = user
					datas.Category = category
					datas.Error = "the size of the file can't be over than 20Mo"
					helper.RenderTemplate(w, "mypage", "mypages", datas)
					return
				}
				imageData, err := ioutil.ReadAll(image)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}
				name := helper.NewName()
				err = ioutil.WriteFile("./static/image/"+name+img_header.Filename, imageData, 0644)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}

				post = models.Post{
					UserID:     user.ID,
					Title:      postTitle,
					Content:    postContent,
					Categories: _postCategories,
					Image:      name + img_header.Filename,
				}
			}
			_, err = controller.CreatePost(db, post)
			if err != nil {
				return
			}
			http.Redirect(w, r, "/mypage", http.StatusSeeOther)
		}

	}
}

// Like post
func LikePoste(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		like := models.PostLike{}

		ok, errorPage := middlewares.CheckRequest(r, "/likepost", "post")
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

		postID, err := helper.StringToUuid(r, "post_id")
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}

		_, err = controller.GetPostByID(db, postID)
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}

		like.PostID = postID
		_, err = controller.CreatePostLike(db, like)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		route := r.FormValue("route")
		//fmt.Println(route)
		http.Redirect(w, r, "/"+route, http.StatusSeeOther)

	}
}

// dDislike posts
func DislikePoste(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dislike := models.PostDislike{}

		ok, errorPage := middlewares.CheckRequest(r, "/dislikepost", "post")
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
			dislike.UserID = User.ID
		}

		postID, err := helper.StringToUuid(r, "post_id")
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}

		_, err = controller.GetPostByID(db, postID)
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}

		dislike.PostID = postID
		_, err = controller.CreatePostDislike(db, dislike)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		route := r.FormValue("route")
		http.Redirect(w, r, "/"+route, http.StatusSeeOther)

	}
}

// Like Comments
func LikeComment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		like := models.CommentLike{}

		ok, errorPage := middlewares.CheckRequest(r, "/likecomment", "post")
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

		commentID, err := helper.StringToUuid(r, "comment_id")
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}

		_, err = controller.GetCommentByID(db, commentID)
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}

		like.CommentID = commentID
		_, err = controller.CreateCommentLike(db, like)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		route := r.FormValue("route")

		http.Redirect(w, r, "/"+route, http.StatusSeeOther)

	}
}

// Dislike comments
func DislikeComment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dislike := models.CommentDislike{}

		ok, errorPage := middlewares.CheckRequest(r, "/dislikecomment", "post")
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
			dislike.UserID = User.ID
		}

		commentID, err := helper.StringToUuid(r, "comment_id")

		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}

		_, err = controller.GetCommentByID(db, commentID)
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}

		dislike.CommentID = commentID
		_, err = controller.CreateCommentDislike(db, dislike)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		route := r.FormValue("route")

		http.Redirect(w, r, "/"+route, http.StatusSeeOther)

	}
}
