package application

import (
	"context"

	"github.com/aboglioli/configd/domain/user"
)

type RegisterUserCommand struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Access   *string `json:"access"`
}

type RegisterUserResponse struct {
	Username string `json:"username"`
	Access   string `json:"access"`
}

type RegisterUser struct {
	userRepo user.UserRepository
}

func NewRegisterUser(
	userRepo user.UserRepository,
) *RegisterUser {
	return &RegisterUser{
		userRepo: userRepo,
	}
}

func (uc *RegisterUser) Exec(
	ctx context.Context,
	cmd *RegisterUserCommand,
) (*RegisterUserResponse, error) {
	username, err := user.NewUsername(cmd.Username)
	if err != nil {
		return nil, err
	}

	password, err := user.NewPassword(cmd.Password)
	if err != nil {
		return nil, err
	}

	access := user.READ_ONLY_ACCESS
	if cmd.Access != nil {
		access, err = user.NewAccess(*cmd.Access)
		if err != nil {
			return nil, err
		}
	}

	u, err := user.NewUser(
		username,
		password,
		access,
	)
	if err != nil {
		return nil, err
	}

	if err := uc.userRepo.Save(ctx, u); err != nil {
		return nil, err
	}

	return &RegisterUserResponse{
		Username: u.Username().Value(),
		Access:   string(u.Access()),
	}, nil
}
