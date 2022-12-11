package users

import "github.com/gin-gonic/gin"

func RouterUsers(r *gin.RouterGroup) {
	r.GET("/", usersIndex)

	//r.GET("/list", getAllUser)
	//r.GET("/user/:addr", getUser)
	//r.POST("/user", createUser)
}
