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
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage the configuration of the quick-random-events tool",
	Long:  `This command allows you to show, edit, and manage the configuration of the quick-random-events tool`,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := ioutil.ReadFile(viper.ConfigFileUsed())
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}
		yamlData, err := yaml.Marshal(string(b))
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}
		color.Cyan(string(yamlData))
	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the current configuration",
	Long:  `This subcommand opens the configuration file in the default text editor for you to make changes`,
	Run: func(cmd *cobra.Command, args []string) {
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vi"
			fmt.Println("$EDITOR environment variable not set, defaulting to vi")
		}

		err := runEditor(editor, viper.ConfigFileUsed())
		if err != nil {
			fmt.Println("Error running editor:", err)
		}
	},
}

func init() {
	configCmd.AddCommand(editCmd)
	rootCmd.AddCommand(configCmd)
}

func runEditor(editor, filePath string) error {
	args := []string{filePath}
	cmd := exec.Command(editor, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
