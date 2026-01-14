package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	generator "github.com/acsellers/go-invoice-generator"
)

const paymentDetails = `Please submit payment through our secure online portal:

<b>https://example.securepayments.cardpointe.com/pay</b>

Enter the <b>Total Amount Due $%s</b> and include the following in the Invoice Number field:
<b>Invoice Ref: INV-%s</b>

If you encounter any issues or need adjustments, please contact:
<b>%s - %s</b>`
const otherNotes = `All amounts shown in United States Dollars (USD). Thank you for partnering with Awesome Company to support your community fundraiser!`

func main() {

	doc, _ := generator.New(generator.Invoice, &generator.Options{
		TextTypeInvoice:    "INVOICE",
		AutoPrint:          true,
		CurrencySymbol:     "$",
		PageSize:           "Letter",
		HideDiscountColumn: true,
		HideTaxColumn:      false,
		HideTextType:       true,
		TextRefTitle:       "Invoice #",
		AdditionalMetas: []generator.TitleValue{
			{"Terms", "Net-30"},
			{"Due", time.Now().AddDate(0, 0, 30).Format("Jan 2, 2006")},
		},
		CompactTotals:    true,
		CompactAddress:   true,
		SingleColumn:     true,
		LogoHeight:       20,
		TextTotalTotal:   "Subtotal",
		TextTotalWithTax: "Total",
		TextTotalTax:     "Tax",
		GreyBgColor:      []int{255, 255, 255},
		DarkBgColor:      []int{255, 255, 255},
		TableBgColor:     []int{32, 95, 183},
		TableTextColor:   []int{255, 255, 255},
	})
	generator.SmallTextFontSize = 8
	doc.SetHeader(&generator.HeaderFooter{
		Text:       "This invoice is provided to Test Company by Awesome Company, if there are any issues, contact Reginald Monocle for corrections.",
		Pagination: true,
	})
	doc.SetRef("INV-12345")

	recNotes := fmt.Sprintf(
		strings.Replace(paymentDetails, "\n", "<br>", -1),
		fmt.Sprintf("%.2f", 497.00),
		"12345",
		"Reginald Monocle",
		"reginald.monocle@awesomecompany.com",
	)
	doc.SetDate(time.Now().Format("Jan 2, 2006"))
	doc.NoteList = []generator.TitleValue{
		{"Payment Details", recNotes},
		{"Notes", otherNotes},
	}

	doc.Options.AdditionalTotals = []generator.TitleValue{
		{"Paid", "$0.00"},
		{"Amount Due", "$497.00"},
	}

	imgPath := "company.png"
	if _, err := os.Stat(imgPath); os.IsNotExist(err) {
		imgPath = "./examples/extracted/company.png"
	}
	var img []byte
	if _, err := os.Stat(imgPath); err == nil {
		img, _ = os.ReadFile(imgPath)
	}
	doc.SetCompany(&generator.Contact{
		Name: "Awesome Company",
		Logo: img,
		Info: &generator.Address{
			Address:    "123 Main St.",
			City:       "Testtown",
			State:      "GA",
			PostalCode: "12345",
			Country:    "United States",
		},
		AdditionalInfo: []string{
			fmt.Sprintf("Contact: Reginald Monocle\nreginald.monocle@awesomecompany.com"),
		},
	})

	doc.SetCustomer(&generator.Contact{
		Name: "Test Company",
		Info: &generator.Address{
			Address:    "123 Main St.",
			Address2:   "Suite 123",
			City:       "Testtown",
			State:      "MA",
			PostalCode: "54321",
			Country:    "United States",
		},
	})

	doc.AppendItem(&generator.Item{
		Name:     "Fundraising Coupon",
		Quantity: "100",
		UnitCost: fmt.Sprintf("%.2f", 4.97),
	})
	built, err := doc.Build()
	if err != nil {
		fmt.Printf("Error building invoice PDF: %v\n", err)
	}
	err = built.OutputFileAndClose("invoice.pdf")
	if err != nil {
		fmt.Printf("Error saving invoice PDF: %v\n", err)
	}
}
