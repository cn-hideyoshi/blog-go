package service

import (
	"blog-go/config"
	"blog-go/dao"
	"blog-go/models"
	"html/template"
)

func GetPostByCategoryId(cId, page, pageSize int) (*models.CategoryResponse, error) {
	categories, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	posts, err := dao.GetPostPageByCId(cId, page, pageSize)
	if err != nil {
		return nil, err
	}

	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.CategoryId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:99]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	total := dao.CountGetPostByCId(cId)
	pageCount := (total-1)/pageSize + 1
	var pages []int
	for i := 0; i < pageCount; i++ {
		pages = append(pages, i+1)
	}

	homeResponse := &models.HomeResponse{
		config.Cfg.Viewer,
		categories,
		postMores,
		total,
		page,
		pages,
		page != pageCount,
	}
	categoryName := dao.GetCategoryNameById(cId)
	categoryResponse := &models.CategoryResponse{
		homeResponse,
		categoryName,
	}
	return categoryResponse, nil
}
