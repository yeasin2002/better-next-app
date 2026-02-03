package cmd

import (
	"embed"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	templatesFS embed.FS
	rootCmd     *cobra.Command
)

func init() {
	rootCmd = &cobra.Command{
		Use:   "better-next-app",
		Short: "A modern, high-performance CLI tool for scaffolding Next.js projects, written in Go",
		Long:  ` A modern, high-performance CLI tool for scaffolding Next.js projects, written in Go. This is a complete rewrite of create-next-app that provides faster startup times, single binary distribution, and feature parity with the original TypeScript implementation. `,
		// Run: func(cmd *cobra.Command, args []string) { },
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Execute(fs embed.FS) error {
	templatesFS = fs
	fmt.Println(templatesFS)
	return rootCmd.Execute()
}
