package user

import (
	logs "github.com/ernesto2108/AP_CreatyHelp/internal/logs"
	"github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain"
	"github.com/lib/pq"
)

func (s UsersStorage) getAll() []*domain.User {
	tx, err := s.PostSqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil
	}

	timeNull := pq.NullTime{}

	query, err := tx.Query("SELECT id, name, nickname, phone, created_at, updated_at FROM users")
	if err != nil {
		logs.Log().Error("cannot search statement")
		_ = tx.Rollback()
		return nil
	}
	defer query.Close()

	var users []*domain.User
	for query.Next(){
		var user domain.User
		err := query.Scan(&user.ID, &user.Name, &user.Nickname, &user.Phone, &user.CreatedAt, &timeNull)
		if err != nil {
			logs.Log().Error("cannot read current row")
			return nil
		}
		users = append(users, &user)
	}

	return users
}
