package views

import (
	"blog-go/common"
	"blog-go/service"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("路径不匹配"))
		return
	}

	page := 1
	pageSize := 10

	pageStr := r.Form.Get("page")
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	categoryResponse, err := service.GetPostByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.Execute(w, categoryResponse)
	fmt.Println(path)
}
