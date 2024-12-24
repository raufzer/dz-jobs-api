package helpers

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
