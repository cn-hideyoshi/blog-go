package dao

import (
	"blog-go/models"
	"log"
)

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select user_name from blog_user where uid =?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}

func GetUser(username, password string) *models.User {
	row := DB.QueryRow(
		"select * from blog_user where user_name =? and passwd = ? limit 1",
		username,
		password,
	)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user = &models.User{}
	err := row.Scan(
		&user.Uid,
		&user.UserName,
		&user.Passwd,
		&user.Avatar,
		&user.CreateAt,
		&user.UpdateAt,
	)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return user
}
