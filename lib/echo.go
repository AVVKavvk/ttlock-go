package lib

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Dict map[string]interface{}

type RequestError struct {
	StatusCode int
	Message    string
	Reason     string
}

func (rr *RequestError) Error() string {
	return fmt.Sprintf("status: %d, err: %s, reason: %s", rr.StatusCode, rr.Message, rr.Reason)
}

func NewErrorCustom(status int, msg, reason string) *RequestError {
	return &RequestError{StatusCode: status, Message: msg, Reason: reason}
}

// 401
func NewErrorUnauthorized() *RequestError {
	return NewErrorCustom(http.StatusUnauthorized, "unauthorized", "invalidAuth")
}

// 403
func NewErrorForbidden() *RequestError {
	return NewErrorCustom(http.StatusForbidden, "forbidden", "forbiddenAction")
}

type V1ResponseWrapper struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}
type V1Response struct {
	V1ResponseWrapper
	Data interface{} `json:"data"`
} //@name Response

func NewV1Response(status int, message, reason string, data interface{}) V1Response {
	return V1Response{
		V1ResponseWrapper: V1ResponseWrapper{
			Status:  status,
			Message: message,
			Reason:  reason,
		},
		Data: data,
	}
}

func NewV1Err(err *RequestError) V1Response {
	return NewV1Response(err.StatusCode, err.Message, err.Reason, nil)
}

func ContextV1Err(ctx echo.Context, err error) error {
	if rerr, ok := err.(*RequestError); ok {
		return ctx.JSON(rerr.StatusCode, NewV1Err(rerr))
	}
	return ctx.JSON(http.StatusBadRequest, NewErrorCustom(
		http.StatusBadRequest, err.Error(), "",
	))
}

func ContextV1Success(ctx echo.Context, message string, data interface{}) error {
	return ctx.JSON(http.StatusOK, NewV1Success(message, data))
}

func NewV1Success(message string, data interface{}) V1Response {
	return NewV1Response(http.StatusOK, message, "", data)
}

func ContextV1Accepted(ctx echo.Context, message string, data interface{}) error {
	return ctx.JSON(http.StatusAccepted, NewV1Response(http.StatusAccepted, message, "", data))
}
