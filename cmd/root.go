/*
Copyright © 2024 Emmanuel Ozeh  github.com/ozedd-ee

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ozedd-ee/json-parser/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var onlyLex bool

var rootCmd = &cobra.Command{
	Use:     "json-parser filename",
	Short:   "A simple JSON parser application",
	Long:    `A simple JSON parser application to demonstrate the Lexical and Syntax analysis stages of compilers.`,
	Example: "json-parser \"test.json\"  \njson-parser -l \"test.json\"",
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		if file[len(file)-5:] != ".json" {
			log.Fatal("Expected a JSON file as argument ")
		}
		fileContent, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		tokens := src.Lex(string(fileContent))
		if onlyLex {
			fmt.Printf("TOKENS:\n %v \n \n", tokens)
		} else {
			ast, _ := src.Parse(tokens)
			fmt.Printf("AST:\n %v \n \n", ast)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&onlyLex, "lex", "l", false, "lex JSON string and return tokens only ")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".json-parser")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
