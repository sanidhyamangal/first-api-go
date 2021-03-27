// Author: Sanidhya Mangal
// github: sanidhyamangal
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sanidhya/firstapi/controllers"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", controllers.GetUsers)
		grp1.POST("user", controllers.CreateUser)
		grp1.GET("user/:id", controllers.GetUserByID)
		grp1.PUT("user/:id", controllers.UpdateUser)
		grp1.DELETE("user/:id", controllers.DeleteUser)
	}

	return r
}
