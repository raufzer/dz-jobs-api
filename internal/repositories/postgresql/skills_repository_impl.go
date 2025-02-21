package postgresql

import (
		"context"
	"database/sql"
	"dz-jobs-api/internal/models"
	repositoryInterfaces "dz-jobs-api/internal/repositories/interfaces"
	"fmt"

	"github.com/google/uuid"
)

type SQLCandidateSkillsRepository struct {
	db *sql.DB
}

func NewCandidateSkillsRepository(db *sql.DB) repositoryInterfaces.CandidateSkillsRepository {
	return &SQLCandidateSkillsRepository{
		db: db,
	}
}

func (r *SQLCandidateSkillsRepository) CreateSkill(ctx context.Context, skill *models.CandidateSkills) error {
	query := `INSERT INTO candidate_skills (candidate_id, skill) VALUES ($1, $2)`
	_, err := r.db.Exec(query, skill.ID, skill.Skill)
	if err != nil {
		return fmt.Errorf("unable to create skill: %w", err)
	}
	return nil
}

func (r *SQLCandidateSkillsRepository) GetSkills(ctx context.Context, candidateID uuid.UUID) ([]models.CandidateSkills, error) {
	rows, err := r.db.Query(`SELECT candidate_id, skill FROM candidate_skills WHERE candidate_id = $1`, candidateID)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch skills: %w", err)
	}
	defer rows.Close()

	var skills []models.CandidateSkills
	for rows.Next() {
		var skill models.CandidateSkills
		if err := rows.Scan(&skill.ID, &skill.Skill); err != nil {
			return nil, fmt.Errorf("unable to scan skill data: %w", err)
		}
		skills = append(skills, skill)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}
	return skills, nil
}

func (r *SQLCandidateSkillsRepository) DeleteSkill(ctx context.Context, candidateID uuid.UUID, skillName string) error {
	query := `DELETE FROM candidate_skills WHERE candidate_id = $1 AND skill = $2`
	_, err := r.db.Exec(query, candidateID, skillName)
	if err != nil {
		return fmt.Errorf("unable to delete skill: %w", err)
	}
	return nil
}
