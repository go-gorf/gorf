/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Create new Gorf App",
	Long:  `Creates a new Gorf app in the current directory`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		CreateNewGorfApp(args[0])
	},
}

type GorfApp struct {
	Name string
}

func (a *GorfApp) Create() {
	gorfTemplateUrl := "https://github.com/go-gorf/template-app.git"
	out, err := exec.Command("git", "clone", gorfTemplateUrl).Output()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename("template", a.Name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func CreateNewGorfApp(name string) {
	gorfApp := &GorfApp{name}
	gorfApp.Create()
	fmt.Println("Successfully created new apps!")
}

func init() {
	rootCmd.AddCommand(appCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
