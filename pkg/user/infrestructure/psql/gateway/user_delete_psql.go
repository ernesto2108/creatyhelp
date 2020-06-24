package user

import (
	logs "github.com/ernesto2108/AP_CreatyHelp/internal/logs"
	"github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain"
)

func (s UsersStorage) delete(id int64) *domain.User {
	tx, err := s.PostSqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil
	}
	var user domain.User

	err = tx.QueryRow(`SELECT id, name, nickname, phone FROM users 
		WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Nickname, &user.Phone)

	if err != nil {
		logs.Log().Error("Unable to execute query. %v")
		_ = tx.Rollback()
		return nil
	}

	_, err = tx.Exec("DELETE FROM users WHERE id = $1", id)

	if err != nil {
		logs.Log().Error("Unable to execute query. %v")
		_ = tx.Rollback()
		return nil
	}

	_ = tx.Commit()

	return &domain.User{
		ID: 			id,
		Name: 			user.Name,
		Nickname: 		user.Nickname,
		Phone: 			user.Phone,
	}
}
