package persist

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"singlecrawler/crawler/engine"
	"singlecrawler/crawler/model"
	"testing"
)

func TestItemSaver(t *testing.T) {
	profile := engine.Item{
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
	const index = "bjw_test"
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.247.130:9200"),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	err = Save(client, profile, index)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index(index).Type(profile.Type).Id(profile.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != profile {
		t.Errorf("Got %v; expected %v", actual, profile)
	}
}
