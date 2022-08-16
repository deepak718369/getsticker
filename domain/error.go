package domain

import (
	"errors"
	"net/http"
	logg "github.com/sirupsen/logrus"
)
// ResponseError make response error to client 
type ResponseError struct {
	Message string `json:"message"`
}
// var bind all the errors 
var (
	// ERR_INTERNAL_SERVER_ERROR will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
     // ErrNotFound will throw if the requested item is not exists
	ErrNotFound      = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict       = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput       = errors.New("given Param is not valid")
)

//GetStatusCode  generate http status code 
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logg.Error(err)
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
