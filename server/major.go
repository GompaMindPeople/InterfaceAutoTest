package server

import (
	"InterfaceAutoTest/model"
	"InterfaceAutoTest/utils"
	"github.com/golang/glog"
	"strconv"
)

type CasePool struct {
	//所有的测试用例
	CaseData map[string]func(*model.HttpClient)
	//测试用例执行前所需要执行的方法
	BeforePool map[string]func(*model.HttpClient)
	//测试用例执行后所需要执行的方法
	AfterPool map[string]func(*model.HttpClient)
}

func New() *CasePool {
	return &CasePool{make(map[string]func(*model.HttpClient), 10), make(map[string]func(*model.HttpClient), 10), make(map[string]func(*model.HttpClient), 10)}
}

func (cp *CasePool) Start() {
	glog.Infoln("====================开始执行测试用例===========================")
	length := len(cp.CaseData)
	glog.Infoln("一共有", strconv.Itoa(length), "用例")

	cp.Execute()

	glog.Infoln("====================测试用例执行结束===========================")
}

//执行所有用例 面向切面
func (cp *CasePool) Execute() {
	for k, v := range cp.CaseData {
		client := &model.HttpClient{}
		client.Initialize(nil)

		glog.Infoln("当前执行的测试用例:", utils.GetFunctionName(v))

		if Before := cp.BeforePool[k]; Before != nil {
			Before(client)
		}
		v(client)
		if After := cp.AfterPool[k]; After != nil {
			After(client)
		}
	}

}

//增加一个简单的 case
func (cp *CasePool) AddCase(funcName string, f func(*model.HttpClient)) {
	cp.CaseData[funcName] = f
}

//注册一个 用例的case
func (cp *CasePool) RegisterCase(funcName string, caseFunc, Before, After func(*model.HttpClient)) {
	cp.CaseData[funcName] = caseFunc
	cp.BeforePool[funcName] = Before
	cp.AfterPool[funcName] = After

}
