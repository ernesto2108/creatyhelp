package user

import (
	model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"
)

func (s UsersCreateGtw) Create(cmd *model.CreateUserCmd) (*model.User ,error) {
	return s.create(cmd)
}
