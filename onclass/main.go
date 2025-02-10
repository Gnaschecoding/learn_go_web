package main

import (
	"fmt"
	"net/http"
)

type SignUpReq struct {
	//tag
	Email             string `form:"email" `
	Password          string `form:"password" `
	ConfirmedPassword string `form:"confirmed_password" `
}
type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "这是订单")
}
func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "这是创建用户")
}
func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "这是用户")
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "这是主菜单")
}

func SignUp(ctx *Context) {
	req := &SignUpReq{}
	err := ctx.ReadJson(req)

	if err != nil {
		ctx.BadRequstJson(err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败:%v\n", err)
	}
}

func main() {
	//http.HandleFunc("/", home)
	//http.HandleFunc("/order", order)
	//http.HandleFunc("/createuser", createUser)
	//http.HandleFunc("/user", user)
	//http.ListenAndServe(":8080", nil)

	server := NewHttpServer("server-test")

	//server.Route("/", home)
	//server.Route("/order", order)
	//server.Route("/createuser", createUser)
	//server.Route("/user", user)
	server.Route(http.MethodGet, "/user/signup", SignUp)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}

}
