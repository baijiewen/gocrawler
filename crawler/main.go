package main

import (
	"singlecrawler/crawler/engine"
	"singlecrawler/crawler/persist"
	"singlecrawler/crawler/scheduler"
	"singlecrawler/crawler/zhenai/parser"
	"singlecrawler/crawler/config"
)

func main() {
	ItemSaver, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    ItemSaver,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
