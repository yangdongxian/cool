package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func JsonMarshal(data interface{}) (string, error) {
	switch data.(type) {
	case string:
		return data.(string), nil
	default:
		bt, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		return string(bt), nil
	}
}
func Marshal(data interface{}) (string, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(data)
	if err != nil {
		return "", err
	}
	s := bf.Bytes() // 去掉末尾的换行符, 恶心的坑啊
	return string(s[0 : len(s)-1]), nil
}

// JSON 操作接口
type JSON interface {
	// 读取JSON对象
	ReadJSON(file string) *jSON
}

type jSON struct {
	Jsob *simplejson.Json
}

// NewJSON 创建JSON对象
func NewJSON() JSON {
	json := &jSON{}
	return json
}

func (json *jSON) ReadJSON(file string) *jSON {
	jstr := Bufio(file)
	js, err := simplejson.NewJson(jstr)
	//er.CheckErr(err)
	if err != nil {
		fmt.Printf("Bufio method -- error:%s ", err.Error())
	}
	json.Jsob = js
	return json
}
