package worker

import (
	"singlecrawler/crawler/engine"
	"imoocGo/crawler/config"
	"singlecrawler/crawler/zhenai/parser"
	config2 "github.com/talent4636/imoocGo/crawler/config"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url string
	Parser SerializedParser
}

type ParserResult struct {
	Items []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser:SerializedParser{
			Name:name,
			Args:args},
	}
}

func SerializeResult(r engine.ParseResult) ParserResult {
	result := ParserResult{
		Items:r.Items,
	}
	for _,req := range r.Requests{
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request,error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil{
		return engine.Request{}, nil
	}
	return engine.Request{Url:r.Url,Parser:parser}, nil
}

func DeserializeResult(r ParserResult) engine.ParseResult {
	result := engine.ParseResult{
		Items:r.Items,
	}
	for _, req := range r.Requests{
		engineReq, err := DeserializeRequest(req)
		if err != nil{
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseProfile:
		if userName, ok:= p.Args.(string);ok{
			return parser.NewProfileParser(userName), nil
		}else{
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	case config2.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}

