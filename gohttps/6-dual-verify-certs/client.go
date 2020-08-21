package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	pool := x509.NewCertPool()
	caCertPath := "ca.crt"

	// 由于client端需要验证server端的数字证书，因此client端需要预先加载ca.crt，以用于服务端数字证书的校验：
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	ok := pool.AppendCertsFromPEM(caCrt) // 添加证书到证书池
	if !ok {
		log.Fatalln(err)
	}

	// 加载client.key和client.crt用于server端连接时的证书校验
	cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool, // 添加已确认的证书池
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8081")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
