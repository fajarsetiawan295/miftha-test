package routes

import (
	users "test-agit/internal/controllers/users"
	handlerUsers "test-agit/internal/handlers/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUsers(db *gorm.DB, route *gin.Engine) {

	repository := users.NewRepository(db)
	service := users.NewService(repository)
	handler := handlerUsers.NewHandler(service)

	gRoute := route.Group("/api/v1/user")

	gRoute.POST("/create", handler.Store)
	gRoute.POST("/update", handler.Update)

	gRoute.GET("/list", handler.List)
	gRoute.GET("/detail", handler.Detail)

	gRoute.DELETE("/delete", handler.Delete)

}
