package rest_errors

import (
	"errors"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}
type restErr struct {
	message string
	status  int
	error   string
	causes  []interface{}
}

func (e restErr) Message() string {
	return e.message
}
func (e restErr) Status() int {
	return e.status
}
func (e restErr) Causes() []interface{} {
	return e.causes
}
func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status:%d -error:%s -Causes:[ %v ]", e.message, e.status, e.error, e.causes)
}
func NewError(msg string) error {
	return errors.New(msg)
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		message: message,
		status:  status,
		error:   err,
		causes:  causes,
	}
}
func NewBadRequestError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}
func NewNotFoundError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusNotFound,
		error:   "not_found",
	}
}
func NewUnauthorizedError(message string) RestErr {
	return restErr{
		message: "unable to retrieve user info. from given access_token",
		status:  http.StatusUnauthorized,
		error:   "unauthorized",
	}
}
func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}
