package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sebastianflor/golang-postgresql-api/internal/data"
)

// New inicialize a new router with configuration.
func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())

	pr := &PostRouter{
		Repository: &data.PostRepository{
			Data: data.New(),
		},
	}

	r.Mount("/posts", pr.Routes())

	return r
}
