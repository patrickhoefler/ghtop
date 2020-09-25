package cmd

import (
	"fmt"
	"io"
	"sort"
	"text/tabwriter"

	"github.com/patrickhoefler/ghtop/internal"
	"github.com/spf13/cobra"
)

func newCmdStats(outWriter io.Writer, repoClient internal.Fetcher) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "defaultbranches",
		Short: "Ranked list of default branch names based on the most starred repos",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			defaultBranchCounts := map[string]int{}

			numberOfRepos, numberOfResults := getPersistentFlags(cmd)

			for _, repo := range repoClient.Fetch(numberOfRepos) {
				defaultBranchCounts[repo.GetDefaultBranch()]++
			}

			defaultBranches := make([]string, 0, len(defaultBranchCounts))
			for key := range defaultBranchCounts {
				defaultBranches = append(defaultBranches, key)
			}

			// First sort alphabetically, this will become the secondary sort order
			sort.Strings(defaultBranches)

			// Then sort by number of stargazers, this is now the primary sort order
			sort.SliceStable(defaultBranches, func(i, j int) bool {
				return defaultBranchCounts[defaultBranches[i]] > defaultBranchCounts[defaultBranches[j]]
			})

			output := tabwriter.NewWriter(outWriter, 0, 0, 3, ' ', 0)
			rankinator := new(internal.Rankinator)

			fmt.Fprintln(output, "Rank\tCount\tBranch Name")
			for _, name := range defaultBranches {
				rank, err := rankinator.Get(defaultBranchCounts[name])
				if err != nil {
					panic(err)
				}

				if numberOfResults > 0 && rank > numberOfResults {
					break
				}

				fmt.Fprintln(
					output,
					fmt.Sprintf(
						"%d\t%d\t%s",
						rank,
						defaultBranchCounts[name],
						name,
					),
				)
			}

			output.Flush()
		},
	}

	return cmd
}
