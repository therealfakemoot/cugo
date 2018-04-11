package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "cugo",
		Short: "Core utilities as a multi-call binary",
		Long:  "cugo is a multi-call binary that provides core utilities for Unix systems.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Printf("cugo: No operands passed\n" +
					"Usage: cugo [COMMAND] [OPTIONS] ARGUMENTS ...\n")
				return
			} else {
				Cugo(args)
			}
		},
	}

	cugoInstall bool
	cugoRemove  bool
)

func init() {
	RootCmd.Flags().SortFlags = false
	// RootCmd.Flags().BoolVarP(&cugoInstall, "install", "i", false,
	//  "Install symlink(s) for target utilities")
	// RootCmd.Flags().BoolVarP(&cugoRemove, "remove", "r", false,
	//  "Remove symlink(s) for target utilities")
}

func Cugo(args []string) {
	//
}
