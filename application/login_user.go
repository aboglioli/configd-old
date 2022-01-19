package application

import (
	"context"

	"github.com/aboglioli/configd/domain/user"
)

type LoginUserCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token string `json:"auth_token"`
}

type LoginUser struct {
	userRepo user.UserRepository
}

func NewLoginUser(
	userRepo user.UserRepository,
) *LoginUser {
	return &LoginUser{
		userRepo: userRepo,
	}
}

func (uc *LoginUser) Exec(
	ctx context.Context,
	cmd *LoginUserCommand,
) (*LoginUserResponse, error) {
	username, err := user.NewUsername(cmd.Username)
	if err != nil {
		return nil, err
	}

	password, err := user.NewPassword(cmd.Password)
	if err != nil {
		return nil, err
	}

	u, err := uc.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	token, err := u.Login(username, password)
	if err != nil {
		return nil, err
	}

	return &LoginUserResponse{
		Token: token.Value(),
	}, nil
}
