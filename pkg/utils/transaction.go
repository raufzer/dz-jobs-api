package utils

import (
	"context"
	"database/sql"
	"fmt"
)

// TxFn represents a function that uses a transaction
type TxFn func(*sql.Tx) error

// WithTransaction executes the given function within a transaction
func WithTransaction(db *sql.DB, fn TxFn) error {
	// Use context for better timeout handling
	ctx := context.Background()

	// Begin transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Defer a rollback in case anything fails
	defer func() {
		if tx != nil {
			tx.Rollback()
		}
	}()

	// Execute the function passing in the transaction
	if err := fn(tx); err != nil {
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
