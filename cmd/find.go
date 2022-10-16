/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	m "recurring-values/pkg/model"
	"recurring-values/pkg/reader"

	"github.com/spf13/cobra"
)

const (
	filesFlagName = "files"
	textFlagName  = "text"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find recurring values",
	Long: `Find recurring values either within specified file or standard input

	Usage:
	find file [file_name1, file_name2]
	find text [some_random_words]`,
	Run: func(cmd *cobra.Command, args []string) {
		input := args[1:]

		switch args[0] {
		case filesFlagName:
			processFiles(&input)
		case textFlagName:
			processText(&input)
		default:
			fmt.Println("Try again!")
		}
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	findCmd.PersistentFlags().String(filesFlagName, "", "Specify files to check")
	findCmd.PersistentFlags().String(textFlagName, "", "Check words separated with spaces")
}

func processFiles(files *[]string) {
	values := &m.NestedValuesWrapper{NestedValues: make(map[string]*m.ValuesWrapper)}

	reader.ReadFiles(files, values)
	values.PrintNestedValues()
}

func processText(words *[]string) {
	values := &m.ValuesWrapper{Values: make(map[string]*m.RecurringValue)}

	reader.ReadWords(words, values)
	values.PrintValues()
}
