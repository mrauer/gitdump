package lib

import (
	"fmt"
	"github.com/google/go-github/github"
)

func GetOrgsPrivateRepository(args []string) error {
	ctx, client := GitLogin()

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
		path, _ := MakeDir(args[0])

		fmt.Println(fmt.Sprintf("Downloading %s", args[1]))
		if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, args[1]), url.String()); err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}

func GetPrivateRepositories(organization string) error {
	ctx, client := GitLogin()

	opt := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	repos, _, err := client.Repositories.ListByOrg(ctx, organization, opt)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Today's directory
	path, _ := MakeDir(organization)

	fmt.Println("\nDownloading all repositories:\n")
	for _, repo := range repos {
		fmt.Println(fmt.Sprintf("Downloading %s", *repo.Name))
		url, _, err := client.Repositories.GetArchiveLink(ctx, organization, *repo.Name, "zipball", nil, true)
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
