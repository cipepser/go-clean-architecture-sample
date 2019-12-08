package usecases

import (
	"errors"
	"github.com/cipepser/go-clean-architecture-sample/entities"
)

type (
	UserRepository interface {
		FindByEmail(string) (entities.User, error)
	}

	UserInputBoundary interface {
		Login(*LoginInput)
	}

	UserOutputBoundary interface {
		Login(*LoginOutput, error)
	}
)

func NewUser(output UserOutputBoundary, repo UserRepository) UserInputBoundary {
	return &UserInteractor{
		UserOutput:     output,
		UserRepository: repo,
	}
}

type UserInteractor struct {
	UserOutput     UserOutputBoundary
	UserRepository UserRepository
}

func (ui *UserInteractor) Login(input *LoginInput) {
	user, err := ui.UserRepository.FindByEmail(input.Email)
	if err != nil {
		ui.UserOutput.Login(&LoginOutput{}, err)
		return
	}

	if ok := user.PasswordVerify(input.Password); !ok {
		ui.UserOutput.Login(&LoginOutput{}, errors.New("verify failed"))
		return
	}

	output := LoginOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	ui.UserOutput.Login(&output, nil)
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	ID    string
	Name  string
	Email string
}
