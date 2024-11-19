package server

import (
	"github.com/Mario-valente/shenlong/internal/server/handler"
	"github.com/spf13/cobra"
)

func Server(rootCmd *cobra.Command, args []string) {
	handler.Server()

}
