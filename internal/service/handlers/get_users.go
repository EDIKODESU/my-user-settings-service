package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"my-user-settings-service/internal/data"
	"my-user-settings-service/internal/resources"
	"my-user-settings-service/internal/service/utils"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	log.Info("Getting users")

	page, perPage := utils.ParsePaginationParams(r)

	users, err := UsersQ(r).Select(page, perPage)
	if err != nil {
		log.Errorf("Failed to fetch profiles from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		ape.Render(w, utils.GetInternalServerError())
		return
	}

	totalCount, err := UsersQ(r).Count()
	if err != nil {
		log.Errorf("Failed to fetch total count of users: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		ape.Render(w, utils.GetInternalServerError())
		return
	}

	totalPages := (totalCount + perPage - 1) / perPage

	selfLink, err := utils.GetSelfPage(r, page, perPage)
	if err != nil {
		log.Errorf("Failed to build pagination link: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		ape.Render(w, utils.GetInternalServerError())
		return
	}

	nextLink, err := utils.GetNextPage(r, page, totalPages, perPage)
	if err != nil {
		log.Errorf("Failed to build pagination link: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		ape.Render(w, utils.GetInternalServerError())
		return
	}

	lastLink, err := utils.GetLastPage(r, totalPages, perPage)
	if err != nil {
		log.Errorf("Failed to build pagination link: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		ape.Render(w, utils.GetInternalServerError())
		return
	}

	response := resources.UserListListResponse{
		Data:  newUserModel(users),
		Links: newLinksModel(selfLink, nextLink, lastLink),
	}
	ape.Render(w, response)
}

func newLinksModel(selfLink, nextLink, lastLink string) *resources.Links {
	return &resources.Links{
		Self: selfLink,
		Next: nextLink,
		Last: lastLink,
	}
}

func newUserModel(users []data.Users) []resources.UserList {
	var userResources []resources.UserList
	for _, u := range users {
		userResources = append(userResources, resources.UserList{
			Key: resources.NewKeyInt64(u.Id, resources.USERS),
			Attributes: &resources.UserListAttributes{
				FirstName:  u.FirstName,
				SecondName: u.SecondName,
				Login:      u.Login,
				Mail:       u.Email,
				Password:   u.Password,
			},
		})
	}
	return userResources
}
