package user

import (
	internal "github.com/ernesto2108/AP_CreatyHelp/internal/storage/psql"
	model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"
)

type UsersStorageGateway interface {
	create(u *model.CreateUserCmd) (*model.User ,error)
	update(id int64) (*model.User, error)
	getId(id int64) (*model.User, error)
	delete(id int64) error
	getAll() []*model.User
}

type UsersStorage struct {
	*internal.PostSqlClient
}
