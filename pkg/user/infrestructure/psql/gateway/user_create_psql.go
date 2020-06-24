package user

import (
	"database/sql"
	logs "github.com/ernesto2108/AP_CreatyHelp/internal/logs"
	"github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain"
)

func (s UsersStorage) create(u *domain.CreateUserCmd) (*domain.User, error) {
	tx, err := s.PostSqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil, err
	}

	// Int Null
	intNull := sql.NullInt64{}

	if u.Phone == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int64 = int64(u.Phone)
	}

	//	String Null
	strNull := sql.NullString{}

	if u.Name == " " {
		strNull.Valid = false
	}else {
		strNull.Valid = true
		strNull.String = u.Name
	}

	if u.Nickname == " " {
		strNull.Valid = false
	} else {
		strNull.Valid = true
		strNull.String = u.Nickname
	}

	query, err := tx.Exec("INSERT INTO users (name, nickname, phone) VALUES ($1, $2, $3)",
		strNull, strNull, intNull)

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

	return &domain.User{
		ID:       id,
		Name:     u.Name,
		Nickname: u.Nickname,
		Phone:    u.Phone,
	}, nil
}