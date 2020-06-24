package user

import (
	logs "github.com/ernesto2108/AP_CreatyHelp/internal/logs"
	"github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain"
	"github.com/lib/pq"
)

func (s UsersStorage) getId(id int64) (*domain.User, error) {
	tx, err := s.PostSqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot search transaction")
		return nil, err
	}

	timeNull := pq.NullTime{}

	var user domain.User

	err = tx.QueryRow(`SELECT id, name, nickname, phone, created_at, updated_at FROM users 
		WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Nickname, &user.Phone, &user.CreatedAt, &timeNull)

	if err != nil {
		logs.Log().Error("cannot execute statement")
		return nil, err
	}

	return &user, nil
}
