package main

import (
	"log"
	"net/http"
	"qasim/base"
)

func main()  {
	eng := new(base.Engine)
	log.Fatal(http.ListenAndServe(":9999", eng))
}
