package cmd

import (
	"fmt"
	"strconv"

	"doviz/internal"

	"github.com/leekchan/accounting"
	"github.com/spf13/cobra"
)

func calculateEurTry(eur float64) float64 {
	eurTryParity := internal.GetEurTryParity()

	return eur * eurTryParity
}

var eurTryCmd = &cobra.Command{
	Use:   "eur-try",
	Short: "EUR degerini TRY cinsinden gosterir",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			value, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				fmt.Println("Yanlis arguman, lutfen bir sayi giriniz.")
				return
			}

			ac := accounting.Accounting{Symbol: "₺", Precision: 2, Thousand: ",", Decimal: "."}
			fmt.Println(ac.FormatMoney(calculateEurTry(value)))
		} else {
			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(eurTryCmd)

}
