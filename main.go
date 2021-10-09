package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
)

func init() {
	//TODO: add param
	fmt.Println("add param init")
	flag.Parse()
}

//type MyServer struct{
//
//}
//
//func (p MyServer)ServeHTTP(resp http.ResponseWriter, req *http.Request) {
//	fmt.Printf("Client ip %s\n", strings.Split(req.RemoteAddr, ":")[0])
//}

func baseInfo(resp http.ResponseWriter, req *http.Request) {
	fmt.Printf("Client ip:port %s\n", req.RemoteAddr)
}

func rootHandle(resp http.ResponseWriter, req *http.Request) {
	baseInfo(resp, req)
	h := resp.Header()
	for key, values := range req.Header {
		value := strings.Join(values, ";")
		glog.V(2).Infof("header %s value %s", key, value)
		h.Add(key, value)
	}

	os.Setenv("VERSION", "V1")
	version := os.Getenv("VERSION")
	h.Add("Version", version)
	glog.V(3).Infof("Version value %s", version)

	resp.Write(([]byte)("hello world"))
}

func healthCheck(resp http.ResponseWriter, req *http.Request) {
	baseInfo(resp, req)
	resp.WriteHeader(200)
	resp.Write(([]byte)("ok"))
}

func routerRegister() {
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/healthz", healthCheck)
}

func serverStart() error {
	//server := MyServer{}
	s := &http.Server{
		Addr:         ":8000",
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.ListenAndServe()
}

func main() {
	glog.Infoln("http server init")

	routerRegister()

	err := serverStart()
	if err != nil {
		glog.Fatalf("Server start failed, %s\n", err)
		os.Exit(-1)
	}

	glog.Infoln("http server start successfully")
}
