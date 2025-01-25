package helpers

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ConvertSalaryRange(min, max float64) (string, []interface{}) {
	if min > 0 && max > 0 {
		return `
			CAST(split_part(salary_range, ' ', 1) AS float) >= $1 
			AND CAST(split_part(salary_range, ' ', 3) AS float) <= $2
		`, []interface{}{min, max}
	} else if min > 0 {
		return `
			CAST(split_part(salary_range, ' ', 1) AS float) >= $1
		`, []interface{}{min}
	} else {
		return `
			CAST(split_part(salary_range, ' ', 3) AS float) <= $1
		`, []interface{}{max}
	}
}
// Global validator instance
var Validate *validator.Validate

func Init() {
	// Initialize the global validator instance and register custom validation
	Validate = validator.New()
	Validate.RegisterValidation("jobSalaryRange", ValidateSalaryRange)
}

// Custom salary range validator
func ValidateSalaryRange(fl validator.FieldLevel) bool {
	// Regex for validating salary range format
	re := regexp.MustCompile(`^\d{5,}-\d{5,}$`) 
	if !re.MatchString(fl.Field().String()) {
		return false
	}

	// Extract the min and max values from the salary range string
	var min, max int
	_, err := fmt.Sscanf(fl.Field().String(), "%d-%d", &min, &max)
	if err != nil {
		return false
	}

	// Ensure that the minimum salary is less than the maximum salary
	return min < max
}