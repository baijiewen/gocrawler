package main

import (
	"singlecrawler/crawler_distributed/rpcsupport"
	"fmt"
	"singlecrawler/crawler_distributed/worker"
	"singlecrawler/crawler_distributed/config"
	"log"
)

func main() {
	err := rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawService{},
	)
	if err != nil{
		log.Fatal(err)
	}
}
