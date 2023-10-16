package views

import (
	"blog-go/common"
	"blog-go/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	if err := r.ParseForm(); err != nil {
		log.Println("Index表单获取失败", err)
		index.WriteError(w, errors.New("系统错误请联系管理员!"))
		return
	}

	page := 1
	pageSize := 10

	pageStr := r.Form.Get("page")
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错", err)
		index.WriteError(w, errors.New("系统错误请联系管理员!"))
	}
	//页面上所有的数据需要有数据
	index.Execute(w, hr)
}
