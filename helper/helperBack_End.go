package helper

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"

	"forum/controller"
	"forum/models"
)

func CorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next(w, r)
    }
}

func SendResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func Debug(str string) {
	fmt.Println(str)
}

func GetPostForHome(db *sql.DB) ([]models.HomeDataPost, error) {
	post, err := controller.GetAllPosts(db)
	if err != nil {
		return nil, err
	}
	var HomeDatas []models.HomeDataPost
	for _, post := range post {
		var HomeData models.HomeDataPost
		comments, err := controller.GetCommentsByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		user, err := controller.GetUserByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}

		var commentdetails []models.CommentDetails
		for _, com := range comments {
			user, err := controller.GetUserByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			var commentdetail models.CommentDetails
			commentdetail.Comment = com
			commentlike, err := controller.GetCommentLikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdislike, err := controller.GetCommentDislikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdetail.CommentLike = len(commentlike)
			commentdetail.CommentDislike = len(commentdislike)
			commentdetail.User = *user
			commentdetails = append(commentdetails, commentdetail)

		}
		likes, err := controller.GetPostLikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrlikes := len(likes)
		dislike, err := controller.GetDislikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrdislikes := len(dislike)

		category, err := controller.GetCategoriesByPost(db, post.ID)
		if err != nil {
			return nil, err
		}

		post.Categories = category
		HomeData.Posts = post
		HomeData.Comment = commentdetails
		HomeData.PostLike = nbrlikes
		HomeData.PostDislike = nbrdislikes
		HomeData.User = *user

		HomeDatas = append(HomeDatas, HomeData)
	}
	return HomeDatas, nil
}

func GetPostDetails(db *sql.DB, postID uuid.UUID) (models.HomeDataPost, error) {
	post, err := controller.GetPostByID(db, postID)
	if err != nil {
		return models.HomeDataPost{}, err
	}

	var HomeData models.HomeDataPost
	comments, err := controller.GetCommentsByPostID(db, post.ID)
	if err != nil {
		return models.HomeDataPost{}, err
	}
	var commentdetails []models.CommentDetails
	for _, com := range comments {
		user, err := controller.GetUserByCommentID(db, com.ID)
		if err != nil {
			return models.HomeDataPost{}, err
		}
		var commentdetail models.CommentDetails
		commentdetail.Comment = com
		commentlike, err := controller.GetCommentLikesByCommentID(db, com.ID)
		if err != nil {
			return models.HomeDataPost{}, err
		}
		commentdislike, err := controller.GetCommentDislikesByCommentID(db, com.ID)
		if err != nil {
			return models.HomeDataPost{}, err
		}
		commentdetail.CommentLike = len(commentlike)
		commentdetail.CommentDislike = len(commentdislike)
		commentdetail.User = *user
		commentdetails = append(commentdetails, commentdetail)
	}
	likes, err := controller.GetPostLikesByPostID(db, post.ID)
	if err != nil {

		return models.HomeDataPost{}, err
	}
	nbrlikes := len(likes)
	dislike, err := controller.GetDislikesByPostID(db, post.ID)
	if err != nil {
		return models.HomeDataPost{}, err
	}
	nbrdislikes := len(dislike)

	category, err := controller.GetCategoriesByPost(db, postID)
	if err != nil {
		return models.HomeDataPost{}, err
	}
	user, err := controller.GetUserByPostID(db, post.ID)
	if err != nil {
		return models.HomeDataPost{}, err
	}

	HomeData.Posts.Categories = category
	HomeData.Posts = post
	HomeData.Comment = commentdetails
	HomeData.PostLike = nbrlikes
	HomeData.PostDislike = nbrdislikes
	HomeData.User = *user

	return HomeData, nil
}

func VerifUser(db *sql.DB, email string, password string) (uuid.UUID, bool) {
	client := new(models.User)
	okEmail, _ := CheckEmail(email)

	if okEmail {
		user, err := controller.GetUserByEmail(db, email)
		if err != nil {
			return uuid.Nil, false
		}
		if user == nil {
			return uuid.Nil, false
		}
		client = user
	} else {
		user, err := controller.GetUserByUsername(db, email)
		if err != nil {
			return uuid.Nil, false
		}
		if user == nil {
			return uuid.Nil, false
		}
		client = user
	}

	if !CheckPasswordHash(password, client.Password) {
		return client.ID, false
	}
	return client.ID, true
}

// CheckPasswordHash compares a password with its hashed version
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckFormAddPost(post models.AddPostRequest, db *sql.DB) error {
	postTitle := post.Title
	postTitle = strings.TrimSpace(postTitle)
	postContent := post.Content
	postContent = strings.TrimSpace(postContent)
	// content and title empty
	if postTitle == "" || postContent == "" {
		return errors.New("all fields must be completed")
	}
	if len(postTitle) > 50 {
		return errors.New("the length of the title is too long")
	}
	if len(postContent) > 1000 {
		return errors.New("the length of the post is too long")
	}
	if postContent == "@$" {
		return errors.New("the number of characters must not exceed 1000")
	}
	_postCategorystring := post.Category
	// No category received
	if len(_postCategorystring) == 0 {
		return errors.New("no category selected")
	}
	// Category not matched
	for _, v := range _postCategorystring {
		catuuid, err := uuid.FromString(v)
		if !VerifCategory(db, catuuid) || err != nil || catuuid == uuid.Nil {
			return errors.New("one of the categories is not compliant")
		}
	}

	return nil
}

