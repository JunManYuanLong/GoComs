package middleware

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"ict.com/public.v1/context"
	"strconv"
)

const (
	DefaultLimit  = 9999
	DefaultOffset = 0

	LimitVar  = "limit"
	OffsetVar = "offset"
)

func ParamParse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &context.CustomContext{Context: c}
		params := c.QueryParams()
		limit, err := strconv.Atoi(params.Get(LimitVar))
		if err != nil {
			limit = DefaultLimit
		}
		offset, er := strconv.Atoi(params.Get(OffsetVar))
		if er != nil {
			offset = DefaultOffset
		}
		cc.Limit = limit
		cc.Offset = offset
		cc.RequestId = uuid.NewV4().String()
		return next(cc)
	}
}
