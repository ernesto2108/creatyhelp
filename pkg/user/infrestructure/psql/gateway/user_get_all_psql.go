package user

import (
	logs "github.com/ernesto2108/AP_CreatyHelp/internal/logs"
	model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"
)

func (s UsersStorage) getAll() []*model.User {
	tx, err := s.PostSqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil
	}

	query, err := tx.Query("SELECT id, name, nickname, phone FROM users")
	if err != nil {
		logs.Log().Error("cannot execute statement")
		_ = tx.Rollback()
		return nil
	}
	defer query.Close()

	var users []*model.User
	for query.Next(){
		var user model.User
		err := query.Scan(&user.ID, &user.Name, &user.Nickname, &user.Phone)
		if err != nil {
			logs.Log().Error("cannot read current row")
			return nil
		}
		users = append(users, &user)
	}

	return users
}
