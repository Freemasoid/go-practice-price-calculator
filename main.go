package main

import (
	"fmt"

	"github.com/Freemasoid/go-practice-price-calculator/cmdmanager"
	"github.com/Freemasoid/go-practice-price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	fmt.Println("4 tax rates are available: 0, 0.21, 0.1, 0.15")
	fmt.Println("Calculations are made for each tax rate individually")
	fmt.Println("When you enter all desired prices, to proceed with calculation hit ENTER without entering new price or submit 0 price")

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.New(cmdm, taxRate)
		priceJob.Process()
	}

}