func VerifCategory(db *sql.DB, v uuid.UUID) bool {
	_, err := controller.GetCategoryByID(db, v)
	if err != nil {
		return false
	}
	return true
}

func StringToUuid(r *http.Request, s string) (uuid.UUID, error) {
	chaine := strings.TrimSpace(r.FormValue(s))
	result, err := uuid.FromString(chaine)
	if err != nil {
		return uuid.Nil, err
	}
	return result, nil
}

func GetPostsForOneUser(db *sql.DB, userID uuid.UUID) ([]models.HomeDataPost, error) {

	post, err := controller.GetPostsByUserID(db, userID)
	if err != nil {
		return nil, err
	}
	var HomeDatas []models.HomeDataPost
	for _, post := range post {
		var HomeData models.HomeDataPost
		comments, err := controller.GetCommentsByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		user, err := controller.GetUserByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		var commentdetails []models.CommentDetails
		for _, com := range comments {
			user, err := controller.GetUserByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			var commentdetail models.CommentDetails
			commentdetail.Comment = com
			commentlike, err := controller.GetCommentLikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdislike, err := controller.GetCommentDislikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdetail.CommentLike = len(commentlike)
			commentdetail.CommentDislike = len(commentdislike)
			commentdetail.User = *user
			commentdetails = append(commentdetails, commentdetail)

		}
		likes, err := controller.GetPostLikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrlikes := len(likes)
		dislike, err := controller.GetDislikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrdislikes := len(dislike)

		category, err := controller.GetCategoriesByPost(db, post.ID)
		if err != nil {
			return nil, err
		}

		post.Categories = category
		HomeData.Posts = post
		HomeData.Comment = commentdetails
		HomeData.PostLike = nbrlikes
		HomeData.PostDislike = nbrdislikes
		HomeData.User = *user

		HomeDatas = append(HomeDatas, HomeData)
	}
	return HomeDatas, nil
}

func GetPostForCategory(db *sql.DB, catID uuid.UUID) ([]models.HomeDataPost, error) {
	post, err := controller.GetPostsByCategory(db, catID)
	if err != nil {
		return nil, err
	}
	var HomeDatas []models.HomeDataPost
	for _, post := range post {
		var HomeData models.HomeDataPost
		comments, err := controller.GetCommentsByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		user, err := controller.GetUserByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		var commentdetails []models.CommentDetails
		for _, com := range comments {
			user, err := controller.GetUserByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			var commentdetail models.CommentDetails
			commentdetail.Comment = com
			commentlike, err := controller.GetCommentLikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdislike, err := controller.GetCommentDislikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdetail.CommentLike = len(commentlike)
			commentdetail.CommentDislike = len(commentdislike)
			commentdetail.User = *user
			commentdetails = append(commentdetails, commentdetail)

		}
		likes, err := controller.GetPostLikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrlikes := len(likes)
		dislike, err := controller.GetDislikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrdislikes := len(dislike)

		category, err := controller.GetCategoriesByPost(db, post.ID)
		if err != nil {
			return nil, err
		}

		post.Categories = category
		HomeData.Posts = post
		HomeData.Comment = commentdetails
		HomeData.PostLike = nbrlikes
		HomeData.PostDislike = nbrdislikes
		HomeData.User = *user

		HomeDatas = append(HomeDatas, HomeData)
	}
	return HomeDatas, nil
}

func GetPostsForOneUserAndCategory(db *sql.DB, userID, catID uuid.UUID) ([]models.HomeDataPost, error) {

	post, err := controller.GetPostsByUserAndCategory(db, userID, catID)
	if err != nil {
		return nil, err
	}
	var HomeDatas []models.HomeDataPost
	for _, post := range post {
		var HomeData models.HomeDataPost
		comments, err := controller.GetCommentsByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		user, err := controller.GetUserByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		var commentdetails []models.CommentDetails
		for _, com := range comments {
			user, err := controller.GetUserByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			var commentdetail models.CommentDetails
			commentdetail.Comment = com
			commentlike, err := controller.GetCommentLikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdislike, err := controller.GetCommentDislikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdetail.CommentLike = len(commentlike)
			commentdetail.CommentDislike = len(commentdislike)
			commentdetail.User = *user
			commentdetails = append(commentdetails, commentdetail)

		}
		likes, err := controller.GetPostLikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrlikes := len(likes)
		dislike, err := controller.GetDislikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrdislikes := len(dislike)

		category, err := controller.GetCategoriesByPost(db, post.ID)
		if err != nil {
			return nil, err
		}

		post.Categories = category
		HomeData.Posts = post
		HomeData.Comment = commentdetails
		HomeData.PostLike = nbrlikes
		HomeData.PostDislike = nbrdislikes
		HomeData.User = *user

		HomeDatas = append(HomeDatas, HomeData)
	}
	return HomeDatas, nil
}

