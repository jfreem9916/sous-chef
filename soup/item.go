package main

// These are items in your inventory
// Not processed
type ScannedItem struct {
	UPCItemData
	// Number of items scanned
	Count int16
	Tags  []string
}

type Ingredient struct {
	Tag   string
	Items []ScannedItem
}
