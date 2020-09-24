# ghtop

[![Go Report Card](https://goreportcard.com/badge/github.com/patrickhoefler/ghtop)](https://goreportcard.com/report/github.com/patrickhoefler/ghtop)
[![Maintainability](https://api.codeclimate.com/v1/badges/bc77f3cc2bd774e8d33f/maintainability)](https://codeclimate.com/github/patrickhoefler/ghtop/maintainability)

ghtop is a command line utility written in Go that displays information about the most starred GitHub repos.

## Build

`go build`

## Example Outputs

```text
% ghtop
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
% ghtop repos --fetch-repos 10
Rank   Stars    Repo                                     Description
1      314955   freeCodeCamp/freeCodeCamp                freeCodeCamp.org's open source codebase and curriculum. Learn to code at home.
2      251745   996icu/996.ICU                           Repo for counting stars and contributing. Press F to pay respect to glorious developers.
3      172815   vuejs/vue                                🖖 Vue.js is a progressive, incrementally-adoptable JavaScript framework for building UI on the web.
4      160037   EbookFoundation/free-programming-books   :books: Freely available programming books
5      156349   facebook/react                           A declarative, efficient, and flexible JavaScript library for building user interfaces.
6      148694   tensorflow/tensorflow                    An Open Source Machine Learning Framework for Everyone
7      144288   twbs/bootstrap                           The most popular HTML, CSS, and JavaScript framework for developing responsive, mobile first projects on the web.
8      141365   sindresorhus/awesome                     😎 Awesome lists about all kinds of interesting topics
9      135524   jwasham/coding-interview-university      A complete computer science study plan to become a software engineer.
10     127687   getify/You-Dont-Know-JS                  A book series on JavaScript. @YDKJS on twitter.
```

### Topics

```text
% ghtop topics --top 10 --fetch-repos 1000
Rank   Count   Topic
1      156     javascript
2      82      python
3      62      react
4      47      java
4      47      nodejs
6      46      go
7      44      android
8      42      awesome
9      37      machine-learning
10     35      awesome-list
```

### Default Branches

```text
% ghtop defaultbranches --top 10 --fetch-repos 1000
Rank   Count   Branch Name
1      875     master
2      27      develop
3      24      main
4      20      dev
5      6       next
6      4       gh-pages
7      3       devel
7      3       trunk
9      2       3.x
9      2       canary
9      2       v2
```

## License

[MIT](https://github.com/patrickhoefler/ghtop/blob/main/LICENSE)
