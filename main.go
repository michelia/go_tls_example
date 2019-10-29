package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	go server()
	client()
	// select {}
}

// 启动https服务
func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, TLS!\n")
	})

	log.Printf("About to listen on 8443. Go to https://127.0.0.1:8443/")
	// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
	// 传入证书(共钥)及私钥 启动tls服务
	err := http.ListenAndServeTLS("127.0.0.1:8443", "cert.pem", "key.pem", nil)
	log.Fatal(err)
}

// 发起https get请求
func client() {
	certBytes, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatalln(err)
	}
	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes) // 添加证书到证书池
	if !ok {
		log.Fatalln(err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: clientCertPool, // 添加已确认的证书池
				// InsecureSkipVerify: true,  // 当然你也可以不认证自签名证书
			},
		},
	}
	resp, err := client.Get("https://127.0.0.1:8443/")
	if err != nil {
		log.Fatalln(999, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(string(body))
}
