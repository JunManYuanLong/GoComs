package server

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	request2 "ict.com/project.v1/request"
	"ict.com/public/context"
	"ict.com/public/utils"
	"strconv"
)

const ProjectId string = "id"

func (s *Server) ProjectAddHandler(c echo.Context) error {
	req := &request2.AddProjectRequest{}
	_ = c.Bind(req)
	ret, err := govalidator.ValidateStruct(req)
	if ret {
		err2 := s.ProjectMgr.Add(req)
		if err2 != nil {
			return utils.ResponseErr(c, err2)
		}
		return utils.ResponseOk(c, nil)
	}
	return utils.ResponseErr(c, err)
}

func (s *Server) ProjectDeleteHandler(c echo.Context) error {
	pId, err := strconv.Atoi(c.Param(ProjectId))
	if err != nil {
		return utils.ResponseErr(c, err)
	}
	err2 := s.ProjectMgr.Delete(pId)
	if err2 != nil {
		return utils.ResponseErr(c, err2)
	}
	return utils.ResponseOk(c, "")
}

func (s *Server) ProjectFindByIdHandler(c echo.Context) error {
	pId, err := strconv.Atoi(c.Param(ProjectId))
	if err != nil {
		return utils.ResponseErr(c, err)
	}
	req := &request2.FindProjectByIdRequest{}
	req.Id = pId
	ret, err := govalidator.ValidateStruct(req)

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
	cc := c.(*context.CustomContext)
	pList, err := s.ProjectMgr.FindAll(cc.Limit, cc.Offset)
	if err != nil {
		return utils.ResponseErr(c, err)
	}
	log.Info(cc)
	return utils.ResponseListOk(c, pList, cc.Limit, cc.Offset)
}

func (s *Server) ProjectUpdateHandler(c echo.Context) error {
	req := &request2.UpdateProjectRequest{}
	_ = c.Bind(req)
	ret, err := govalidator.ValidateStruct(req)
	if ret {
		if err2 := s.ProjectMgr.Update(req); err2 != nil {
			return utils.ResponseErr(c, err2)
		}
		return utils.ResponseOk(c, nil)
	}
	return utils.ResponseErr(c, err)
}
