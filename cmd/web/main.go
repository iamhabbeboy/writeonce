package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo/v4"
	"github.com/theterminalguy/writeonce/internal/router"
)

func main() {
	e := echo.New()
	e = router.DefineRoutes(e)
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("routes.json", data, 0644)
	e.Logger.Fatal(e.Start(":3000"))
}
