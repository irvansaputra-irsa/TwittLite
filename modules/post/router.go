package post

import (
	"twittlite/databases/connection"
	"twittlite/helpers/common"
	"twittlite/middlewares"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/posts")
	api.Use(middlewares.VerifyToken())
	{
		api.POST("", CreatePostRouter)
		api.GET("user/:id", GetUserPostsRouter)
		api.PUT("", UpdatePostRouter)
		api.DELETE(":id", DeletePostRouter)
		api.GET(":id", GetDetailPostRouter)
		api.GET("timeline", GetTimelineRouter)
	}
}

func CreatePostRouter(ctx *gin.Context) {
	var (
		postRepo = NewRepository(connection.DBConnections)
		postServ = NewService(postRepo)
	)
	err := postServ.CreatePostService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully create new post")
}

func GetUserPostsRouter(ctx *gin.Context) {
	var (
		postRepo = NewRepository(connection.DBConnections)
		postServ = NewService(postRepo)
	)
	res, err := postServ.GetUserPostsService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	total := int64(len(res))
	common.GenerateSuccessResponseWithListData(ctx, "successfully get user posts", total, res)
}

func UpdatePostRouter(ctx *gin.Context) {
	var (
		postRepo = NewRepository(connection.DBConnections)
		postServ = NewService(postRepo)
	)

	err := postServ.UpdatePostService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponse(ctx, "successfully update post")
}

func DeletePostRouter(ctx *gin.Context) {
	var (
		postRepo = NewRepository(connection.DBConnections)
		postServ = NewService(postRepo)
	)
	err := postServ.DeletePostService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponse(ctx, "successfully delete post")

}

func GetDetailPostRouter(ctx *gin.Context) {
	var (
		postRepo = NewRepository(connection.DBConnections)
		postServ = NewService(postRepo)
	)

	res, err := postServ.GetDetailPostService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "successfully get book data", res)
}

func GetTimelineRouter(ctx *gin.Context) {
	var (
		postRepo = NewRepository(connection.DBConnections)
		postServ = NewService(postRepo)
	)

	res, err := postServ.GetTimelineService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	total := int64(len(res))

	common.GenerateSuccessResponseWithListData(ctx, "successfully get book data", total, res)
}
