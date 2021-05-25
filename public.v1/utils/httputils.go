package utils

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

const (
	ACTIVE  = 0
	DISABLE = 1
)

type (
	Code struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func (c *Code) Error() string {
	return c.Message
}

func NewCode(code int, msg string) *Code {
	return &Code{
		Code:    code,
		Message: msg,
	}
}

var (
	CodeSucc               = NewCode(2000, "success")
	CodeRecordNotExist     = NewCode(2000, "没有查询到记录")
	CodeCreateErr          = NewCode(2001, "创建失败")
	CodeUpdateErr          = NewCode(2002, "更新失败")
	CodeSaveErr            = NewCode(2003, "保存失败")
	CodeOperatErr          = NewCode(2004, "操作失败")
	CodePermissionNotFound = NewCode(4001, "没有权限")
	CodeObjectNotFound     = NewCode(4004, "对象不存在")
	CodeObjectHasBeExsit   = NewCode(4005, "对象已存在")
	CodeServerError        = NewCode(5000, "服务异常")
	CodePassWordError      = NewCode(6001, "password error")
	CodeLoginFailed        = NewCode(6002, "username or password wrong")
)

func response(c echo.Context, code *Code, data interface{}, limit int, offset int) error {
	var realData interface{}
	if code.Code != 2000 || data == nil {
		log.Printf("dataway data empty", realData)
		realData = map[string]interface{}{
			"code":    code.Code,
			"message": code.Message,
			"data":    data,
		}
		return c.JSON(http.StatusOK, realData)
	} else {
		if v, ok := data.(map[string]interface{}); ok {
			v["code"] = code.Code
			realData = v
		} else {
			realData = map[string]interface{}{
				"code":    code.Code,
				"message": code.Message,
				"data":    data,
			}
			if limit != 0 || offset != 0 {
				realData.(map[string]interface{})["limit"] = limit
				realData.(map[string]interface{})["offset"] = offset
			}
		}
		return c.JSON(http.StatusOK, realData)
	}
}

func ResponseListOk(c echo.Context, data interface{}, limit int, offset int) error {
	return response(c, CodeSucc, data, limit, offset)
}

func responseValidate(c echo.Context, code govalidator.Errors) error {
	realData := make(map[string]interface{})
	realData["message"] = code.Error()
	realData["code"] = http.StatusBadGateway
	return c.JSON(http.StatusBadGateway, realData)

}

func ResponseOk(c echo.Context, data interface{}) error {
	return response(c, CodeSucc, data, 0, 0)
}

func ResponseErr(c echo.Context, err error) error {
	switch e := err.(type) {
	case *Code:
		return response(c, e, nil, 0, 0)
	case govalidator.Errors:
		return responseValidate(c, e)
	}
	realData := make(map[string]interface{})
	realData["message"] = err.Error()
	realData["data"] = make(map[string]interface{})
	realData["code"] = http.StatusServiceUnavailable
	return c.JSON(http.StatusServiceUnavailable, realData)
}
