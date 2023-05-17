package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.apilayer.com/financelayer/news?tickers=tickers&tags=tags&sources=sources&sort=sort&offset=offset&limit=limit&keywords=keywords&fallback=fallback&date=date"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", "Z3xFtlug4p2GbfKKQgel9jX8CdBtBffw")

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
