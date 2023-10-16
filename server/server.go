package server

import (
	"blog-go/router"
	"log"
	"net/http"
)

type MsServer struct {
}

var App = &MsServer{}

func (*MsServer) Start(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
