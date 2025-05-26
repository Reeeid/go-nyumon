package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// set flag for command argument
var images int

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Show a random dog image",
	Long:  `This command fetches a random dog image from the Dog API`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < images; i++ {
			resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			var dog DogResponse
			if err := json.NewDecoder(resp.Body).Decode(&dog); err != nil {
				fmt.Println("Error decoding JSON:", err)
				return
			}
			fmt.Println(dog.Message)
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.Flags().IntVarP(&images, "images", "i", 1, "Number of images to fetch")

}
