package main

import (
	"flag"
	"fmt"
	"time"

	"samh_ticket_rank_0.1/recommend/go/collection_api"
)

func main() {
	ComicId := flag.String("c", "291", "--漫画id")
	fmt.Println(*ComicId)
	flag.Parse()
	flag.Usage()
	start := time.Now()
	Json := collection_api.CidTicketApi(*ComicId)
	fmt.Println(Json)
	cost := time.Since(start)
	fmt.Println("cost=", cost)
}
