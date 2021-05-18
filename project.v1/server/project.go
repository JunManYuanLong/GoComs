package server

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	context2 "ict.com/project.v1/context"
	request2 "ict.com/project.v1/request"
	"ict.com/public.v1/utils"
)

func (s *Server) ProjectAddHandler(c echo.Context) error {
	req := &request2.AddProjectRequest{}
	ret, err := govalidator.ValidateStruct(req)
	c.Bind(req)
	cc := c.(context2.CustomContext)
	ctx := cc.Request().Context()

	if ret {
		projectAddReplay, err2 := s.ProjectMgr.Add(ctx, req)
		if err2 != nil {
			return utils.ResponseErr(cc, err2)
		}
		return utils.ResponseOk(cc, projectAddReplay)
	}
	return utils.ResponseErr(cc, err)
}

func (s *Server) ProjectFindByIdHandler(c echo.Context) error {
	req := &request2.FindProjectByIdRequest{}
	_ = c.Bind(req)
	ret, err := govalidator.ValidateStruct(req)
	//cc, err3 := c.(context2.CustomContext)
	//fmt.Println(err3)
	//ctx := cc.Request().Context()

	if ret {
		projectFindByIdReplay, err2 := s.ProjectMgr.FindById(req)
		if err2 != nil {
			return utils.ResponseErr(c, err2)
		}
		return utils.ResponseOk(c, projectFindByIdReplay)
	}
	return utils.ResponseErr(c, err)
}
