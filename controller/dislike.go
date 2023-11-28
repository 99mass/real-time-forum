package controller

import (
	"database/sql"
	"errors"
	"time"

	"github.com/gofrs/uuid"

	"forum/models"
)


func GetPostDislikesCount(db *sql.DB, postID uuid.UUID) (int, error) {
	query := `
    SELECT COUNT(*) FROM comment_dislikes WHERE post_id = ?;
    `

	var count int
	err := db.QueryRow(query, postID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetCommentDislikesCount(db *sql.DB, PostID uuid.UUID, commentID uuid.UUID) (int, error) {
	query := `
    SELECT COUNT(*) FROM comment_dislikes WHERE comment_id = ? AND post_id = ?;
    `

	var count int
	err := db.QueryRow(query, PostID, commentID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func CreatePostDislike(db *sql.DB, dislike models.PostDislike) (uuid.UUID, error) {
	_, errdislike := GetPostDislikeByUserID(db, dislike)
	if errdislike == nil {
		return uuid.UUID{}, errdislike
	}

	like := models.PostLike{
		UserID: dislike.UserID,
		PostID: dislike.PostID,
	}
	lik, errlike := GetPostLikeByUserID(db, like)
	if errlike == nil {
		RemovePostLike(db, lik.ID)
	}

	query := `
        INSERT INTO post_dislikes (id, user_id, post_id, created_at)
        VALUES (?, ?, ?, ?);
    `

	newUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = db.Exec(query, newUUID.String(), dislike.UserID, dislike.PostID, time.Now())
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUUID, nil
}

func CreateCommentDislike(db *sql.DB, dislike models.CommentDislike) (uuid.UUID, error) {
	_, errdislike := GetCommentDislikeByUserID(db, dislike)
	if errdislike == nil {
		return uuid.UUID{}, errdislike
	}

	like := models.CommentLike{
		UserID:    dislike.UserID,
		CommentID: dislike.CommentID,
	}
	lik, errlike := GetCommentLikeByUserID(db, like)
	if errlike == nil {
		RemoveCommentLike(db, lik.ID)
	}
	
	query := `
        INSERT INTO comment_dislikes (id, user_id, comment_id, created_at)
        VALUES (?, ?, ?, ?);
    `

	newUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = db.Exec(query, newUUID.String(), dislike.UserID, dislike.CommentID, time.Now())
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUUID, nil
}

func GetPostDislikeByID(db *sql.DB, dislikeID uuid.UUID) (models.PostDislike, error) {
	var dislike models.PostDislike
	query := `
        SELECT id, user_id, post_id, created_at
        FROM post_dislikes
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, dislikeID).Scan(&dislike.ID, &dislike.UserID, &dislike.PostID, &dislike.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.PostDislike{}, errors.New("dislike de publication non trouvé")
		}
		return models.PostDislike{}, err
	}

	return dislike, nil
}

func GetCommentDislikeByID(db *sql.DB, dislikeID uuid.UUID) (models.CommentDislike, error) {
	var dislike models.CommentDislike
	query := `
        SELECT id, user_id, comment_id, created_at
        FROM comment_dislikes
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, dislikeID).Scan(&dislike.ID, &dislike.UserID, &dislike.CommentID, &dislike.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.CommentDislike{}, errors.New("dislike de commentaire non trouvé")
		}
		return models.CommentDislike{}, err
	}

	return dislike, nil
}

func UpdatePostDislike(db *sql.DB, dislike models.PostDislike) error {
	query := `
        UPDATE post_dislikes
        SET user_id = ?, post_id = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, dislike.UserID, dislike.PostID, dislike.ID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCommentDislike(db *sql.DB, dislike models.CommentDislike) error {
	query := `
        UPDATE comment_dislikes
        SET user_id = ?, comment_id = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, dislike.UserID, dislike.CommentID, dislike.ID)
	if err != nil {
		return err
	}

	return nil
}

func RemovePostDislike(db *sql.DB, dislikeID uuid.UUID) error {
	query := `
        DELETE FROM post_dislikes
        WHERE id = ?;
    `

	_, err := db.Exec(query, dislikeID)
	if err != nil {
		return err
	}

	return nil
}

func RemoveCommentDislike(db *sql.DB, dislikeID uuid.UUID) error {
	query := `
        DELETE FROM comment_dislikes
        WHERE id = ?;
    `

	_, err := db.Exec(query, dislikeID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllDislikes(db *sql.DB) ([]interface{}, error) {
	postDislikes, err := GetAllPostDislikes(db)
	if err != nil {
		return nil, err
	}

	commentDislikes, err := GetAllCommentDislikes(db)
	if err != nil {
		return nil, err
	}

	dislikes := append([]interface{}{}, postDislikes)
	dislikes = append(dislikes, commentDislikes)

	return dislikes, nil
}

func GetAllPostDislikes(db *sql.DB) ([]models.PostDislike, error) {
	query := `
        SELECT id, user_id, post_id, created_at
        FROM post_dislikes;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postDislikes []models.PostDislike
	for rows.Next() {
		var dislike models.PostDislike
		err := rows.Scan(&dislike.ID, &dislike.UserID, &dislike.PostID, &dislike.CreatedAt)
		if err != nil {
			return nil, err
		}
		postDislikes = append(postDislikes, dislike)
	}

	return postDislikes, nil
}

func GetAllCommentDislikes(db *sql.DB) ([]models.CommentDislike, error) {
	query := `
        SELECT id, user_id, comment_id, created_at
        FROM comment_dislikes;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentDislikes []models.CommentDislike
	for rows.Next() {
		var dislike models.CommentDislike
		err := rows.Scan(&dislike.ID, &dislike.UserID, &dislike.CommentID, &dislike.CreatedAt)
		if err != nil {
			return nil, err
		}
		commentDislikes = append(commentDislikes, dislike)
	}

	return commentDislikes, nil
}

// GetDislikesByPostID retrieves all dislikes for a specific post by post ID.
func GetDislikesByPostID(db *sql.DB, postID uuid.UUID) ([]models.PostDislike, error) {
	query := `
        SELECT id, user_id, post_id, created_at
        FROM post_dislikes
        WHERE post_id = ?;
    `

	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dislikes []models.PostDislike
	for rows.Next() {
		var dislike models.PostDislike
		err := rows.Scan(&dislike.ID, &dislike.UserID, &dislike.PostID, &dislike.CreatedAt)
		if err != nil {
			return nil, err
		}
		dislikes = append(dislikes, dislike)
	}

	return dislikes, nil
}

// GetCommentDislikesByCommentID retrieves all dislikes for a specific comment by comment ID.
func GetCommentDislikesByCommentID(db *sql.DB, commentID uuid.UUID) ([]models.CommentDislike, error) {
	query := `
        SELECT id, user_id, comment_id, created_at
        FROM comment_dislikes
        WHERE comment_id = ?;
    `

	rows, err := db.Query(query, commentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentDislikes []models.CommentDislike
	for rows.Next() {
		var commentDislike models.CommentDislike
		err := rows.Scan(&commentDislike.ID, &commentDislike.UserID, &commentDislike.CommentID, &commentDislike.CreatedAt)
		if err != nil {
			return nil, err
		}
		commentDislikes = append(commentDislikes, commentDislike)
	}

	return commentDislikes, nil
}

func GetPostDislikeByUserID(db *sql.DB, dislike models.PostDislike) (models.PostDislike, error) {
	query := `
        SELECT id, user_id, post_id, created_at
        FROM post_dislikes
        WHERE user_id = ? AND post_id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, dislike.UserID, dislike.PostID).Scan(&dislike.ID, &dislike.UserID, &dislike.PostID, &dislike.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.PostDislike{}, errors.New("like de publication non trouvé")
		}
		return models.PostDislike{}, err
	}
	return dislike, nil
}

func GetCommentDislikeByUserID(db *sql.DB, dislike models.CommentDislike) (models.CommentDislike, error) {
	query := `
        SELECT id, user_id, comment_id, created_at
        FROM comment_dislikes
        WHERE user_id = ? AND comment_id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, dislike.UserID, dislike.CommentID).Scan(&dislike.ID, &dislike.UserID, &dislike.CommentID, &dislike.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.CommentDislike{}, errors.New("like de commentaire non trouvé")
		}
		return models.CommentDislike{}, err
	}
	return dislike, nil
}
