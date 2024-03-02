package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/kit/kv"
	"my-user-settings-service/internal/config"
	"my-user-settings-service/internal/data/pg"
	"my-user-settings-service/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	cfg := config.New(kv.MustFromEnv())

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxUsersQ(pg.NewUsersQ(cfg.DB())),
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
