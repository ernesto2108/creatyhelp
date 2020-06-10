package user

import (
	internal "github.com/ernesto2108/AP_CreatyHelp/internal/storage/psql"
	model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"
	ift "github.com/ernesto2108/AP_CreatyHelp/pkg/user/infrestructure/psql/gateway"
)

type UsersCreateGateway interface {
	Create(u *model.CreateUserCmd) (*model.User ,error)
}

type UsersCreateGtw struct {
	ift.UsersStorageGateway
}

func NewUsersCreateGateway(c *internal.PostSqlClient) UsersCreateGateway {
	return &UsersCreateGtw{&ift.UsersStorage{c}}
}

