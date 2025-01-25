package helpers

// import (
// 	"fmt"
// 	"regexp"

// 	"github.com/go-playground/validator/v10"
// )

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

// var Validate *validator.Validate

// func Init() {

// 	Validate = validator.New()
// 	Validate.RegisterValidation("jobSalaryRange", ValidateSalaryRange)
// }

// func ValidateSalaryRange(fl validator.FieldLevel) bool {

// 	re := regexp.MustCompile(`^\d{5,}-\d{5,}$`)
// 	if !re.MatchString(fl.Field().String()) {
// 		return false
// 	}

// 	var min, max int
// 	_, err := fmt.Sscanf(fl.Field().String(), "%d-%d", &min, &max)
// 	if err != nil {
// 		return false
// 	}

// 	return min < max
// }
