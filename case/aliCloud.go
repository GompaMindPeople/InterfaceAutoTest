package testCase

import (
	"InterfaceAutoTest/service"
	"net/http"
	"strings"
)

var cloudService = service.ALiCloudService{}

func Start(hc *http.Client, id string) string {
	return cloudService.Start(hc, id)
}

func Stop(hc *http.Client, id string) string {
	return cloudService.Stop(hc, id)
}

func Restart(hc *http.Client, id string) string {
	return cloudService.Restart(hc, id)
}

func ResetPassword(hc *http.Client, id, newPassword string) string {
	return cloudService.ResetPassword(hc, id, newPassword)

}

func ReinstallSystem(hc *http.Client, id, systemId string) string {
	return cloudService.ReinstallSystem(hc, id, systemId)
}
func ResetControlPassword(hc *http.Client, id, password string) string {
	return cloudService.ResetControlPassword(hc, id, password)
}

func GetServerList(hc *http.Client) (result []string) {
	list := cloudService.GetList(hc)
	split := strings.Split(list, "type=\\\"checkbox\\\" name=\\\"item[]\\\" value=\\\"")
	for i := 1; i < len(split); i++ {
		index := strings.Index(split[i], "\\\"")
		result = append(result, split[i][:index])
	}
	return result
}

//<dd data-value=\"

func GetSystemList(hc *http.Client, id string) (result []string) {
	list := cloudService.GetSystemList(hc, id)
	split := strings.Split(list, "<dd data-value=\\\"")
	for i := 1; i < len(split); i++ {
		index := strings.Index(split[i], "\\\">")
		result = append(result, split[i][:index])
	}
	return result
}

func GetServerStatus(hc *http.Client, id string) string {
	return cloudService.GetServerStatus(hc, id)
}

func GetServerInfo(hc *http.Client, id string) string {
	return cloudService.GetServerInfo(hc, id)
}
