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

	cli.RegisterCreatejobCmd(rootCmd)
	cli.RegisterServerCmd(rootCmd)
	cli.RegisterGetJobCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
