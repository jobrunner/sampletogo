package main

import (
	"net/http"

	"github.com/jobrunner/sampletogo/libs/hello"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/two/greet/:audience", Greet)
	_ = e.Start(":8020")
}

func Greet(ctx echo.Context) error {
	audience := ctx.Param("audience")
	listener, err := hello.NewAudience(audience)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.String(http.StatusOK, listener.Greet())
}
