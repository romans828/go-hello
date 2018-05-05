package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"encoding/json"
	"time"
)

var logger *log.Logger

func main() {
	fmt.Println("Hello World Started.")

	filePath := fmt.Sprintf("%s/%s", "log", "hello.log")
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("[FATAL] Can't open log file. (path:%s, error:%s)", filePath, err.Error())
		os.Exit(1)
	}
	defer logFile.Close()
	logger = log.New(logFile, "", 0)

	h := &HelloHandler{}
	http.Handle("/", h)
	http.ListenAndServe(":5000", nil)
}

type HelloHandler struct {
	http.Handler
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	str := "Your Request Path is " + r.URL.Path
	w.Write([]byte(str))
	req := NewRequest(r.URL.Path)
	logger.Println(req.ToJSONString())
	res := NewResponse(str)
	b, err := json.Marshal(res)
	if err != nil {
		w.Write([]byte(err.Error()))
		logger.Println(err)
	} else {
		w.Write(b)
		logger.Println(res.ToJSONString())
	}
}

type Request struct {
	Type string `json:"type"`
	Timestamp string `json:"timestamp"`
	RequestPath string `json:"request_path"`
}

func NewRequest(reqPath string) *Request {
	return &Request{
		Type: "REQUEST",
		Timestamp:   time.Now().Format("2006-01-02T15:04:05.000"),
		RequestPath: reqPath,
	}
}

func (req *Request) ToJSONString() string {
	return ToJSONString(req)
}

type Response struct {
	Type string `json:"type"`
	Timestamp string `json:"timestamp"`
	ResponseMessage string `json:"response_message"`
}

func NewResponse(resMsg string) *Response {
	return &Response{
		Type: "RESPONSE",
		Timestamp:   time.Now().Format("2006-01-02T15:04:05.000"),
		ResponseMessage: resMsg,
	}
}

func (res *Response) ToJSONString() string {
	return ToJSONString(res)
}

func ToJSONString(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		errMsg := fmt.Sprintf("{ 'error': %s }", err.Error())
		return errMsg
	}
	return string(b)
}
