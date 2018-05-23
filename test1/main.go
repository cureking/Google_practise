package main

import (
	"net/http"
	"fmt"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"regexp"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func main() {
	resp,err:=http.Get("http://www.zhenai.com/zhenghun")
	if err!=nil{
		panic(err)
	}
	if resp.StatusCode!=http.StatusOK{
		fmt.Printf("Error" +
			"statusCode:%d",resp.StatusCode)
	}
	e:=determineEncoding(resp.Body)
	utf8Reader:=transform.NewReader(resp.Body,e.NewDecoder())
	all,err:=ioutil.ReadAll(utf8Reader)
	if err!=nil{
		panic(all)
	}
	printCityList(all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	p,err:=bufio.NewReader(r).Peek(1024)
	if err!=nil{
		panic(err)
	}
	e,_,_:=charset.DetermineEncoding(p,"")
	return e
}

func printCityList(contents []byte) [][][]byte {
	re:=regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]+>([^<]+)</a>`)
	matches:=re.FindAllSubmatch(contents,-1)
	for _,m:=range matches{
		fmt.Printf("City:%s URL:%s \n",m[2],m[1])
	}
	return matches
}