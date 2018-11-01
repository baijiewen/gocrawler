package parser

import (
	"fmt"
	"regexp"
	"singlecrawler/crawler/engine"
	"singlecrawler/crawler/config"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	fmt.Printf("find all matches: %s\n", matches)
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items,"city " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser:  engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	fmt.Printf("test result: %s\n", result)
	return result
}
