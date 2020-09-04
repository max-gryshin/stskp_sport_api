package auth_service

import (
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/repository"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	user, err := repository.FindUserByUsername(a.Username)
	if err != nil {
		return false, err
	}

	return user.Username == "maxim", nil
}
