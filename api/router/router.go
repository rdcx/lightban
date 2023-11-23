package router

import (
	"lightban/api/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUp(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", h.Login)
	r.POST("/register", h.Register)

	r.GET("/user", h.Auth(h.GetUser))

	r.GET("/projects", h.Auth(h.GetProjects))
	r.GET("/projects/:id", h.Auth(h.GetProject))

	r.POST("/projects", h.Auth(h.CreateProject))
	r.PUT("/projects/:id", h.Auth(h.UpdateProject))
	r.DELETE("/projects/:id", h.Auth(h.DeleteProject))

	r.POST("/projects/:id/lists", h.Auth(h.CreateList))
	r.PUT("/projects/:id/lists/:id", h.Auth(h.UpdateList))
	r.DELETE("/projects/:id/lists/:id", h.Auth(h.DeleteList))

	r.POST("/lists/:id/tasks", h.Auth(h.CreateTask))
	r.PUT("/tasks/:id", h.Auth(h.UpdateTask))
	r.DELETE("/tasks/:id", h.Auth(h.DeleteTask))

	return r
}
