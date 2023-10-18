package handlers

import "github.com/gin-gonic/gin"

type Handler struct {}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

    userManagement := router.Group("/users")
    {
        userManagement.GET("/", h.getAllUsers)
        userManagement.GET("/:id", h.getUserByID)
        userManagement.POST("/", h.createUser)
        userManagement.PUT("/:id", h.updateUser)
        userManagement.DELETE("/:id", h.deleteUser)
    }

	return router
}