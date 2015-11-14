# watson

[![Build Status](https://travis-ci.org/andygrunwald/watson.svg?branch=master)](https://travis-ci.org/andygrunwald/watson)

A crawler for Googles code review system [Gerrit](https://www.gerritcodereview.com/).

## Features

* List all projects of a Gerrit instance
* Crawl a Gerrit instance and st
* Various storage backends (currently [MySQL](https://www.mysql.com/) only)
* Special storage backends for identites (currently [sortinghat](https://github.com/MetricsGrimoire/sortinghat) only)

## Installation

TODO

## Usage

```sh
$ ./watson -help
NAME:
   Watson - Crawl your Gerrit!

USAGE:
   ./watson [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR(S):
   Andy Grunwald <andygrunwald@gmail.com>

COMMANDS:
   list-projects, lp	Lists all projects of a Gerrit instance
   crawl, c		Crawls a Gerrit instance
   help, h		Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --instance, -i 	URL for the Gerrit instance [$WATSON_INSTANCE]
   --username, -u 	Username for the Gerrit instance [$WATSON_AUTH_USERNAME]
   --password, -p 	Password for the Gerrit instance [$WATSON_AUTH_PASSWORD]
   --auth-mode, --am 	Mode for Gerrit authentication (basic or cookie) [$WATSON_AUTH_MODE]
   --help, -h		show help
   --version, -v	print the version
```

### List projects

Get all projects of [Android](ttps://android-review.googlesource.com/):

```sh
$ ./watson -i "https://android-review.googlesource.com/" list-projects
```

List all extensions with descriptions of [TYPO3](https://review.typo3.org/):

```sh
$ ./watson -i "https://review.typo3.org/" list-projects \
           -template "{{ .IDEscaped | printf \"%-60s\" }} {{ .Description }}" \
           -filter "TYPO3CMS/Extensions/.*"
```

### Crawl

TODO

## License

This project is released under the terms of the [Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).