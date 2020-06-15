package user

import (
	internal "github.com/ernesto2108/AP_CreatyHelp/internal/storage/psql"
	model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"
)

type UsersCreateGateway interface {
	Create(u *model.CreateUserCmd) (*model.User ,error)
	GetId(id int64) (*model.User, error)
}

type UsersCreateGtw struct {
	UsersStorageGateway
}

func NewUsersCreateGateway(c *internal.PostSqlClient) UsersCreateGateway {
	return &UsersCreateGtw{&UsersStorage{c}}
}

