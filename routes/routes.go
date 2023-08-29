package routes

import (
	"github.com/Ephmeral/TodoList/api"
	"github.com/Ephmeral/TodoList/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-srcret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.GET("task/all", api.ShowAllTasks)
			authed.DELETE("task/:id", api.DeleteTask)
			authed.PUT("task/:id", api.UpdateTask)
		}
	}
	return r
}
