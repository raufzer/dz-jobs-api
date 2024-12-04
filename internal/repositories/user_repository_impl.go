package repositories

import (
	"database/sql"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"errors"
)

// SQLUserRepository is the SQL implementation of the UserRepository interface.
type SQLUserRepository struct {
	db *sql.DB
}

// NewUserRepository initializes a new UserRepository with a database connection.
func NewUserRepository(db *sql.DB) repositoryInterfaces.UserRepository {
	return &SQLUserRepository{
		db: db,
	}
}

// Create a new user in the database.
func (r *SQLUserRepository) Create(user *models.User) error {
	query := "INSERT INTO users (name, email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id"

	// Use QueryRow instead of Prepare for simple inserts
	var id int
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.Role).Scan(&id)
	if err != nil {
		return err
	}

	user.ID = id
	return nil
}

// GetByEmail fetches a user by name from the database.
func (r *SQLUserRepository) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// GetByID fetches a user by ID from the database.
func (r *SQLUserRepository) GetByID(id int) (*models.User, error) {
	query := "SELECT id, name, email, role, created_at, updated_at FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// GetAll fetches all users from the database.
func (r *SQLUserRepository) GetAll() ([]*models.User, error) {
	query := "SELECT id, name, email, role, created_at, updated_at FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Update updates an existing user in the database.
func (r *SQLUserRepository) Update(id int, user *models.User) error {
	query := "UPDATE users SET name = $1, email = $2, password = $3, role = $4, updated_at = NOW() WHERE id = $5"

	result, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.Role, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no user found or updated")
	}

	return nil
}

// Delete removes a user from the database by name.
func (r *SQLUserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted")
	}

	return nil
}
