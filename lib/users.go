package lib

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
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
		path, _ := MakeDir(args[0])

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
	path, _ := MakeDir(account)

	fmt.Println("\nDownloading all repositories:\n")
	for i := 0; i < len(data); i++ {
		fmt.Println(fmt.Sprintf("Downloading %s", data[i].Name))
		url := fmt.Sprintf("%s/archive/master.zip", data[i].Url)
		if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, data[i].Name), url); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func GetUsersPrivateRepository(args []string) {
	ctx, client := GitLogin()

	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}, Type: "owner"}

	repos, _, err := client.Repositories.List(ctx, "", opt)

	if err != nil {
		fmt.Println(err.Error())
	}
	for _, repo := range repos {
		fmt.Println(fmt.Sprintf("- %s", *repo.Name))
	}
	fmt.Println("\nTo download: gitdump orgs private <OWNER> <REPOSITORY>\n")
	if len(args) == 2 {
		owner := args[0]
		repository := args[1]
		url, _, err := client.Repositories.GetArchiveLink(ctx, owner, repository, "zipball", nil, true)
		if err != nil {
			fmt.Println(err.Error())
		}
		// Today's directory
		path, _ := MakeDir(owner)

		fmt.Println(fmt.Sprintf("Downloading %s", repository))
		if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, repository), url.String()); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func GetUsersPrivateRepositories(owner string) error {
	ctx, client := GitLogin()

	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}, Type: "owner"}

	repos, _, err := client.Repositories.List(ctx, "", opt)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Today's directory
	path, _ := MakeDir(owner)

	fmt.Println("\nDownloading all repositories:\n")
	for _, repo := range repos {
		fmt.Println(fmt.Sprintf("Downloading %s", *repo.Name))
		url, _, err := client.Repositories.GetArchiveLink(ctx, owner, *repo.Name, "zipball", nil, true)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, *repo.Name), url.String()); err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}
