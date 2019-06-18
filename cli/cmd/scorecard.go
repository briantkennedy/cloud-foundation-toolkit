package cmd

import (
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/cli/scorecard"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scorecardCmd)
	scorecardCmd.Flags().StringVar(&flags.scorecard.policyPath, "policy-path", "", "Path to directory containing validation policies")
	scorecardCmd.MarkFlagRequired("policy-path")
}

var scorecardCmd = &cobra.Command{
	Use:   "scorecard",
	Short: "Print a scorecard of your GCP envirment",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println("Generating CFT scorecard")
		var err error

		inventory := scorecard.NewInventory("gcp-foundation-shared-devops")

		if false {
			err = scorecard.ExportInventory(inventory)
			if err != nil {
				return errors.Wrap(err, "Error exporting asset inventory")
			}
		}

		config := &scorecard.ScoringConfig{
			PolicyPath: flags.scorecard.policyPath,
		}
		err = scorecard.ScoreInventory(inventory, config)
		if err != nil {
			return err
		}

		return nil
	},
}