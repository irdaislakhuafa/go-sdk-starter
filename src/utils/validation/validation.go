package validation

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/strformat"
)

var customMessages = map[string]string{
	"required": "Field '{{ .Field }}' is required",
	"email":    "Field '{{ .Field }}' must be a valid email format",
	"max":      "Field '{{ .Field }}' cannot exceed {{ .Param }} characters",
	"min":      "Field '{{ .Field }}' must be at least {{ .Param }} characters",
	"gte":      "Field '{{ .Field }}' must be greater than or equal to {{ .Param }}",
	"lte":      "Field '{{ .Field }}' must be less than or equal to {{ .Param }}",
	"oneof":    "Field '{{ .Field }}' must be one of [{{ .Param }}]",
	"number":   "Field '{{ .Field }}' must be a number format",
}

func ExtractError(err error, val any) error {
	if err, isOk := err.(validator.ValidationErrors); isOk {
		err := err[0]
		msg, isExist := customMessages[err.Tag()]
		if !isExist {
			return errors.NewWithCode(codes.CodeBadRequest, "%s", err.Error())
		}

		msg = strformat.TWE(msg, map[string]string{
			"Field": getJSONTag(reflect.TypeOf(val), err.Field()),
			"Param": err.Param(),
		})
		return errors.NewWithCode(codes.CodeBadRequest, "%s", msg)
	}
	return errors.NewWithCode(codes.CodeBadRequest, "%s", err.Error())
}

func getJSONTag(ref reflect.Type, fieldName string) string {
	sf, isOk := ref.FieldByName(fieldName)
	if !isOk {
		return fieldName
	}

	tagVal := sf.Tag.Get("json")
	if tagVal == "" {
		return fieldName
	}

	return tagVal
}
