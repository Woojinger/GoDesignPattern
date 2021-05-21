package main

func main() {
	creator := DefaultItemCreator{}

	item1 := creator.create("sword")
	item2 := creator.create("shield")
	item3 := creator.create("armor")

	item1.use()
	item2.use()
	item3.use()
}
