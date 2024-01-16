package main

import (
	"github.com/gin-gonic/gin"

	// "github.com/Kami0rn/Carne/controller"
	Authcontroller "github.com/Kami0rn/Carne/controller/auth"

	"github.com/Kami0rn/Carne/entity"
	"github.com/gin-contrib/cors"
)

const PORT = "8080"

func main() {

	entity.SetupDatabase()

	r := gin.Default()

	r.Use(cors.Default())
	r.POST("/register", Authcontroller.Register)

	r.POST("/login", Authcontroller.Login)

	r.Run("localhost: " + PORT)

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
