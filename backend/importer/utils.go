package importer

import (
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"net/url"

	"log"
	"strconv"
)

// 内部辅助函数，在转换失败时返回-1
func toInt(str []byte) int {
	num, err := strconv.Atoi(string(str))
	if err != nil {
		return -1
	}
	return num
}
func toFloat(str []byte) float32 {
	num, err := strconv.ParseFloat(string(str), 32)
	if err != nil {
		return -1
	}
	return float32(num)
}

// 发送HTTP/GET请求，并解码为UTF8
func httpGet(requrl []byte) []byte {
	res, err := http.Get(string(requrl))
	if err != nil {
		log.Printf("无法连接到服务器`%s`", requrl)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("读取回应消息失败")
		return nil
	}

	gbk_decoder := mahonia.NewDecoder("gbk")
	return []byte(gbk_decoder.ConvertString(string(body)))
}

// 发送HTTP/POST请求，并解码为UTF8
func httpPost(requrl []byte, args *url.Values) []byte {
	gbk_encoder := mahonia.NewEncoder("gbk")
	for k := range *args {
		for v := range (*args)[k] {
			(*args)[k][v] = gbk_encoder.ConvertString((*args)[k][v])
		}
	}

	res, err := http.PostForm(string(requrl), *args)
	if err != nil {
		log.Printf("无法连接到服务器`%s`", requrl)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("读取回应消息失败")
		return nil
	}

	gbk_decoder := mahonia.NewDecoder("gbk")
	return []byte(gbk_decoder.ConvertString(string(body)))
}

// 特殊查找字符串函数，匹配失败时直接放弃该位置的匹配。
// 返回第一个匹配的首位置
func find(src []byte, dst string, pos int) int {
	srcLen := len(src)
	dstLen := len(dst)
	for i, j := pos, 0; i < srcLen; i++ {
		if src[i] == dst[j] {
			j++
			if j == dstLen {
				return i - j + 1
			}
		} else {
			j = 0
		}
	}
	return -1
}

// 获取两个关键字中间的字符串
// 返回该字符串的首末位置 [start, end)
func between(src []byte, dst1 string, dst2 string, pos int) [2]int {
	srcLen := len(src)
	dstLen := len(dst1)
	for pos < srcLen {
		lpos := find(src, dst1, pos)
		if lpos == -1 {
			break
		}
		pos = lpos + dstLen
		rpos := find(src, dst2, pos)
		if rpos != -1 {
			return [2]int{lpos + dstLen, rpos}
		}
	}
	return [2]int{-1, -1}
}

// 获取所有“两个关键字中间的字符串”
// 返回包含所有字符串的数组
func allBetween(src []byte, dst1 string, dst2 string, pos int) [][]byte {
	res := make([][]byte, 0, 40)
	for {
		index := between(src, dst1, dst2, pos)
		if index[0] == -1 {
			break
		}
		res = append(res, src[index[0]:index[1]])
		pos = index[1]
	}
	return res
}
