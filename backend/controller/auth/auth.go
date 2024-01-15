package auth

import (
	"fmt"
	"github.com/Kami0rn/Carne/entity"
	"os"
	"time"
	"strconv"

	// "go/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

// Binding from JSON
type RegisterBody struct {
	FirstName    string `json:"FirstName" binding:"required"`
	LastName    string `json:"LastName" binding:"required"`
	Email       string `json:"Email" binding:"required"`
	PhoneNumber int    `json:"PhoneNumber"`
	Password         string `json:"Password"`

}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check user Exist
	var userExist entity.User
	db := entity.DB()
	db.Where("email = ?", json.Email).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User already exists"})
		return
	}


	//create user
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := entity.User{
		FirstName:    json.FirstName,
		Password:    string(encryptedPassword),
		LastName:    json.LastName,
		Email:       json.Email,
		PhoneNumber: json.PhoneNumber,

	}
	db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "User Registered",
			"userId":  user.ID})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "User Unregistered"})
	}
}

type LoginBody struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	db := entity.DB()
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userExist entity.User
	db.Where("email = ?", json.Email).First(&userExist)
	if userExist.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User does not exists"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	if err == nil {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))

		timeoutStr := os.Getenv("TIME_OUT")
		timeout, err := strconv.Atoi(timeoutStr)
		if err != nil {
			// Handle error if TIME_OUT is not set or not a valid integer
			// You can set a default timeout value in case of an error
			timeout = 60 // Defaulting to 7 minutes if TIME_OUT is not set or not a valid integer
		}
	
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExist.ID,
			"exp":    time.Now().Add(time.Minute * time.Duration(timeout)).Unix(),
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Login Success", "token": tokenString , "userId":  userExist.ID ,})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Login failed ,"})
	}
}
