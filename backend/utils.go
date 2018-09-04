package main

import (
	"log"
	
	"net/http"
	"net/url"
	"io/ioutil"
	
	"strconv"
	
	"github.com/axgle/mahonia"
)

var (
	gbk_encoder = mahonia.NewEncoder("gbk")
	gbk_decoder = mahonia.NewDecoder("gbk")
)


func StringToFloat(str string) float32 {
	num, err := strconv.ParseFloat(str, 32)
	if err != nil {
		num = 0
	}
	return float32(num)
}

func BytesToFloat(str []byte) float32 {
	return StringToFloat(string(str))
}
	
func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		num = 0
	}
	return num
}

func BytesToInt(str []byte) int {
	return StringToInt(string(str))
}

func HttpGet(requrl string) []byte {
	res, err := http.Get(requrl)
	if err != nil {
		log.Panic("无法连接至服务器：", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic("无法读取响应数据：", err)
	}
	return []byte(gbk_decoder.ConvertString(string(body)))
}

func HttpPost(requrl string, args url.Values) []byte {
	for k, _ := range args {
		for v, _ := range args[k] {
			args[k][v] = gbk_encoder.ConvertString(args[k][v])
		}
	}
	res, err := http.PostForm(requrl, args)
    if err != nil {
		log.Panic("无法连接至服务器：", err)
    }
    defer res.Body.Close()
	
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
		log.Panic("无法读取响应数据：", err)
    }
	return []byte(gbk_decoder.ConvertString(string(body)))
}
