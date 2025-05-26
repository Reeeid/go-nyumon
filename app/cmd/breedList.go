/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// breedListCmd represents the breedList command
var breedListCmd = &cobra.Command{
	Use:   "breed-list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://dog.ceo/api/breeds/list")
		if err != nil {
			fmt.Printf("Error: HTTP Request failed: %v\n", err)
		}
		defer resp.Body.Close()
		var dogList DogListResponse
		if err := json.NewDecoder(resp.Body).Decode(&dogList); err != nil {
			fmt.Printf("Error: JSON Decode failed: %v\n", err)
			return
		}
		for i := 0; i < len(dogList.Message); i++ {
			fmt.Println(dogList.Message[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(breedListCmd)
}