func GetDetailPost(db *sql.DB, post models.Post) (models.HomeDataPost, error) {

	var HomeData models.HomeDataPost
	comments, err := controller.GetCommentsByPostID(db, post.ID)
	if err != nil {
		return models.HomeDataPost{}, err
	}
	var commentdetails []models.CommentDetails
	for _, com := range comments {
		user, err := controller.GetUserByCommentID(db, com.ID)
		if err != nil {
			return models.HomeDataPost{}, err
		}
		var commentdetail models.CommentDetails
		commentdetail.Comment = com
		commentlike, err := controller.GetCommentLikesByCommentID(db, com.ID)
		if err != nil {
			return models.HomeDataPost{}, err
		}
		commentdislike, err := controller.GetCommentDislikesByCommentID(db, com.ID)
		if err != nil {
			return models.HomeDataPost{}, err
		}
		commentdetail.CommentLike = len(commentlike)
		commentdetail.CommentDislike = len(commentdislike)
		commentdetail.User = *user
		commentdetails = append(commentdetails, commentdetail)
	}
	likes, err := controller.GetPostLikesByPostID(db, post.ID)
	if err != nil {

		return models.HomeDataPost{}, err
	}
	nbrlikes := len(likes)
	dislike, err := controller.GetDislikesByPostID(db, post.ID)
	if err != nil {
		return models.HomeDataPost{}, err
	}
	nbrdislikes := len(dislike)

	category, err := controller.GetCategoriesByPost(db, post.ID)
	if err != nil {
		return models.HomeDataPost{}, err
	}
	user, err := controller.GetUserByPostID(db, post.ID)
	if err != nil {
		return models.HomeDataPost{}, err
	}

	HomeData.Posts.Categories = category
	HomeData.Posts = post
	HomeData.Comment = commentdetails
	HomeData.PostLike = nbrlikes
	HomeData.PostDislike = nbrdislikes
	HomeData.User = *user

	return HomeData, nil
}

func SetLikesAndDislikes(User models.User, datas []models.HomeDataPost, db *sql.DB) ([]models.HomeDataPost, error) {
	//Get if liked
	dataliked := []models.HomeDataPost{}

	for _, post := range datas {
		liked, err := IsPostliked(db, User.ID, post.Posts.ID)
		if err != nil {
			return datas, err
		}
		//Get if disliked
		disliked, errdis := IsPostDisliked(db, User.ID, post.Posts.ID)
		if errdis != nil {
			return datas, errdis
		}
		//fmt.Println(liked)
		post.Liked = liked
		post.Disliked = disliked
		dataliked = append(dataliked, post)

		for i, comment := range post.Comment {
			liked, err := IsCommentliked(db, User.ID, comment.Comment.ID)
			if err != nil {
				return datas, err
			}

			disliked, err := IsCommentDisliked(db, User.ID, comment.Comment.ID)
			if err != nil {
				return datas, err
			}

			post.Comment[i].Liked = liked
			post.Comment[i].Disliked = disliked
			//fmt.Println("like:", post.Comment[i].Liked, "dislike:", post.Comment[i].Disliked,post.Comment[i].Comment.Content)

			// comment.Liked = liked
			// comment.Disliked = disliked
		}

	}

	return dataliked, nil
}

func GetPostForFilter(db *sql.DB, post []models.Post) ([]models.HomeDataPost, error) {
	var HomeDatas []models.HomeDataPost
	for _, post := range post {
		var HomeData models.HomeDataPost
		comments, err := controller.GetCommentsByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		user, err := controller.GetUserByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}

		var commentdetails []models.CommentDetails
		for _, com := range comments {
			user, err := controller.GetUserByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			var commentdetail models.CommentDetails
			commentdetail.Comment = com
			commentlike, err := controller.GetCommentLikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdislike, err := controller.GetCommentDislikesByCommentID(db, com.ID)
			if err != nil {
				return nil, err
			}
			commentdetail.CommentLike = len(commentlike)
			commentdetail.CommentDislike = len(commentdislike)
			commentdetail.User = *user
			commentdetails = append(commentdetails, commentdetail)

		}
		likes, err := controller.GetPostLikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrlikes := len(likes)
		dislike, err := controller.GetDislikesByPostID(db, post.ID)
		if err != nil {
			return nil, err
		}
		nbrdislikes := len(dislike)

		category, err := controller.GetCategoriesByPost(db, post.ID)
		if err != nil {
			return nil, err
		}

		post.Categories = category
		HomeData.Posts = post
		HomeData.Comment = commentdetails
		HomeData.PostLike = nbrlikes
		HomeData.PostDislike = nbrdislikes
		HomeData.User = *user

		HomeDatas = append(HomeDatas, HomeData)
	}
	return HomeDatas, nil
}

func NewName() string {
	newUUID, _ := uuid.NewV4()

	return newUUID.String()
}

func VerifImage(str string) bool {
	str = strings.ToLower(str)
	if strings.HasSuffix(str, ".png") || strings.HasSuffix(str, ".jpeg") || strings.HasSuffix(str, ".gif") {
		return true
	}
	return false
}
