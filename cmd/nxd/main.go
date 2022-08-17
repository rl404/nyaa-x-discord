package main

import (
	"github.com/rl404/nyaa-x-discord/internal/utils"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "nxd",
		Short: "Nyaa x Discord",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "bot",
		Short: "Run bot",
		RunE: func(*cobra.Command, []string) error {
			return bot()
		},
	})

	cronCmd := cobra.Command{
		Use:   "cron",
		Short: "Cron",
	}

	cronCmd.AddCommand(&cobra.Command{
		Use:   "check",
		Short: "Check update",
		RunE: func(*cobra.Command, []string) error {
			return cronCheck()
		},
	})

	cmd.AddCommand(&cronCmd)

	if err := cmd.Execute(); err != nil {
		utils.Fatal(err.Error())
	}
}
