package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/mrauer/repos")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var repos []map[string]interface{}
	if err := json.Unmarshal(body, &repos); err != nil {
		panic(err)
	}
	fmt.Println(repos[0]["url"])
}
