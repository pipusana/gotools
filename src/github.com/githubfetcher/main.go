package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ReposInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	resp, err := http.Get("http://www.mocky.io/v2/5b52fd522f0000510d3bb683")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var repos []ReposInfo
	error := json.Unmarshal(body, &repos)

	if error != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		fmt.Printf("%+v\n", repo)
	}

}
