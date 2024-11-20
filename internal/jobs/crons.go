package jobs

import (
	"fmt"

	"github.com/Mario-valente/shenlong/internal/k8s"
	"github.com/spf13/cobra"
)

func CreateCron(rootCmd *cobra.Command, args []string) {
	pathKubeconfig, _ := rootCmd.Flags().GetString("pathKubeconfig")
	command, _ := rootCmd.Flags().GetStringSlice("command")
	image, _ := rootCmd.Flags().GetString("image")
	name, _ := rootCmd.Flags().GetString("name")
	namespace, _ := rootCmd.Flags().GetString("namespace")
	ttl, _ := rootCmd.Flags().GetInt32("ttl")
	schedule, _ := rootCmd.Flags().GetString("schedule")

	_, err := k8s.CreateCronsK8s(name, namespace, image, command, pathKubeconfig, ttl, schedule)
	if err != nil {
		fmt.Println("error to create cron in k8s")
	}
}
