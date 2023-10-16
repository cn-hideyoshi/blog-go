package main

import (
	"blog-go/common"
	"blog-go/server"
)

func init() {
	common.LoadTemplate()
}
func main() {
	//程序入口，一个项目只能有一个入口
	//web程序，http协议 ip port
	server.App.Start("127.0.0.1", "8081")
}

//	func index(w http.ResponseWriter, r *http.Request) {
//		//w.Write([]byte("hello hideyoshi.top go blog"))
//
//		w.Header().Set("Content-Type", "application/json")
//		indexData := IndexData{
//			Title: "hideyoshi's blog",
//			Desc:  "现在是入门教程",
//		}
//		jsonStr, _ := json.Marshal(indexData)
//		w.Write(jsonStr)
//
// }
