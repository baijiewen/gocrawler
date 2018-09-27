package parser

import (
	"fmt"
	"regexp"
	"singlecrawler/crawler/engine"
	"singlecrawler/crawler/model"
	"strconv"
)

var (
	ageRe      = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
	marrageRe  = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	weightRe   = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
	heightRe   = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
	incomeRe   = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	genderRe   = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	eduRe      = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	jobRe      = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
	hometownRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	starRe     = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
	houseRe    = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe      = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
	idUrlRe    = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name
	profile.Age = extracInt(contents, ageRe)
	profile.Height = extracInt(contents, heightRe)
	profile.Weight = extracInt(contents, weightRe)
	profile.Marriage = extractString(contents, marrageRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, eduRe)
	profile.Occupation = extractString(contents, jobRe)
	profile.Hukou = extractString(contents, hometownRe)
	profile.Xinzuo = extractString(contents, starRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{Url: url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile},
		},
	}
	return result
}

func extracInt(contents []byte, re *regexp.Regexp) int {
	match := extractString(contents, re)
	tmp, err := strconv.Atoi(match)
	if err != nil {
		fmt.Errorf("convert string to int failed!!!!!!%s", err)
	}
	return tmp
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
