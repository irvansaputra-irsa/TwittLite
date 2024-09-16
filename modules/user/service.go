package user

import (
	"errors"
	"strconv"
	"twittlite/helpers/common"
	"twittlite/middlewares"

	"github.com/gin-gonic/gin"
)

type Service interface {
	RegisterService(ctx *gin.Context) (err error)
	LoginService(ctx *gin.Context) (result *LoginResponse, err error)
	GetDetailUserService(ctx *gin.Context) (result UserProfileCheck, err error)
	UpdateProfileService(ctx *gin.Context) (err error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (service *userService) RegisterService(ctx *gin.Context) (err error) {
	var newUser User
	err = ctx.ShouldBind(&newUser)
	if err != nil {
		return err
	}

	hashedPassword, err := common.HashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = hashedPassword

	err = service.repository.RegisterRepository(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (service *userService) LoginService(ctx *gin.Context) (result *LoginResponse, err error) {
	var userReq LoginRequest
	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return result, err
	}

	user, err := service.repository.LoginRepository(userReq)
	if err != nil {
		return result, err
	}

	match := common.VerifyPassword(userReq.Password, user.Password)
	if !match {
		err = errors.New("invalid credential")
		return
	}

	token, err := middlewares.CreateToken(user.Id)
	if err != nil {
		return result, err
	}

	return &LoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}, nil
}

func (s *userService) GetDetailUserService(ctx *gin.Context) (result UserProfileCheck, err error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return result, err
	}
	return s.repository.GetDetailUserRepository(id)
}

func (s *userService) UpdateProfileService(ctx *gin.Context) (err error) {
	var user UserUpdateProfile
	err = ctx.ShouldBind(&user)
	if err != nil {
		return err
	}

	// get authorization of the owner
	userId, err := middlewares.EncryptToken(ctx)
	if err != nil {
		return err
	}
	user.Id = userId

	return s.repository.UpdateProfileRepository(user)
}
