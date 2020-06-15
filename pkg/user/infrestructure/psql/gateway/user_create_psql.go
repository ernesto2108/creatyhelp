package user

import (
	logs "github.com/ernesto2108/AP_CreatyHelp/internal/logs"
	model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"
)

func (s UsersStorage) create(u *model.CreateUserCmd) (*model.User, error) {
	tx, err := s.PostSqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil, err
	}

	query, err := tx.Exec(`INSERT INTO users (name, nickname, phone) VALUES ($1, $2, $3)`,
		u.Name, u.Nickname, u.Phone)

	if err != nil {
		logs.Log().Error("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := query.RowsAffected()

	if err != nil {
		logs.Log().Error("cannot fetch last id")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &model.User{
		ID:       id,
		Name:     u.Name,
		Nickname: u.Nickname,
		Phone:    u.Phone,
	}, nil
}