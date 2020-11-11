package auth

import (
	"fmt"
	"net/http"

	"github.com/erdedigital/go-users/config"
	"github.com/gin-gonic/gin"
)

func CreateUser(p *ModuleUser, s AuthService) error {
	err := s.ServiceSave(p)
	if err != nil {
		return err
	}
	return nil
}

func Register(c *gin.Context) {
	user := &ModuleUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  fmt.Sprintf("Please make sure payload. [%s]", err),
			"data": nil,
		})
		return
	}

	db, err := config.PostgresDB()
	if err != nil {
		fmt.Println("[Error]", err)

	}
	params := NewModuleUser()
	params.Email = user.Email
	authDbUser := NewAuthService(db)
	err = CreateUser(user, authDbUser)
	if err != nil {
		fmt.Println(err)
	}

}
