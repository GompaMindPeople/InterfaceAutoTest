package testCase

import (
	"InterfaceAutoTest/service"
	"net/http"
)

var ser = service.LoginService{}

//所有的case 都需要登录操作
func Login(hc *http.Client, username, password string) string {
	result, _ := ser.Login(hc, username, password)
	return result
}
