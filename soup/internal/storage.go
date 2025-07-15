package main

import "soup/internal/upc"

type UPCItemDatastore interface {
	//Returns error if something bad or item has no entry yet
	GetItem(upc string) (upc.UPCItemData, error)
	CreateItem(upc.UPCItemData) error
}
