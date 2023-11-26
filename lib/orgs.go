package lib

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/google/go-github/github"
)

const (
	zipballFormat = "zipball"
)

// ListOrganizations fetches and logs a list of GitHub organizations.
func ListOrganizations(args []string) error {
	ctx, client := GitLogin()

	opt := &github.ListOptions{PerPage: perPage}
	var allOrganizations []*github.Organization

	for {
		organizations, resp, err := client.Organizations.List(ctx, "", opt)
		if err != nil {
			return err
		}
		allOrganizations = append(allOrganizations, organizations...)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	log.Println("\nOrganizations:")
	for _, organization := range allOrganizations {
		log.Println(fmt.Sprintf("  %s", *organization.Login))
	}

	return nil
}

// ListOrganizationRepositories fetches and logs a list of repositories for a GitHub organization.
func ListOrganizationRepositories(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("organization name not provided")
	}

	ctx, client := GitLogin()

	org := args[0]
	opt := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: perPage}}
	var allRepos []*github.Repository

	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			return err
		}
		allRepos = append(allRepos, repos...)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	log.Println("\nOrganization repositories:")
	for _, repo := range allRepos {
		log.Println(fmt.Sprintf("  %s", *repo.Name))
	}

	return nil
}

// GetOrganizationRepository downloads the zipball of a specific repository in an organization.
func GetOrganizationRepository(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Requires precisely 2 arguments")
	}

	ctx, client := GitLogin()
	owner, repo := args[0], args[1]

	url, _, err := client.Repositories.GetArchiveLink(ctx, owner, repo, zipballFormat, nil)
	if err != nil {
		return err
	}

	path, _ := MakeDir(owner)
	log.Printf("Downloading %s/%s", owner, repo)

	if err := DownloadFile(filepath.Join(path, fmt.Sprintf("%s.zip", repo)), url.String()); err != nil {
		return err
	}

	return nil
}

// DumpOrganizationRepositories downloads the zipballs of all repositories in an organization.
func DumpOrganizationRepositories(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Requires precisely 1 argument")
	}

	org := args[0]
	ctx, client := GitLogin()

	opt := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: perPage}}

	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			return err
		}

		path, _ := MakeDir(org)

		for _, repo := range repos {
			log.Printf("Downloading %s/%s", org, *repo.Name)

			url, _, err := client.Repositories.GetArchiveLink(ctx, org, *repo.Name, zipballFormat, nil)
			if err != nil {
				log.Printf("Error fetching archive link for %s/%s: %v", org, *repo.Name, err)
				continue
			}

			if err := DownloadFile(filepath.Join(path, fmt.Sprintf("%s.zip", *repo.Name)), url.String()); err != nil {
				log.Printf("Error downloading %s/%s: %v", org, *repo.Name, err)
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
