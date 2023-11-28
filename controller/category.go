package controller

import (
	"database/sql"
	"errors"

	"forum/models"

	"github.com/gofrs/uuid"
)

func CreateCategory(db *sql.DB, category models.Category) (uuid.UUID, error) {
	// Vérifiez d'abord si la catégorie existe déjà
	query := `
        SELECT id FROM categories WHERE name_category = ? LIMIT 1;
    `

	var existingID string
	err := db.QueryRow(query, category.NameCategory).Scan(&existingID)

	switch {
	case err == sql.ErrNoRows:
		// La catégorie n'existe pas, alors créez-la
		insertQuery := `
            INSERT INTO categories (id, name_category)
            VALUES (?, ?);
        `

		newUUID, err := uuid.NewV4()
		if err != nil {
			return uuid.UUID{}, err
		}

		_, err = db.Exec(insertQuery, newUUID.String(), category.NameCategory)
		if err != nil {
			return uuid.UUID{}, err
		}

		return newUUID, nil

	case err != nil:
		// Une erreur s'est produite lors de la requête
		return uuid.UUID{}, err

	default:
		// La catégorie existe déjà, retournez l'ID existant
		existingUUID, err := uuid.FromString(existingID)
		if err != nil {
			return uuid.UUID{}, err
		}
		return existingUUID, nil
	}
}

func GetAllCategories(db *sql.DB) ([]models.Category, error) {
	query := `
        SELECT id, name_category
        FROM categories;
    `

	rows, err := db.Query(query)
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

func GetCategoryByID(db *sql.DB, categoryID uuid.UUID) (models.Category, error) {
	var category models.Category
	query := `
        SELECT id, name_category
        FROM categories
        WHERE id = ?
        LIMIT 1;
    `

	err := db.QueryRow(query, categoryID).Scan(&category.ID, &category.NameCategory)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Category{}, errors.New("catégorie non trouvée")
		}
		return models.Category{}, err
	}

	return category, nil
}

func UpdateCategory(db *sql.DB, category models.Category) error {
	query := `
        UPDATE categories
        SET name_category = ?
        WHERE id = ?;
    `

	_, err := db.Exec(query, category.NameCategory, category.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategory(db *sql.DB, categoryID uuid.UUID) error {
	query := `
        DELETE FROM categories
        WHERE id = ?;
    `

	_, err := db.Exec(query, categoryID)
	if err != nil {
		return err
	}

	return nil
}
