package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var client = &http.Client{Timeout: 3 * time.Second}
var results = make(map[string]ApiResult)

const API_ROOT_USERS = "https://api.github.com/users/"

type ApiResult struct {
	Name string `json:"name"`
	Url  string `json:"svn_url"`
}

func GetPublicRepository(args []string) {
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

	if len(args) == 1 {
		fmt.Println("\nAvailable repositories:\n")
		for i := 0; i < len(data); i++ {
			results[data[i].Name] = data[i]
			fmt.Println(fmt.Sprintf("- %s", data[i].Name))
		}
		fmt.Println("\nTo download: gitdump users get <USERNAME> <REPOSITORY>\n")
	}

	if len(args) == 2 {
		repository := args[1]

		for i := 0; i < len(data); i++ {
			results[data[i].Name] = data[i]
		}

		fmt.Println(fmt.Sprintf("\nDownloading %s\n", repository))

		// Today's directory
		path := fmt.Sprintf("data/%s", time.Now().Format("2006-01-02"))
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		}

		url := fmt.Sprintf("%s/archive/master.zip", results[repository].Url)

		if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, results[repository].Name), url); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func GetPublicRepositories(account string) {
	r, err := client.Get(API_ROOT_USERS + account + "/repos?per_page=1000")
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
	path := fmt.Sprintf("data/%s/%s", account, time.Now().Format("2006-01-02"))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}

	fmt.Println("\nDownloading all repositories:\n")
	for i := 0; i < len(data); i++ {
		fmt.Println(fmt.Sprintf("Downloading %s", data[i].Name))
		url := fmt.Sprintf("%s/archive/master.zip", data[i].Url)
		if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, data[i].Name), url); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
