package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"my-user-settings-service/internal/data"
	"my-user-settings-service/internal/service/utils"
	"net/http"
)

//type User struct {
//	ID         int64  `jsonapi:"primary,users"`
//	FirstName  string `db:"first_name" jsonapi:"attr,first_name"`
//	SecondName string `db:"second_name" jsonapi:"attr,second_name"`
//	Login      string `db:"login" jsonapi:"attr,login"`
//	Email      string `db:"mail" jsonapi:"attr,mail"`
//	Pass       string `db:"password" jsonapi:"attr,password"`
//}

type UsersListResponse struct {
	Data  []*data.Users `json:"data"`
	Links utils.Links   `json:"links"`
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	log.Info("Getting users")

	// Отримання параметрів пагінації з запиту
	//page, perPage := parsePaginationParams(r)
	page, perPage := utils.ParsePaginationParams(r)

	// Виконання запиту до бази даних для отримання користувачів
	users, err := UsersQ(r).Select(page, perPage)
	if err != nil {
		log.Errorf("Failed to fetch profiles from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Формування відповіді у форматі JSON:API
	w.Header().Set("Content-Type", "application/vnd.api+json")

	var userResources []*data.Users
	for _, u := range users {
		userResources = append(userResources, &data.Users{
			ID:         u.ID,
			FirstName:  u.FirstName,
			SecondName: u.SecondName,
			Login:      u.Login,
			Email:      u.Email,
			Pass:       u.Pass,
		})
	}

	totalCount, err := UsersQ(r).Count()
	if err != nil {
		log.Errorf("Failed to fetch total count of users: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	totalPages := (totalCount + perPage - 1) / perPage
	println(totalPages)
	println(totalCount)

	nextLink, err := utils.GetNextPage(r.URL.String(), page, totalPages, perPage)
	if err != nil {
		log.Errorf("Failed to build pagination link: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	lastLink, err := utils.GetLastPage(r.URL.String(), totalPages, perPage)
	if err != nil {
		log.Errorf("Failed to build pagination link: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	linkResponse := utils.Links{
		Self: r.URL.String(),
		Next: nextLink,
		Last: lastLink,
	}

	response := UsersListResponse{
		Data:  userResources,
		Links: linkResponse,
	}

	ape.Render(w, response)

	//if err := jsonapi.MarshalPayload(w, response); err != nil {
	//	log.Errorf("Failed to marshal JSON API response: %v", err)
	//	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	return
	//}
}

//func parsePaginationParams(r *http.Request) (int, int) {
//	pageStr := r.URL.Query().Get("page")
//	perPageStr := r.URL.Query().Get("per_page")
//
//	// Перетворення рядків параметрів у цілі числа
//	page, err := strconv.Atoi(pageStr)
//	if err != nil || page <= 0 {
//		page = 1
//	}
//
//	perPage, err := strconv.Atoi(perPageStr)
//	if err != nil || perPage <= 0 {
//		perPage = 10
//	}
//	return page, perPage
//}
