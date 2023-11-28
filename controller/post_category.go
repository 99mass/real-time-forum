package controller

import (
	"database/sql"

	"forum/models"

	"github.com/gofrs/uuid"
)

// Créer un PostCategory en utlisant les deux clés étrangers comme clé primaire
func CreatePostCategory(db *sql.DB, postID, categoryID uuid.UUID) error {
	query := `
        INSERT INTO posts_categories (post_id, category_id)
        VALUES (?, ?);
    `

	_, err := db.Exec(query, postID, categoryID)
	if err != nil {
		return err
	}

	return nil
}

// supprimer un postCategory en utilisant les IDs de post et de category
func DeletePostCategory(db *sql.DB, postID, categoryID uuid.UUID) error {
	query := `
        DELETE FROM posts_categories
        WHERE post_id = ? AND category_id = ?;
    `

	_, err := db.Exec(query, postID, categoryID)
	if err != nil {
		return err
	}

	return nil
}

// mettre à jour un postCategory en utilisant les IDs de post et de category
func UpdatePostCategory(db *sql.DB, postID, categoryID uuid.UUID) error {
	query := `
        UPDATE posts_categories
        SET category_id = ?
        WHERE post_id = ?;
    `

	_, err := db.Exec(query, categoryID, postID)
	if err != nil {
		return err
	}

	return nil
}

// GetPostsByCategory récupère les posts associés à une catégorie donnée
func GetPostsByCategory(db *sql.DB, categoryID uuid.UUID) ([]models.Post, error) {
	query := `
        SELECT p.id, p.user_id, p.title, p.content, p.image_post, p.created_at
        FROM posts p
        JOIN posts_categories pc ON p.id = pc.post_id
        WHERE pc.category_id = ?
		ORDER BY created_at DESC;
    `

	rows, err := db.Query(query, categoryID)
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

// GetCategoriesByPost récupère les catégories associées à un post donné
func GetCategoriesByPost(db *sql.DB, postID uuid.UUID) ([]models.Category, error) {
	query := `
        SELECT c.id, c.name_category
        FROM categories c
        JOIN posts_categories pc ON c.id = pc.category_id
        WHERE pc.post_id = ?;
    `

	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.NameCategory)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// GetAllPostCategories retrieves all records from the PostCategory table.
func GetAllPostCategories(db *sql.DB) ([]models.PostCategory, error) {
	query := `
        SELECT category_id, post_id
        FROM posts_categories;
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postCategories []models.PostCategory
	for rows.Next() {
		var pc models.PostCategory
		err := rows.Scan(&pc.CategoryID, &pc.PostID)
		if err != nil {
			return nil, err
		}
		postCategories = append(postCategories, pc)
	}

	return postCategories, nil
}
