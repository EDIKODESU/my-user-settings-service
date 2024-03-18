package utils

import "my-user-settings-service/internal/resources"

type BadRequestResource struct {
	Status string `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type BadRequestResponse struct {
	Errors []*BadRequestResource `json:"errors"`
}

func GetBadRequest() resources.RequestError {
	errData := resources.RequestErrorResource{
		Status: "400",
		Title:  "Bad Request",
		Detail: "Your request parameters are invalid.",
	}
	return resources.RequestError{
		Errors: &[]resources.RequestErrorResource{errData},
	}
}

func GetInternalServerError() resources.RequestError {
	errData := resources.RequestErrorResource{
		Status: "500",
		Title:  "Internal Server Error",
		Detail: "An unexpected error occurred on the server.",
	}
	return resources.RequestError{
		Errors: &[]resources.RequestErrorResource{errData},
	}
}
