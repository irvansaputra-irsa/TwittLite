package follow

import (
	"database/sql"
	"errors"
	"twittlite/helpers/constant"
)

type Repository interface {
	FollowRepository(f Follow) (err error)
	IsAlreadyFollowRepository(f Follow) (follow FollowWithName, err error)
	GetFollowingListRepository(userId int) (followings []FollowingList, err error)
	GetFollowerListRepository(userId int) (followers []FollowerList, err error)
}

type followRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &followRepository{
		db: database,
	}
}

func (r *followRepository) FollowRepository(f Follow) (err error) {
	sqlStmt := "INSERT INTO " + constant.FollowTableName.String() + " (follower_id, following_id) VALUES ($1, $2)"

	params := []interface{}{
		f.FollowerId,
		f.FollowingId,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (r *followRepository) IsAlreadyFollowRepository(f Follow) (follow FollowWithName, err error) {
	sqlStmt := "SELECT f.id, f.following_id, u.username, f.follower_id, us.username FROM " + constant.FollowTableName.String() + " f JOIN " + constant.UserTableName.String() + " u ON u.id = f.following_id JOIN " + constant.UserTableName.String() + " us ON f.follower_id = us.id WHERE f.following_id = $1 AND f.follower_id = $2"
	params := []interface{}{
		f.FollowingId,
		f.FollowerId,
	}

	err = r.db.QueryRow(sqlStmt, params...).Scan(&follow.Id, &follow.FollowingId, &follow.FollowingUsername, &follow.FollowerId, &follow.FollowerUsername)
	if err != nil {
		if err == sql.ErrNoRows {
			return follow, nil
		}
		return follow, err
	}
	return follow, nil
}

func (r *followRepository) GetFollowingListRepository(userId int) (followings []FollowingList, err error) {
	sqlStmt := "SELECT f.id, following_id, u.username FROM " + constant.FollowTableName.String() + " f JOIN " + constant.UserTableName.String() + " u ON f.following_id = u.id WHERE f.follower_id = $1"
	rows, err := r.db.Query(sqlStmt, userId)
	if err != nil {
		return followings, err
	}

	for rows.Next() {
		var usr FollowingList
		err := rows.Scan(&usr.Id, &usr.FollowingId, &usr.FollowingName)
		if err != nil {
			return followings, err
		}
		followings = append(followings, usr)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(followings) == 0 {
		return nil, errors.New("followings list is empty")
	}

	return followings, nil
}

func (r *followRepository) GetFollowerListRepository(userId int) (followers []FollowerList, err error) {
	sqlStmt := "SELECT f.id, follower_id, u.username FROM " + constant.FollowTableName.String() + " f JOIN " + constant.UserTableName.String() + " u ON f.follower_id = u.id WHERE f.following_id = $1"
	rows, err := r.db.Query(sqlStmt, userId)
	if err != nil {
		return followers, err
	}
	for rows.Next() {
		var usr FollowerList
		err := rows.Scan(&usr.Id, &usr.FollowerId, &usr.FollowerName)
		if err != nil {
			return followers, err
		}
		followers = append(followers, usr)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(followers) == 0 {
		return nil, errors.New("followers list is empty")
	}

	return followers, nil
}
