package server

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"ict.com/auth.v1/request"
	"ict.com/public/context"
	"ict.com/public/utils"
	"strconv"
)

const (
	UserId = "id"
)

func (s *Server) UserAddHandler(c echo.Context) error {
	req := &request.AddUserRequest{}
	_ = c.Bind(req)
	ctx := c.Request().Context()
	ret, err := govalidator.ValidateStruct(req)
	if ret {
		err2 := s.UserMgr.Add(ctx, req)
		if err2 != nil {
			return utils.ResponseErr(c, err2)
		}
		return utils.ResponseOk(c, nil)
	}
	return utils.ResponseErr(c, err)
}

func (s *Server) UserDeleteHandler(c echo.Context) error {
	ctx := c.Request().Context()
	uId, err := strconv.Atoi(c.Param(UserId))
	if err != nil {
		return utils.ResponseErr(c, err)
	}
	err2 := s.UserMgr.Delete(ctx, uId)
	if err2 != nil {
		return utils.ResponseErr(c, err2)
	}
	return utils.ResponseOk(c, nil)
}

func (s *Server) UserUpdateHandler(c echo.Context) error {
	req := &request.UpdateUserRequest{}
	_ = c.Bind(req)
	ctx := c.Request().Context()
	ret, err := govalidator.ValidateStruct(req)
	if ret {
		if err2 := s.UserMgr.Update(ctx, req); err != nil {
			return utils.ResponseErr(c, err2)
		}
		return utils.ResponseOk(c, nil)
	}
	return utils.ResponseErr(c, err)
}

func (s *Server) UserFindAllHandler(c echo.Context) error {
	cc := c.(*context.CustomContext)
	ctx := cc.Request().Context()
	uList, err := s.UserMgr.FindAll(ctx, cc.Limit, cc.Offset)
	if err != nil {
		return utils.ResponseErr(c, err)
	}
	return utils.ResponseListOk(c, uList, cc.Limit, cc.Offset)
}

func (s *Server) UserFindByIdHandler(c echo.Context) error {
	ctx := c.Request().Context()
	uId, err := strconv.Atoi(c.Param(UserId))
	if err != nil {
		return utils.ResponseErr(c, err)
	}
	userFindByIdReplay, err2 := s.UserMgr.FindById(ctx, uId)
	if err2 != nil {
		return utils.ResponseErr(c, err2)
	}
	return utils.ResponseOk(c, userFindByIdReplay)
}

func (s *Server) UserIsValidHandler(c echo.Context) error {
	return nil
}

func (s *Server) ParsePasswordHandler(c echo.Context) error {
	return nil
}

func (s *Server) UserBindUserGroupHandler(c echo.Context) error {
	return nil
}

func (s *Server) UpdatePasswordHandler(c echo.Context) error {
	return nil
}

func (s *Server) ResetPasswordHandler(c echo.Context) error {
	return nil
}
