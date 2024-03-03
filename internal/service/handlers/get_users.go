package handlers

import (
	"gitlab.com/distributed_lab/ape"
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
		return
	}

	var userResources []*utils.UserResource
	for _, u := range users {
		userResources = append(userResources, &utils.UserResource{
			Type: "users",
			ID:   u.ID,
			Attributes: map[string]interface{}{
				"first_name":  u.FirstName,
				"second_name": u.SecondName,
				"login":       u.Login,
				"email":       u.Email,
				"password":    u.Password,
			},
		})
	}

	totalCount, err := UsersQ(r).Count()
	if err != nil {
		log.Errorf("Failed to fetch total count of users: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	totalPages := (totalCount + perPage - 1) / perPage
	
	selfLink, err := utils.GetSelfPage(r, page, perPage)
	if err != nil {
		log.Errorf("Failed to build pagination link: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	nextLink, err := utils.GetNextPage(r, page, totalPages, perPage)
	if err != nil {
		log.Errorf("Failed to build pagination link: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	lastLink, err := utils.GetLastPage(r, totalPages, perPage)
	if err != nil {
		log.Errorf("Failed to build pagination link: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	linkResponse := utils.Links{
		Self: selfLink,
		Next: nextLink,
		Last: lastLink,
	}

	response := utils.BuildResponceList(userResources, linkResponse)

	ape.Render(w, response)
}
