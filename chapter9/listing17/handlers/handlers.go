// Package handlers 为 web 服务提供端点。
package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes 设置 web 服务的路由。
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON 返回一个简单的 JSON 文档。
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@ardanstudios.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
