package main

type player struct {
	location      iLocation
	backpack      iBag
	inventory     iInventory
	cmd           iCmd
	actions       iPlayerComplexAction
	checker       iCheckerLocation
	simpleActions iPlayersSimpleActions
}

func NewPlayer() *player {
	return &player{
		location:      nil,
		backpack:      nil,
		inventory:     nil,
		cmd:           nil,
		actions:       nil,
		checker:       nil,
		simpleActions: nil,
	}
}

func (p *player) createBag(keyLocation string) iBag { //надеть рюкзак
	itemInLocation := p.location.showItemsInLocation(keyLocation)
	var backpack iBag
	for item := range itemInLocation {
		if Action == "надеть" && ItemOne == itemInLocation[item].name {
			backpack = NewBag(itemInLocation[item].name)
			p.location.deleteItemInLocation(ItemOne, keyLocation)
		}
	}
	return backpack
}

func (p *player) goToLocation(location map[string]location) string {
	for k, _ := range location {
		if k == ItemOne {
			checkStatus := p.checker.checker(location)
			if checkStatus == true {
				answer = p.location.giveDescription(ItemOne)
				PreviousLocation = ItemOne
			} else {
				answer = "нет пути в " + ItemOne
			}
		}
	}
	return answer
}

func (p *player) lookAround(keyLocation string) string {

	answer = p.location.giveMoreDescription(keyLocation)

	return answer
}

func (p *player) takeSomethingFromLocation(keyLocation string, bag iBag) string {
	if Action == "взять" {
		items := p.location.showItemsInLocation(keyLocation)
		for item := range items {
			if ItemOne == items[item].name {
				p.actions.addItemInBag(ItemOne, bag)
				answer = "предмет добавлен в инвентарь: " + ItemOne
			} else {
				answer = "нет такого"
			}
		}
	}
	return answer
}

func (p *player) applySomething(bag iBag) string {
	s := bag.showItemsInBag()
	items := s.([]string)
	for item := range items {
		if Action == "применить" && ItemOne == items[item] && ItemTwo == "дверь" {
			answer = "дверь открыта"
		}
		if ItemOne != items[item] {
			answer = "нет такого"
		}
		if ItemTwo != "дверь" {
			answer = "не к чему применить"
		}
	}
	return answer
}

// создает предмет и добавляет в рюкзак
func (p *player) addItemInBag(item string, bag iBag) {
	i := GetItem(p.inventory, item)
	s := bag.showItemsInBag()
	it := i.(string)
	items := s.([]string)
	items = append(items, it)
}
