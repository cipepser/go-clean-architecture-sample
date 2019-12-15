package frameWorksAndDrivers

import (
	"encoding/json"
	"github.com/cipepser/go-clean-architecture-sample/interfaceAdapter/presenter"
	controllers "github.com/cipepser/go-clean-architecture-sample/interfaceAdapter/user"
	"github.com/cipepser/go-clean-architecture-sample/usecases"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type context struct {
	w http.ResponseWriter
	r *http.Request
}

func (c context) Error(status int) {
	http.Error(c.w, http.StatusText(status), status)
}

func (c context) Bind(v interface{}) error {
	return json.NewDecoder(c.r.Body).Decode(&v)
}

func (c context) JSON(status int, v interface{}) error {
	res, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.w.WriteHeader(status)
	c.w.Header().Set("Content-Type", "application/json")
	c.w.Write(res)
	return nil
}

func NewContext(w http.ResponseWriter, r *http.Request) context {
	return context{
		w: w,
		r: r,
	}
}

var Router chi.Router

func init() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			ctx := NewContext(w, r)
			userController := controllers.NewUserController(&presenter.User{Context: ctx}, &usecases.UserInteractor{})
			userController.Login(ctx)
		})
	})

}
