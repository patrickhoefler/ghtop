package cmd

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/patrickhoefler/ghtop/internal"
	"github.com/spf13/cobra"
)

func newCmdRepos(outWriter io.Writer, fetchRepos internal.RepoFetcher) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "repos",
		Short: "Ranked list of the most starred repos",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			numberOfRepos, err := cmd.Flags().GetInt("fetch-repos")
			if err != nil {
				panic(err)
			}
			numberOfResults, err := cmd.Flags().GetInt("top")
			if err != nil {
				panic(err)
			}

			output := tabwriter.NewWriter(outWriter, 0, 0, 3, ' ', 0)
			rankinator := new(internal.Rankinator)

			fmt.Fprintln(output, "Rank\tStars\tRepo\tDescription")
			for _, repo := range fetchRepos.Fetch(numberOfRepos) {
				rank, err := rankinator.Get(repo.GetStargazersCount())
				if err != nil {
					panic(err)
				}

				if numberOfResults > 0 && rank > numberOfResults {
					break
				}

				fmt.Fprintln(
					output,
					fmt.Sprintf(
						"%d\t%d\t%s\t%s",
						rank,
						repo.GetStargazersCount(),
						repo.GetFullName(),
						repo.GetDescription(),
					),
				)
			}

			output.Flush()
		},
	}

	return
}
