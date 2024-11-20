package jobs

import (
	"fmt"

	"github.com/Mario-valente/shenlong/internal/k8s"
	"github.com/spf13/cobra"
)

func CreateJob(rootCmd *cobra.Command, args []string) {

	pathKubeconfig, _ := rootCmd.Flags().GetString("pathKubeconfig")
	command, _ := rootCmd.Flags().GetStringSlice("command")
	image, _ := rootCmd.Flags().GetString("image")
	name, _ := rootCmd.Flags().GetString("name")
	namespace, _ := rootCmd.Flags().GetString("namespace")
	ttl, _ := rootCmd.Flags().GetInt32("ttl")

	_, err := k8s.CreateJobsK8s(name, namespace, image, command, ttl, pathKubeconfig)
	if err != nil {
		fmt.Println("error to create job in k8s")
	}
}

func GetJob(rootCmd *cobra.Command, args []string) {
	pathKubeconfig, _ := rootCmd.Flags().GetString("pathKubeconfig")
	name, _ := rootCmd.Flags().GetString("name")
	namespace, _ := rootCmd.Flags().GetString("namespace")

	_, err := k8s.GetJobsK8s(name, namespace, pathKubeconfig)
	if err != nil {
		fmt.Println("error to get job in k8s")
	}
}
