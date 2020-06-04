package testCase

import (
	"InterfaceAutoTest/common/assert"
	"InterfaceAutoTest/model"

	"math/rand"
	"strconv"

	"github.com/golang/glog"
	"time"
)

func TestStart() {
	httpClient := model.HttpClient{}
	var client = httpClient.Initialize(nil)
	cc := client.Client
	loginResult := Login(cc, "kkwlcs", "a123456")
	glog.Infoln(loginResult)
	list := GetServerList(cc)
	for _, v := range list {
		status := GetServerStatus(cc, v)
		time.Sleep(time.Second * 5)
		status = GetServerStatus(cc, v)
		reg := assert.AssertsByReg(status, "已过期")
		if !reg {
			result := Start(cc, v)
			if assert.AssertsByReg(status, "运行中") {
				if assert.AssertsByReg(result, "当前状态不允许此操作，请稍候重试") {
					glog.Infoln("用例执行正确!result->" + result)
				} else {
					glog.Error("用例执行结果错误!result->" + result + " 结果应该是-->当前状态不允许此操作，请稍候重试")
				}
				continue
			}
			if assert.AssertsByReg(status, "已停止") {
				if assert.AssertsByReg(result, "云主机开机成功") {
					glog.Infoln("用例执行正确!result->" + result)
				} else {

					glog.Errorln("用例执行结果错误!result->" + result + " 当前状态-->" + status + " 结果应该是-->云主机开机成功")
				}
				continue
			}
		}
		glog.Infoln("start时错误的状态--->" + status + " id:" + v)
	}

}

func TestStop() {
	var client = httpClient.NewClient(nil)
	cc := client.Client
	loginResult := Login(cc, "kkwlcs", "a123456")
	glog.Infoln(loginResult)
	list := GetServerList(cc)
	for _, v := range list {
		status := GetServerStatus(cc, v)
		time.Sleep(time.Second * 5)
		status = GetServerStatus(cc, v)
		reg := assert.AssertsByReg(status, "已过期")
		if !reg {
			result := Stop(cc, v)
			if assert.AssertsByReg(status, "已停止") {
				if assert.AssertsByReg(result, "当前状态不允许此操作，请稍候重试") {
					glog.Infoln("用例执行正确!\nresult->" + result)
				} else {
					glog.Errorln("用例执行结果错误!\nresult->" + result + " 结果应该是-->当前状态不允许此操作，请稍候重试")
				}
				continue
			}
			if assert.AssertsByReg(status, "运行中") {
				if assert.AssertsByReg(result, "云主机停止成功") {
					glog.Infoln("用例执行正确!\nresult->" + result)
				} else {
					glog.Errorln("用例执行结果错误!\nresult->" + result + " 当前状态-->" + status + " 结果应该是-->云主机停止成功")
				}
				continue
			}
			glog.Infoln("stop时 错误的状态--->" + status + " id:" + v)
		}

	}
}

func TestRestart() {
	var client = httpClient.NewClient(nil)
	cc := client.Client
	loginResult := Login(cc, "kkwlcs", "a123456")
	glog.Infoln(loginResult)
	list := GetServerList(cc)
	for _, v := range list {
		status := GetServerStatus(cc, v)
		reg := assert.AssertsByReg(status, "已过期")
		if !reg {
			result := Restart(cc, v)
			time.Sleep(time.Second * 5)
			status = GetServerStatus(cc, v)
			if assert.AssertsByReg(status, "已停止") {
				if assert.AssertsByReg(result, "当前状态不允许此操作，请稍候重试") {
					glog.Infoln("用例执行正确!\nresult->" + result)
				} else {
					glog.Errorln("用例执行结果错误!\nresult->" + result + " 结果应该是-->当前状态不允许此操作，请稍候重试")
				}
				continue
			}
			if assert.AssertsByReg(status, "运行中") {
				if assert.AssertsByReg(result, "云主机重启成功") {
					glog.Infoln("用例执行正确!\nresult->" + result)
				} else {
					glog.Errorln("用例执行结果错误!\nresult->" + result + " 当前状态-->" + status + " 结果应该是-->云主机停止成功")
				}
				continue
			}
			glog.Infoln("Restart时 错误的状态--->" + status + " id:" + v)
		}

	}
}

func TestResetPassword() {
	var client = httpClient.NewClient(nil)
	cc := client.Client
	loginResult := Login(cc, "kkwlcs", "a123456")
	glog.Infoln(loginResult)

	list := GetServerList(cc)
	for _, v := range list {
		status := GetServerStatus(cc, v)
		reg := assert.AssertsByReg(status, "已过期")
		if !reg {

			result := ResetPassword(cc, v, "123456")
			assert.AssertLog(assert.AssertsByReg(result, "密码重置失败"), result)

			result = ResetPassword(cc, v, "xuzhimin123456789")
			assert.AssertLog(assert.AssertsByReg(result, "密码重置失败"), result)

			intn := rand.Intn(10000000)
			itoa := strconv.Itoa(intn)
			password := "a." + itoa
			result = ResetPassword(cc, v, password)
			if assert.AssertsByReg(result, "密码重置成功") {
				assert.AssertLog(assert.AssertsByReg(GetServerInfo(cc, v), password), result)
			} else {
				glog.Errorln("测试用来执行失败:" + result)
			}

			//assert.AssertLog(,result)

			continue
		}
		glog.Infoln("TestResetPassword时 错误的状态--->" + status + " id:" + v)

	}
}

func TestReinstallSystem() {
	var client = httpClient.NewClient(nil)
	cc := client.Client
	loginResult := Login(cc, "kkwlcs", "a123456")
	glog.Infoln(loginResult)
	list := GetServerList(cc)
	for _, v := range list {
		status := GetServerStatus(cc, v)
		reg := assert.AssertsByReg(status, "已过期")
		if !reg {
			systemList := GetSystemList(cc, v)
			if assert.AssertsByReg(status, "运行中") {
				glog.Infoln("运行中的系统是不允许进行重装系统!,正在关闭关闭系统.并等待8s,服务器id->" + v)
				Stop(cc, v)
				time.Sleep(time.Second * 8)
			}
			intn := rand.Intn(len(systemList))
			result := ReinstallSystem(cc, v, systemList[intn])
			assert.AssertLog(assert.AssertsByReg(result, "云主机正在重装系统"), result)
			continue
		}
		glog.Infoln("TestReinstallSystem 错误的状态--->" + status + " id:" + v)

	}

}

func TestResetControlPassword() {
	var client = httpClient.NewClient(nil)
	cc := client.Client
	loginResult := Login(cc, "kkwlcs", "a123456")
	glog.Infoln(loginResult)
	list := GetServerList(cc)
	for _, v := range list {

		intn := rand.Intn(10000000)
		itoa := strconv.Itoa(intn)

		result := ResetControlPassword(cc, v, itoa)
		if assert.AssertsByReg(result, "密码重置成功") {
			assert.AssertLog(assert.AssertsByReg(GetServerInfo(cc, v), itoa), result)
		} else {
			glog.Errorln("测试用来执行失败:" + result)
		}

	}
}
