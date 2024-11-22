package main

import (
	"os"
	"path"

	"github.com/Mario-valente/shenlong/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   path.Base(os.Args[0]),
		Short: "A CLI tool for create jobs in k8s",
	}

	cli.RegisterCreateJobCmd(rootCmd)
	cli.RegisterServerCmd(rootCmd)
	cli.RegisterGetCmd(rootCmd)
	cli.RegisterDeleteCmd(rootCmd)
	cli.RegisterCreateCronCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
