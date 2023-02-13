/*
Copyright © 2023 Xudong Guo <guoxudong.dev@gmail.com>

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
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a sample config file",
	Long:  `Create a sample config file that contains a list of activities and foods`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		if _, err := os.Stat(configPath); err == nil {
			fmt.Printf("The config file already exists at %s\n", configPath)
			return
		}

		file, err := os.Create(configPath)
		if err != nil {
			fmt.Printf("Error creating config file: %s\n", err)
			return
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		defer writer.Flush()

		sampleConfig := `
activities:
- name: Visit the Museum
  description: Take a tour of the local museum
  weight: 10

- name: Go Hiking
  description: Explore the local mountains
  weight: 5

foods:
- name: Sushi
  description: Enjoy some Japanese cuisine
  weight: 15

- name: Italian Food
  description: Try some delicious pasta and pizza
  weight: 8
`
		_, err = writer.WriteString(sampleConfig)
		if err != nil {
			fmt.Printf("Error writing to config file: %s\n", err)
			return
		}

		fmt.Printf("Sample config file created at %s\n", configPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	homeDir, _ := os.UserHomeDir()
	defaultConfigPath := filepath.Join(homeDir, ".config", "quick-random-events", "config.yaml")
	initCmd.Flags().String("config", defaultConfigPath, "Path to the config file")
}
