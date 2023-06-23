package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	go writeToEs(1)
	go writeToEs(2)
	go writeToEs(3)
	go writeToEs(4)
	go writeToEs(5)
	go writeToEs(6)
	go writeToEs(7)
	go writeToEs(8)
	//fmt.Println(GetGID)
	//queryES()
	wg.Add(1)
	wg.Wait()
	//rand.Seed(time.Now().UnixNano())
	//character := GenerateChineseCharacter(rand.Intn(100))
	//s := string(character)
	//fmt.Println(s)
}

type name struct {
	Age      int    `json:"age"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Context1 string `json:"context1"`
	Context2 string `json:"context2"`
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var sex = [2]string{"男", "女"}
var ops uint64 = 0
var wg sync.WaitGroup

func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func writeToEs(i int) {
	fmt.Println(GetGID)
	rand.Seed(time.Now().UnixNano())
	context1 := strings.Builder{}
	context1.WriteString(GenerateChineseCharacter(rand.Intn(5)))
	context1.WriteString("连运动到")
	context1.WriteString(string(rand.Intn(50)))
	context1.WriteString("经度")
	context1.WriteString(string(rand.Intn(50)))
	context1.WriteString("纬度")
	context2 := strings.Builder{}
	context2.WriteString(GenerateChineseCharacter(rand.Intn(5)))
	context2.WriteString("连执行打击")
	context2.WriteString(GenerateChineseCharacter(rand.Intn(5)))
	context2.WriteString("目标任务")
	fmt.Println(context1.String())
	for {
		time.Sleep(100 * time.Millisecond)
		name := &name{
			Age:  rand.Intn(100),
			Name: RandomString(rand.Intn(20), defaultLetters),
			//Name:     "是个大染缸的然后突然很痛很痛回家",
			Sex:      sex[rand.Intn(2)],
			Context1: context1.String(),
			Context2: context2.String(),
		}
		b, _ := json.Marshal(name)
		fmt.Println(string(b))
		//resp, err := http.Post("http://192.168.1.28:9200/student", "application/json", bytes.NewBuffer(b))
		//request, _ := http.NewRequest("PUT", "http://192.168.1.28:9200/student", bytes.NewBuffer(b))
		//strconv.Itoa(rand.Intn(5)
		//url := strings.Join([]string{"http://192.168.1.28:9200/intelligence/type", "1"}, "")
		//fmt.Println(url)
		go func() {
			request, _ := http.NewRequest("POST", "http://192.168.1.28:9200/intelligence/_doc", bytes.NewBuffer(b))
			request.Header.Set("Content-Type", "application/json")
			fmt.Println("请求开始", time.Now().UnixNano()/1e6)
			//resp, err := http.DefaultClient.Do(request)
			_, _ = http.DefaultClient.Do(request)
			//defer resp.Body.Close()
			//if err != nil {
			//	fmt.Println(err)
			//} else {
			//f, _ := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
			//defer func() {
			//	f.Close()
			//}()
			//log.SetOutput(f)
			//body, _ := ioutil.ReadAll(resp.Body)
			//fmt.Println("请求完成", string(body), time.Now().UnixNano()/1e6)
			//log.Println("请求完成", string(body), time.Now().UnixNano()/1e6)
			//}
		}()
		//ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))
		//atomic.AddUint64(&ops, 1)
		//runtime.Gosched()
		fmt.Println("gogogogogogog", i)
	}
	wg.Done()
}

func queryES() {
	var jsonStr = []byte(`{
	"query":{
		"match":{
			"name":"是个大染缸的然后突然很痛很痛回家"
		}
	}
}`)
	for {
		i := 0
		//go func() {
		request, _ := http.NewRequest("GET", "http://192.168.1.28:9200/intelligence/_search", bytes.NewBuffer(jsonStr))
		request.Header.Set("Content-Type", "application/json")
		fmt.Println("查询开始", i, time.Now().UnixNano()/1e6)
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		//defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		s := string(body)
		if strings.Index(s, "是个大染缸的然后突然很痛很痛回家") != -1 {
			fmt.Println("查询结束", i, time.Now().UnixNano()/1e6, string(body))
			return
		}
		//}()
		i++
	}
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func GenerateChineseCharacter(length int) string {
	a := make([]rune, length)
	for i := range a {
		a[i] = rune(RandInt(19968, 40869))
	}
	//fmt.Println(string(a))
	//fmt.Println(len(a))
	return string(a)
}

func RandInt(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}
