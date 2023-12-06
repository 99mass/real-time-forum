package helper

import (
	"database/sql"
	"encoding/json"
	"forum/controller"
	"forum/models"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func CreateDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		return nil, err
	}
	DB = db
	// defer db.Close()
	return db, nil
}

func CreateTables(db *sql.DB) error {
	schema, err := ioutil.ReadFile("./database/structure.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}

func Comment(content string, id int) {

}

// The function give all data the the template page needs
func GetDataTemplate(postID string, db *sql.DB, r *http.Request, User, Post, Posts, ErrAuth, Category bool) (models.Home, error) {
	var datas models.Home

	//---Get All Posts-------//
	if Posts {
		posts, err := GetPostForHome(db)
		if err != nil {
			//ErrorPage(w, http.StatusInternalServerError)
			return datas, err
		}

		datas.Datas = posts
	}

	//---Get One Post-------//
	if Post {

		ID, err := uuid.FromString(postID)
		if err != nil {
			//ErrorPage(w, http.StatusInternalServerError)
			return datas, err
		}
		postData, errPD := GetPostDetails(db, ID)
		if errPD != nil {
			//ErrorPage(w, http.StatusInternalServerError)
			return datas, errPD
		}
		postid := postID
		for i := range postData.Comment {
			postData.Comment[i].Route = "post?post_id=" + postid
		}

		for i := range postData.Comment {
			postData.Comment[i].Route = "post?post_id=" + postid
		}

		datas.PostData = postData

	}

	//---Get the User-------//
	if User {
		var sessiondata bool
		sessionID, errsess := GetSessionRequest(r)
		if errsess != nil {
			sessiondata = false
		} else {
			sessiondata = true
			session, errgets := controller.GetSessionByID(db, sessionID)
			if errgets != nil || &session == nil {
				sessiondata = false
			}
			datas.User, errgets = controller.GetUserBySessionId(sessionID, db)
			if errgets != nil {
				sessiondata = false
			}
		}
		datas.Session = sessiondata
	}
	//Set likes and dislikes for One Poste
	DataslikedONe, err := SetLikesAndDislikes(datas.User, []models.HomeDataPost{datas.PostData}, db)
	if err != nil {
		return datas, err
	}
	datas.PostData = DataslikedONe[0]
	// fmt.Println(len(DataslikedOne))

	//Set likes and dislikes
	Datasliked, err := SetLikesAndDislikes(datas.User, datas.Datas, db)
	if err != nil {
		return datas, err
	}

	datas.Datas = Datasliked

	//---Get All Categories-------//
	if Category {
		category, err := controller.GetAllCategories(db)
		if err != nil {
			return datas, err
			//ErrorPage(w, http.StatusInternalServerError)
		}
		datas.Category = category
	}

	//---Get Error autification---//
	if ErrAuth {
		var loginReq models.LoginRequest

		err := json.NewDecoder(r.Body).Decode(&loginReq)
		if err != nil {
			datas.ErrorAuth.GeneralError = "incorrect request "
			return datas, nil
		}
		email := loginReq.Email
		email = strings.TrimSpace(email)
		password := loginReq.Motdepasse
		password = strings.TrimSpace(password)

		userID, toConnect := VerifUser(db, email, password)

		if toConnect {
			datas.User.ID = userID
			return datas, nil
		} else {
			datas.ErrorAuth.GeneralError = "Incorrect email address or password"
			return datas, nil
			//RenderTemplate(w, "signin", "auth", datas)
		}
	}

	return datas, nil
}

func IsPostliked(db *sql.DB, UserId, PostId uuid.UUID) (bool, error) {
	var like models.PostLike
	query := `
        SELECT id,user_id, post_id
        FROM post_likes
        WHERE user_id = ? AND post_id = ?
        LIMIT 1;
    `
	err := db.QueryRow(query, UserId, PostId).Scan(&like.UserID, &like.PostID)
	if err == sql.ErrNoRows {

		return false, nil
	}
	return true, nil
}

func IsPostDisliked(db *sql.DB, UserId, PostId uuid.UUID) (bool, error) {
	var dislike models.PostDislike
	query := `
        SELECT id,user_id, post_id
        FROM post_dislikes
        WHERE user_id = ? AND post_id = ?
        LIMIT 1;
    `
	err := db.QueryRow(query, UserId, PostId).Scan(&dislike.UserID, &dislike.PostID)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}

func IsCommentliked(db *sql.DB, UserId, CommentId uuid.UUID) (bool, error) {
	var like models.CommentLike
	query := `
        SELECT id,user_id, comment_id
        FROM comment_likes
        WHERE user_id = ? AND comment_id = ?
        LIMIT 1;
    `
	err := db.QueryRow(query, UserId, CommentId).Scan(&like.UserID, &like.CommentID)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}

func IsCommentDisliked(db *sql.DB, UserId, CommentId uuid.UUID) (bool, error) {
	var like models.CommentDislike
	query := `
        SELECT id,user_id, comment_id
        FROM comment_dislikes
        WHERE user_id = ? AND comment_id = ?
        LIMIT 1;
    `
	err := db.QueryRow(query, UserId, CommentId).Scan(&like.UserID, &like.CommentID)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}
