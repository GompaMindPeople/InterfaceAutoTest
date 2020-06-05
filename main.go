package main

import (
	testCase "InterfaceAutoTest/case"
	"InterfaceAutoTest/server"
)

func main() {

	pool := server.New()
	RegisterCase(pool)
	pool.Start()

}

func RegisterCase(cp *server.CasePool) {

	cp.RegisterCase("TestStart", testCase.TestStart, testCase.Login, nil)
	cp.RegisterCase("TestStop", testCase.TestStop, testCase.Login, nil)
	cp.RegisterCase("TestRestart", testCase.TestRestart, testCase.Login, nil)
	cp.RegisterCase("TestResetPassword", testCase.TestResetPassword, testCase.Login, nil)
	cp.RegisterCase("TestReinstallSystem", testCase.TestReinstallSystem, testCase.Login, nil)
	cp.RegisterCase("TestResetControlPassword", testCase.TestResetControlPassword, testCase.Login, nil)

}
