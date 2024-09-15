package post

import (
	"errors"
	"strconv"
	"twittlite/middlewares"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreatePostService(ctx *gin.Context) (err error)
	GetUserPostsService(ctx *gin.Context) (res []PostResponse, err error)
	UpdatePostService(ctx *gin.Context) (err error)
	DeletePostService(ctx *gin.Context) (err error)
	GetDetailPostService(ctx *gin.Context) (result PostResponseWithUsername, err error)
	GetTimelineService(ctx *gin.Context) (result []PostResponseWithUsername, err error)
}

type postService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &postService{
		repository,
	}
}

func (s *postService) CreatePostService(ctx *gin.Context) (err error) {
	var newPost PostRequest
	err = ctx.ShouldBind(&newPost)
	if err != nil {
		return err
	}

	userId, err := middlewares.EncryptToken(ctx)
	if err != nil {
		return err
	}
	newPost.UserId = userId

	err = s.repository.CreatePostRepository(newPost)
	if err != nil {
		return err
	}
	return nil
}

func (s *postService) GetUserPostsService(ctx *gin.Context) (res []PostResponse, err error) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return res, err
	}
	res, err = s.repository.GetUserPostsRepository(userId)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *postService) UpdatePostService(ctx *gin.Context) (err error) {
	var post PostUpdateRequest
	err = ctx.ShouldBind(&post)
	if err != nil {
		return err
	}

	currentPost, err := s.repository.CheckPostRepository(post.Id)
	if err != nil {
		return err
	}

	// check authorization of the post owner
	userId, err := middlewares.EncryptToken(ctx)
	if err != nil {
		return err
	}

	if userId != currentPost.UserId {
		return errors.New("cannot update post, you are not the owner")
	}

	return s.repository.UpdatePostRepository(post)
}

func (s *postService) DeletePostService(ctx *gin.Context) (err error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	currentPost, err := s.repository.CheckPostRepository(id)
	if err != nil {
		return err
	}

	// check authorization of the post owner
	userId, err := middlewares.EncryptToken(ctx)
	if err != nil {
		return err
	}

	if userId != currentPost.UserId {
		return errors.New("cannot delete post, you are not the owner")
	}

	return s.repository.DeletePostRepository(id)
}

func (s *postService) GetDetailPostService(ctx *gin.Context) (result PostResponseWithUsername, err error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return result, err
	}
	return s.repository.CheckPostRepository(id)
}

func (s *postService) GetTimelineService(ctx *gin.Context) (result []PostResponseWithUsername, err error) {
	userId, err := middlewares.EncryptToken(ctx)
	if err != nil {
		return result, err
	}
	return s.repository.GetTimelineRepository(userId)
}
