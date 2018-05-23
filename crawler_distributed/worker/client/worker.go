package client

import (
	"net/rpc"

	"learning/TheGoProgrammingLanguage/Google_practise/learngo/crawler/engine"
	"learning/TheGoProgrammingLanguage/Google_practise/learngo/crawler_distributed/config"
	"learning/TheGoProgrammingLanguage/Google_practise/learngo/crawler_distributed/worker"
)

func CreateProcessor(
	clientChan chan *rpc.Client) engine.Processor {

	return func(
		req engine.Request) (
		engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc,
			sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult),
			nil
	}
}
