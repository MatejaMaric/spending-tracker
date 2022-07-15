package gui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MatejaMaric/spending-tracker/services"
)

func createTransactionTable(w *fyne.Window) fyne.CanvasObject {
	transactions := services.GetTransactions()

	tableHeader := func() fyne.CanvasObject {
		transactionLabels := container.NewGridWithColumns(7,
			widget.NewLabel("Made At"),
			widget.NewLabel("Processed At"),
			widget.NewLabel("Description"),
			widget.NewLabel("Paid"),
			widget.NewLabel("Received"),
			widget.NewLabel("Balance"),
			widget.NewLabel("Personal Description"),
		)
		return container.NewBorder(nil, nil, nil, widget.NewLabel("Actions"), transactionLabels)
	}

	newListItem := func() fyne.CanvasObject {
		madeAtLbl := widget.NewLabel("MadeAt")
		processedAtLbl := widget.NewLabel("ProcessedAt")
		descriptionLbl := widget.NewLabel("Description")
		paidLbl := widget.NewLabel("Paid")
		receivedLbl := widget.NewLabel("Received")
		balanceLbl := widget.NewLabel("Balance")
		personalDescriptionLbl := widget.NewLabel("PersonalDescription")

		madeAtLbl.Wrapping = fyne.TextWrapBreak
		processedAtLbl.Wrapping = fyne.TextWrapBreak
		descriptionLbl.Wrapping = fyne.TextWrapBreak
		paidLbl.Wrapping = fyne.TextWrapBreak
		receivedLbl.Wrapping = fyne.TextWrapBreak
		balanceLbl.Wrapping = fyne.TextWrapBreak
		personalDescriptionLbl.Wrapping = fyne.TextWrapBreak

		transactionLabels := container.NewGridWithColumns(7,
			madeAtLbl,
			processedAtLbl,
			descriptionLbl,
			paidLbl,
			receivedLbl,
			balanceLbl,
			personalDescriptionLbl,
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
				widget.NewFormItem("Made At", widget.NewLabel(madeAtLbl.Text)),
				widget.NewFormItem("Processed At", widget.NewLabel(processedAtLbl.Text)),
				widget.NewFormItem("Description", widget.NewLabel(descriptionLbl.Text)),
				widget.NewFormItem("Paid", widget.NewLabel(paidLbl.Text)),
				widget.NewFormItem("Received", widget.NewLabel(receivedLbl.Text)),
				widget.NewFormItem("Balance", widget.NewLabel(balanceLbl.Text)),
				widget.NewFormItem("Personal Description", personalDescriptionEntry),
			}

			formCallback := func(b bool) {
				if b {
					err := services.UpdateTransactionPersonalDescription(transactions[i].ID, personalDescriptionEntry.Text)
					if err != nil {
						dialog.ShowError(err, *w)
					} else {
						dialog.ShowInformation("Success!", "Transaction successfully edited!", *w)
					}
				}
			}

			formDialog := dialog.NewForm("Edit Transaction", "Save", "Cancel", formItems, formCallback, *w)

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
