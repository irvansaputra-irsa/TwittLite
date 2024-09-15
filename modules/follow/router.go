package follow

import (
	"twittlite/databases/connection"
	"twittlite/helpers/common"
	"twittlite/middlewares"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/follows")
	api.Use(middlewares.VerifyToken())
	{
		api.POST(":id", FollowRouter)
		api.GET("following/:id", GetFollowingListRouter)
		api.GET("follower/:id", GetFollowerListRouter)
	}
}

func FollowRouter(ctx *gin.Context) {
	var (
		followRepo = NewRepository(connection.DBConnections)
		followServ = NewService(followRepo)
	)
	err := followServ.FollowService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully following new user")
}

func GetFollowingListRouter(ctx *gin.Context) {
	var (
		followRepo = NewRepository(connection.DBConnections)
		followServ = NewService(followRepo)
	)
	res, err := followServ.GetFollowingListService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	total := int64(len(res))
	common.GenerateSuccessResponseWithListData(ctx, "successfully get following list", total, res)
}

func GetFollowerListRouter(ctx *gin.Context) {
	var (
		followRepo = NewRepository(connection.DBConnections)
		followServ = NewService(followRepo)
	)
	res, err := followServ.GetFollowerListService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	total := int64(len(res))
	common.GenerateSuccessResponseWithListData(ctx, "successfully get follower list", total, res)
}
