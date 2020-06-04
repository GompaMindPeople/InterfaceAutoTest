package service

import (
	"InterfaceAutoTest/model"
	"net/http"
)

type LoginService struct {
}

func (ls *LoginService) Login(hc *http.Client, username, password string) (string, error) {
	//model1 := model.RequestModel{Url: , Method: "POST", RequestHead: m, RequestBody: "username=" + username + "&password=" + password + "&verify=&phone=&msg_code="}
	model1 := model.MakeRequest("http://test.kkclouds.com/passport/cloudlogin.html", "POST", nil, "username="+username+"&password="+password+"&verify=&phone=&msg_code=")
	body := model1.SendRequestForResponseBody(hc, nil)
	//fmt.Print(body)
	return body, nil

}
