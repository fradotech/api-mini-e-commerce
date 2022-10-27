package view

import "net/http"

type Response struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	GeneralInfo string      `json:"general_info"`
	Payload     interface{} `json:"payload"`
	Error       interface{} `json:"error"`
}

func SuccessCreated(payload interface{}) *Response {
	return &Response{
		Status:      http.StatusCreated,
		Message:     "success create",
		GeneralInfo: "Kelompok 1",
		Payload:     payload,
	}
}

func SuccessUpdated(payload interface{}) *Response {
	return &Response{
		Status:      http.StatusOK,
		Message:     "success updated",
		GeneralInfo: "Kelompok 1",
		Payload:     payload,
	}
}

func SuccessDeleted(payload interface{}) *Response {
	return &Response{
		Status:      http.StatusOK,
		Message:     "success delete",
		GeneralInfo: "Kelompok 1",
		Payload:     payload,
	}
}

func SuccessFindAll(payload interface{}) *Response {
	return &Response{
		Status:      http.StatusOK,
		Message:     "success get data",
		GeneralInfo: "Kelompok 1",
		Payload:     payload,
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
