package biz

import (
	"github.com/gin-gonic/gin"
	"yujian-backend/pkg/biz/auth"

	"yujian-backend/pkg/biz/user"
)

// SetupRouter 设置路由
func SetupRouter(r *gin.Engine) {
	// 用户相关的路由
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", user.CreateUser())
		userGroup.GET("/:id", user.GetUserById())
		userGroup.PUT("/:id", user.UpdateUser())
		userGroup.DELETE("/:id", user.DeleteUser())
	}

	// 登录相关的路由
	r.POST("/login", auth.UserLogin())
	r.POST("/register", auth.UserLogin())

}
