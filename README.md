# Gitdump
Simple tool for downloading Github repositories.

### Available commands
* gitdump users ls **<USER>** (no auth)
* gitdump users get **<USER> <REPO>** (no auth)
* gitdump users dump **<USER>** (no auth)
* gitdump owners ls (auth)
* gitdump owners get **<OWNER> <REPO>** (auth)
* gitdump owners dump **<OWNER>** (auth)
* gitdump orgs ls (auth)
* gitdump orgs ls **<ORG>** (auth)
* gitdump orgs get **<ORG> <REPO>** (auth)
* gitdump orgs dump **<ORG>** (auth)

### Commands
* **ls** list the organizations or repositories
* **get** download a single repository
* **dump** download all the repositories

### Arguments
* **USER** public username
* **OWNER** account associated with the auth token
* **ORG** name of one organization that has been granted
* **REPO** repository name

### Scopes
* **users** any public GitHub repository
* **owners** GitHub account associated with the auth token
* **orgs** granted organization(s) commands

### Example
To download one repository in the [https://github.com/GoogleContainerTools](https://github.com/GoogleContainerTools) account, you can use the following set of commands:

* gitdump users ls GoogleContainerTools
* gitdump users get GoogleContainerTools skaffold

### Authentication
TBD