package presenter

import "github.com/cipepser/go-clean-architecture-sample/usecases"

type Context interface {
	JSON(int, interface{}) error
	Error(int)
}

type User struct {
	Context Context
}

func (u *User) Login(output *usecases.LoginOutput, err error) {
	if err != nil {
		u.Context.Error(404)
	}
	u.Context.JSON(200, output)
}
