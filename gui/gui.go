package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	a := app.New()
	w := a.NewWindow("Spending Tracker")

	transactionTable := createTransactionTable(&w)

	addTransactionsBtn := createAddTransactionBtn(&w)

	exitBtn := widget.NewButtonWithIcon("Exit", theme.CancelIcon(), func() {
		w.Close()
	})

	w.SetContent(container.NewBorder(nil, container.NewVBox(addTransactionsBtn, exitBtn), nil, nil, transactionTable))

	w.Resize(fyne.NewSize(1000, 700))
	w.CenterOnScreen()
	w.Show()
	a.Run()
}
