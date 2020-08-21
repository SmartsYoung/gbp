package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
)

var addr2 = flag.String("addr", "127.0.0.1:8080", "ws service addr")
var key string

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("error : ", err)
		}
	}()
	flag.Parse()
	log.SetFlags(0)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{
		Scheme: "ws",
		Host:   *addr2,
		Path:   "/echo",
	}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("websocket connect error : ", err)
	}
	defer c.Close()
	receiveData := make(chan string)
	responseMessage := make(chan string)
	go func() {
		for {
			fmt.Print("请输入聊天内容 : ")
			fmt.Scan(&key)
			if key != "" {
				receiveData <- key
			}
			data := <-responseMessage
			fmt.Println("I receive message : ", data)
		}
	}()
	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			close(receiveData)
		case data := <-receiveData:
			err := c.WriteMessage(websocket.TextMessage, []byte(data))
			if err != nil {
				log.Println("write:", err)
				return
			}
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read from server err:", err)
			} else {
				responseMessage <- string(message)
			}
		}
	}
}
