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
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

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
        // Check if input is a JSON file
		if file[len(file)-5:] != ".json" {
			log.Fatal("Expected a JSON file as argument ")
		}
		openFile, err := os.Open(file)
		if err != nil {
			log.Fatal("Error opening file:", err)
		}
		defer openFile.Close()
        // Remove carriage returns before passing to lexer
		fileContent := removeCarriageReturns(openFile)

		tokens := src.Lex(fileContent)
		if onlyLex {
			fmt.Printf("TOKENS:\n %v \n \n", tokens)
		} else {
			abstractSyntaxTree, _ := src.Parse(tokens)
			jsonEncodedAST := marshalJSON(&abstractSyntaxTree)
			fmt.Printf("AST:\n %v \n \n", jsonEncodedAST)
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

func removeCarriageReturns(openFile *os.File) string {
	scanner := bufio.NewScanner(openFile)
	var cleanedLines []string

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Remove carriage return characters but leave newline (\n)
		cleanedLine := strings.ReplaceAll(line, "\r", "")
		cleanedLines = append(cleanedLines, cleanedLine)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}
	fileContent := strings.Join(cleanedLines, "")
	return fileContent
}

func marshalJSON(abstractSyntaxTree *src.Node) string {
	jsonEncodedAST, err := json.MarshalIndent(abstractSyntaxTree, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonEncodedAST)
}
