package service

import (
	// "fmt"
	"net/http"

	"github.com/unrolled/render"
)

func loginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.Method == "GET" {
			formatter.HTML(w, http.StatusOK, "login", nil)
		} else {
			//请求的是登录数据，那么执行登录的逻辑判断
			// fmt.Println("username:", req.Form["username"])
			// fmt.Println("password:", req.Form["password"])
			formatter.HTML(w, http.StatusOK, "login-info", struct {
				Header   map[string][]string `json:"header"`
				Username string `json:"username"`
				Password string `json:"password"`
			}{Header: req.Header, Username: req.Form["username"][0], Password: req.Form["password"][0]})
		}
	}
}
