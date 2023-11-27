package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dummy entity
type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetUserByID(ctx *gin.Context) {
	id := ctx.Param("userId")

	ctx.JSON(http.StatusOK, User{
		Id:   id,
		Name: fmt.Sprintf("name-%s", id),
	})
}
