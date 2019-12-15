package controllers

import "github.com/cipepser/go-clean-architecture-sample/usecases"

type Context interface {
	Bind(interface{}) error
}

type UserController struct {
	Interactoer usecases.UserInputBoundary
}

func (controller *UserController) Login(c Context) {
	type userLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	req := userLoginRequest{}
	c.Bind(&req)

	input := usecases.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	controller.Interactoer.Login(&input)
}

func NewUserController(out usecases.UserOutputBoundary, repo usecases.UserRepository) *UserController {
	interactoer := usecases.NewUser(out, repo)
	return &UserController{
		Interactoer: interactoer,
	}
}
