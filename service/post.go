package service

import (
	"blog-go/config"
	"blog-go/dao"
	"blog-go/models"
	"html/template"
	"log"
)

func GetPostDetail(pId int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pId)

	if err != nil {
		return nil, err
	}

	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.CategoryId)
	var postMore = models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	var postDetailRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	return postDetailRes, nil
}

func GetPostById(pId int) (*models.Post, error) {
	post, err := dao.GetPostById(pId)

	if err != nil {
		return nil, err
	}
	return post, nil
}

func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	categories, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}
	wr.Categorys = categories
	return
}

func SavePost(post *models.Post) {
	dao.SavePost(post)
}

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func SearchPost(condition string) []models.SearchResp {
	var searchResp = []models.SearchResp{}
	posts, _ := dao.GetPostSearch(condition)
	for _, post := range posts {
		searchResp = append(searchResp, models.SearchResp{
			post.Pid,
			post.Title,
		})
	}
	return searchResp

}
