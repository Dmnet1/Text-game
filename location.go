package main

type location struct {
	name            string
	description     string
	moreDescription string
	items           []Item
	ways            []string
}

func (l *location) checker(location map[string]location) (checkStatus bool) {
	for i := range location[PreviousLocation].ways {
		if ItemOne == location[PreviousLocation].ways[i] {
			checkStatus = true
			break
		} else {
			checkStatus = false
		}
	}
	return checkStatus
}

func (l *location) giveDescription(keyLocation string) string {
	answer = ""
	answer = Locations[keyLocation].description
	if Locations[keyLocation].ways != nil {
		for i := range Locations[keyLocation].ways {
			answer = answer + Locations[keyLocation].ways[i]
		}
	}
	return answer
}

func (l *location) giveMoreDescription(keyLocation string) string {
	answer = ""
	answer = Locations[keyLocation].moreDescription
	if Locations[keyLocation].ways != nil {
		for i := range Locations[keyLocation].ways {
			answer = answer + Locations[keyLocation].ways[i]
		}
	}

	return answer
}

func (l *location) showItemsInLocation(keyLocation string) []Item {
	return Locations[keyLocation].items
}

func (l *location) descriptionAfterTakingItem() {

}

func (l *location) deleteItemInLocation(item, keyLocation string) {
	for i := range Locations[keyLocation].items {
		if item == Locations[keyLocation].items[i].name {
			l.items = append(Locations[keyLocation].items[:i], Locations[keyLocation].items[i+1:]...) //удаление предмета из комнаты
			break
		}
	}
}

var Locations = map[string]location{}
var AnotherLocation = map[string]location{}

func GameInit() {
	Locations["кухня"] = location{
		name:            "кухня",
		description:     "кухня, ничего интересного. можно пройти - ",
		moreDescription: "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - ",
		items:           nil,
		ways:            []string{"коридор"},
	}
	Locations["коридор"] = location{
		name:            "коридор",
		description:     "ничего интересного. можно пройти - ",
		moreDescription: "",
		items:           nil,
		ways:            []string{"кухня", "комната", "улица"},
	}
	Locations["комната"] = location{
		name:            "комната",
		description:     "ты в своей комнате. можно пройти - ",
		moreDescription: "на столе: ключи, конспекты, на стуле - рюкзак. можно пройти - ",
		items:           []Item{{name: "рюкзак"}, {name: "ключи"}, {name: "конспекты"}},
		ways:            []string{"коридор"},
	}
	Locations["улица"] = location{
		name:            "улица",
		description:     "на улице весна. можно пройти - ",
		moreDescription: "",
		items:           nil,
		ways:            []string{"домой"},
	}
	AnotherLocation["универ"] = location{
		name:            "универ",
		description:     "универ, тут нужно учиться",
		moreDescription: "можно пойти домой",
		items:           nil,
		ways:            []string{"улица"},
	}
}
