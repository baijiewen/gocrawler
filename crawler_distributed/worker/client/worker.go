package client

import (
	"net/rpc"
	"singlecrawler/crawler/engine"
	"singlecrawler/crawler_distributed/worker"
	"singlecrawler/crawler_distributed/config"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error){
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParserResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil{
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}

