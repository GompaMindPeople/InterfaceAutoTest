package service

import "net/http"

type CloudService interface {
	GetList(hc *http.Client) string
	Start(hc *http.Client, id string) string
	Stop(hc *http.Client, id string) string
	Restart(hc *http.Client, id string) string
	GetServerStatus(hc *http.Client, id string) string
	GetServerInfo(hc *http.Client, id string) string
}
