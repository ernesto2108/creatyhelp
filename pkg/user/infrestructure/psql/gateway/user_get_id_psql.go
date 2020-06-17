package user

import (
	logs "github.com/ernesto2108/AP_CreatyHelp/internal/logs"
	model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"
)

func (s UsersStorage) getId(id int64) (*model.User, error) {
	tx, err := s.PostSqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot search transaction")
		return nil, err
	}

	var user model.User

	err = tx.QueryRow(`SELECT id, name, nickname, phone FROM users 
		WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Nickname, &user.Phone)

	if err != nil {
		logs.Log().Error("cannot execute statement")
		return nil, err
	}

	return &user, nil
}
