package service

import (
	"my-user-settings-service/internal/service/handlers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/integrations/my-user-settings-service", func(r chi.Router) {
		r.Get("/profile", handlers.GetUsersHandler)
		r.Post("/profile", handlers.CreateUserHandler)
		r.Put("/profile/{id}", handlers.UpdateUserHandler)
		r.Delete("/profile/{id}", handlers.DeleteUserHandler)
	})

	return r
}
