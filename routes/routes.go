package routes

import (
	"fmt"
	"net/http"

	"github.com/erdedigital/go-users/auth"
	"github.com/gin-gonic/gin"
)

func SetupServer(port string) {

	router := gin.Default()
	router.GET("/", Defautl)
	authGroup := router.Group("/auth")
	authGroup.POST("/register", auth.Register)
	router.Run(fmt.Sprintf(":%s", port))
}

func Defautl(c *gin.Context) {
	c.String(http.StatusNotFound, "%s", "404 page not found")
}
