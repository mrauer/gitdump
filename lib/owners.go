package lib

import (
	"fmt"
	"log"

	"github.com/google/go-github/github"
)

const (
	repoType = "owner"
)

// ListPrivateRepositories fetches and logs a list of private repositories for the authenticated user.
func ListPrivateRepositories() error {
	ctx, client := GitLogin()

	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: perPage}, Type: repoType}

	for {
		repos, resp, err := client.Repositories.List(ctx, "", opt)
		if err != nil {
			return err
		}

		log.Println("\nPrivate repositories:")
		for _, repo := range repos {
			log.Println(fmt.Sprintf("  %s", *repo.Name))
		}

		// Check if there are more pages to retrieve
		if resp.NextPage == 0 {
			break
		}

		// Update the options to fetch the next page
		opt.Page = resp.NextPage
	}

	return nil
}

// GetPrivateRepository downloads the zipball of a specific private repository.
func GetPrivateRepository(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Requires precisely 2 arguments")
	}

	ctx, client := GitLogin()
	owner, repo := args[0], args[1]

	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: perPage}}

	for {
		repos, resp, err := client.Repositories.List(ctx, owner, opt)
		if err != nil {
			return err
		}

		// Find the specified repository in the list
		var targetRepo *github.Repository
		for _, r := range repos {
			if *r.Name == repo {
				targetRepo = r
				break
			}
		}

		if targetRepo == nil {
			return fmt.Errorf("repository %s not found in user %s's account", repo, owner)
		}

		url, _, err := client.Repositories.GetArchiveLink(ctx, owner, *targetRepo.Name, github.Zipball, nil)
		if err != nil {
			return err
		}

		path, _ := MakeDir(owner)
		log.Printf("Downloading %s", repo)

		if err := DownloadFile(fmt.Sprintf("%s/%s.zip", path, repo), url.String()); err != nil {
			return err
		}

		// Check if there are more pages to retrieve
		if resp.NextPage == 0 {
			break
		}

		// Update the options to fetch the next page
		opt.Page = resp.NextPage
	}

	return nil
}

// DumpPrivateRepositories downloads the zipballs of all private repositories for the authenticated user.
func DumpPrivateRepositories(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Requires precisely 1 argument")
	}

	owner := args[0]
	ctx, client := GitLogin()

	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: perPage}, Type: repoType}

	for {
		repos, resp, err := client.Repositories.List(ctx, "", opt)
		if err != nil {
			return err
		}

		path, _ := MakeDir(owner)

		for _, repo := range repos {
			log.Printf("Downloading %s", *repo.Name)

			url, _, err := client.Repositories.GetArchiveLink(ctx, owner, *repo.Name, github.Zipball, nil)
			if err != nil {
				log.Printf("Error fetching archive link for %s: %v", *repo.Name, err)
				continue
			}

			if err := DownloadFile(fmt.Sprintf("%s/%s.zip", path, *repo.Name), url.String()); err != nil {
				log.Printf("Error downloading %s: %v", *repo.Name, err)
			}
		}

		// Check if there are more pages to retrieve
		if resp.NextPage == 0 {
			break
		}

		// Update the options to fetch the next page
		opt.Page = resp.NextPage
	}

	return nil
}
