package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var breedCmd = &cobra.Command{
	Use:   "breed [犬種名]",
	Short: "指定した犬種の画像を取得",
	Long:  `指定した犬種のランダムな画像を取得する`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: 犬種が指定されていません")
			os.Exit(1) // エラー終了コードを設定
			return
		}

		breed := args[0]
		resp, err := http.Get("https://dog.ceo/api/breed/" + breed + "/images/random")
		if err != nil {
			fmt.Printf("Error: HTTPリクエストエラー: %v\n", err)
			os.Exit(1)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			fmt.Println("Error: 無効な犬種です")
			os.Exit(1)
			return
		}

		var dog DogResponse
		if err := json.NewDecoder(resp.Body).Decode(&dog); err != nil {
			fmt.Printf("Error: JSONデコードエラー: %v\n", err)
			os.Exit(1)
			return
		}

		fmt.Println(dog.Message)
	},
}

func init() {
	rootCmd.AddCommand(breedCmd)
}
