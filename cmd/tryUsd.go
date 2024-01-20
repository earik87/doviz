package cmd

import (
	"fmt"
	"strconv"

	"doviz/internal"

	"github.com/spf13/cobra"
)

func calculateTryUsd(try float64) float64 {
	usdTryParity := internal.GetUsdTryParity()

	return try / usdTryParity
}

var tryUsdCmd = &cobra.Command{
	Use:   "try-usd",
	Short: "TRY degerini USD cinsinden gosterir",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			value, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				fmt.Println("Yanlis arguman, lutfen bir sayi giriniz.")
				return
			}

			usdTry := fmt.Sprintf("%.2f", calculateTryUsd(value))
			fmt.Println(usdTry)
		} else {
			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(tryUsdCmd)

}
