package user

import (
	model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"
)

func (s UsersCreateGtw) Create(cmd *model.CreateUserCmd) (*model.User ,error) {
	return s.create(cmd)
}

func (s UsersCreateGtw) Update(id int64) (*model.User, error) {
	return s.update(id)
}

func (s UsersCreateGtw) GetId(id int64) (*model.User, error) {
	return s.getId(id)
}

func (s UsersCreateGtw) Delete(id int64) error {
	return s.delete(id)
}

func (s UsersCreateGtw) GetAll() []*model.User {
	return s.getAll()
}

