package cmd

/*
Copyright © 2023 Vivek Singh viveksingh0143@gmail.com

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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wms",
	Short: "Warehouse Management System",
	Long: `It is a software application designed to facilitate efficient management and tracking of inventory within a warehouse setting. 
This server acts as the backend infrastructure, providing REST APIs that enable seamless communication between the front-end applications and the warehouse management system

Key features of the "Warehouse Management System Rest Resource API Server" include:
1. Inventory Management: The API server allows users to create, update, and retrieve information about products and their quantities stored in the warehouse. It enables real-time tracking of inventory levels, ensuring accurate stock management.
2. Order Processing: The server handles order processing, allowing users to place orders, manage order status, and generate order reports. It ensures smooth order fulfillment and provides insights into order history.
3. Location Tracking: The API server facilitates location tracking within the warehouse, helping users quickly locate specific items and optimize storage space.
4. Data Analytics: It offers data analytics capabilities to generate reports, visualize trends, and make data-driven decisions to improve warehouse operations and efficiency.
5. REST API Design: The server follows REST API design principles, making it easy for developers to interact with the system using standard HTTP methods and well-defined endpoints.
6. Scalability and Performance: The API server is built with scalability and performance in mind, ensuring that it can handle a large number of concurrent requests and maintain high availability.
7. Integration Capabilities: It supports integration with other systems like ERP (Enterprise Resource Planning) software, point-of-sale systems, and shipping carriers, enabling seamless data synchronization and efficient workflows.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/wms.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name "wms.yaml" (without extension).
		viper.AddConfigPath(home)

		viper.AddConfigPath("/etc/wms/")
		viper.AddConfigPath("$HOME/.wms")
		viper.AddConfigPath("./.wms")

		viper.SetConfigType("yaml")
		viper.SetConfigName("wms.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
