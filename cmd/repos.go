package cmd

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/patrickhoefler/ghtop/internal/fetching"
	"github.com/patrickhoefler/ghtop/internal/ranking"
	"github.com/spf13/cobra"
)

func newCmdRepos(outWriter io.Writer, repoClient fetching.Fetcher) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "repos",
		Short: "Ranked list of the most starred repos",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			numberOfRepos, numberOfResults := getPersistentFlags(cmd)

			output := tabwriter.NewWriter(outWriter, 0, 0, 3, ' ', 0)
			rankinator := new(ranking.Rankinator)

			fmt.Fprintln(output, "Rank\tStars\tRepo\tDescription")
			for _, repo := range repoClient.Fetch(numberOfRepos) {
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
