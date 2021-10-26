# ghtop

[![Go Report Card](https://goreportcard.com/badge/github.com/patrickhoefler/ghtop)](https://goreportcard.com/report/github.com/patrickhoefler/ghtop)
[![Maintainability](https://api.codeclimate.com/v1/badges/bc77f3cc2bd774e8d33f/maintainability)](https://codeclimate.com/github/patrickhoefler/ghtop/maintainability)
[![codecov](https://codecov.io/gh/patrickhoefler/ghtop/branch/main/graph/badge.svg)](https://codecov.io/gh/patrickhoefler/ghtop)

`ghtop` lists the most starred GitHub repos and counts their topics and default branches.

## Getting Started

### Docker / [nerdctl](https://github.com/containerd/nerdctl)

```shell
docker run --rm ghcr.io/patrickhoefler/ghtop
```

### [Homebrew](https://brew.sh/)

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
1      334056   freeCodeCamp/freeCodeCamp                freeCodeCamp.org's open-source codebase and curriculum. Learn to code for free.
2      259491   996icu/996.ICU                           Repo for counting stars and contributing. Press F to pay respect to glorious developers.
3      208841   EbookFoundation/free-programming-books   :books: Freely available programming books
4      195187   jwasham/coding-interview-university      A complete computer science study plan to become a software engineer.
5      189527   vuejs/vue                                🖖 Vue.js is a progressive, incrementally-adoptable JavaScript framework for building UI on the web.
6      176439   facebook/react                           A declarative, efficient, and flexible JavaScript library for building user interfaces.
7      174549   kamranahmedse/developer-roadmap          Roadmap to becoming a web developer in 2021
8      172905   sindresorhus/awesome                     😎 Awesome lists about all kinds of interesting topics
9      164934   public-apis/public-apis                  A collective list of free APIs
10     160011   tensorflow/tensorflow                    An Open Source Machine Learning Framework for Everyone
```

### Topics

```text
❯ ghtop topics --top 10 --fetch-repos 1000
Rank   Count   Topic
1      155     javascript
2      121     hacktoberfest
3      87      python
4      69      react
5      50      nodejs
6      49      go
7      46      java
8      45      awesome
9      43      machine-learning
10     41      android
```

### Default Branches

```text
❯ ghtop defaultbranches --top 10 --fetch-repos 1000
Rank   Count   Branch Name
1      733     master
2      161     main
3      32      develop
4      21      dev
5      7       next
6      3       devel
6      3       gh-pages
6      3       trunk
6      3       v2
10     2       3.x
10     2       8.x
10     2       canary
10     2       development
```

## License

[MIT](https://github.com/patrickhoefler/ghtop/blob/main/LICENSE)
