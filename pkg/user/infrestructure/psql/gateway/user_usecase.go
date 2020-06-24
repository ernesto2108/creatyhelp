package user

import (
	"github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain"
)

func (s UsersCreateGtw) Create(cmd *domain.CreateUserCmd) (*domain.User,error) {
	return s.create(cmd)
}

func (s UsersCreateGtw) Update(u *domain.UpdateUserCmd) *domain.User {
	return s.update(u)
}

func (s UsersCreateGtw) GetId(id int64) (*domain.User, error) {
	return s.getId(id)
}

func (s UsersCreateGtw) Delete(id int64) *domain.User {
	return s.delete(id)
}

func (s UsersCreateGtw) GetAll() []*domain.User {
	return s.getAll()
}

