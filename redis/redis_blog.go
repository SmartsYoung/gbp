package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"net/http"

	"strings"
)

const RedisAddress = "127.0.0.1:6379"
const RedisDb = 0

const AllowRequestUrlH = "*"
const AllowRequestUrlW = "*"
const IllegalCharacters = "?"
const DefaultReadCount = "1"

var (
	// 定义常量
	RedisClient *redis.Pool
)

func main() {
	// 初始化redis连接池
	initRedisPool()

	// 启动web服务监听
	http.HandleFunc("/hello", blogReadCountIncr) //设置访问的路由
	err := http.ListenAndServe(":9401", nil)     //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func blogReadCountIncr(responseWriter http.ResponseWriter, request *http.Request) {

	// 解析参数，默认不解析
	request.ParseForm()

	blogId := request.Form.Get("blogId")

	log.Println(">>>>>> method blogReadCountIncr exec , request params is : ", blogId)

	// 判断请求参数是否为空
	if "" == blogId {
		result := ResultCode{
			Code: 200,
			Msg:  "success",
		}

		ret, _ := json.Marshal(result)
		fmt.Fprintf(responseWriter, string(ret)) //这个写入到w的是输出到客户端的
	}

	readCount := redisGet(blogId)
	if "" == readCount {
		// 不符合规则，直接返回
		flag := strings.Index(blogId, AllowRequestUrlH) != 0 || strings.Index(blogId, AllowRequestUrlW) != 0 || strings.Contains(blogId, IllegalCharacters)
		if !flag {
			result := ResultCode{
				Code: 200,
				Msg:  "success",
			}

			ret, _ := json.Marshal(result)
			fmt.Fprintf(responseWriter, string(ret)) //这个写入到w的是输出到客户端的
		}

		redisSet(blogId, DefaultReadCount)
		readCount = DefaultReadCount
	} else {
		readCount = redisIncr(blogId)
	}
	log.Println(">>>>>> readCount is : ", readCount)
	result := ResultCode{
		Code: 200,
		Msg:  "success",
		Data: readCount,
	}
	ret, _ := json.Marshal(result)
	fmt.Fprintf(responseWriter, string(ret)) //这个写入到w的是输出到客户端的
}

// 结构体定义返回值
type ResultCode struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data string `json:"data"`
}
