package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	/*exInfo, err := getExchangeInfo()
	if err != nil {
		println(err.Error())

	}
	getCurrPrices(exInfo)
	for key,value := range priceMap{
		fmt.Printf("%v, %f\n",key,value)
	}*/
	/*var input string
	for{
		_, err := fmt.Scanln(&input)
		if err != nil{
			print(err.Error())
		}

	}*/
	app := app.New()
	w := app.NewWindow("GoCryptoMon")
	w.SetContent(widget.NewLabel("Coins:"))

	w.ShowAndRun()
}
