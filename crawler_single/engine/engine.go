package engine

import (
	"learning/TheGoProgrammingLanguage/Google_practise/crawler/fetcher"
	"log"
)

func Run(seeds ...Request)  {
	requests:=[]Request{}
	for _,r:=range seeds{
		requests= append(requests, r)
	}

	for len(requests)>0{
		r:=requests[0]
		requests=requests[1:]

		log.Printf("Fetching %s\n",r.Url)

		body,err:=fetcher.Fetch(r.Url)
		if err!=nil{
			log.Printf("Error fetch" +
				"Fetching Url %s : %v\n",r.Url,err)
			continue
		}

		ParseResult:= r.ParserFunc(body)

		requests= append(requests, ParseResult.Requests...)

		for _,item:=range ParseResult.Items{
			log.Printf("Got item :%v\n",item)
		}
	}
}

























//func Run(seeds ...Request)  {
//	var requests []Request
//	for _,r:=range seeds{
//		requests=append(requests,r)
//	}
//
//	for len(requests)>0{
//		r:=requests[0]
//		requests=requests[1:]
//
//		log.Printf("Fetching %s",r.Url)
//
//		body,err:=fetcher.Fetch(r.Url)
//		if err!=nil{
//			log.Printf("Fetcher: error" +
//				"fetching url %s: %v ",r.Url,err)
//			continue
//		}
//
//		ParseResult:=r.ParserFunc(body)
//		requests=append(requests,
//			ParseResult.Requests...)
//
//		for _,item:=range ParseResult.Items{
//			log.Printf("Got item %v",item)
//		}
//
//	}
//}
