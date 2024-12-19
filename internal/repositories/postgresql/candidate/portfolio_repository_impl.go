package candidate

import (
	"database/sql"
	models "dz-jobs-api/internal/models/candidate"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces/candidate"
	"fmt"
	"github.com/google/uuid"
)

type SQLCandidatePortfolioRepository struct {
	db *sql.DB
}

func NewCandidatePortfolioRepository(db *sql.DB) repositoryInterfaces.CandidatePortfolioRepository {
	return &SQLCandidatePortfolioRepository{
		db: db,
	}
}

func (r *SQLCandidatePortfolioRepository) CreatePortfolio(portfolio models.CandidatePortfolio) error {
	query := `INSERT INTO candidate_portfolio (project_id, candidate_id, project_name, project_link, category, description) 
			VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, portfolio.ProjectID, portfolio.CandidateID, portfolio.ProjectName, portfolio.ProjectLink, portfolio.Category, portfolio.Description)
	if err != nil {
		return fmt.Errorf("unable to create portfolio: %w", err)
	}
	return nil
}

func (r *SQLCandidatePortfolioRepository) GetPortfolioByCandidateID(id uuid.UUID) ([]models.CandidatePortfolio, error) {
	rows, err := r.db.Query(`SELECT project_id, candidate_id, project_name, project_link, category, description FROM candidate_portfolio WHERE candidate_id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch portfolio: %w", err)
	}
	defer rows.Close()

	var portfolios []models.CandidatePortfolio
	for rows.Next() {
		var portfolio models.CandidatePortfolio
		if err := rows.Scan(&portfolio.ProjectID, &portfolio.CandidateID, &portfolio.ProjectName, &portfolio.ProjectLink, &portfolio.Category, &portfolio.Description); err != nil {
			return nil, fmt.Errorf("unable to scan portfolio data: %w", err)
		}
		portfolios = append(portfolios, portfolio)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}
	return portfolios, nil
}
func (r *SQLCandidatePortfolioRepository) UpdatePortfolio(portfolio models.CandidatePortfolio) error {
	
	query := `
		UPDATE candidate_portfolio
		SET project_id = $1, project_name = $2, project_link = $3, category = $4, description = $5
		WHERE candidate_id = $6 AND project_name = $7
	`
	_, err := r.db.Exec(query, portfolio.ProjectID, portfolio.ProjectName, portfolio.ProjectLink, portfolio.Category, portfolio.Description, portfolio.CandidateID, portfolio.ProjectName)
	if err != nil {
		return fmt.Errorf("unable to update portfolio: %w", err)
	}
	return nil
}

func (r *SQLCandidatePortfolioRepository) DeletePortfolio(id uuid.UUID, projectName string) error {
	query := `DELETE FROM candidate_portfolio WHERE candidate_id = $1 AND project_name = $2`
	_, err := r.db.Exec(query, id, projectName)
	if err != nil {
		return fmt.Errorf("unable to delete portfolio: %w", err)
	}
	return nil
}
