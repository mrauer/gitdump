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
