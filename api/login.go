package api

import (
	"blog-go/common"
	"blog-go/service"
	"net/http"
)

func (*APIHandler) Login(w http.ResponseWriter, r *http.Request) {
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}

	common.Success(w, loginRes)
}
