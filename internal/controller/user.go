package controller

import (
	"estudosGo/internal/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	users, err := repository.GetUsers()
	if err != nil {
		fmt.Printf("Tipo de users: %T\n", users)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
