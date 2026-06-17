package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

var invalidExcelSheetNameChars = regexp.MustCompile(`[\[\]\*:/\\?]`)

// zone don't allow "[ ] * : / \ ?" these characters for invalid excel sheet name when export
func init() {
	Validate.RegisterValidation("excel_sheet_name", func(fl validator.FieldLevel) bool {
		return !invalidExcelSheetNameChars.MatchString(fl.Field().String())
	})
}
