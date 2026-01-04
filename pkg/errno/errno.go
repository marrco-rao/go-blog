package errno

import (
	"net/http"
)

// 统一错误码
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	HTTP    int    `json:"-"`
}

func (e *Error) Error() string {
	return e.Message
}

var (
	OK = &Error{
		0,
		"success",
		http.StatusOK}
	InternalServerError = &Error{
		50000,
		"internal server error",
		http.StatusInternalServerError}
	InvalidParams = &Error{
		40400,
		"invalid parameters",
		http.StatusBadRequest}
	Unauthorized = &Error{
		40401,
		"unauthorized",
		http.StatusUnauthorized}
	UserNotFound = &Error{
		20401,
		"user not found",
		http.StatusUnauthorized}

	PostNotFound = &Error{
		Code:    40404,
		Message: "post not found",
		HTTP:    http.StatusNotFound,
	}

	Timeout = &Error{
		Code:    40408,
		Message: "Request Timeout",
		HTTP:    http.StatusRequestTimeout,
	}
)
