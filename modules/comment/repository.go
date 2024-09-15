package comment

import (
	"database/sql"
	"errors"
	"time"
	"twittlite/helpers/constant"
)

type Repository interface {
	CreateCommentRepository(uId int, c Comment) (err error)
	GetPostCommentsRepository(pId int) (comments []Comment, err error)
	DeleteCommentRepository(commentId int) (err error)
	CheckCommentRepository(cId int) (comment Comment, err error)
	UpdateCommentRepository(c Comment) (err error)
	GetUserCommentsRepository(uId int) (comments []Comment, err error)
}

type commentRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &commentRepository{
		db: database,
	}
}

func (r *commentRepository) CreateCommentRepository(uId int, c Comment) (err error) {
	sqlStmt := "INSERT INTO " + constant.CommentTableName.String() + " (content, user_id, post_id) VALUES ($1, $2, $3)"

	params := []interface{}{
		c.Content,
		uId,
		c.PostId,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return err
	}

	return nil
}

func (r *commentRepository) GetPostCommentsRepository(pId int) (comments []Comment, err error) {
	sqlStmt := "SELECT id, content, user_id, post_id, created_at, modified_at FROM " + constant.CommentTableName.String() + " WHERE post_id = $1 ORDER BY created_at ASC"

	rows, err := r.db.Query(sqlStmt, pId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var _comment Comment
		err := rows.Scan(&_comment.Id, &_comment.Content, &_comment.UserId, &_comment.PostId, &_comment.CreatedAt, &_comment.ModifiedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, _comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, errors.New("post comments is empty")
	}

	return comments, nil
}

func (r *commentRepository) GetUserCommentsRepository(uId int) (comments []Comment, err error) {
	sqlStmt := "SELECT id, content, user_id, post_id, created_at, modified_at FROM " + constant.CommentTableName.String() + " WHERE user_id = $1 ORDER BY created_at DESC"

	rows, err := r.db.Query(sqlStmt, uId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var _comment Comment
		err := rows.Scan(&_comment.Id, &_comment.Content, &_comment.UserId, &_comment.PostId, &_comment.CreatedAt, &_comment.ModifiedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, _comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, errors.New("user comments is empty")
	}

	return comments, nil
}

func (r *commentRepository) DeleteCommentRepository(commentId int) (err error) {
	sqlStmt := "DELETE FROM " + constant.CommentTableName.String() + " WHERE id = $1"
	_, err = r.db.Exec(sqlStmt, commentId)
	if err != nil {
		return err
	}
	return nil
}

func (r *commentRepository) CheckCommentRepository(cId int) (comment Comment, err error) {
	sqlStmt := "SELECT id, content, user_id, post_id, created_at, modified_at FROM " + constant.CommentTableName.String() + " WHERE id = $1"
	err = r.db.
		QueryRow(sqlStmt, cId).
		Scan(&comment.Id, &comment.Content, &comment.UserId, &comment.PostId, &comment.CreatedAt, &comment.ModifiedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return comment, errors.New("comment is not exist")
		}
		return comment, err
	}
	return comment, nil
}

func (r *commentRepository) UpdateCommentRepository(c Comment) (err error) {
	sqlStmt := "UPDATE " + constant.CommentTableName.String() + " SET content = $1, modified_at = $2 WHERE id = $3"
	params := []interface{}{
		c.Content,
		time.Now(),
		c.Id,
	}
	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return err
	}

	return nil
}
