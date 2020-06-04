//配置文件工具
package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ReadConfig(fileName string) map[string]interface{} {
	file, err := ioutil.ReadFile(fileName)
	m := map[string]interface{}{}
	if err != nil {
		log.Fatal("读取配置文件发生错误-->", err)
	}
	err = yaml.Unmarshal(file, m)
	if err != nil {
		log.Fatal("解析配置文件的时候发生错误-->", err)
	}
	return m
}
