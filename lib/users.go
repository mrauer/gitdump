package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	apiRootUsers   = "https://api.github.com/users/"
	pageQueryParam = "page"
)

var (
	client  = &http.Client{Timeout: 3 * time.Second}
	results = make(map[string]ApiResult)
)

// ApiResult represents the structure of the API response.
type ApiResult struct {
	Name string `json:"name"`
	Url  string `json:"svn_url"`
}

// ListPublicRepositories fetches and logs a list of public repositories for a GitHub user.
func ListPublicRepositories(args []string) {
	if len(args) != 1 {
		log.Fatal("Requires precisely 1 argument")
	}

	username := args[0]
	page := 1

	for {
		// Make the API request for the current page
		r, err := client.Get(fmt.Sprintf("%s%s/repos?per_page=%d&%s=%d", apiRootUsers, username, perPage, pageQueryParam, page))
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()

		// Check for non-successful HTTP status code
		if r.StatusCode != http.StatusOK {
			log.Fatalf("Error: %s", r.Status)
		}

		// Read and parse the response body
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		var data []ApiResult
		err = json.Unmarshal(b, &data)
		if err != nil {
			log.Printf("Error decoding JSON: %s", err)
			break
		}

		// Break the loop if there are no more pages
		if len(data) == 0 {
			break
		}

		// Process the repositories for the current page
		log.Printf("\nPublic repositories - Page %d:\n", page)
		for i := 0; i < len(data); i++ {
			results[data[i].Name] = data[i]
			log.Println(fmt.Sprintf("  %s", data[i].Name))
		}

		// Move to the next page
		page++
	}
}

// GetPublicRepository downloads the zipball of a specific public repository for a GitHub user.
func GetPublicRepository(args []string) {
	if len(args) == 2 {
		user := args[0]
		repo := args[1]
		page := 1

		for {
			// Make the API request for the current page
			r, err := client.Get(fmt.Sprintf("%s%s/repos?per_page=%d&%s=%d", apiRootUsers, user, perPage, pageQueryParam, page))
			if err != nil {
				log.Fatal(err)
			}
			defer r.Body.Close()

			// Read and parse the response body
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			data := []ApiResult{}
			err = json.Unmarshal([]byte(b), &data)
			if err != nil {
				log.Println(err.Error())
			}

			// Process the repositories for the current page
			for i := 0; i < len(data); i++ {
				results[data[i].Name] = data[i]

				// Check if the current repository is the one we're looking for
				if data[i].Name == repo {
					log.Printf("Downloading %s\n", repo)

					// Today's directory
					path, _ := MakeDir(user)
					url := fmt.Sprintf("%s/archive/master.zip", results[repo].Url)
					if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, results[repo].Name), url); err != nil {
						log.Println(err.Error())
					}

					// Repository found, so return from the function
					return
				}
			}

			// Move to the next page
			page++

			// Break the loop if there are no more pages
			if len(data) == 0 {
				log.Printf("Repository '%s' not found.\n", repo)
				break
			}
		}
	}
}

// DumpPublicRepositories downloads the zipballs of all public repositories for a GitHub user.
func DumpPublicRepositories(args []string) {
	if len(args) == 1 {
		user := args[0]
		page := 1

		for {
			// Make the API request for the current page
			r, err := client.Get(fmt.Sprintf("%s%s/repos?per_page=%d&%s=%d", apiRootUsers, user, perPage, pageQueryParam, page))
			if err != nil {
				log.Fatal(err)
			}
			defer r.Body.Close()

			// Read and parse the response body
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			data := []ApiResult{}
			err = json.Unmarshal([]byte(b), &data)
			if err != nil {
				log.Println(err.Error())
			}

			// Break the loop if there are no more pages
			if len(data) == 0 {
				break
			}

			// Today's directory
			path, _ := MakeDir(user)
			for i := 0; i < len(data); i++ {
				log.Printf("Downloading %s\n", data[i].Name)
				url := fmt.Sprintf("%s/archive/master.zip", data[i].Url)
				if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, data[i].Name), url); err != nil {
					log.Println(err.Error())
				}
			}

			// Move to the next page
			page++
		}
	}
}
