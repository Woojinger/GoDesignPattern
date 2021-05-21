package main

import "fmt"

type creator interface {
	end(itemName string) string
	init(itemName string) string
	createItem(itemName string) Item
}

type DefaultItemCreator struct {
	creator
}

func (c *DefaultItemCreator) create(itemName string) Item {
	c.init(itemName)
	item := c.createItem(itemName)
	c.end(itemName)
	return item
}

func (c *DefaultItemCreator) end(itemName string) string {
	fmt.Println("end!")
	return itemName
}

func (c *DefaultItemCreator) init(itemName string) string {
	fmt.Println("init Task!")
	return itemName
}

func (c *DefaultItemCreator) createItem(itemName string) Item {
	return &DefaultItem{itemName: itemName}
}
