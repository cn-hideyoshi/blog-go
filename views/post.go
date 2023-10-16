package views

import (
	"blog-go/common"
	"blog-go/service"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	path := r.URL.Path
	pidStr := strings.TrimPrefix(path, "/p/")
	pidStr = strings.TrimSuffix(pidStr, ".html")
	pId, err := strconv.Atoi(pidStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别路径"))
		return
	}

	postDetail, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postDetail)
	//detail.Execute(w, data)
}
