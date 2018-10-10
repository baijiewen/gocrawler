package controller

import (
	"singlecrawler/crawler/frontend/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"strings"
	"strconv"
	"singlecrawler/crawler/frontend/model"
	"context"
	"reflect"
	"singlecrawler/crawler/engine"
	"regexp"
	"fmt"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func (h SearchResultHandler) ServeHTTP (w http.ResponseWriter, req *http.Request){
	q := strings.TrimSpace(req.FormValue("q"))

	from, err:=strconv.Atoi(req.FormValue("from"))
	if err!=nil{
		from =0
	}

	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil{
		http.Error(w, err.Error(),http.StatusBadRequest)
	}
	err = h.view.Render(w, page)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//fmt.Fprintf(w, "q=%s, from=%d",q,from)
}

func (h SearchResultHandler) getSearchResult(q string,from int) (model.SearchResult, error){
	var result model.SearchResult
	result.Query = q
	fmt.Println(rewriteQueryString(q))
	resp, err := h.client.Search("dating_profile").Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).Do(context.Background())

	if err != nil{
		return result,err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}


func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.247.130:9200"),
		elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}
	return SearchResultHandler{
		view: view.CreateSearchResultView(template),
		client:client,
	}
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}