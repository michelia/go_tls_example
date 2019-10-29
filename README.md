# go_tls_example

使用`go run $GOROOT/src/crypto/tls/generate_cert.go --ca  --duration 8760h --host 127.0.0.1,::1,example.com`生成自签名证书, 会在当前目录下生成两个文件`cert.pem`, `key.pem`, 分别是证书(共钥)和私钥  
`--duration` 证书有效期(默认一年)  
`--host` 逗号分隔的域名或服务端ip, 如果服务端是多机部署, 记得把所有ip都填进去. 但还是建议使用内网域名.  

## 示例代码
见[main.go](main.go)

## 参考
- [https://colobu.com/2016/06/07/simple-golang-tls-examples/]()
- [https://golang.google.cn/pkg/net/http/#example_ListenAndServeTLS]()

