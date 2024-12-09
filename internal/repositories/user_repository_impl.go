package repositories

import (
	"database/sql"
	"dz-jobs-api/internal/helpers"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"net/http"
)

type SQLUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositoryInterfaces.UserRepository {
	return &SQLUserRepository{
		db: db,
	}
}

func (r *SQLUserRepository) Create(user *models.User) error {
	query := "INSERT INTO users (name, email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING userid"

	var userid int
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.Role).Scan(&userid)
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to create user")
	}

	user.ID = userid
	return nil
}

func (r *SQLUserRepository) GetByEmail(email string) (*models.User, error) {
	query := "SELECT userid, name, email, password, role, created_at, updated_at FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Failed to fetch user by email")
	}

	return user, nil
}

func (r *SQLUserRepository) GetByID(userid int) (*models.User, error) {
	query := "SELECT userid, name, email, role, created_at, updated_at FROM users WHERE userid = $1"
	row := r.db.QueryRow(query, userid)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.NewCustomError(http.StatusNotFound, "User not found")
		}
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Failed to fetch user by ID")
	}
	return user, nil
}

func (r *SQLUserRepository) GetAll() ([]*models.User, error) {
	query := "SELECT userid, name, email, role, created_at, updated_at FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Failed to fetch users")
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, helpers.NewCustomError(http.StatusInternalServerError, "Failed to scan user data")
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, helpers.NewCustomError(http.StatusInternalServerError, "Error occurred while iterating users")
	}

	return users, nil
}

func (r *SQLUserRepository) Update(userid int, user *models.User) error {
	query := "UPDATE users SET name = $1, email = $2, password = $3, role = $4, updated_at = NOW() WHERE userid = $5"

	result, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.Role, userid)
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to update user")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to check rows affected")
	}
	if rowsAffected == 0 {
		return helpers.NewCustomError(http.StatusNotFound, "No user found or updated")
	}

	return nil
}

func (r *SQLUserRepository) Delete(userid int) error {
	query := "DELETE FROM users WHERE userid = $1"
	result, err := r.db.Exec(query, userid)
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to delete user")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return helpers.NewCustomError(http.StatusInternalServerError, "Failed to check rows affected")
	}
	if rowsAffected == 0 {
		return helpers.NewCustomError(http.StatusNotFound, "No rows deleted")
	}

	return nil
}
