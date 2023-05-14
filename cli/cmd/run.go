/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run Gorf Application",
	Long: `Run Command will run the go application by invoking
go run . cmd. For passing additional args to gorf use --  For example:

gorf run -- port
`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Gorf...")
		out, err := exec.Command("go", "run", ".").Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
