package formvalidation

import (
	"net/http"
	"strconv"
)

type FormValidation struct {
	*http.Request
}

func NewFormValidation(r *http.Request) *FormValidation {
	return &FormValidation{
		Request: r,
	}
}

func (fv *FormValidation) ValidateID(pathName string) (int, bool) {
	var isValid bool
	pathValue := fv.Request.PathValue(pathName)
	if pathValue == "" {
		return 0, false
	}
	id, err := strconv.Atoi(pathValue)
	if err == nil {
		isValid = true
	}
	return id, isValid
}
