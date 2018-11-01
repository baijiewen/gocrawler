package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	Submit(Request)
	//ConfigureMasterWorkerChan(chan Request)
	ReadyNotifier
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	//var isDupmap map[string]bool

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		fmt.Printf("test bjw :%s\n", r.Url)
		if isDuplicate(r.Url) {
			fmt.Printf("has been visited %s\n", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				fmt.Printf("has been visited %s\n", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var isDupmap = make(map[string]bool)

func isDuplicate(url string) bool {
	if isDupmap[url] {
		return true
	}
	isDupmap[url] = true
	return false
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
