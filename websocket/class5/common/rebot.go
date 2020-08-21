package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var err error

const (
	// 你的图灵机器人 apikey
	APIKEY = "8df35c9e99b54c95a6941e9af24e69c0"
	// 你的 userId
	USERID = "123456"
	// 图灵的接口请求地址
	URL = "http://openapi.tuling123.com/openapi/api/v2"
)

type ReqData struct {
	ReqType    int                    `json:"reqType"`
	Perception map[string]interface{} `json:"perception"`
	UserInfo   map[string]string      `json:"userInfo"`
}
type Result struct {
	ResultType string                 `json:"resultType"`
	Values     map[string]interface{} `json:"values"`
	GroupType  int                    `json:"groupType"`
}
type RespData struct {
	Intent  map[string]interface{} `json:"intent"`
	Results []Result               `json:"results"`
}

func NewRobot() *ReqData {
	info := map[string]string{
		"apiKey": APIKEY,
		"userId": USERID,
	}
	return &ReqData{
		ReqType:    0,
		Perception: nil,
		UserInfo:   info,
	}
}
func (r *ReqData) Chat(v interface{}) (error, []interface{}) {
	if val, ok := v.(string); ok {
		inputText := map[string]string{
			"text": val,
		}
		r.Perception = make(map[string]interface{})
		r.Perception["inputText"] = inputText
		jsonData, err := json.Marshal(r)
		if err == nil {
			return r.Post(jsonData)
		}
	} else {
		err = errors.New("format error")
	}
	return err, nil
}
func (r *ReqData) Post(data []byte) (error, []interface{}) {
	defer func() {
		if err != nil {
			log.Print("service_robot:post", err)
		}
	}()
	body := bytes.NewBuffer([]byte(data))
	req, err := http.NewRequest("POST", URL, body)
	if err == nil {
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		defer resp.Body.Close()
		if err != nil {
			return err, nil
		}
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err, nil
		}
		log.Println(string(result))
		var response RespData
		err = json.Unmarshal(result, &response)
		if err != nil {
			return err, nil
		}
		var outPut []interface{}
		for _, v := range response.Results {
			for _, value := range v.Values {
				outPut = append(outPut, value)
			}
		}
		return nil, outPut
	}
	return err, nil
}
