package server

import (
	"github.com/facebookgo/grace/gracehttp"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	business2 "ict.com/auth.v1/business"
	"ict.com/project.v1/business"
	middleware2 "ict.com/public/middleware"
)

type (
	Server struct {
		Addr   string
		App    echo.Echo
		Db     gorm.DB
		DbUri  string
		Logger log.Logger

		ProjectMgr business.ProjectMgr
		UserMgr    business2.UserMgr
	}

	Config struct {
		Addr   string
		Db     gorm.DB
		DbUri  string
		Logger log.Logger
	}
)

func NewServer(cfg *Config) *Server {
	app := echo.New()
	projectMgr, err := business.NewProjectBss(cfg.DbUri)
	if err != nil {
		log.Error("new project err:=>", err)
	}
	userMgr, err := business2.NewUserBss(cfg.DbUri)
	if err != nil {
		log.Error("new user err:=>", err)
	}
	s := &Server{
		Addr:       cfg.Addr,
		App:        *app,
		Db:         cfg.Db,
		DbUri:      cfg.DbUri,
		Logger:     cfg.Logger,
		ProjectMgr: projectMgr,
		UserMgr:    userMgr,
	}
	s.Db.SingularTable(true)
	s.Db.LogMode(true)
	configureHandler(s)
	return s
}

func configureMiddle(group *echo.Group, opName string, s *Server) {
	group.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.DELETE, echo.POST, echo.OPTIONS, echo.PUT, echo.HEAD},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))
	group.Use(middleware2.ParamParse)
	group.Use(middleware.Logger())
	group.Use(middleware.Recover())
}

func configureHandler(s *Server) {
	v := s.App.Group("/v1")
	{
		projectGroup := v.Group("/project")
		configureMiddle(projectGroup, "project", s)
		projectGroup.POST("/", s.ProjectAddHandler)
		projectGroup.DELETE("/:id", s.ProjectDeleteHandler)
		projectGroup.POST("/update", s.ProjectUpdateHandler)
		projectGroup.GET("/", s.ProjectFindAllHandler)
		projectGroup.GET("/:id", s.ProjectFindByIdHandler)
	}
	{
		userGroup := v.Group("/user")
		configureMiddle(userGroup, "user", s)
		userGroup.POST("/", s.UserAddHandler)
		userGroup.DELETE("/:id", s.UserDeleteHandler)
		userGroup.POST("/update", s.UserUpdateHandler)
		userGroup.GET("/", s.UserFindAllHandler)
		userGroup.GET("/:id", s.UserFindByIdHandler)
		userGroup.POST("/valid", s.UserIsValidHandler)
		userGroup.POST("/parsepassword", s.ParsePasswordHandler)
		userGroup.POST("/bindusergroup", s.UserBindUserGroupHandler)
		userGroup.POST("/updatepassword", s.UpdatePasswordHandler)
		userGroup.POST("/resetpassword", s.ResetPasswordHandler)
	}
}

func (s *Server) Start() {
	s.App.Server.Addr = s.Addr
	s.App.Logger.Fatal(gracehttp.Serve(s.App.Server))
}
