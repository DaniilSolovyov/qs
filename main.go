package main

import (
	"fmt"

	cobra "github.com/spf13/cobra"
	v "github.com/untillpro/qg/vcs"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "qg",
		Short: "Quick git wrapper",
		Run: func(cmd *cobra.Command, args []string) {
			v.Status()
		},
	}

	var uploadCmd = &cobra.Command{
		Use:   "u",
		Short: "Upload sources",
		Run: func(cmd *cobra.Command, args []string) {
			v.Upload()
		},
	}

	var downloadCmd = &cobra.Command{
		Use:   "d",
		Short: "Download sources",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside download with args: %v\n", args)
		},
	}

	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(downloadCmd)

	rootCmd.Execute()

}