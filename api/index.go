package main

import (
	"twittlite/configs"
	"twittlite/databases/connection"
	migration "twittlite/migrations"
	"twittlite/modules/comment"
	"twittlite/modules/follow"
	"twittlite/modules/post"
	"twittlite/modules/user"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Initiator()
	connection.Initiator()
	defer connection.DBConnections.Close()
	migration.Initiator(connection.DBConnections)

	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()

	user.Initiator(router)
	post.Initiator(router)
	follow.Initiator(router)
	comment.Initiator(router)

	router.Run(":8080")
}
