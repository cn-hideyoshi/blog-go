package service

import (
	"blog-go/dao"
	"blog-go/models"
	"blog-go/utils"
	"errors"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "hideyoshi")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}

	uid := user.Uid
	award, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token生成失败")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar

	var lr = &models.LoginRes{
		award,
		userInfo,
	}

	return lr, nil
}
