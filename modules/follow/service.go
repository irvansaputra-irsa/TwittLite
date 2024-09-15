package follow

import (
	"errors"
	"strconv"
	"twittlite/middlewares"

	"github.com/gin-gonic/gin"
)

type Service interface {
	FollowService(ctx *gin.Context) (err error)
	GetFollowingListService(ctx *gin.Context) (res []FollowingList, err error)
	GetFollowerListService(ctx *gin.Context) (res []FollowerList, err error)
}

type followService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &followService{
		repository,
	}
}

func (s *followService) FollowService(ctx *gin.Context) (err error) {
	followingId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	userId, err := middlewares.EncryptToken(ctx)
	if err != nil {
		return err
	}

	//validate if user cannot follow themself
	if followingId == userId {
		return errors.New("you cannot follow yourself, please follow another")
	}

	var follows = Follow{
		FollowerId:  userId,
		FollowingId: followingId,
	}

	//validate if user already followed target user id
	checkFollow, err := s.repository.IsAlreadyFollowRepository(follows)
	if err != nil {
		return err
	}

	if checkFollow.Id != 0 {
		return errors.New(checkFollow.FollowerUsername + " already following " + checkFollow.FollowingUsername)
	}

	return s.repository.FollowRepository(follows)
}

func (s *followService) GetFollowingListService(ctx *gin.Context) (res []FollowingList, err error) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return res, err
	}
	res, err = s.repository.GetFollowingListRepository(userId)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *followService) GetFollowerListService(ctx *gin.Context) (res []FollowerList, err error) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return res, err
	}
	res, err = s.repository.GetFollowerListRepository(userId)
	if err != nil {
		return res, err
	}

	return res, nil
}
