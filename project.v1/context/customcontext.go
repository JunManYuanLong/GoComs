package context

import (
	"github.com/labstack/echo"
)

type (
	CustomContext struct {
		echo.Context
		Limit       int
		Offset      int
		UserId      int
		UserName    string
		UserGroup   []string
		Role        []int64
		IsAdmin     bool
		GroupIDList []int32
	}
)
