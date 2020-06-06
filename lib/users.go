package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 3 * time.Second}
var results = make(map[string]ApiResult)

const API_ROOT_USERS = "https://api.github.com/users/"

type ApiResult struct {
	Name string `json:"name"`
	Url  string `json:"svn_url"`
}

func ListPublicRepositories(args []string) {
	if len(args) == 1 {
		r, err := client.Get(API_ROOT_USERS + args[0] + "/repos?per_page=1000")
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		data := []ApiResult{}
		err = json.Unmarshal([]byte(b), &data)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("\nPublic repositories:")
		for i := 0; i < len(data); i++ {
			results[data[i].Name] = data[i]
			fmt.Println(fmt.Sprintf("  %s", data[i].Name))
		}
	}
}

func GetPublicRepository(args []string) {
	if len(args) == 2 {
		user := args[0]
		repo := args[1]
		r, err := client.Get(API_ROOT_USERS + args[0] + "/repos?per_page=1000")
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		data := []ApiResult{}
		err = json.Unmarshal([]byte(b), &data)
		if err != nil {
			fmt.Println(err.Error())
		}
		for i := 0; i < len(data); i++ {
			results[data[i].Name] = data[i]
		}
		fmt.Println(fmt.Sprintf("Downloading %s", repo))
		// Today's directory
		path, _ := MakeDir(user)
		url := fmt.Sprintf("%s/archive/master.zip", results[repo].Url)
		if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, results[repo].Name), url); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func DumpPublicRepositories(args []string) {
	if len(args) == 1 {
		user := args[0]
		r, err := client.Get(API_ROOT_USERS + user + "/repos?per_page=1000")
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		data := []ApiResult{}
		err = json.Unmarshal([]byte(b), &data)
		if err != nil {
			fmt.Println(err.Error())
		}
		// Today's directory
		path, _ := MakeDir(user)
		for i := 0; i < len(data); i++ {
			fmt.Println(fmt.Sprintf("Downloading %s", data[i].Name))
			url := fmt.Sprintf("%s/archive/master.zip", data[i].Url)
			if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, data[i].Name), url); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
