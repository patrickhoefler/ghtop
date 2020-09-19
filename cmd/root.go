package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/patrickhoefler/ghtop/internal"
	"github.com/spf13/cobra"
)

const (
	defaultNumberOfRepos = 100
	minNumberOfRepos     = 1
	maxNumberOfRepos     = 1000
	minNumberOfResults   = 1
	maxNumberOfResults   = 1000
)

type numberOfReposValue int

func (n *numberOfReposValue) Set(s string) (err error) {
	value, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return
	}

	if (value < minNumberOfRepos) || (value > maxNumberOfRepos) {
		err = fmt.Errorf("number of repos must be between %d and %d", minNumberOfRepos, maxNumberOfRepos)
		return
	}

	*n = numberOfReposValue(value)
	return
}
func (n *numberOfReposValue) String() string { return strconv.Itoa(int(*n)) }
func (n *numberOfReposValue) Type() string   { return "int" }

type numberOfResultsValue int

func (n *numberOfResultsValue) Set(s string) (err error) {
	value, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return
	}

	if (value < minNumberOfResults) || (value > maxNumberOfResults) {
		err = fmt.Errorf("number of Results must be between %d and %d", minNumberOfResults, maxNumberOfResults)
		return
	}

	*n = numberOfResultsValue(value)
	return
}
func (n *numberOfResultsValue) String() string { return strconv.Itoa(int(*n)) }
func (n *numberOfResultsValue) Type() string   { return "int" }

func newRootCmd(repoFetcher internal.RepoFetcher, out io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{
		Short: "ghtop displays information about the most starred GitHub repos.",
		Use:   "ghtop",
	}

	var numberOfRepos = numberOfReposValue(defaultNumberOfRepos)
	rootCmd.PersistentFlags().VarP(
		&numberOfRepos,
		"fetch-repos",
		"f",
		fmt.Sprintf("Number of repos to fetch, between %d and %d", minNumberOfRepos, maxNumberOfRepos),
	)

	var numberOfResults = *new(numberOfResultsValue)
	rootCmd.PersistentFlags().VarP(
		&numberOfResults,
		"top",
		"t",
		fmt.Sprintf("Number of results to display, between %d and %d", minNumberOfResults, maxNumberOfResults),
	)

	rootCmd.AddCommand(
		newCmdRepos(out, repoFetcher),
		newCmdStats(out, repoFetcher),
		newCmdTopics(out, repoFetcher),
	)

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	newRootCmd(new(internal.GitHubRepoFetcher), os.Stdout).Execute()
}
