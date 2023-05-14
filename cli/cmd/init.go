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

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Gorf project",
	Long: `Initialize a new Gorf project
Usage example:

gorf init myproject
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Creating project (%v)...", args[0])
		CreateNewGorfProject(args[0])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

type Project struct {
	Name string
}

func (p *Project) Create() {
	gorfTemplateUrl := "https://github.com/go-gorf/template.git"
	out, err := exec.Command("git", "clone", gorfTemplateUrl).Output()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename("template", p.Name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func CreateNewGorfProject(name string) {
	project := &Project{name}
	project.Create()
	fmt.Println("Successfully created project!")
	fmt.Printf("cd %v\ngorf-cli run\n", project.Name)
}
