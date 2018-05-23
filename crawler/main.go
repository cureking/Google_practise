package main

import (
	"learning/TheGoProgrammingLanguage/Google_practise/learngo/crawler/config"
	"learning/TheGoProgrammingLanguage/Google_practise/learngo/crawler/engine"
	"learning/TheGoProgrammingLanguage/Google_practise/learngo/crawler/persist"
	"learning/TheGoProgrammingLanguage/Google_practise/learngo/crawler/scheduler"
	"learning/TheGoProgrammingLanguage/Google_practise/learngo/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.starter.url.here",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
