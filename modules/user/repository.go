package user

import (
	"database/sql"
	"errors"
	"twittlite/helpers/constant"
)

type Repository interface {
	RegisterRepository(u User) (err error)
	LoginRepository(u LoginRequest) (result User, err error)
	GetDetailUserRepository(uId int) (result UserProfileCheck, err error)
	UpdateProfileRepository(u UserUpdateProfile) (err error)
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &userRepository{
		db: database,
	}
}

func (r *userRepository) RegisterRepository(u User) (err error) {
	sqlStmt := "INSERT INTO " + constant.UserTableName.String() + " (username, password, email) VALUES ($1, $2, $3)"

	params := []interface{}{
		u.Username,
		u.Password,
		u.Email,
	}
	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) LoginRepository(u LoginRequest) (result User, err error) {
	sqlStmt := "SELECT id, username, email, password FROM " + constant.UserTableName.String() + " WHERE email = $1"
	err = r.db.QueryRow(sqlStmt, u.Email).Scan(&result.Id, &result.Username, &result.Email, &result.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New("invalid email")
		}
		return result, err
	}
	return result, nil
}

func (r *userRepository) GetDetailUserRepository(uId int) (result UserProfileCheck, err error) {
	sqlStmt := "SELECT id, username, bio, location FROM " + constant.UserTableName.String() + " WHERE id = $1"
	err = r.db.QueryRow(sqlStmt, uId).Scan(&result.Id, &result.Username, &result.Bio, &result.Location)

	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New("user is not exist")
		}
		return result, err
	}
	return result, nil
}

func (r *userRepository) UpdateProfileRepository(u UserUpdateProfile) (err error) {
	sqlStmt := "UPDATE " + constant.UserTableName.String() + " SET bio = $1, location = $2 WHERE id = $3"
	params := []interface{}{
		u.Bio,
		u.Location,
		u.Id,
	}
	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return err
	}

	return nil
}
