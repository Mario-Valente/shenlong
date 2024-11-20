package cmd

import (
	"fmt"

	"github.com/Mario-valente/shenlong/internal/k8s"
	"github.com/spf13/cobra"
)

func Get(rootCmd *cobra.Command, args []string) {
	pathKubeconfig, _ := rootCmd.Flags().GetString("pathKubeconfig")
	name, _ := rootCmd.Flags().GetString("name")
	namespace, _ := rootCmd.Flags().GetString("namespace")
	job, _ := rootCmd.Flags().GetString("job")
	cron, _ := rootCmd.Flags().GetString("cron")

	switch {
	case job != "":
		_, err := k8s.GetJobsK8s(name, namespace, pathKubeconfig)
		if err != nil {
			fmt.Println("error to get job in k8s")
		}
	case cron != "":
		_, err := k8s.GetCronsK8s(name, namespace, pathKubeconfig)
		if err != nil {
			fmt.Println("error to get cron in k8s")
		}
	}
}
