package view

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
	Error   interface{} `json:"error"`
}

func SuccessCreated(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Payload: payload,
		Message: "success create",
	}
}

func SuccessUpdated(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Payload: payload,
		Message: "success updated",
	}
}

func SuccessDeleted(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Payload: payload,
		Message: "success delete",
	}
}

func SuccessFindAll(payload interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Payload: payload,
		Message: "success get data",
	}
}

func ErrBadRequest(err interface{}) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: "Bad request",
		Error:   err,
	}
}

func ErrInternalServer(err interface{}) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
		Error:   err,
	}
}
func ErrNotFound() *Response {
	return &Response{
		Status:  http.StatusNotFound,
		Message: "Data not found",
		Error:   "NO_DATA",
	}
}
func ErrUnauthorized() *Response {
	return &Response{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
		Error:   "UNAUTHORIZED",
	}
}
