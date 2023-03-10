package main

// абстрактная сумка для инвентаря
type Bag struct {
	items []Item
}

func NewBag(typeOfBag string) iBag {
	var bag iBag
	if typeOfBag == "рюкзак" {
		bag = &Bag{items: nil}
	}
	return bag
}

func (b *Bag) showItemsInBag() interface{} {
	var items []string
	for i := range b.items {
		items = append(items, b.items[i].name)
	}
	return items
}

type Item struct {
	name string
}

func (i *Item) newItem(item string) interface{} {
	it := &Item{name: item}
	return it.name
}

func GetItem(inventory iInventory, item string) interface{} {
	return inventory.newItem(item)
}
