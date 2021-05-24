package business

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"ict.com/project.v1/model"
	"ict.com/project.v1/request"
	model2 "ict.com/public.v1/model"
	"ict.com/public.v1/utils"
)

const (
	PROJECT_STATUS           = "status"
	PROJECT_STATUS_CONDITION = "status =? "
)

type (
	ProjectMgr interface {
		Add(req *request.AddProjectRequest) error
		Delete(req *request.DeleteProjectRequest) error
		Update(req *request.UpdateProjectRequest) error
		FindAll(req *request.FindAllProjectRequest) ([]model.Project, error)
		FindById(req *request.FindProjectByIdRequest) (model.Project, error)
	}

	ProjectBss struct {
		Conn gorm.DB
	}
)

func (p *ProjectBss) Add(req *request.AddProjectRequest) error {
	log.Info("into project add function...")
	project := &model.Project{
		Name:        req.Name,
		Description: req.Description,
		Logo:        req.Logo,
		Category:    req.Category,
		Level:       req.Level,
		PM:          req.PM,
		TD:          req.TD,
		Background:  req.Background,
		Worth:       req.Worth,
		Target:      req.Target,
		Milestone:   req.Milestone,
		Budget:      req.Budget,
		IsShow:      req.IsShow,
	}

	if err := p.Conn.Create(project).Error; err != nil {
		p.Conn.Rollback()
		log.Error("err", err.Error())
		return err
	}
	return nil
}

func (p *ProjectBss) Delete(req *request.DeleteProjectRequest) error {
	log.Info("into project delete function")
	pro := &model.Project{EntityModel: model2.EntityModel{ID: req.Id}}
	if err := p.Conn.First(pro).Update(PROJECT_STATUS, utils.DISABLE).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProjectBss) Update(req *request.UpdateProjectRequest) error {
	log.Info("into project Update function")

	pro := &model.Project{
		EntityModel: model2.EntityModel{
			ID: req.Id,
		},
	}
	p.Conn.First(pro)
	pro.Status = req.Status
	pro.Name = req.Name
	pro.Description = req.Description
	pro.Logo = req.Logo
	pro.Weight = req.Weight
	pro.Category = req.Category
	pro.Level = req.Level
	pro.PM = req.PM
	pro.TD = req.TD
	pro.Background = req.Background
	pro.Worth = req.Worth
	pro.Target = req.Target
	pro.Milestone = req.Milestone
	pro.Budget = req.Budget
	pro.IsShow = req.IsShow

	if err := p.Conn.Save(pro).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProjectBss) FindAll(req *request.FindAllProjectRequest) ([]model.Project, error) {
	log.Info("into project FindAll function...")
	var ret []model.Project
	if err := p.Conn.Limit(req.Limit).Offset(req.Offset).Where(PROJECT_STATUS, utils.ACTIVE).Find(&ret).Error; err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *ProjectBss) FindById(req *request.FindProjectByIdRequest) (model.Project, error) {
	log.Info("into project FindById function...")
	pro := model.Project{EntityModel: model2.EntityModel{ID: req.Id}}
	if err := p.Conn.Where(PROJECT_STATUS_CONDITION, utils.ACTIVE).First(&pro).Error; err != nil {
		log.Info("查询信息:====>", err)
		return model.Project{}, err
	}
	return pro, nil
}

func NewProjectBss(dbUri string) (ProjectMgr, error) {
	db, err := gorm.Open("mysql", dbUri)
	if err != nil {
		return &ProjectBss{}, err
	}
	db.SingularTable(true)
	db.LogMode(true)
	return &ProjectBss{
		Conn: *db,
	}, err
}
