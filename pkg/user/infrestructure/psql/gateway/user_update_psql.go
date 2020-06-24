package user

import (
	logs "github.com/ernesto2108/AP_CreatyHelp/internal/logs"
	"github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain"
)

func (s UsersStorage) update(u *domain.UpdateUserCmd) *domain.User {
	tx, err := s.PostSqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot update transaction")
		return nil
	}

	_, err = tx.Exec(`UPDATE users SET name=$1 ,nickname=$2 ,phone=$3, updated_at = NOW()
		WHERE id = $4`, u.Name, u.Nickname, u.Phone, u.ID)

	if err != nil {
		logs.Log().Error("Unable to execute query. %v")
		_ = tx.Rollback()
		return nil
	}

	_ = tx.Commit()

	return &domain.User{
		ID: 		u.ID,
		Name: 		u.Name,
		Nickname: 	u.Nickname,
		Phone: 		u.Phone,
	}
}
