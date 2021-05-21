package main

import "fmt"

type Item interface {
	use()
}

type DefaultItem struct {
	itemName string
}

func (item *DefaultItem) use() {
	fmt.Println(item.itemName + " use it")
}
