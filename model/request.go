package model

import (
	"InterfaceAutoTest/common/httpTool"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type RequestModel struct {
	Url         string
	Method      string
	RequestHead map[string]string
	RequestBody string
}

//发送请求, 返回响应提的 字符串数据.
func (rm *RequestModel) SendRequestForResponseBody(hc *http.Client, cookie map[string]string) string {
	request, err := http.NewRequest(rm.Method, rm.Url, strings.NewReader(rm.RequestBody))
	if err != nil {
		log.Print("构造请求的时候发生错误,-->url,", rm.Url, ".请求体,", rm.RequestBody, "..err-->", err)
		return ""
	}
	//填充请求头...
	setHead(request, rm.RequestHead)
	//填充cookie
	setCookie(request, cookie)
	response, err := hc.Do(request)
	list := setCookieByCookieList(response.Header["Set-Cookie"])
	hc.Jar.SetCookies(request.URL, list)
	if err != nil {
		log.Print("发生请求时,发生错误,-->url,", rm.Url, ".请求体,", rm.RequestBody, "..err-->", err)
		return ""
	}
	defer func() {
		if response != nil {
			err := response.Body.Close()
			if err != nil {
				log.Print("释放响应体流的时候发生错误.", err)
			}
		}
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print("读取响应体流时,发生错误,-->url,", rm.Url, ".请求体,", rm.RequestBody, "..err-->", err)
		return ""
	}
	return string(body)
}

func setCookie(r *http.Request, cookie map[string]string) {
	for k, v := range cookie {
		c := &http.Cookie{Name: k, Value: v}
		r.AddCookie(c)
	}

}

func setCookieByCookieList(cookies []string) []*http.Cookie {
	var result []*http.Cookie
	for _, v := range cookies {
		split := strings.Split(v, ";")
		cookie := http.Cookie{}
		for _, v1 := range split {
			split1 := strings.Split(v1, "=")
			if split1[0] == " path" {
				cookie.Path = split1[1]
				break
			}
			if split1[0] == " Max-Age" {
				i, err := strconv.Atoi(split1[1])
				if err != nil {
					log.Print("字符串转整形时,发生错误.->", err)
					break
				}
				cookie.MaxAge = i
				break
			}
			if split1[0] == " HttpOnly" {
				cookie.HttpOnly = true
				break

			}
			cookie.Name = split1[0]
			cookie.Value = split1[1]

		}
		result = append(result, &cookie)
	}

	return result

}

func setHead(r *http.Request, head map[string]string) {
	for k, v := range head {
		r.Header.Set(k, v)
	}
}

func MakeRequest(Url string, Method string, RequestHead map[string]string, RequestBody string) *RequestModel {
	m := httpTool.MakeHeader()
	if RequestHead != nil {
		m = RequestHead
	}
	model := RequestModel{Url, Method, m, RequestBody}
	return &model

}
