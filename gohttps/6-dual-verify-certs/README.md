### ca:
```
openssl genrsa -out ca.key 2048 
openssl req -x509 -new -nodes -key ca.key -subj "/CN=tonybai.com" -days 5000 -out ca.crt
```
### server:
```
openssl genrsa -out server.key 2048 
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000
```

### client:
```
openssl genrsa -out client.key 2048 
openssl req -new -key client.key -subj "/CN=tonybai_cn" -out client.csr
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 5000
```

### client:
`openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile client.ext -out client.crt -days 5000`


## CA:
    私钥文件 ca.key
    数字证书 ca.crt

## Server:
    私钥文件       server.key
    证书请求文件         .csr
    数字证书      server.crt

## 格式说明
.cer/.crt是用于存放证书，它是2进制形式存放的，不含私钥。
.pem跟crt/cer的区别是它以Ascii来表示。_



### CA
#### 为了保证证书的可靠性和有效性，在这里可引入 CA 颁发的根证书的概念。其遵守 X.509 标准

#### 生成 Key
`openssl genrsa -out ca.key 2048`

#### 生成密钥
`openssl req -new -x509 -days 7200 -key ca.key -out ca.pem`

### Server
#### 生成 CSR
`openssl req -new -key server.key -out server.csr`

#### 填写信息
```
Country Name (2 letter code) []:
State or Province Name (full name) []:
Locality Name (eg, city) []:
Organization Name (eg, company) []:
Organizational Unit Name (eg, section) []:
Common Name (eg, fully qualified host name) []:go-grpc-example
Email Address []:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
```

#### 基于 CA 签发
`openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem`


### Client
#### 生成 Key
`openssl ecparam -genkey -name secp384r1 -out client.key`
#### 生成 CSR
`openssl req -new -key client.key -out client.csr`
#### 基于 CA 签发
`openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem`

