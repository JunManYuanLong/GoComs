package business

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"ict.com/auth.v1/config"
	"ict.com/auth.v1/model"
	"ict.com/auth.v1/request"
	model2 "ict.com/public.v1/model"
	"ict.com/public.v1/utils"
	"time"
)

const (
	UserStatus                         = "status"
	UserStatusCondition                = "status=?"
	UserNameAndPasswordVerifyCondition = "name = ? and password = ?"

	OrderIdDesc = "id desc"
)

type (
	UserMgr interface {
		Add(ctx context.Context, req request.AddUserRequest) error
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, req request.UpdateUserRequest) error
		FindAll(ctx context.Context, limit int, offset int) ([]model.User, error)
		FindById(ctx context.Context, id int) (model.User, error)
		IsValid(ctx context.Context, username string, password string) (int, bool)
		ParsePassword(ctx context.Context, password string) string
		BindUserGroup(ctx context.Context, req request.UserBindGroupRequest) error
		UpdatePassword(ctx context.Context, req request.UpdatePasswordRequest) error
		ResetPassword(ctx context.Context, req request.ResetPasswordRequest) error
	}

	UserBss struct {
		Conn gorm.DB
	}
)

func (u *UserBss) Add(ctx context.Context, req request.AddUserRequest) error {
	user := &model.User{EntityModel: model2.EntityModel{
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	},
		Nickname:  req.Nickname,
		Password:  u.ParsePassword(ctx, req.Password),
		Name:      req.Name,
		Email:     req.Email,
		Telephone: req.Telephone,
		Picture:   req.Picture,
	}
	if err := u.Conn.Create(user).Error; err != nil {
		u.Conn.Rollback()
		log.Error("error===>", err.Error())
		return err
	}
	return nil
}

func (u *UserBss) Delete(_ context.Context, id int) error {
	user := &model.User{EntityModel: model2.EntityModel{ID: id}}
	if err := u.Conn.First(user).Update(UserStatus, utils.DISABLE).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserBss) Update(_ context.Context, req request.UpdateUserRequest) error {
	user := &model.User{
		EntityModel: model2.EntityModel{
			ID: req.Id,
		},
	}
	if err := u.Conn.First(user, UserStatusCondition, utils.ACTIVE).Error; err != nil {
		return err
	}
	user.EntityModel.UpdateTime = time.Now()
	user.Nickname = req.Nickname
	user.Picture = req.Picture
	user.Email = req.Email
	user.Telephone = req.Telephone
	if err := u.Conn.Save(user).Error; err != nil {
		u.Conn.Rollback()
		return err
	}
	return nil
}

func (u *UserBss) FindAll(_ context.Context, limit int, offset int) ([]model.User, error) {
	var ret []model.User
	if err := u.Conn.Order(OrderIdDesc).Limit(limit).Offset(offset).Where(UserStatusCondition, utils.ACTIVE).Find(&ret).Error; err != nil {
		log.Error(err)
		return ret, err
	}
	return ret, nil
}

func (u *UserBss) FindById(_ context.Context, id int) (model.User, error) {
	var user model.User
	if err := u.Conn.Where(UserStatusCondition, utils.ACTIVE).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserBss) IsValid(_ context.Context, username string, password string) (int, bool) {
	var user model.User
	if err := u.Conn.Where(UserNameAndPasswordVerifyCondition, username, u.ParsePassword(context.Background(), password)).First(&user).Error; err != nil {
		if user.ID != 0 {
			return user.ID, true
		}
	}
	return 0, false
}

func (u *UserBss) ParsePassword(_ context.Context, password string) string {
	return fmt.Sprintf("x%", md5.Sum([]byte(config.Salt+password)))
}

func (u *UserBss) BindUserGroup(_ context.Context, req request.UserBindGroupRequest) error {
	return nil
}

func (u *UserBss) UpdatePassword(ctx context.Context, req request.UpdatePasswordRequest) error {
	user := model.User{
		EntityModel: model2.EntityModel{
			ID: req.UserId,
		},
	}
	if err := u.Conn.Where(UserStatusCondition, utils.ACTIVE).First(&user).Error; err != nil {
		return err
	}
	if user.Password != u.ParsePassword(ctx, req.OldPassword) {
		return utils.CodePassWordError
	}

	user.Password = u.ParsePassword(ctx, req.NewPassword)
	user.UpdateTime = time.Now()
	if err := u.Conn.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserBss) ResetPassword(ctx context.Context, req request.ResetPasswordRequest) error {
	user := model.User{
		EntityModel: model2.EntityModel{
			ID: req.UserId,
		},
	}
	if err := u.Conn.Where(UserStatusCondition, utils.ACTIVE).First(&user).Error; err != nil {
		log.Error("error:====>", err)
		return err
	}
	user.Password = u.ParsePassword(ctx, req.NewPassword)
	user.UpdateTime = time.Now()
	if err := u.Conn.Save(user).Error; err != nil {
		log.Error("save password err=====>", err)
		u.Conn.Rollback()
		return err
	}
	return nil
}

func NewUserBss(dbUri string) (UserMgr, error) {
	db, err := gorm.Open("mysql", dbUri)
	if err != nil {
		return &UserBss{}, err
	}
	db.SingularTable(true)
	db.LogMode(true)
	return &UserBss{
		Conn: *db,
	}, err
}
