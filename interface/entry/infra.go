package main

import (
	"fmt"
	"learngo/interface/mock"
	"learngo/interface/retr"
	"time"
)

/**
定义接口
*/
type Retriever interface {
	Get(url string) string
}

type Post interface {
	Post(url string, param map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Post
}

const url = "http://www.baidu.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"content": "fake content"})
	return s.Get(url)
}

func getRetriever() Retriever {
	return &retr.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
}

func main() {
	var r Retriever = getRetriever()
	fmt.Printf("%T\n", r)

	if retireRetriever, ok := r.(*retr.Retriever); ok {
		fmt.Printf("%T %v\n", retireRetriever, retireRetriever)
	}

	switch v := r.(type) {
	case *retr.Retriever:
		v.Get("https://www.baidu.com")
	case mock.Retriever:
		v.Get("fake content")

	}

	//data := r.Get("https://www.baidu.com")
	//fmt.Println(data)

	mockRetriever := mock.Retriever{
		Content: "this is a fake content"}

	fmt.Println(session(&mockRetriever))

}
