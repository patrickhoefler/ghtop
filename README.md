# ghtop

[![Go Report Card](https://goreportcard.com/badge/github.com/patrickhoefler/ghtop)](https://goreportcard.com/report/github.com/patrickhoefler/ghtop)
[![Maintainability](https://api.codeclimate.com/v1/badges/bc77f3cc2bd774e8d33f/maintainability)](https://codeclimate.com/github/patrickhoefler/ghtop/maintainability)

`ghtop` lists the most starred GitHub repos and counts their topics and default branches.

## Getting Started

### Docker / [nerdctl](https://github.com/containerd/nerdctl)

```text
docker run --rm ghcr.io/patrickhoefler/ghtop
```

### [Homebrew](https://brew.sh/)

```text
brew install patrickhoefler/tap/ghtop
ghtop
```

### Build from Source

```text
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
  completion      Generate the autocompletion script for the specified shell
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
1      343825   freeCodeCamp/freeCodeCamp                freeCodeCamp.org's open-source codebase and curriculum. Learn to code for free.
2      261722   996icu/996.ICU                           Repo for counting stars and contributing. Press F to pay respect to glorious developers.
3      230502   EbookFoundation/free-programming-books   :books: Freely available programming books
4      217074   jwasham/coding-interview-university      A complete computer science study plan to become a software engineer.
5      196848   sindresorhus/awesome                     😎 Awesome lists about all kinds of interesting topics
6      194912   vuejs/vue                                🖖 Vue.js is a progressive, incrementally-adoptable JavaScript framework for building UI on the web.
7      191669   kamranahmedse/developer-roadmap          Roadmap to becoming a developer in 2022
8      189239   public-apis/public-apis                  A collective list of free APIs
9      186251   facebook/react                           A declarative, efficient, and flexible JavaScript library for building user interfaces.
10     172937   donnemartin/system-design-primer         Learn how to design large-scale systems. Prep for the system design interview.  Includes Anki flashcards.
```

### Topics

```text
❯ ghtop topics --top 10 --fetch-repos 1000
Rank   Count   Topic
1      155     javascript
2      106     hacktoberfest
3      87      python
4      66      react
5      52      go
6      48      nodejs
7      47      machine-learning
8      45      java
9      44      awesome
10     42      android
```

### Default Branches

```text
❯ ghtop defaultbranches --top 10 --fetch-repos 1000
Rank   Count   Branch Name
1      692     master
2      195     main
3      35      develop
4      22      dev
5      7       next
6      3       devel
6      3       gh-pages
6      3       trunk
6      3       v2
10     2       3.x
10     2       9.x
10     2       canary
10     2       development
```

## License

[MIT](https://github.com/patrickhoefler/ghtop/blob/main/LICENSE)
