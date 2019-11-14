// 模拟跨域问题
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Entrance)
	http.ListenAndServe(":8000", nil)
}

// CORS，服务器端配置
func Entrance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type, token, Access-Control-Allow-Origin") //header的类型
	w.Header().Set("Access-Control-Allow-Credentials", "false")
	w.Header().Set("content-type", "application/json") //返回数据格式是json
	w.Write([]byte("hello"))
}
