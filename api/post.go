package api

import (
	"blog-go/common"
	"blog-go/models"
	"blog-go/service"
	"blog-go/utils"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*APIHandler) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claim.Uid

	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)

		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)

		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidFloat := params["pid"].(float64)
		pType := int(postType)
		pid := int(pidFloat)
		post := &models.Post{
			pid,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}
}

func (*APIHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pidStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		common.Error(w, errors.New("参数错误"))
		return
	}

	detail, err := service.GetPostById(pid)
	if err != nil {
		common.Error(w, errors.New("查询失败"))
	}

	common.Success(w, detail)
}

func (*APIHandler) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")

	searchResp := service.SearchPost(condition)
	common.Success(w, searchResp)
}
