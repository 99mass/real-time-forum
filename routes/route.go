package routes

import (
	"database/sql"
	"net/http"

	"forum/handler"
	"forum/helper"
)

func Route(db *sql.DB) {
	server := &handler.Server{
		Clients: make(map[string]*handler.Client),
	}
	http.HandleFunc("/ws", helper.CorsMiddleware(server.HandleConnections))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	http.HandleFunc("/verifySession", helper.CorsMiddleware(handler.Index(db)))
	http.HandleFunc("/signin", helper.CorsMiddleware(handler.SinginHandler(db)))
	http.HandleFunc("/register", helper.CorsMiddleware(handler.RegisterHandler(db)))
	http.HandleFunc("/mypage", handler.GetMypage(db))
	http.HandleFunc("/post", helper.CorsMiddleware(handler.GetOnePost(db)))
	http.HandleFunc("/addcomment", helper.CorsMiddleware(handler.AddComment(db)))
	http.HandleFunc("/profil", handler.GetProfil(db))
	http.HandleFunc("/signout", helper.CorsMiddleware(handler.SignOutHandler(db)))
	http.HandleFunc("/addpost", helper.CorsMiddleware(handler.AddPostHandler(db)))
	//http.HandleFunc("/addpostmypage", handler.AddPostHandlerForMyPage(db))
	http.HandleFunc("/category", helper.CorsMiddleware(handler.GetPostCategory(db)))
	http.HandleFunc("/likepost", handler.LikePoste(db))
	http.HandleFunc("/dislikepost", handler.DislikePoste(db))
	http.HandleFunc("/likecomment", handler.LikeComment(db))
	http.HandleFunc("/dislikecomment", handler.DislikeComment(db))
	http.HandleFunc("/search", handler.Search(db))
	http.HandleFunc("/filter", handler.Filter(db))
	http.HandleFunc("/filtermypage", handler.FilterMyPage(db))
}
