/*
Copyright © 2024 Asciifaceman

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Attempt to read a single packet from EyeToy",
	Long:  `Attempt to read a single packet from EyeToy.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			hobocode.HeaderLeft("Startup")
			et, err := eyetoy.New(10 * time.Second)
			if err != nil {
				hobocode.Errorf("Startup error: %v", err)
				os.Exit(1)
			}

			defer et.Close()

			err = et.GetImage()
			if err != nil {
				hobocode.Errorf("Failed to read image: %v", err)
			}
		*/
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
