# GitDump
<!-- PROJECT LOGO -->
<p align="center">
  <a href="https://github.com/mrauer/gitdump">
    <img src="images/logo.png" alt="Logo">
  </a>

  <h3 align="center">GitDump</h3>

  <p align="center">
    Tool for downloading GitHub repositories.
    <br />
    <br />
    <a href="https://github.com/mrauer/gitdump/issues">Report Bug</a>
    <a href="https://github.com/mrauer/gitdump/issues">Request Feature</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
* [Installation](#installation)
  * [Linux](#linux)
  * [Mac](#mac)
* [Configuration](#configuration)
* [Usage](#usage)
* [Contribute](#contribute)
* [license](#license)

<!-- ABOUT THE PROJECT -->
## About The Project

GitDump is a tool that will give you the ability to download one or all the repositories in a GitHub account. This could be achieved for a public, private or an organization account. That way you get an archive of your work or somebody else work.

<!-- INSTALLATION -->
## Installation

This software can be used on Linux, Mac and Windows.

Just go to the releases page if you want to download the binaries at [https://github.com/mrauer/gitdump/releases](https://github.com/mrauer/gitdump/releases) or copy/paste the following commands if you're on Linux or Mac.

<!-- LINUX -->
### Linux

```sh
curl -Lo gitdump https://github.com/mrauer/gitdump/releases/download/v0.2.0/gitdump_0.2.0_linux_amd64 && chmod +x gitdump && sudo mv gitdump /usr/local/bin
```

<!-- MAC -->
### Mac

```sh
curl -Lo gitdump https://github.com/mrauer/gitdump/releases/download/v0.2.0/gitdump_0.2.0_darwin_amd64 && chmod +x gitdump && sudo mv gitdump /usr/local/bin
```

<!-- CONFIGURATION -->
## Configuration

Once the software can be run from your machine the next step is to configure it.

First you need to specify where the repositories will be downloaded in your machine:

```sh
gitdump config path <path>
```

Then we need a token to be authenticated on GitHub.

On your GitHub account, go to **Settings** > **Developer settings** > **Personal access tokens**

In there, click on **Generate new token**. Check the **repo** scope (Full control of private repositories).

Copy the token and type the following command:

```sh
gitdump config token <github_token>
```

Example:

```sh
gitdump config path /tmp
gitdump config token aid6acceae52aa987303421fd757ff684d4f2b9d
```

<!-- USAGE -->
## Usage

Those are the main commands you can use:

```sh
gitdump users ls <user> (list the repositories of a public user)
gitdump users get <user> <repository> (download a public repository)
gitdump users dump <user> (download all repositories)

gitdump owners ls <owner> (list the repositories of a private user)
gitdump owners get <owner> <repository> (download a private repository)
gitdump owners dump <owner> (download all repositories)

gitdump orgs ls (list all organizations in your account)
gitdump orgs ls <organization> (list the repositories of an organization)
gitdump orgs get <organization> <repository> (download an organization repository)
gitdump orgs dump <organization> (download all repositories)
```

Example:

```sh
gitdump users ls Netflix (list all public repositories of Netflix)
```

<!-- CONTRIBUTE -->
## Contribute

If you have any ideas or find a bug, you could always submit on the "Issues" section but you can also directly help since the project is opensource. Those are the commands you need to run a local machine:

```sh
make dev (run a dev machine assuming you have docker installed)
make binary (generate a binary)
make releases (generate the releases for all environments)
```

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.
