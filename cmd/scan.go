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
	"github.com/asciifaceman/hobocode"
	"github.com/asciifaceman/scwvision/pkg/eyetoy"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		e := &eyetoy.EyeToy{}
		e.GetContext()
		err := e.Open()
		if err != nil {
			hobocode.Errorf("Error opening device: %v", err)
		}
		defer e.Close()

		_, done, ep, err := e.GetInterfaceEndpoint(1)
		if err != nil {
			hobocode.Errorf("Failed to get interface and endpoint: %v", err)
		}
		defer done()

		readBytes, data, err := e.ReadEndpoint(ep)
		if err != nil {
			hobocode.Errorf("Error reading from endpoint [%s]: %v", ep.String(), err)
			return
		}

		hobocode.Infof("Read %d bytes", readBytes)

		if readBytes > 0 {
			hobocode.Infof("Received data:")
			spew.Dump(data)
		}

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
