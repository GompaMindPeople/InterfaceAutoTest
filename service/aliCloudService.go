package service

import (
	"InterfaceAutoTest/common/httpTool"
	"InterfaceAutoTest/model"
	"net/http"
)

type ALiCloudService struct {
}

///aliyun/aliyun_product.html
//获取阿里云的的服务列表
func (tx *ALiCloudService) GetList(hc *http.Client) string {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/aliyun_product.html", Method: "GET", RequestHead: m, RequestBody: ""}
	return model1.SendRequestForResponseBody(hc, nil)
}

//开启一台服务器
func (tx *ALiCloudService) Start(hc *http.Client, id string) string {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/do_aliyun_host_status/id/" + id + "/type/1.html", Method: "POST", RequestHead: m, RequestBody: "items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil)
}

//关闭一台服务器
func (tx *ALiCloudService) Stop(hc *http.Client, id string) string {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/do_aliyun_host_status/id/" + id + "/type/2.html", Method: "POST", RequestHead: m, RequestBody: "items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil)
}

//重启一台服务器
func (tx *ALiCloudService) Restart(hc *http.Client, id string) string {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/do_aliyun_host_status/id/" + id + "/type/4.html", Method: "POST", RequestHead: m, RequestBody: "items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil)
}

//重新设置密码
func (tx *ALiCloudService) ResetPassword(hc *http.Client, id, password string) string {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/do_aliyun_host_status/id/" + id + "/type/6.html", Method: "POST", RequestHead: m, RequestBody: "password=" + password + "&items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil)
}

//imageId	centos_7_05_64_20G_alibase_20181210.vhd
//items[]	704
//重新安装系统
func (tx *ALiCloudService) ReinstallSystem(hc *http.Client, id, systemName string) string {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/do_aliyun_host_status/id/" + id + "/type/7.html", Method: "POST", RequestHead: m, RequestBody: "imageId=" + systemName + "&items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil)
}

//重置控制台密码
func (tx *ALiCloudService) ResetControlPassword(hc *http.Client, id, password string) string {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/do_aliyun_host_status/id/" + id + "/type/8.html", Method: "POST", RequestHead: m, RequestBody: "password=" + password + "&items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil)
}

//获取当前服务器的状态
func (tx *ALiCloudService) GetServerStatus(hc *http.Client, id string) string {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/to_aliyun_host_status.html?id=" + id, Method: "POST", RequestHead: m, RequestBody: "items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil)
}

//获取当前服务器可以安装的系统列表
func (tx *ALiCloudService) GetSystemList(hc *http.Client, id string) string {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/do_aliyun_host_status/id/" + id + "/type/7.html", Method: "GET", RequestHead: m, RequestBody: ""}
	return model1.SendRequestForResponseBody(hc, nil)
}

//
//获取服务器的详细配置
func (tx *ALiCloudService) GetServerInfo(hc *http.Client, id string) string {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/aliyun_product_show/id/" + id + ".html", Method: "GET", RequestHead: m, RequestBody: ""}
	return model1.SendRequestForResponseBody(hc, nil)
}
