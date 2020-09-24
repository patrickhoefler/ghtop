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

	minNumberOfResults = 1
	maxNumberOfResults = 1000
)

type intMinMax struct {
	value int64
	min   int64
	max   int64
}

func (i *intMinMax) Set(s string) (err error) {
	value, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return
	}

	if (value < i.min) || (value > i.max) {
		err = fmt.Errorf("value must be between %d and %d", i.min, i.max)
		return
	}

	i.value = value
	return
}
func (i *intMinMax) String() string { return strconv.Itoa(int(i.value)) }
func (i *intMinMax) Type() string   { return "int" }

func newRootCmd(repoFetcher internal.RepoFetcher, out io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{
		Short: "ghtop displays information about the most starred GitHub repos.",
		Use:   "ghtop",
	}

	var numberOfRepos = intMinMax{value: defaultNumberOfRepos, min: minNumberOfRepos, max: maxNumberOfRepos}
	rootCmd.PersistentFlags().VarP(
		&numberOfRepos,
		"fetch-repos",
		"f",
		fmt.Sprintf("Number of repos to fetch, between %d and %d", minNumberOfRepos, maxNumberOfRepos),
	)

	var numberOfResults = intMinMax{min: minNumberOfResults, max: maxNumberOfResults}
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
