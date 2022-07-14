package main

import (
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MatejaMaric/spending-tracker/models"
	"github.com/MatejaMaric/spending-tracker/services"
	"github.com/adrg/xdg"
)

func main() {
	dbPath, err := xdg.DataFile("spending-tracker/data.db")
	if err != nil {
		log.Fatal(err)
	}

	models.Connect(dbPath)

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

func createTransactionTable(w *fyne.Window) fyne.CanvasObject {
	transactions := services.GetTransactions()

	tableHeader := func() fyne.CanvasObject {
		transactionLabels := container.NewGridWithColumns(7,
			widget.NewLabel("MadeAt"),
			widget.NewLabel("ProcessedAt"),
			widget.NewLabel("Description"),
			widget.NewLabel("Paid"),
			widget.NewLabel("Received"),
			widget.NewLabel("Balance"),
			widget.NewLabel("PersonalDescription"),
		)
		return container.NewBorder(nil, nil, nil, widget.NewLabel("Actions"), transactionLabels)
	}

	newListItem := func() fyne.CanvasObject {
		transactionLabels := container.NewGridWithColumns(7,
			widget.NewLabel("MadeAt"),
			widget.NewLabel("ProcessedAt"),
			widget.NewLabel("Description"),
			widget.NewLabel("Paid"),
			widget.NewLabel("Received"),
			widget.NewLabel("Balance"),
			widget.NewLabel("PersonalDescription"),
		)
		editBtn := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), nil)
		return container.NewBorder(nil, nil, nil, editBtn, transactionLabels)
	}

	updateListItem := func(i widget.ListItemID, o fyne.CanvasObject) {
		c := o.(*fyne.Container)
		labels := c.Objects[0].(*fyne.Container)
		editBtn := c.Objects[1].(*widget.Button)

		madeAtLbl := labels.Objects[0].(*widget.Label)
		processedAtLbl := labels.Objects[1].(*widget.Label)
		descriptionLbl := labels.Objects[2].(*widget.Label)
		paidLbl := labels.Objects[3].(*widget.Label)
		receivedLbl := labels.Objects[4].(*widget.Label)
		balanceLbl := labels.Objects[5].(*widget.Label)
		personalDescriptionLbl := labels.Objects[6].(*widget.Label)

		madeAtLbl.SetText(transactions[i].MadeAt.Format("02.01.2006."))
		processedAtLbl.SetText(transactions[i].ProcessedAt.Format("02.01.2006."))
		descriptionLbl.SetText(transactions[i].Description)
		paidLbl.SetText(strconv.FormatFloat(transactions[i].Paid, 'f', 2, 64))
		receivedLbl.SetText(strconv.FormatFloat(transactions[i].Received, 'f', 2, 64))
		balanceLbl.SetText(strconv.FormatFloat(transactions[i].Balance, 'f', 2, 64))
		personalDescriptionLbl.SetText(transactions[i].PersonalDescription)

		editBtn.OnTapped = func() {
			personalDescriptionEntry := widget.NewEntry()
			personalDescriptionEntry.SetText(personalDescriptionLbl.Text)

			formItems := []*widget.FormItem{
				widget.NewFormItem("Made At", madeAtLbl),
				widget.NewFormItem("Processed At", processedAtLbl),
				widget.NewFormItem("Description", descriptionLbl),
				widget.NewFormItem("Paid", paidLbl),
				widget.NewFormItem("Received", receivedLbl),
				widget.NewFormItem("Balance", balanceLbl),
				widget.NewFormItem("Personal Description", personalDescriptionEntry),
			}

			formDialog := dialog.NewForm("Edit Transaction", "Save", "Cancel", formItems, nil, *w)

			formDialog.Show()
		}
	}

	return container.NewBorder(tableHeader(), nil, nil, nil, widget.NewList(
		func() int { return len(transactions) },
		newListItem,
		updateListItem,
	))
}

func createAddTransactionBtn(w *fyne.Window) *widget.Button {
	return widget.NewButtonWithIcon("Add Transactions", theme.ContentAddIcon(), func() {
		inputEntry := widget.NewMultiLineEntry()

		formItems := []*widget.FormItem{
			widget.NewFormItem("Input HTML", inputEntry),
		}

		formCallback := func(b bool) {
			if b {
				addedRows, err := services.CreateTransactions(inputEntry.Text)
				if err != nil {
					dialog.ShowError(err, *w)
				} else {
					dialog.ShowInformation("New Transactions Added", fmt.Sprintf("Number of added transactions: %d\n", addedRows), *w)
				}
			}
		}

		formDialog := dialog.NewForm("Add New Transactions", "Parse", "Cancel", formItems, formCallback, *w)
		formDialog.Show()
	})
}
