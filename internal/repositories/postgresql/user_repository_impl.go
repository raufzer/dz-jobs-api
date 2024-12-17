package postgresql

import (
	"database/sql"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"errors"
	"fmt"

	"github.com/google/uuid"

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
	var userid uuid.UUID
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.Role).Scan(&userid)
	if err != nil {
		return fmt.Errorf("repository: failed to create user: %w", err)
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("repository: failed to fetch user by email: %w", err)
	}
	return user, nil
}
func (r *SQLUserRepository) GetByID(userid uuid.UUID) (*models.User, error) {
	query := "SELECT userid, name, email, role, created_at, updated_at FROM users WHERE userid = $1"
	row := r.db.QueryRow(query, userid)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("repository: failed to fetch user by ID: %w", err)
	}
	return user, nil
}
func (r *SQLUserRepository) GetAll() ([]*models.User, error) {
	query := "SELECT userid, name, email, role, created_at, updated_at FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("repository: failed to fetch users: %w", err)
	}
	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("repository: failed to scan user data: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repository: error occurred while iterating users: %w", err)
	}
	return users, nil
}
func (r *SQLUserRepository) Update(userid uuid.UUID, user *models.User) error {
	query := "UPDATE users SET name = $1, email = $2, password = $3, role = $4, updated_at = NOW() WHERE userid = $5"
	result, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.Role, userid)
	if err != nil {
		return fmt.Errorf("repository: failed to update user: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
func (r *SQLUserRepository) UpdatePassword(email, hashedPassword string) error {
	query := "UPDATE users SET password = $1, updated_at = NOW() WHERE email = $2"
	result, err := r.db.Exec(query, hashedPassword, email)
	if err != nil {
		return fmt.Errorf("repository: failed to update password: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
func (r *SQLUserRepository) Delete(userid uuid.UUID) error {
	query := "DELETE FROM users WHERE userid = $1"
	result, err := r.db.Exec(query, userid)
	if err != nil {
		return fmt.Errorf("repository: failed to delete user: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
