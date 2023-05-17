package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var apis map[int]string

// channel to store map[int]interface{}
var channel chan map[int]interface{}

func fetchData(API int) {
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {

			var result map[string]interface{}
			json.Unmarshal([]byte(body), &result)

			var re = make(map[int]interface{})

			switch API {
			case 1:
				if result["success"] == true {
					re[API] = result["rates"].(map[string]interface{})["USD"]
				} else {
					re[API] = result["error"].(map[string]interface{})["info"]
				}
				// store the result into the channel
				channel <- re
				fmt.Println("Result for API 1 stored")

			case 2:
				if result["main"] != nil {
					re[API] = result["main"].(map[string]interface{})["temp"]
				} else {
					re[API] = result["message"]
				}
				// store the result into the channel
				channel <- re
				fmt.Println("Result for API 2 stored")
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func main() {

	//channels are the pipes that connect the concurrent go routines
	// creates a channel to store the results from the
	// API calls
	channel = make(chan map[int]interface{})

	apis = make(map[int]string)
	apis[1] =
		"http://data.fixer.io/api/latest?access_key=" +
			"<access_key>"
	apis[2] =
		"http://api.openweathermap.org/data/2.5/" +
			"weather?q=SINGAPORE&appid=<app_id>"

	go fetchData(1)
	go fetchData(2)

	// we expect two results in the channel
	for i := 0; i < 2; i++ {
		fmt.Println(<-channel)
	}
	fmt.Println("Done!")

	fmt.Scanln()
}
