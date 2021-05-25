package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	e := echo.New()
	//路由
	//String
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "响应String格式")
	})
	e.GET("html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>响应HTML格式</h1>")
	})
	e.GET("/json", func(c echo.Context) error {
		//p := &Person{
		//	Name: "哈哈哈",
		//	Age:  18,
		//}
		return c.JSON(http.StatusOK, nil)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
