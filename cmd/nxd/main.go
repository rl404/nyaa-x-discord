package main

import (
	"log"

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

	cmd.AddCommand(&cobra.Command{
		Use:   "cron",
		Short: "Run check command but with scheduler",
		RunE: func(*cobra.Command, []string) error {
			return cron()
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "check",
		Short: "Run Nyaa checker and notify user",
		RunE: func(*cobra.Command, []string) error {
			return check()
		},
	})

	if err := cmd.Execute(); err != nil {
		log.Println(err)
	}
}
