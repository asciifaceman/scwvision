/*
Copyright © 2024 Asciifaceman
*/
package cmd

import (
	"os"

	"github.com/asciifaceman/hobocode"
	"github.com/asciifaceman/scwvision/pkg/eyetoy"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Tests connection to Eyetoy",
	Long: `Tests the connection to the Eyetoy by connecting to it
and initializing it. Blinks the red light briefly. 

Will shut it off when complete.`,
	Run: func(cmd *cobra.Command, args []string) {
		e, err := eyetoy.New()
		if err != nil {
			hobocode.Errorf("Error starting up: %v", err)
			os.Exit(1)
		}

		err = e.Test(3)
		if err != nil {
			hobocode.Errorf("Error running test: %v", err)
			os.Exit(1)
		}

		hobocode.Success("Test ran successfully")

	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
