package service

import (
	"InterfaceAutoTest/model"
	"net/http"
)

type KeyPair interface {
	AddKeyPair(hc *http.Client, name, region, type1, publicKey string) (string, *model.RequestModel)
	GetKeyPairByName(hc *http.Client, name string) (string, *model.RequestModel)
	UnBindingKeyPairByInstanceId(hc *http.Client, keyId, instanceIds string) (string, *model.RequestModel)
	BindingKeyPairByInstanceId(hc *http.Client, keyId, instanceIds string) (string, *model.RequestModel)
	DeleteKeyPair(hc *http.Client, keyId string) (string, *model.RequestModel)
	ListKeyPair(hc *http.Client) (string, *model.RequestModel)
}
