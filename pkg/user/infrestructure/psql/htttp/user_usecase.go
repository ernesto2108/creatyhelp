package user

import model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"

func (s UsersCreateGtw) Create(u *model.CreateUserCmd) (*model.User ,error) {
	return s.Create(u)
}
