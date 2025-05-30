package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

// these models are for the different post request
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type SessionRequest struct {
	Session string `json:"session"`
}
type LoginRequest struct {
	Email      string `json:"email"`
	Motdepasse string `json:"motdepasse"`
}

type RegisterRequest struct {
	UserName     string `json:"username"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Gender       string `json:"gender"`
	Age          string `json:"age"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Confpassword string `json:"confpassword"`
}

type AddPostRequest struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category []string `json:"category"`
	Image    string   `json:"image"`
}

type OnePostRequest struct {
	PostID string `json:"postid"`
}

type OneCommentRequest struct {
	CommentID string `json:"commentid"`
}

type AddCommentRequest struct {
	PostID  string `json:"postid"`
	UserID  string `json:"userid"`
	Content string `json:"content"`
}

type GetPostsByCategoryRequest struct {
	CategoryID string `json:"categoryid"`
}

// These models are for Database

type User struct {
	Conn      *websocket.Conn
	ID        uuid.UUID
	Username  string `json:"Username"`
	FirstName string
	LastName  string
	Gender    string
	Age       int
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ExpiresAt time.Time
	CreatedAt time.Time
}

type Category struct {
	ID           uuid.UUID
	NameCategory string
}

type Comment struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	PostID    uuid.UUID
	Content   string
	CreatedAt string
}

type PostLike struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	PostID    uuid.UUID
	CreatedAt time.Time
}

type PostDislike struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	PostID    uuid.UUID
	CreatedAt time.Time
}

type Home struct {
	Session     bool
	Category    []Category
	Datas       []HomeDataPost
	User        User
	ErrorAuth   ErrorAuth
	PostData    HomeDataPost
	DataProfil  DataMyProfil
	Error       string
	ErrorFilter string
}
type HomeDataPost struct {
	Posts        Post
	Comment      []CommentDetails
	CommentCount int
	PostLike     int
	PostDislike  int
	User         User
	Liked        bool
	Disliked     bool
	Route        string
}
type Post struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Title      string
	Content    string
	Image      string
	CategoryID []uuid.UUID
	Categories []Category
	CreatedAt  string
	//CreatedAt  time.Time
}
type PostCategory struct {
	CategoryID uuid.UUID
	PostID     uuid.UUID
}
type CommentDetails struct {
	Comment        Comment
	CommentLike    int
	CommentDislike int
	User           User
	Liked          bool
	Disliked       bool
	Route          string
}

type CommentLike struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CommentID uuid.UUID
	CreatedAt time.Time
}

type CommentDislike struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CommentID uuid.UUID
	CreatedAt time.Time
}

type ErrorAuth struct {
	EmailError    string
	UserNameError string
	PasswordError string
	GeneralError  string
}

type DataMypage struct {
	Session     bool
	Datas       []HomeDataPost
	User        User
	Category    []Category
	CategoryID  uuid.UUID
	Error       string
	ErrorFilter string
}

type DataMyProfil struct {
	User       User
	Categories map[string]int
	Posts      []HomeDataPost
}
