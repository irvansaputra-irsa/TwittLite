package post

import (
	"database/sql"
	"errors"
	"time"
	"twittlite/helpers/constant"
)

type Repository interface {
	CreatePostRepository(p PostRequest) (err error)
	GetUserPostsRepository(idUser int) (result []PostResponse, err error)
	UpdatePostRepository(p PostUpdateRequest) (err error)
	DeletePostRepository(postId int) (err error)
	CheckPostRepository(cId int) (result PostResponseWithUsername, err error)
	GetTimelineRepository(uId int) (posts []PostResponseWithUsername, err error)
}

type postRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &postRepository{
		db: database,
	}
}

func (r *postRepository) CreatePostRepository(p PostRequest) (err error) {
	sqlStmt := "INSERT INTO " + constant.PostTableName.String() + " (content, user_id) VALUES ($1, $2)"
	params := []interface{}{
		p.Content,
		p.UserId,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return err
	}
	return nil
}

func (r *postRepository) GetUserPostsRepository(idUser int) (posts []PostResponse, err error) {
	sqlStmt := "SELECT id, content, user_id, created_at, modified_at FROM " + constant.PostTableName.String() + " WHERE user_id = $1 ORDER BY created_at DESC"

	rows, err := r.db.Query(sqlStmt, idUser)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var _post PostResponse
		err := rows.Scan(&_post.Id, &_post.Content, &_post.UserId, &_post.CreatedAt, &_post.ModifiedAt)

		if err != nil {
			return posts, err
		}
		posts = append(posts, _post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, errors.New("posts list is empty")
	}

	return posts, nil
}

func (r *postRepository) UpdatePostRepository(p PostUpdateRequest) (err error) {
	sqlStmt := "UPDATE " + constant.PostTableName.String() + " SET content = $1, modified_at = $2 WHERE id = $3"
	params := []interface{}{
		p.Content,
		time.Now(),
		p.Id,
	}
	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (r *postRepository) DeletePostRepository(postId int) (err error) {
	sqlStmt := "DELETE FROM " + constant.PostTableName.String() + " WHERE id = $1"
	_, err = r.db.Exec(sqlStmt, postId)
	if err != nil {
		return err
	}
	return nil
}

func (r *postRepository) CheckPostRepository(cId int) (result PostResponseWithUsername, err error) {
	sqlStmt := "SELECT p.id, content, user_id, username, p.created_at, p.modified_at FROM " + constant.PostTableName.String() + " p JOIN " + constant.UserTableName.String() + " usr ON p.user_id = usr.id WHERE p.id = $1"
	err = r.db.
		QueryRow(sqlStmt, cId).
		Scan(&result.Id, &result.Content, &result.UserId, &result.Username, &result.CreatedAt, &result.ModifiedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New("post is not exist")
		}
		return result, err
	}
	return result, nil
}

func (r *postRepository) GetTimelineRepository(uId int) (posts []PostResponseWithUsername, err error) {
	sqlStmt := "SELECT p.id, p.content, p.user_id, u.username, p.created_at, p.modified_at FROM " + constant.PostTableName.String() + " p LEFT JOIN " + constant.UserTableName.String() + " u ON p.user_id = u.id LEFT JOIN " + constant.FollowTableName.String() + " f ON f.following_id = p.user_id WHERE f.follower_id = $1  OR p.user_id = $1 ORDER BY p.created_at DESC"

	rows, err := r.db.Query(sqlStmt, uId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post PostResponseWithUsername
		err := rows.Scan(&post.Id, &post.Content, &post.UserId, &post.Username, &post.CreatedAt, &post.ModifiedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return nil, errors.New("timeline posts list is empty")
	}

	return posts, nil
}
