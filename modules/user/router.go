package user

import (
	"twittlite/databases/connection"
	"twittlite/helpers/common"
	"twittlite/middlewares"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/users")
	{
		api.POST("/register", RegisterRouter)
		api.POST("/login", LoginRouter)
		api.GET(":id", middlewares.VerifyToken(), GetDetailUserRouter)
	}
}

func RegisterRouter(ctx *gin.Context) {
	var (
		userRepo = NewRepository(connection.DBConnections)
		userServ = NewService(userRepo)
	)

	err := userServ.RegisterService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully register new user")
}

func LoginRouter(ctx *gin.Context) {
	var (
		userRepo = NewRepository(connection.DBConnections)
		userServ = NewService(userRepo)
	)

	res, err := userServ.LoginService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully login user", res)
}

func GetDetailUserRouter(ctx *gin.Context) {
	var (
		userRepo = NewRepository(connection.DBConnections)
		userServ = NewService(userRepo)
	)

	res, err := userServ.GetDetailUserService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "successfully get user data", res)
}
