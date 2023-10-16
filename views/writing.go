package views

import (
	"blog-go/common"
	"blog-go/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()

	writing.WriteData(w, wr)
}
