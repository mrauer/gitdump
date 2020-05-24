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

	// list all repositories for the authenticated user
	if len(args) < 1 {
		opt := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 1000}}
		repos, _, err := client.Repositories.ListByOrg(ctx, "EENCloud", opt)

		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		fmt.Println("\nUsage: gitdump orgs get <REPOSITORY>\n")

		for _, repo := range repos {
			fmt.Println(fmt.Sprintf("- %s", *repo.Name))
		}

		url, _, _ := client.Repositories.GetArchiveLink(ctx, "EENCloud", "probe", "zipball", nil, true)

		// Today's directory
		path := fmt.Sprintf("data/%s", time.Now().Format("2006-01-02"))
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		}

		if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, "probe"), url.String()); err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Will download here")
	}
	return nil
}
