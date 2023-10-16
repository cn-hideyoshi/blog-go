package service

import (
	"blog-go/config"
	"blog-go/dao"
	"blog-go/models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	category, _ := dao.GetAllCategory()
	posts, _ := dao.GetAllPost()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	return models.PigeonholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		category,
		pigeonholeMap,
	}
}
