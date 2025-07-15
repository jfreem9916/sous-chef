package main

import "soup/upc"

type UPCItemDatastore interface {
	//Returns error if something bad or item has no entry yet
	GetItem(upc string) (upc.UPCItemData, error)
	CreateItem(upc.UPCItemData) error
}
