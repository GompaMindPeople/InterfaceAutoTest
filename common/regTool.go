//正则表达式 工具包
package common

import (
	"log"
	"regexp"
)

func MatchString(pattern, value string) bool {
	result, err := regexp.MatchString(pattern, value)
	if err != nil {
		log.Print("正则匹配时,发生报错.表达式-->", pattern)
		return false
	}
	return result
}
