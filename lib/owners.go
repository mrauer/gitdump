package lib

import (
	"fmt"
	"github.com/google/go-github/github"
)

func ListPrivateRepositories() {
	ctx, client := GitLogin()
	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}, Type: "owner"}
	repos, _, err := client.Repositories.List(ctx, "", opt)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("\nPrivate repositories:")
	for _, repo := range repos {
		fmt.Println(fmt.Sprintf("  %s", *repo.Name))
	}
}

func GetPrivateRepository(args []string) {
	if len(args) == 2 {
		ctx, client := GitLogin()
		owner := args[0]
		repo := args[1]
		url, _, err := client.Repositories.GetArchiveLink(ctx, owner, repo, github.Zipball, nil)
		if err != nil {
			fmt.Println(err.Error())
		}
		// Today's directory
		path, _ := MakeDir(owner)
		fmt.Println(fmt.Sprintf("Downloading %s", repo))
		if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, repo), url.String()); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func DumpPrivateRepositories(args []string) error {
	if len(args) == 1 {
		owner := args[0]
		ctx, client := GitLogin()
		opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}, Type: "owner"}
		repos, _, err := client.Repositories.List(ctx, "", opt)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		// Today's directory
		path, _ := MakeDir(owner)
		for _, repo := range repos {
			fmt.Println(fmt.Sprintf("Downloading %s", *repo.Name))
			url, _, err := client.Repositories.GetArchiveLink(ctx, owner, *repo.Name, github.Zipball, nil)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, *repo.Name), url.String()); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	return nil
}
