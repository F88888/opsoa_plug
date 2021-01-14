package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"opsoa_plug/config/global"
	"strconv"
	"strings"
)

type addLabel struct {
	Code int    `json:"id"`
	Data string `json:"data"`
}

// @Tags Func
// @Summary 启动插件
// @Security Post
// @Success error
func Post(url string, data map[string]string) (content []byte, err error) {
	// 要访问的Url地址
	var r http.Request
	var request *http.Response
	contentType := "application/x-www-form-urlencoded"
	_ = r.ParseForm()
	// 发起请求
	for key, value := range data {
		r.Form.Add(key, value)
	}
	// 判断是否nil
	if request, err = http.Post(url, contentType, strings.NewReader(r.Form.Encode())); err == nil {
		defer request.Body.Close()
		content, err = ioutil.ReadAll(request.Body)
	}
	// 返回
	return content, err
}

// @Tags Func
// @Summary 添加标签
// @Security Post
// @Success error
func AddLabel(typeId int, label string, value string) (content []byte, err error) {
	var data addLabel
	var sqlLink bytes.Buffer
	postData := make(map[string]string)
	// 设置map
	postData["label"] = label
	postData["content"] = value
	postData["key"] = global.Key
	postData["type"] = strconv.Itoa(typeId)
	postData["id"] = strconv.Itoa(global.ID)
	sqlLink.WriteString("http://127.0.0.1:")
	sqlLink.WriteString(strconv.Itoa(global.Port))
	sqlLink.WriteString("/addLabel")
	// 判断是否链接到客户端
	if content, err = Post(sqlLink.String(), postData); err == nil {
		if err = json.Unmarshal(content, &data); err == nil {
			if data.Code == 200 {
				err = nil
			} else {
				err = errors.New(data.Data)
			}
		}
	}
	return content, err
}
