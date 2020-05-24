package lib

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

func GetPrivateRepository(args []string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GIT_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	if len(args) < 1 {
		opts := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 30}}
		repos, _, err := client.Repositories.List(ctx, "", opts)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		fmt.Println("\nUsage: gitdump orgs get <REPOSITORY>\n")

		for _, repo := range repos {
			fmt.Println(fmt.Sprintf("- %s", *repo.Name))
		}
	} else {
		fmt.Println("Will download here")
	}
	return nil
}
