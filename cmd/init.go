package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

func newProject(projectPath string, goIdentifier string) {
	const GIT_REPO = "https://github.com/ManitVig/GoTHAM-starter-app.git"
	fmt.Println("Initializing Project...")
	cmd := exec.Command("git", "clone", GIT_REPO, projectPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error while initiliazing project.")
		return
	}
	fmt.Println("Cloned Successfully.")
	err = os.RemoveAll(filepath.Join(projectPath, ".git"))
	if err != nil {
		fmt.Println("Error while resetting .git directory.")
	}

	err = os.Remove(filepath.Join(projectPath, ".gitignore"))
	if err != nil {
		fmt.Println("Error while resetting .gitignore.")
	}
}

func overwrite(projectPath string) bool {
	fmt.Print("A directory with that name already exist do you want to overwrite? (Y/n): ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Couldn't read input. Cancelling operation.")
		return false
	}
	if input == "y" || input == "Y" || input == "" {
		err := os.RemoveAll(projectPath)
		if err != nil {
			fmt.Println("Couldn't delete project at path.")
			return false
		}
		return true
	}

	if input == "n" || input == "N" {
		return false
	}

	return false
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:        "init",
	Short:      "Start a new Application",
	Long:       `Initialize a new GoTHAM application using either official templates or community templates.`,
	Args:       cobra.ExactArgs(2),
	ArgAliases: []string{"project-path", "go-identifier"},
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting the current working directory.")
			return
		}
		projectPath := filepath.Join(cwd, args[0])
		goIdentifier := args[1]
		fmt.Printf("Project Path: %s\n", projectPath)
		_, exists := os.Stat(projectPath)
		if exists != nil {
			newProject(projectPath, goIdentifier)
		} else {
			overwrite := overwrite(projectPath)
			if !overwrite {
				fmt.Println("Exiting...")
				return
			} else {
				newProject(projectPath, goIdentifier)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
