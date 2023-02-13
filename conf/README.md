## 参考

[code](https://linzyblog.netlify.app/2022/11/03/grpc-tls/#4、mutual-TLS)
[conf](https://blog.csdn.net/qq_28356505/article/details/120828080)

## 服务端TLS(server-side tls)

1.生成ca私钥

```shell
openssl genrsa -out ca.key 4096
```

2.生成ca证书签发请求，得到ca.csr

```shell
openssl req -new -sha256 -out ca.csr -key ca.key -config ca.conf
```

3.生成根证书
> openssl req：生成自签名证书，-new指生成证书请求、-sha256指使用sha256加密、-key指定私钥文件、-x509指输出证书、-days 3650为有效期，此后则输入证书拥有者信息

```shell
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.crt
```

4.Server端生成私钥

```shell
openssl genrsa -out server.key 2048
```

5.生成证书签发请求

```shell
openssl req -new -sha256 -out server.csr -key server.key -config server.conf
```

6.使用ca证书签发服务端证书

```shell
openssl x509 -req -days 3650 -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.pem -extensions req_ext -extfile server.conf
```

## 双向TLS(mutual tls)

1.生成ca私钥

```shell
openssl genrsa -out ca.key 4096
```

2.生成ca证书签发请求，得到ca.csr

```shell
openssl req -new -sha256 -out ca.csr -key ca.key -config ca.conf
```
> openssl req：生成自签名证书，-new指生成证书请求、-sha256指使用sha256加密、-key指定私钥文件、-x509指输出证书、-days 3650为有效期，此后则输入证书拥有者信息

3.生成根证书

```shell
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.crt
```

4.服务端私钥、证书

```shell
# 1. 生成私钥 得到server.key
openssl genrsa -out server.key 2048

# 2. 生成证书签发请求 得到server.csr
openssl req -new -sha256 -out server.csr -key server.key -config server.conf

# 3. 用CA证书生成终端用户证书 得到server.crt
openssl x509 -req -sha256 -CA ca.crt -CAkey ca.key -CAcreateserial -days 365 -in server.csr -out server.crt -extensions req_ext -extfile server.conf
```

5.客户端私钥、证书

```shell
# 1. 生成私钥 得到client.key
openssl genrsa -out client.key 2048

# 2. 生成证书签发请求 得到client.csr
openssl req -new -key client.key -out client.csr

# 3. 用CA证书生成客户端证书 得到client.crt
openssl x509 -req -sha256 -CA ca.crt -CAkey ca.key -CAcreateserial -days 365  -in client.csr -out client.crt
```