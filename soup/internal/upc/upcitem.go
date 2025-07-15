package upc

// For a given barcode, all related information
// unique per barcode
type UPCItemData struct {
	//  The barcode
	UPC          string
	Name         string
	Weight_grams float32
	Volume_ml    float32
	//if volume or weight is defined, the weight / volume
	//is the sum of the weight/volume of the quant.
	Quantity float32
	Price    float32
}
