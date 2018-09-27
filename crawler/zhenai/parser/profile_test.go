package parser

import (
	"fmt"
	"io/ioutil"
	"singlecrawler/crawler/engine"
	"singlecrawler/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "http://album.zhenai.com/u/1687783581", "冷颜")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
	fmt.Printf("TEST BJW %v", result)
	actual := result.Items[0]

	expected := engine.Item{
		"http://album.zhenai.com/u/1687783581",
		"zhenai",
		"1687783581",
		model.Profile{"冷颜", "女", 29, 155, 0,
			"3000元以下", "离异", "高中及以下", "后勤",
			"云南曲靖", "", "租房", "未购车"},
	}
	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}

}
