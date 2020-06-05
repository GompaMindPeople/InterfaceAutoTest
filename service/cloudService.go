package service

import (
	"InterfaceAutoTest/model"
	"net/http"
)

type CloudService interface {
	GetList(hc *http.Client) (string, *model.RequestModel)
	Start(hc *http.Client, id string) (string, *model.RequestModel)
	Stop(hc *http.Client, id string) (string, *model.RequestModel)
	Restart(hc *http.Client, id string) (string, *model.RequestModel)
	GetServerStatus(hc *http.Client, id string) (string, *model.RequestModel)
	GetServerInfo(hc *http.Client, id string) (string, *model.RequestModel)
	ResetPassword(hc *http.Client, id, password string) (string, *model.RequestModel)
	ReinstallSystem(hc *http.Client, id, systemName string) (string, *model.RequestModel)
	ResetControlPassword(hc *http.Client, id, password string) (string, *model.RequestModel)
	GetSystemList(hc *http.Client, id string) (string, *model.RequestModel)
}
