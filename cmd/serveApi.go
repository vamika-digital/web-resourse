package cmd

/*
Copyright © 2023 Vivek Kumar <viveksingh0143@gmail.com>

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

import (
	"github.com/spf13/cobra"
	"github.com/vamika-digital/wms-resourse/app"
)

// serveApiCmd represents the serveApi command
var serveApiCmd = &cobra.Command{
	Use:   "serve-api",
	Short: "Start API Resource Server",
	Long:  `It serve all the APIs related to WMS`,
	Run: func(cmd *cobra.Command, args []string) {
		app.StartRestServer()
	},
}

func init() {
	rootCmd.AddCommand(serveApiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveApiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveApiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
