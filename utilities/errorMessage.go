package utilities

import (
	"strings"
)

func FieldContainSpaces(fields []Field) []ApiError {
	var errors []ApiError

	for _, field := range fields {
		if !strings.Contains(field.Value, " ") {
			continue
		}
		errors = append(errors, ApiError{
			Field:  field.Name,
			Tag:    "spaces",
			Message: field.Name + " tidak boleh mengandung spasi",
		})
	}

	return errors
}

func UsernameContainSpecialChar() response {
	return ApiResponse(
		"Username tidak boleh mengandung karakter spesial,"+
		" kecuali garis bawah(_) dan titik(.)",
		false,
		nil,
	)
}