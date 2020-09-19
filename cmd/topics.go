package cmd

import (
	"fmt"
	"io"
	"sort"
	"text/tabwriter"

	"github.com/patrickhoefler/ghtop/internal"
	"github.com/spf13/cobra"
)

func newCmdTopics(outWriter io.Writer, fetchRepos internal.RepoFetcher) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "topics",
		Short: "Ranked list of topics based on the most starred repos",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			topicCounts := map[string]int{}

			numberOfRepos, err := cmd.Flags().GetInt("fetch-repos")
			if err != nil {
				panic(err)
			}
			numberOfResults, err := cmd.Flags().GetInt("top")
			if err != nil {
				panic(err)
			}

			for _, repo := range fetchRepos.Fetch(numberOfRepos) {
				if repo == nil {
					continue
				}

				for _, topic := range repo.Topics {
					topicCounts[topic]++
				}
			}

			topics := make([]string, 0, len(topicCounts))
			for key := range topicCounts {
				topics = append(topics, key)
			}

			// First sort alphabetically, this will become the secondary sort order
			sort.Strings(topics)

			// Then sort by number of stargazers, this is now the primary sort order
			sort.SliceStable(topics, func(i, j int) bool {
				return topicCounts[topics[i]] > topicCounts[topics[j]]
			})

			output := tabwriter.NewWriter(outWriter, 0, 0, 3, ' ', 0)
			rankinator := new(internal.Rankinator)

			fmt.Fprintln(output, "Rank\tCount\tTopic")
			for _, name := range topics {
				rank, err := rankinator.Get(topicCounts[name])
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
						topicCounts[name],
						name,
					),
				)
			}

			output.Flush()
		},
	}

	return cmd
}
