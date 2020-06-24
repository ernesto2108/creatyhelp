package user

import (
	internal "github.com/ernesto2108/AP_CreatyHelp/internal/storage/psql"
	"github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain"
)

type UsersStorageGateway interface {
	create(u *domain.CreateUserCmd) (*domain.User,error)
	update(u *domain.UpdateUserCmd) *domain.User
	getId(id int64) (*domain.User, error)
	delete(id int64) *domain.User
	getAll() []*domain.User
}

type UsersStorage struct {
	*internal.PostSqlClient
}
