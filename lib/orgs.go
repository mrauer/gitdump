package lib

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"time"
)

func GetPrivateRepository(args []string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GIT_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	fmt.Println("\nUsage: gitdump orgs get <ORG> <REPOSITORY>\n")
	if len(args) == 0 {
		organizations, _, err := client.Organizations.List(ctx, "", nil)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		for _, organization := range organizations {
			fmt.Println(fmt.Sprintf("- %s", *organization.Login))
		}
	}

	if len(args) == 1 {
		// list all repositories for the authenticated user
		opt := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 1000}}
		repos, _, err := client.Repositories.ListByOrg(ctx, args[0], opt)

		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		for _, repo := range repos {
			fmt.Println(fmt.Sprintf("- %s", *repo.Name))
		}
	}

	if len(args) == 2 {
		url, _, err := client.Repositories.GetArchiveLink(ctx, args[0], args[1], "zipball", nil, true)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		// Today's directory
		path := fmt.Sprintf("data/%s/%s", args[0], time.Now().Format("2006-01-02"))
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.MkdirAll(path, 0700)
		}

		fmt.Println(fmt.Sprintf("Downloading %s", args[1]))
		if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, args[1]), url.String()); err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}
