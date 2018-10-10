package model

type SearchResult struct {
	Query string
	Hits int64
	Start int
	PrevFrom int
	NextFrom int
	Items []interface{}
}
