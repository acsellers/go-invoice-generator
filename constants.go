package generator

const (
	// Invoice define the "invoice" document type
	Invoice string = "INVOICE"

	// Quotation define the "quotation" document type
	Quotation string = "QUOTATION"

	// DeliveryNote define the "delievry note" document type
	DeliveryNote string = "DELIVERY_NOTE"

	// BaseMargin define base margin used in documents
	BaseMargin float64 = 10

	// BaseMarginTop define base margin top used in documents
	BaseMarginTop float64 = 20

	// HeaderMarginTop define base header margin top used in documents
	HeaderMarginTop float64 = 5

	// MaxPageHeight define the maximum height for a single page
	MaxPageHeight float64 = 260
)

// Cols offsets, add up to 190
var (
	// ItemColNameOffset ...
	ItemColNameOffset float64 = 10
	ItemColNameWidth  float64 = 70
	// ItemColUnitPriceOffset ...
	ItemColUnitPriceOffset float64 = ItemColNameOffset + ItemColNameWidth
	ItemColUnitPriceWidth  float64 = 33

	// ItemColQuantityOffset ...
	ItemColQuantityOffset float64 = ItemColUnitPriceOffset + ItemColUnitPriceWidth
	ItemColQuantityWidth  float64 = 10

	// ItemColTotalHTOffset ...
	ItemColTotalHTOffset float64 = ItemColQuantityOffset + ItemColQuantityWidth
	ItemColTotalHTWidth  float64 = 17

	// ItemColDiscountOffset ...
	ItemColDiscountOffset float64 = ItemColTotalHTOffset + ItemColTotalHTWidth
	ItemColDiscountWidth  float64 = 17

	// ItemColTaxOffset ...
	ItemColTaxOffset float64 = ItemColDiscountOffset + ItemColDiscountWidth
	ItemColTaxWidth  float64 = 18

	// ItemColTotalTTCOffset ...
	ItemColTotalTTCOffset float64 = ItemColTaxOffset + ItemColTaxWidth
	ItemColTotalTTCWidth  float64 = 23

	ItemRowWidth               float64 = 180
	ItemNameWidthNoTaxDiscount float64 = ItemRowWidth - (ItemColUnitPriceWidth + ItemColQuantityWidth + ItemColTotalTTCWidth)
	ItemNameWidthNoTax         float64 = ItemRowWidth - (ItemColUnitPriceWidth + ItemColQuantityWidth + ItemColDiscountWidth + ItemColTotalTTCWidth)
	ItemNameWidthNoDiscount    float64 = ItemRowWidth - (ItemColUnitPriceWidth + ItemColQuantityWidth + ItemColTotalHTWidth + ItemColTaxWidth + ItemColTotalTTCWidth)
)

var (
	// BaseTextFontSize define the base font size for text in document
	BaseTextFontSize float64 = 8

	// SmallTextFontSize define the small font size for text in document
	SmallTextFontSize float64 = 7

	// ExtraSmallTextFontSize define the extra small font size for text in document
	ExtraSmallTextFontSize float64 = 6

	// LargeTextFontSize define the large font size for text in document
	LargeTextFontSize float64 = 10
)
