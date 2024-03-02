package handlers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"my-user-settings-service/internal/data"
	"net/http"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	usersCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxUsersQ(db data.UsersQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, usersCtxKey, db)
	}
}

func UsersQ(r *http.Request) data.UsersQ {
	return r.Context().Value(usersCtxKey).(data.UsersQ).New()
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}
