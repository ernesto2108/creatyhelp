package user

import (
	internal "github.com/ernesto2108/AP_CreatyHelp/internal/storage/psql"
	"github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain"
)

type UsersCreateGateway interface {
	Create(u *domain.CreateUserCmd) (*domain.User,error)
	Update(u *domain.UpdateUserCmd) *domain.User
	GetId(id int64) (*domain.User, error)
	Delete(id int64) *domain.User
	GetAll() []*domain.User
}

type UsersCreateGtw struct {
	UsersStorageGateway
}

func NewUsersCreateGateway(c *internal.PostSqlClient) UsersCreateGateway {
	return &UsersCreateGtw{&UsersStorage{c}}
}

