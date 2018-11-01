package main

import (
	"testing"
	"singlecrawler/crawler_distributed/rpcsupport"
	"singlecrawler/crawler_distributed/worker"
	"singlecrawler/crawler_distributed/config"
	"time"
	"fmt"
)

func TestCrawlService(t *testing.T)  {
	const host  = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil{
		panic(err)
	}

	req := worker.Request{
		Url:"http://album.zhenai.com/u/1687783581",
		Parser: worker.SerializedParser{
			Name:config.ParseProfile,
			Args:"冷颜",
		},
	}
	var result worker.ParserResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil{
		t.Error(err)
	}else{
		fmt.Println(result)
	}
}