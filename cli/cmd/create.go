/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/waldrey/go-expert/cli/internal/database"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples`,
		RunE:  runCreate(categoryDb),
	}
}

func init() {
	createCmd := newCreateCmd(getCategoryDB(getDB()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Category name")
	createCmd.Flags().StringP("description", "d", "", "Category description")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}

		fmt.Println("Category created")
		return nil
	}
}
