package server

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	request2 "ict.com/project.v1/request"
	"ict.com/public.v1/utils"
	"strconv"
)

const PROJECT_ID string = "id"

func (s *Server) ProjectAddHandler(c echo.Context) error {
	req := &request2.AddProjectRequest{}
	_ = c.Bind(req)
	ret, err := govalidator.ValidateStruct(req)

	if ret {
		err2 := s.ProjectMgr.Add(req)
		if err2 != nil {
			return utils.ResponseErr(c, err2)
		}
		return utils.ResponseOk(c, "")
	}
	return utils.ResponseErr(c, err)
}

func (s *Server) ProjectDeleteHandler(c echo.Context) error {
	pId, err := strconv.Atoi(c.Param(PROJECT_ID))
	req := &request2.DeleteProjectRequest{}
	req.Id = pId
	if err != nil {
		return utils.ResponseErr(c, err)
	}
	err2 := s.ProjectMgr.Delete(req)
	if err2 != nil {
		return utils.ResponseErr(c, err2)
	}
	return utils.ResponseOk(c, "")
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

func (s *Server) ProjectFindAllHandler(c echo.Context) error {
	return nil
}

func (s *Server) ProjectUpdateHandler(c echo.Context) error {
	return nil
}
