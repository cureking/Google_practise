package main

import (
	"fmt"

	"log"

	"flag"

	"learning/TheGoProgrammingLanguage/Google_practise/learngo/crawler/fetcher"
	"learning/TheGoProgrammingLanguage/Google_practise/learngo/crawler_distributed/rpcsupport"
	"learning/TheGoProgrammingLanguage/Google_practise/learngo/crawler_distributed/worker"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	fetcher.SetVerboseLogging()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
