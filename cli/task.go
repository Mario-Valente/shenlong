package cli

import (
	"github.com/Mario-valente/shenlong/internal/cmd"
	"github.com/Mario-valente/shenlong/internal/jobs"
	"github.com/Mario-valente/shenlong/internal/server"
	"github.com/spf13/cobra"
)

func RegisterCreateJobCmd(rootCmd *cobra.Command) {
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

func RegisterGetCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "get",
		Run: cmd.Get,
	}

	cmd.PersistentFlags().String("name", "", "name of the job")
	cmd.PersistentFlags().String("namespace", "default", "namespace of the job")
	cmd.PersistentFlags().String("kubeconfig", "", "path to kubeconfig file")
	cmd.PersistentFlags().String("job", "", "job to get")
	cmd.PersistentFlags().String("cron", "", "cron to get")

	cmd.MarkPersistentFlagRequired("name")

	rootCmd.AddCommand(cmd)
}

func RegisterCreateCronCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "create-cron",
		Run: jobs.CreateCron,
	}

	cmd.PersistentFlags().String("image", "busybox:latest", "image to run the cron in k8s")
	cmd.PersistentFlags().String("cpu", "", "quantity of cpu to run the cron")
	cmd.PersistentFlags().String("memory", "", "quantity of memory to run the cron")
	cmd.PersistentFlags().String("namespace", "default", "namespace to run the cron")
	cmd.PersistentFlags().String("name", "", "name of the cron")
	cmd.PersistentFlags().StringSlice("command", []string{}, "command to run in the cron")
	cmd.PersistentFlags().String("kubeconfig", "", "path to kubeconfig file")
	cmd.PersistentFlags().Int32("ttl", 100, "time to live of the cron")
	cmd.PersistentFlags().String("schedule", "*/1 * * * *", "schedule to run the cron")

	cmd.MarkPersistentFlagRequired("command")
	cmd.MarkPersistentFlagRequired("name")
	cmd.MarkPersistentFlagRequired("schedule")

	rootCmd.AddCommand(cmd)

}

func RegisterDeleteCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "delete",
		Run: cmd.Delete,
	}

	cmd.PersistentFlags().String("name", "", "name of the job")
	cmd.PersistentFlags().String("namespace", "default", "namespace of the job")
	cmd.PersistentFlags().String("kubeconfig", "", "path to kubeconfig file")
	cmd.PersistentFlags().String("job", "", "job to delete")
	cmd.PersistentFlags().String("cron", "", "cron to delete")

	cmd.MarkPersistentFlagRequired("name")

	rootCmd.AddCommand(cmd)
}
