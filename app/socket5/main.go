package main

import (
	"fmt"
	"github.com/leiqD/go-socket5/app"
	"github.com/leiqD/go-socket5/pkg/svcutil"
	"github.com/spf13/cobra"
)

var (
	Version = "unknown"
	Build   = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Build: %s\nVersion: %s\n", Build, Version)
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start socket5 service",
	Run: func(cmd *cobra.Command, args []string) {
		var app = &app.Program{}
		svcutil.Run(app)
	},
}


func main(){
	var rootCmd = &cobra.Command{Use: "rgb-workerd"}
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.Execute()
}
