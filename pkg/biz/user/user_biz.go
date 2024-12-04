package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"yujian-backend/pkg/db"
	"yujian-backend/pkg/model"
)

// CreateUser 创建用户的处理函数
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRepository := db.GetUserRepository()
		var userDTO model.UserDTO
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if id, err := userRepository.CreateUser(&userDTO); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "id": id})
		}
	}
}

// GetUserById 根据ID获取用户的处理函数
func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRepository := db.GetUserRepository()
		id := c.Param("id")
		userId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		userDO, err := userRepository.GetUserById(userId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": userDO})
	}
}

// UpdateUser 更新用户的处理函数
func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRepository := db.GetUserRepository()

		id := c.Param("id")
		userId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var userDTO model.UserDTO
		if err := c.ShouldBindJSON(&userDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userDO := userDTO.Transfer()
		userDO.Id = userId
		if err := userRepository.UpdateUser(userDO); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
}

// DeleteUser 删除用户的处理函数
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRepository := db.GetUserRepository()

		id := c.Param("id")
		userId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		if err := userRepository.DeleteUser(userId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
