package user

import (
	internal "github.com/ernesto2108/AP_CreatyHelp/internal/storage/psql"
	model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"
)

type UsersStorageGateway interface {
	create(u *model.CreateUserCmd) (*model.User ,error)
	update(u *model.UpdateUserCmd) *model.User
	getId(id int64) (*model.User, error)
	delete(id int64) *model.User
	getAll() []*model.User
}

type UsersStorage struct {
	*internal.PostSqlClient
}
