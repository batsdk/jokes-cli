/*
Copyright © 2023 Shemil Kaweesha <shemil.business@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {

	var joke Joke

	resBytes := getJokesData("https://icanhazdadjoke.com/")
	err := json.Unmarshal(resBytes, &joke)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(joke.Joke)

}

func getJokesData(baseUrl string) []byte {
	req, err := http.NewRequest(http.MethodGet, baseUrl, nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Jokes CLI")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	resBytes, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	return resBytes

}
