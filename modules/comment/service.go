package comment

import (
	"errors"
	"strconv"
	"twittlite/middlewares"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateCommentService(ctx *gin.Context) (err error)
	GetPostCommentsService(ctx *gin.Context) (result []Comment, err error)
	DeleteCommentService(ctx *gin.Context) (err error)
	UpdateCommentService(ctx *gin.Context) (err error)
	GetUserCommentsService(ctx *gin.Context) (result []Comment, err error)
	GetDetailCommentService(ctx *gin.Context) (result Comment, err error)
}

type commentService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &commentService{
		repository,
	}
}

func (s *commentService) CreateCommentService(ctx *gin.Context) (err error) {
	var comment Comment
	err = ctx.ShouldBind(&comment)
	if err != nil {
		return err
	}

	userId, err := middlewares.EncryptToken(ctx)
	if err != nil {
		return err
	}

	return s.repository.CreateCommentRepository(userId, comment)
}

func (s *commentService) GetPostCommentsService(ctx *gin.Context) (result []Comment, err error) {
	commentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, err
	}
	return s.repository.GetPostCommentsRepository(commentId)
}

func (s *commentService) GetUserCommentsService(ctx *gin.Context) (result []Comment, err error) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, err
	}
	return s.repository.GetUserCommentsRepository(userId)
}

func (s *commentService) DeleteCommentService(ctx *gin.Context) (err error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	currentComment, err := s.repository.CheckCommentRepository(id)
	if err != nil {
		return err
	}

	// check authorization of the post owner
	userId, err := middlewares.EncryptToken(ctx)
	if err != nil {
		return err
	}

	if userId != currentComment.UserId {
		return errors.New("cannot delete comment, you are not the owner")
	}

	return s.repository.DeleteCommentRepository(id)
}

func (s *commentService) UpdateCommentService(ctx *gin.Context) (err error) {
	var comment Comment
	err = ctx.ShouldBind(&comment)
	if err != nil {
		return err
	}

	currentComment, err := s.repository.CheckCommentRepository(comment.Id)
	if err != nil {
		return err
	}

	// check authorization of the post owner
	userId, err := middlewares.EncryptToken(ctx)
	if err != nil {
		return err
	}

	if userId != currentComment.UserId {
		return errors.New("cannot update comment, you are not the owner")
	}

	return s.repository.UpdateCommentRepository(comment)
}

func (s *commentService) GetDetailCommentService(ctx *gin.Context) (result Comment, err error) {
	commentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return result, err
	}
	return s.repository.CheckCommentRepository(commentId)
}
