package main

import (
	"fmt"
	"learning/TheGoProgrammingLanguage/ch4/retriver/real"
)

type Retriever interface {
	Get(string) string
}

type Poster interface {
	Post(url string, string map[string]string) string
}

func download(r Retriever, url string) string {
	return r.Get(url)
}

func post(poster Poster, url string) {
	poster.Post(url, map[string]string{})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func Session(s RetrieverPoster) {
	//s.Get()
	//s.Post()
}

func main() {
	var r Retriever
	r = real.Retriever{"", 0, "Ending"}
	fmt.Println(download(r, "http://www.sohu.com"))

}

//type Retriever interface {
//	Get(string) string
//}
//
//func download(r Retriever, url string) string {
//	return r.Get(url)
//}
//
//func main() {
//	var r Retriever
//	r = mock.Restriever{"this is test at localhost!"}
//	fmt.Println(download(r, "http://www.sohu.com"))
//	r = real.Retriever{"", 0, "Ending"}
//	fmt.Println(download(r, "http://www.sohu.com"))
//	fmt.Print()
//}
