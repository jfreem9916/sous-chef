package main

type UPCItemDatastore interface {
	//Returns error if something bad or item has no entry yet
	GetItem(upc string) (UPCItemData, error)
	CreateItem(UPCItemData) error
}
