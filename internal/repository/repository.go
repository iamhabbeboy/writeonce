package repository

import "github.com/go-playground/validator/v10"

func ValidateParams(s interface{}, fields ...string) error {
	validate := validator.New()
	if len(fields) > 0 {
		if err := validate.StructPartial(s, fields...); err != nil {
			return err
		}
	} else {
		if err := validate.Struct(s); err != nil {
			return err
		}
	}
	return nil
}
