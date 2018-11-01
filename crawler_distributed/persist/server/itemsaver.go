package main

import (
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"singlecrawler/crawler_distributed/rpcsupport"
	"singlecrawler/crawler_distributed/persist"
	"singlecrawler/crawler_distributed/config"
	"github.com/gpmgo/gopm/modules/log"
)

var port = flag.Int("port", 0,"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0{
		fmt.Println("must specify a port")
		return
	}
	err := serveRpc(
		fmt.Sprintf(":%d", *port),
		config.ElasticIndex)
	if err!= nil{
		log.Fatal("Server run failed")
	}
}

func serveRpc(host,index string) error {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.247.130:9200"),
		elastic.SetSniff(false))
	if err!= nil{
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client:client,
		Index:index,
	})
}