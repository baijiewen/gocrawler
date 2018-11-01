package persist

import (
	"context"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"singlecrawler/crawler/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.247.130:9200"),
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+"#%d: %v", itemCount, item)
			itemCount++

			err := Save(client, item, index)
			if err != nil {
				log.Print("Item Saver: error "+"saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) (err error) {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().Index(index).Type(item.Type).Id(item.Id).BodyJson(item)
	if item.Id != "" {
		indexService.Do(context.Background())
	}
	if err != nil {
		return err
	}
	return nil
}
