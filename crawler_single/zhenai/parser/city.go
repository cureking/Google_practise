package parser

import (
	"learning/TheGoProgrammingLanguage/Google_practise/crawler/engine"
	"regexp"
)

const cityRe=`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(
	contents []byte) engine.ParseResult {
	re:=regexp.MustCompile(cityRe)
	matches:=re.FindAllSubmatch(contents,-1)
	//fmt.Printf("%s",matches)

	result:=engine.ParseResult{}
	for _,m:=range matches{
		name:=m[2]
		result.Items=append(
			result.Items,"User "+string (name))
		result.Requests=append(
			result.Requests,engine.Request{
				Url:string(m[1]),
				ParserFunc: func(c []byte) engine.ParseResult {
					return  ParseProfile(c,string(name))
				},
		})
	}
	return result
}