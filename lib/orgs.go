package lib

import (
	"fmt"
	"github.com/google/go-github/github"
)

func ListOrganizations(args []string) error {
	ctx, client := GitLogin()
	organizations, _, err := client.Organizations.List(ctx, "", nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("\nOrganizations:")
	for _, organization := range organizations {
		fmt.Println(fmt.Sprintf("  %s", *organization.Login))
	}
	return nil
}

func ListOrganizationRepositories(args []string) error {
	ctx, client := GitLogin()
	org := args[0]
	opt := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 1000}}
	repos, _, err := client.Repositories.ListByOrg(ctx, org, opt)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("\nOrganization repositories:")
	for _, repo := range repos {
		fmt.Println(fmt.Sprintf("  %s", *repo.Name))
	}
	return nil
}

func GetOrganizationRepository(args []string) {
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
		fmt.Println(fmt.Sprintf("Downloading %s", args[1]))
		if err = DownloadFile(fmt.Sprintf("%s/%s.zip", path, repo), url.String()); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func DumpOrganizationRepositories(args []string) error {
	if len(args) == 1 {
		org := args[0]
		ctx, client := GitLogin()
		opt := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 1000}}
		repos, _, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		// Today's directory
		path, _ := MakeDir(org)
		for _, repo := range repos {
			fmt.Println(fmt.Sprintf("Downloading %s", *repo.Name))
			url, _, err := client.Repositories.GetArchiveLink(ctx, org, *repo.Name, github.Zipball, nil)
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
