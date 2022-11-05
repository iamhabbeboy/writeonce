package main

import (
	"github.com/labstack/echo/v4"
	"github.com/theterminalguy/writeonce/internal/router"
)

func main() {
	e := echo.New()
	router.DefineRoutes(e)
	e.Logger.Fatal(e.Start(":3000"))
}
