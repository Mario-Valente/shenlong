package cli

import (
	"github.com/Mario-valente/shenlong/internal/jobs"
	"github.com/Mario-valente/shenlong/internal/server"
	"github.com/spf13/cobra"
)

func RegisterCreatejobCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "create-job",
		Run: jobs.CreateJob,
	}

	cmd.PersistentFlags().String("image", "busybox:latest", "image to run the job in k8s")
	cmd.PersistentFlags().String("cpu", "", "quantity of cpu to run the job")
	cmd.PersistentFlags().String("memory", "", "quantity of memory to run the job")
	cmd.PersistentFlags().String("namespace", "default", "namespace to run the job")
	cmd.PersistentFlags().String("name", "", "name of the job")
	cmd.PersistentFlags().StringSlice("command", []string{}, "command to run in the job")
	cmd.PersistentFlags().String("kubeconfig", "", "path to kubeconfig file")
	cmd.PersistentFlags().Int32("ttl", 100, "time to live of the job")

	cmd.MarkPersistentFlagRequired("command")
	cmd.MarkPersistentFlagRequired("name")

	rootCmd.AddCommand(cmd)

}

func RegisterServerCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "server",
		Run: server.Server,
	}

	cmd.PersistentFlags().String("run", "", "run the server")

	rootCmd.AddCommand(cmd)
}

func RegisterGetJobCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "get-job",
		Run: jobs.GetJob,
	}

	cmd.PersistentFlags().String("name", "", "name of the job")
	cmd.PersistentFlags().String("namespace", "default", "namespace of the job")
	cmd.PersistentFlags().String("kubeconfig", "", "path to kubeconfig file")

	cmd.MarkPersistentFlagRequired("name")

	rootCmd.AddCommand(cmd)
}
