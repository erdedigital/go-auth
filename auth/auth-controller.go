package auth

import (
	"fmt"
	"net/http"

	"github.com/erdedigital/go-users/auth/model"
	"github.com/erdedigital/go-users/auth/service"
	"github.com/erdedigital/go-users/config"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	auth := &model.Auth{}
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  fmt.Sprintf("Please make sure payload. [%s]", err),
			"data": nil,
		})
	}
	db, err := config.PostgresDB()
	if err != nil {
		fmt.Println("[Error]", err.Error())
	}
	params := model.NewAuth()
	params.Email = auth.Email
	configAuth := service.NewAuth(db)
	err = CreateAuth(auth, configAuth)
	if err != nil {
		fmt.Println(err)
	}

}
