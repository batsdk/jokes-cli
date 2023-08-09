/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random Joke",
	Long:  "Get a random Joke",
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke map[string][]string

func getRandomJoke() {

	rand.Seed(time.Now().UnixNano())

	jsonFilePath := "./jokes.json"
	jsonData, err := os.ReadFile(jsonFilePath)

	if err != nil {
		fmt.Println(err)
	}

	var jokes Joke

	err = json.Unmarshal(jsonData, &jokes)
	if err != nil {
		fmt.Println(err)
		return
	}

	maxNumber := len(jokes["jokes"])
	randomNumber := rand.Intn(maxNumber)

	fmt.Println(jokes["jokes"][randomNumber])
}
