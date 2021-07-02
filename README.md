# ghtop

[![Go Report Card](https://goreportcard.com/badge/github.com/patrickhoefler/ghtop)](https://goreportcard.com/report/github.com/patrickhoefler/ghtop)
[![Maintainability](https://api.codeclimate.com/v1/badges/bc77f3cc2bd774e8d33f/maintainability)](https://codeclimate.com/github/patrickhoefler/ghtop/maintainability)
[![codecov](https://codecov.io/gh/patrickhoefler/ghtop/branch/main/graph/badge.svg)](https://codecov.io/gh/patrickhoefler/ghtop)

ghtop is a command line utility written in Go that displays information about the most starred GitHub repos.

## Getting Started

### Docker

```shell
docker run --rm ghcr.io/patrickhoefler/ghtop
```

### Homebrew

```shell
brew install patrickhoefler/tap/ghtop
ghtop
```

### Build from Source

```shell
go build
./ghtop
```

## Example Outputs

```text
❯ ghtop
ghtop displays information about the most starred GitHub repos.

Usage:
  ghtop [command]

Available Commands:
  defaultbranches Ranked list of default branch names based on the most starred repos
  help            Help about any command
  repos           Ranked list of the most starred repos
  topics          Ranked list of topics based on the most starred repos

Flags:
  -f, --fetch-repos int   Number of repos to fetch, between 1 and 1000 (default 100)
  -h, --help              help for ghtop
  -t, --top int           Number of results to display, between 1 and 1000

Use "ghtop [command] --help" for more information about a command.
```

### Repos

```text
❯ ghtop repos --fetch-repos 10
Rank   Stars    Repo                                     Description
1      324809   freeCodeCamp/freeCodeCamp                freeCodeCamp.org's open-source codebase and curriculum. Learn to code for free.
2      257304   996icu/996.ICU                           Repo for counting stars and contributing. Press F to pay respect to glorious developers.
3      191342   EbookFoundation/free-programming-books   :books: Freely available programming books
4      184155   vuejs/vue                                🖖 Vue.js is a progressive, incrementally-adoptable JavaScript framework for building UI on the web.
5      175156   jwasham/coding-interview-university      A complete computer science study plan to become a software engineer.
6      169397   facebook/react                           A declarative, efficient, and flexible JavaScript library for building user interfaces.
7      162589   sindresorhus/awesome                     😎 Awesome lists about all kinds of interesting topics
8      161774   kamranahmedse/developer-roadmap          Roadmap to becoming a web developer in 2021
9      156348   tensorflow/tensorflow                    An Open Source Machine Learning Framework for Everyone
10     150739   twbs/bootstrap                           The most popular HTML, CSS, and JavaScript framework for developing responsive, mobile first projects on the web.
```

### Topics

```text
❯ ghtop topics --top 10 --fetch-repos 1000
Rank   Count   Topic
1      155     javascript
2      81      python
3      77      hacktoberfest
4      67      react
5      48      go
5      48      nodejs
7      47      java
8      43      awesome
9      42      android
10     37      machine-learning
```

### Default Branches

```text
❯ ghtop defaultbranches --top 10 --fetch-repos 1000
Rank   Count   Branch Name
1      783     master
2      116     main
3      29      develop
4      22      dev
5      8       next
6      3       devel
6      3       gh-pages
6      3       trunk
9      2       3.x
9      2       8.x
9      2       canary
9      2       development
9      2       release
9      2       v2
```

## License

[MIT](https://github.com/patrickhoefler/ghtop/blob/main/LICENSE)
