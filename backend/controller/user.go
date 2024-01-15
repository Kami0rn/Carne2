package controller

import (
	"net/http"

	"github.com/Kami0rn/Carne/entity"
	"github.com/gin-gonic/gin"

)


func ReadAll(c *gin.Context) {
	db := entity.DB()
	var users []entity.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User read success", "users": users})

}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	db := entity.DB()
	var user entity.User
	db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "user": user})
	return
}

