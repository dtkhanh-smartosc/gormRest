package public

import (
	"github.com/HiBang15/sample-gorm.git/internal/module/user/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func SetRouter(router *gin.RouterGroup) {
	log.Print("Start init public router  BE SAMPLE GORM.....")

	users := router.Group("users")
	{
		users.POST("/", controller.CreateUser)
		users.GET("/", controller.GetUsers)
		users.GET("/:id", controller.GetUser) // users/1?queryName=queryValue
		users.GET("", controller.GetUsersWithPagination)
		users.DELETE("/:id", controller.DeleteUser)
		users.PUT("/:id", controller.UpdateUser)
		users.GET("/search", controller.GetUserWithSearch)
	}

	log.Print("Finish init public router BE SAMPLE GORM ....")
}
