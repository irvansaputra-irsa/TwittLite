package comment

import (
	"twittlite/databases/connection"
	"twittlite/helpers/common"
	"twittlite/middlewares"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/comments")
	api.Use(middlewares.VerifyToken())
	{
		api.POST("", CreateCommentRouter)
		api.GET(":id", GetDetailCommentRouter)
		api.GET("post/:id", GetPostCommentsRouter)
		api.GET("user/:id", GetUserCommentsRouter)
		api.DELETE(":id", DeleteCommentRouter)
		api.PUT("", UpdateCommentRouter)
	}
}

func CreateCommentRouter(ctx *gin.Context) {
	var (
		cmntepo  = NewRepository(connection.DBConnections)
		cmntServ = NewService(cmntepo)
	)
	err := cmntServ.CreateCommentService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully create new comment")
}

func GetPostCommentsRouter(ctx *gin.Context) {
	var (
		cmntepo  = NewRepository(connection.DBConnections)
		cmntServ = NewService(cmntepo)
	)
	res, err := cmntServ.GetPostCommentsService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	total := int64(len(res))
	common.GenerateSuccessResponseWithListData(ctx, "successfully get post comments", total, res)
}

func GetUserCommentsRouter(ctx *gin.Context) {
	var (
		cmntepo  = NewRepository(connection.DBConnections)
		cmntServ = NewService(cmntepo)
	)
	res, err := cmntServ.GetUserCommentsService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	total := int64(len(res))
	common.GenerateSuccessResponseWithListData(ctx, "successfully get user comments", total, res)
}

func DeleteCommentRouter(ctx *gin.Context) {
	var (
		cmntepo  = NewRepository(connection.DBConnections)
		cmntServ = NewService(cmntepo)
	)
	err := cmntServ.DeleteCommentService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponse(ctx, "successfully delete comment")
}

func UpdateCommentRouter(ctx *gin.Context) {
	var (
		cmntepo  = NewRepository(connection.DBConnections)
		cmntServ = NewService(cmntepo)
	)

	err := cmntServ.UpdateCommentService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponse(ctx, "successfully update comment")
}

func GetDetailCommentRouter(ctx *gin.Context) {
	var (
		cmntepo  = NewRepository(connection.DBConnections)
		cmntServ = NewService(cmntepo)
	)
	res, err := cmntServ.GetDetailCommentService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get detail comments", res)
}
