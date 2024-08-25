package cmd

import (
	"fmt"
	"os"
	"smart-mess/config"
	"smart-mess/infra/conn"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use: "smart-mess",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

// Execute executes the root command
func Execute() {
	// load application configuration
	if err := config.Load(); err != nil {
		//log.Error(err)
		os.Exit(1)
	}

	conn.ConnectDB()
	//conn.ConnectCache()
	//conn.InitWorkerClient()
	//conn.InitWorkerInspector()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
