package controller

import (
	"database/sql"
	"errors"
	"time"

	"forum/models"

	"github.com/gofrs/uuid"
)

func FormatCreatedAt(createdTimeStr string) (string, error) {
	createdTime, err := time.Parse(time.RFC3339Nano, createdTimeStr)
	if err != nil {
		return "", err
	}
	return createdTime.Format("2006-01-02 15:04:05"), nil
}
func CreatePost(db *sql.DB, post models.Post) (uuid.UUID, error) {

	newUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}, err
	}
	for _, v := range post.Categories {
		err := CreatePostCategory(db, newUUID, v.ID)
		if err != nil {
			return v.ID, errors.New("")
		}
	}

	query := `
        INSERT INTO posts (id, user_id, title,  content, image_post, created_at)
        VALUES (?, ?, ?, ?, ?, ?);
    	`

	_, err = db.Exec(query, newUUID.String(), post.UserID, post.Title, post.Content, post.Image, time.Now())
	if err != nil {
		return uuid.UUID{}, err
	}

	return newUUID, nil
}

func GetPostByID(db *sql.DB, postID uuid.UUID) (models.Post, error) {
	var post models.Post
	query := `
        SELECT id, user_id, title, content, image_post, created_at
        FROM posts
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, postID).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Image, &post.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Post{}, errors.New("publication non trouvée")
		}
		return models.Post{}, err
	}
	timeformated, _ := FormatCreatedAt(post.CreatedAt)
	post.CreatedAt = timeformated

	return post, nil
}

func UpdatePost(db *sql.DB, post models.Post) error {
	query := `
        UPDATE posts
        SET title = ?, content = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeletePost(db *sql.DB, postID uuid.UUID) error {
	query := `
        DELETE FROM posts
        WHERE id = ?;
    `

	_, err := db.Exec(query, postID)
	if err != nil {
		return err
	}

	return nil
}

func GetAllPosts(db *sql.DB) ([]models.Post, error) {
	query := `
        SELECT id, user_id, title, content, image_post, created_at
        FROM posts
		ORDER BY created_at DESC;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Image, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		timeformated, _ := FormatCreatedAt(post.CreatedAt)
		post.CreatedAt = timeformated
		posts = append(posts, post)
	}

	return posts, nil
}

// Function to get user by post ID
func GetUserByPostID(db *sql.DB, postID uuid.UUID) (*models.User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.created_at
		FROM users u
		INNER JOIN posts p ON u.id = p.user_id
		WHERE p.id = ?;
	`

	var user models.User
	err := db.QueryRow(query, postID).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetPostsByUserID(db *sql.DB, userID uuid.UUID) ([]models.Post, error) {
	query := `
		SELECT id, user_id, title, content, image_post, created_at
		FROM posts
		WHERE user_id = ?
		ORDER BY created_at DESC;
	`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Image, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		timeformated, _ := FormatCreatedAt(post.CreatedAt)
		post.CreatedAt = timeformated
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostsByUserAndCategory(db *sql.DB, userID uuid.UUID, categoryID uuid.UUID) ([]models.Post, error) {
	query := `
        SELECT p.id, p.user_id, p.title, p.content, p.image_post, p.created_at
        FROM posts p
        JOIN posts_categories pc ON p.id = pc.post_id
        WHERE p.user_id = ? AND pc.category_id = ?
		ORDER BY created_at DESC;
    `

	rows, err := db.Query(query, userID.String(), categoryID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Image, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		timeformated, _ := FormatCreatedAt(post.CreatedAt)
		post.CreatedAt = timeformated
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostsByTitle(db *sql.DB, title string) ([]models.Post, error) {
	query := `
        SELECT id, title, content, image_post, created_at
        FROM posts
        WHERE title LIKE ?;
    `

	rows, err := db.Query(query, "%"+title+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		timeformated, _ := FormatCreatedAt(post.CreatedAt)
		post.CreatedAt = timeformated
		posts = append(posts, post)
	}

	return posts, nil
}
