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
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"math/rand"
	"os"
	"time"
)

var cfgFile string

type Config struct {
	Events []Event `yaml:"events"`
}

type Event struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	List        []Category `yaml:"list"`
}

type Category struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Weight      int    `yaml:"weight"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qres",
	Short: "Quickly generate a random result based on a configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := LoadConfig(viper.ConfigFileUsed())
		if err != nil {
			fmt.Println("Error loading config file:", err)
			os.Exit(1)
		}
		return config.SelectRandomEvent()
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/quick-random-events/config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home + "/.config/quick-random-events")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) SelectRandomEvent() error {
	for _, event := range c.Events {
		if len(event.List) == 0 {
			return errors.New("no more items to choose from")
		}
		item := event.ChooseCategory()
		println(event.Name + ": " + color.GreenString(item.Name))
	}

	// Save the updated config with decreased weight
	if err := c.Save(); err != nil {
		return err
	}

	return nil
}

func (e Event) ChooseCategory() Category {
	var totalWeight int
	for _, item := range e.List {
		totalWeight += item.Weight
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(totalWeight)

	var category Category
	for _, item := range e.List {
		if randomNumber < item.Weight {
			category = item
			break
		}
		randomNumber -= item.Weight
	}

	return category
}

func (c *Config) Save() error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(viper.ConfigFileUsed(), data, 0644)
}
