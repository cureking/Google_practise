package main

import (
	"learning/TheGoProgrammingLanguage/Google_practise/crawler/engine"
	"learning/TheGoProgrammingLanguage/Google_practise/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}