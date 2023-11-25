# GitDump
<!-- PROJECT LOGO -->
<p align="center">
  <a href="https://github.com/mrauer/gitdump">
    <img src="images/logo.png" alt="Logo" width="225" height="225">
  </a>

  <h3 align="center">GitDump</h3>

  <p align="center">
    Tool for downloading GitHub repositories.
    <br />
    <br />
    <a href="https://github.com/mrauer/gitdump/issues">Report Bug</a>
    |
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
* [License](#license)

<!-- ABOUT THE PROJECT -->
## About the Project

GitDump is a tool created to enable you to download either specific repositories or the complete repository set associated with a GitHub account. Whether the account is public, private, or affiliated with an organization, GitDump simplifies the process of creating an archive. This empowers you to preserve your own projects or those of others.

<!-- INSTALLATION -->
## Installation

This application is compatible with Linux, Mac, and Windows operating systems. To obtain the binaries, you can visit the releases page at [https://github.com/mrauer/gitdump/releases](https://github.com/mrauer/gitdump/releases). Alternatively, if you're using Linux or Mac, you can simply copy and paste the provided commands.

<!-- LINUX -->
### Linux

>curl -Lo gitdump https://github.com/mrauer/gitdump/releases/download/v0.3.0/gitdump_0.3.0_linux_amd64 \
&& chmod +x gitdump \
&& sudo mv gitdump /usr/local/bin

<!-- MAC -->
### Mac

>curl -Lo gitdump https://github.com/mrauer/gitdump/releases/download/v0.3.0/gitdump_0.3.0_darwin_amd64 \
&& chmod +x gitdump \
&& sudo mv gitdump /usr/local/bin

<!-- CONFIGURATION -->
## Configuration

After successfully running the software on your machine, the next step is configuration.

Start by specifying the directory where repositories will be downloaded on your machine:

>gitdump config path <path>

Next, obtain an authentication token for GitHub.

Navigate to your GitHub account, go to **Settings** > **Developer settings** > **Personal access tokens**.

Click on **Generate new token**, ensure the **repo** scope is selected (Full control of private repositories).

Copy the token and use the following command:

>gitdump config token <github_token>

Example:

```sh
gitdump config path /tmp
gitdump config token aid6acceae52aa987303421fd757ff684d4f2b9d
```

<!-- USAGE -->
## Usage

Here, you'll find the key commands for using GitDump:

```sh
gitdump users ls <user> (list repositories of a public user)
gitdump users get <user> <repository> (download a public repository)
gitdump users dump <user> (download all repositories from a public user)

gitdump owners ls <owner> (list repositories of a private user)
gitdump owners get <owner> <repository> (download a private repository)
gitdump owners dump <owner> (download all repositories from a private user)

gitdump orgs ls (list all organizations in your account)
gitdump orgs ls <organization> (list repositories of an organization)
gitdump orgs get <organization> <repository> (download an organization repository)
gitdump orgs dump <organization> (download all repositories from an organization)
```

Example:

>gitdump users ls netflix (list all public repositories from Netflix)

<!-- CONTRIBUTE -->
## Contribute

If you come up with any ideas or discover a bug, feel free to submit them in the `issues` section. Moreover, you have the opportunity to contribute directly to the project, given its open-source nature. To get started, execute the following commands on your local machine:

```sh
make dev (to run a development machine assuming Docker is installed)
make binary (to generate a binary)
make releases (to generate releases for all environments)
```

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.
