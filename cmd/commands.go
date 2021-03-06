package cmd

import (
	"amireshoon/gow/gow"
	"amireshoon/gow/markdown"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// GowVersion current version of gow
var GowVersion = "1.0.0"

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(works)
	rootCmd.AddCommand(_init)
	rootCmd.AddCommand(mdAdd)
	rootCmd.AddCommand(mdCheck)
	rootCmd.AddCommand(mdAddToReadMe)
	works.AddCommand(add)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of gow you are using.",
	Long:  `Print version number of gow you are currently using.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gow version ", gow.GetVersion())
	},
}

var works = &cobra.Command{
	Use:   "works",
	Short: "Get work by name or id",
	Long:  `Get work or todo by calling it by name or id`,
	Run: func(cmd *cobra.Command, args []string) {
		gow.Parse()
	},
}

var add = &cobra.Command{
	Use:   "add [string name] [string desc]",
	Short: "Add new work",
	Long:  `Add new main work or todo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Right on way", args)
	},
}

var _init = &cobra.Command{
	Use:   "init [string project name]",
	Short: "Initialize new todo workspace",
	Long:  `Adds TODO.md if not exists and initialize new workspace`,
	Run: func(cmd *cobra.Command, args []string) {

		err := gow.ParseTodo(PathFlag, DescriptionFlag, args...)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var mdAdd = &cobra.Command{
	Use:   "add [string title]",
	Short: "Add new todo to TODO.md",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !markdown.HasTodo(PathFlag) {
			return errors.New("This project has no TODO.md.\nTry generating it by gow intit [name] -d [description:optional] -p [path:optional]")
		}
		desc := ""
		if args[0] != "" {
			desc = string(strings.Join(args[:], " "))
		} else {
			desc = DescriptionFlag
		}
		markdown.AddTodo(desc, PathFlag)
		return nil
	},
}

var mdCheck = &cobra.Command{
	Use:   "check [int index]",
	Short: "Check existing todo",
	RunE: func(cmd *cobra.Command, args []string) error {
		i, _ := strconv.Atoi(args[0])
		markdown.CheckTodo(i, PathFlag)
		return nil
	},
}

var mdAddToReadMe = &cobra.Command{
	Use:   "readme",
	Short: "Add TODO.md content to bottom of README.md",
	Run: func(cmd *cobra.Command, args []string) {
		markdown.AddToReadme(PathFlag)
	},
}
