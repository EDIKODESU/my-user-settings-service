package utils

type UserResource struct {
	Type       string      `json:"type"`
	ID         int64       `json:"id"`
	Attributes interface{} `json:"attributes"`
}

type UsersListResponse struct {
	Data  []*UserResource `json:"data"`
	Links Links           `json:"links"`
}

func BuildResponceList(dataResponce []*UserResource, linkResponse Links) UsersListResponse {
	response := UsersListResponse{
		Data:  dataResponce,
		Links: linkResponse,
	}
	return response
}
