package validation

import (
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/irdaislakhuafa/go-sdk/header"
)

func maxFileSize(fl validator.FieldLevel) bool {
	f, isOk := fl.Field().Interface().(multipart.FileHeader)
	if !isOk {
		return true
	}

	param := fl.Param()
	if rg := regexp.MustCompile(`(?i)\d+\s?(MB|KB)`); rg.MatchString(param) {
		param = param[:len(param)-2]
	}

	size, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return false
	}

	if strings.HasSuffix(fl.Param(), "KB") {
		size = size * 1024
	}
	if strings.HasSuffix(fl.Param(), "MB") {
		size = size * 1024 * 1024
	}

	return f.Size <= size
}

func mimeType(fl validator.FieldLevel) bool {
	fh, isOk := fl.Field().Interface().(multipart.FileHeader)
	if !isOk {
		return true
	}

	allowed := map[string]bool{}
	for _, v := range strings.Split(fl.Param(), " ") {
		allowed[v] = true
	}

	f, err := fh.Open()
	if err != nil {
		return false
	}
	defer f.Close()

	contentType := fh.Header.Get(header.KeyContentType)
	return allowed[contentType]
}
