package main

import (
	"testing"
	"singlecrawler/crawler_distributed/rpcsupport"
	"singlecrawler/crawler/engine"
	"singlecrawler/crawler/model"
	"imoocGo/crawler_distributed/config"
	"time"
)

func TestItemSaver(t *testing.T)  {
	const host = ":1234"

	go serveRpc(host, "test1")
	time.Sleep(time.Second)
	client, err:=rpcsupport.NewClient(host)
	if err != nil{
		panic(err)
	}

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1687783581",
		Type: "zhenai",
		Id:   "1687783581",
		Payload: model.Profile{Name: "冷颜",
			Gender:     "女",
			Age:        29,
			Height:     155,
			Weight:     0,
			Income:     "3000元以下",
			Marriage:   "离异",
			Education:  "高中及以下",
			Occupation: "后勤",
			Hukou:      "云南曲靖",
			Xinzuo:     "牡羊座",
			House:      "未购房",
			Car:        "未购车"},
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err!= nil || result != "ok"{
		t.Errorf("result: %s; err: %s",result, err)
	}

}