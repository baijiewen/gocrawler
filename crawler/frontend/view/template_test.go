package view

import (
	"testing"
	"os"
	"singlecrawler/crawler/frontend/model"
	"singlecrawler/crawler/engine"
common "singlecrawler/crawler/model"
)

func TestTemplate(t *testing.T)  {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url: "http://album.zhenai.com/u/106297033",
		Type: "zhenai",
		Id:"106297033",
		Payload: common.Profile{
			Name : "晨曦",
			Gender: "女",
			Age : 39,
			Height : 158,
			Weight : 53,
			Income : "5001-8000元",
			Marriage : "未婚",
			Education : "大学本科",
			Occupation : "教育/科研",
			Hukou : "海南海口",
			Xinzuo : "狮子座",
			House : "和家人同住",
			Car : "未购车",
		}}
		for i:=0;i<10;i++{
			page.Items = append(page.Items, item)
		}

		err = view.Render(out, page)
		if err!=nil{
			panic(err)
		}
	}

