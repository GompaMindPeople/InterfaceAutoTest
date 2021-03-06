package txCloudService

import (
	"InterfaceAutoTest/common/httpTool"
	"InterfaceAutoTest/model"
	"net/http"
)

type TXCloudService struct{}

//
//获取tx云的的服务列表
//http://test.kkclouds.com/tencent_cloud/tencent_product.html
func (tx *TXCloudService) GetList(hc *http.Client) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/tencent_cloud/tencent_product.html", Method: "GET", RequestHead: m, RequestBody: ""}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//开启一台服务器
//http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/2025.html
func (tx *TXCloudService) Start(hc *http.Client, id string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/" + id + "/type/1.html", Method: "POST", RequestHead: m, RequestBody: "items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//关闭一台服务器
//http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/2025/type/2.html
func (tx *TXCloudService) Stop(hc *http.Client, id string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/" + id + "/type/2.html", Method: "POST", RequestHead: m, RequestBody: "items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//重启一台服务器
//http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/2025/type/4.html
func (tx *TXCloudService) Restart(hc *http.Client, id string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/" + id + "/type/4.html", Method: "POST", RequestHead: m, RequestBody: "items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//重新设置密码
//http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/2025/type/6.html
func (tx *TXCloudService) ResetPassword(hc *http.Client, id, password string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/" + id + "/type/6.html", Method: "POST", RequestHead: m, RequestBody: "password=" + password + "&items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//imageId	centos_7_05_64_20G_alibase_20181210.vhd
//items[]	704
//重新安装系统fResetControlPassword
//http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/2025/type/7.html
func (tx *TXCloudService) ReinstallSystem(hc *http.Client, id, systemName string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/" + id + "/type/7.html", Method: "POST", RequestHead: m, RequestBody: "imageId=" + systemName + "&items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//重置控制台密码
func (tx *TXCloudService) ResetControlPassword(hc *http.Client, id, password string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/do_aliyun_host_status/id/" + id + "/type/8.html", Method: "POST", RequestHead: m, RequestBody: "password=" + password + "&items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//获取当前服务器的状态
//http://test.kkclouds.com/tencent_cloud/to_tencent_cloud_host_status.html?id=2025
func (tx *TXCloudService) GetServerStatus(hc *http.Client, id string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/tencent_cloud/to_tencent_cloud_host_status.html?id=" + id, Method: "POST", RequestHead: m, RequestBody: "items[]=" + id}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//获取当前服务器可以安装的系统列表
//http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/2025/type/7.html
func (tx *TXCloudService) GetSystemList(hc *http.Client, id string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/tencent_cloud/do_tencent_host_status/id/" + id + "/type/7.html", Method: "GET", RequestHead: m, RequestBody: ""}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//
//获取服务器的详细配置
//http://test.kkclouds.com/tencent_cloud/tencent_product_show/id/2025.html
func (tx *TXCloudService) GetServerInfo(hc *http.Client, id string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/tencent_cloud/tencent_product_show/id/" + id + ".html", Method: "GET", RequestHead: m, RequestBody: ""}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//=================================================密钥对===================================================================
//http://test.kkclouds.com/aliyun/add_key_pair.html
//name	dfgfd
//region	cn-hangzhou
//type	1,2   类型1 为自动新建密钥对,  类型2 为导入已有密钥对
//public_key
//
func (tx *TXCloudService) AddKeyPair(hc *http.Client, name, region, type1, publicKey string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/add_key_pair.html", Method: "POST", RequestHead: m, RequestBody: "name=" + name + "&region=" + region + "&type=" + type1 + "&public_key=" + publicKey}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//http://test.kkclouds.com/aliyun/key_pair.html?name=dddddddd
func (tx *TXCloudService) GetKeyPairByName(hc *http.Client, name string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/key_pair.html?name=" + name, Method: "GET", RequestHead: m, RequestBody: ""}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//http://test.kkclouds.com/aliyun/untying_key_pair/id/168.html
//instance_ids	i-bp16htuqeqnvgisgc5q6
func (tx *TXCloudService) UnBindingKeyPairByInstanceId(hc *http.Client, keyId, instanceIds string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/untying_key_pair/id/" + keyId + ".html", Method: "POST", RequestHead: m, RequestBody: "instance_ids=" + instanceIds}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//http://test.kkclouds.com/aliyun/binding_key_pair/id/168.html
//instance_ids	i-bp16htuqeqnvgisgc5q6
func (tx *TXCloudService) BindingKeyPairByInstanceId(hc *http.Client, keyId, instanceIds string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/binding_key_pair/id/" + keyId + ".html", Method: "POST", RequestHead: m, RequestBody: "instance_ids=" + instanceIds}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//http://test.kkclouds.com/aliyun/delete_key_pair/id/168.html
func (tx *TXCloudService) DeleteKeyPair(hc *http.Client, keyId string) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/delete_key_pair/id/" + keyId + ".html", Method: "POST", RequestHead: m, RequestBody: ""}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}

//http://test.kkclouds.com/aliyun/key_pair.html
//列出所有的 密钥对
func (tx *TXCloudService) ListKeyPair(hc *http.Client) (string, *model.RequestModel) {
	m := httpTool.MakeHeader()
	model1 := model.RequestModel{Url: "http://test.kkclouds.com/aliyun/key_pair.html", Method: "GET", RequestHead: m, RequestBody: ""}
	return model1.SendRequestForResponseBody(hc, nil), &model1
}
