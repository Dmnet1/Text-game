package main

import (
	"fmt"
	"strings"
)

type Location struct {
	Corridor, Kitchen, Street, Room string
	Answer                          map[string]string
	StatusLocation                  map[string]bool
}

func newLocation(Corridor, Kitchen, Street, Room string) *Location {
	return &Location{
		Corridor:       Corridor,
		Kitchen:        Kitchen,
		Street:         Street,
		Room:           Room,
		Answer:         nil,
		StatusLocation: nil,
	}
}

type Items struct {
	Key, Abstracts, Backpack, Door, onTable string
	BackpackTrigger                         bool
}

func newItems(BackpackTrigger bool, Key, Abstracts, Backpack, Door, onTable string) *Items {
	return &Items{
		Key:             Key,
		Abstracts:       Abstracts,
		Backpack:        Backpack,
		Door:            Door,
		onTable:         onTable,
		BackpackTrigger: BackpackTrigger,
	}
}

type Commands struct {
	Answer, LookAround, Go, PutOn, Take, Apply string
}

func newCommands(Answer, LookAround, Go, PutOn, Take, Apply string) *Commands {
	return &Commands{
		Answer:     Answer,
		LookAround: LookAround,
		Go:         Go,
		PutOn:      PutOn,
		Take:       Take,
		Apply:      Apply,
	}
}

type ResultsAfterSplit struct {
	Action, ItemOne, ItemTwo string
}

func newResultsAfterSplit(Action, ItemOne, ItemTwo string) *ResultsAfterSplit {
	return &ResultsAfterSplit{
		Action:  Action,
		ItemOne: ItemOne,
		ItemTwo: ItemTwo,
	}
}

func (r *ResultsAfterSplit) splitTheCommand(command string) {
	splitCmd := strings.Fields(command)
	for i := 0; i < len(splitCmd); i++ {
		if splitCmd[i] == "осмотреться" || splitCmd[i] == "идти" || splitCmd[i] == "взять" ||
			splitCmd[i] == "применить" || splitCmd[i] == "надеть" {
			r.Action = splitCmd[i]
		}
		if splitCmd[i] == "комната" || splitCmd[i] == "коридор" || splitCmd[i] == "кухня" ||
			splitCmd[i] == "улица" || splitCmd[i] == "ключи" || splitCmd[i] == "конспекты" ||
			splitCmd[i] == "дверь" || splitCmd[i] == "рюкзак" {
			r.ItemOne = splitCmd[i]
		}
		if splitCmd[i] == "дверь" {
			r.ItemTwo = splitCmd[i]
		}
		/*switch splitCmd[i] {
		case "осмотреться":
			r.Action = splitCmd[i]

		case "идти":
			r.Action = splitCmd[i]
		case "комната":
			r.ItemOne = splitCmd[i]
		case "коридор":
			r.ItemOne = splitCmd[i]
		case "кухня":
			r.ItemOne = splitCmd[i]
		case "улица":
			r.ItemOne = splitCmd[i]

		case "взять":
			r.Action = splitCmd[i]
		case "применить":
			r.Action = splitCmd[i]
		case "ключи":
			r.ItemOne = splitCmd[i]
		case "конспекты":
			r.ItemOne = splitCmd[i]
		case "дверь":
			r.ItemTwo = splitCmd[i]

		case "надеть":
			r.Action = splitCmd[i]
		case "рюкзак":
			r.ItemOne = splitCmd[i]
		}*/
	}
}

// 1
func lookAround(l Location, r ResultsAfterSplit, c *Commands, i Items) {
	if r.Action == c.LookAround {
		if l.StatusLocation[l.Kitchen] == true {
			c.Answer = "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
		}
		if l.StatusLocation[l.Room] == true {
			if i.Backpack != "" && i.Key != "" && i.Abstracts != "" {
				c.Answer = "на столе: ключи, конспекты, на стуле - рюкзак. можно пройти - коридор"
			}
			if i.Backpack == "" && i.Key != "" && i.Abstracts != "" {
				c.Answer = "на столе: ключи, конспекты. можно пройти - коридор"
			}
			if i.Backpack == "" && i.Key == "" && i.Abstracts != "" {
				c.Answer = "на столе: конспекты. можно пройти - коридор"
			}
			if i.Backpack == "" && i.Key == "" && i.Abstracts == "" {
				c.Answer = "пустая комната. можно пройти - коридор"
			}
		}
	}
}

func initializeMap(l *Location) {
	l.Answer = make(map[string]string)
	l.Answer[l.Room] = "ты в своей комнате. можно пройти - " + l.Corridor
	l.Answer[l.Kitchen] = l.Kitchen + ", ничего интересного. можно пройти - " + l.Corridor
	l.Answer[l.Corridor] = "ничего интересного. можно пройти - " + l.Kitchen + ", " + l.Room + ", " + l.Street
	l.Answer[l.Street] = "на улице весна. можно пройти - домой"
	l.Answer["нет пути"] = "нет пути в "

	l.StatusLocation = make(map[string]bool)
	l.StatusLocation[l.Room] = false
	l.StatusLocation[l.Kitchen] = true
	l.StatusLocation[l.Corridor] = false
	l.StatusLocation[l.Street] = false
}

// 2
func playerGoToLocation(r ResultsAfterSplit, c *Commands, l *Location, i Items) {
	if r.Action == c.Go && r.ItemOne != l.Street {
		for keyLoc, keyVal := range l.Answer {
			if keyLoc == r.ItemOne && r.ItemOne != l.Corridor && l.StatusLocation[l.Corridor] == false { //In corridor
				c.Answer = l.Answer["нет пути"] + r.ItemOne
			} else if keyLoc == r.ItemOne && r.ItemOne == l.Corridor && l.StatusLocation[l.Corridor] == false {
				c.Answer = l.Answer[l.Corridor]
				l.StatusLocation[l.Corridor] = true
				l.StatusLocation[l.Kitchen] = false
				l.StatusLocation[l.Room] = false
				l.StatusLocation[l.Street] = false
				break
			}
			if keyLoc == r.ItemOne && r.ItemOne != l.Corridor && l.StatusLocation[l.Corridor] == true {
				l.StatusLocation[l.Corridor] = false
				for keyLocStat, _ := range l.StatusLocation {
					if keyLocStat == r.ItemOne {
						l.StatusLocation[keyLocStat] = true
						break
					}
				}
				c.Answer = keyVal
			}
		}
	}
	if r.ItemOne == l.Street && i.Door == "" {
		c.Answer = l.Answer[l.Street]
		l.StatusLocation[l.Street] = true
	} else if r.ItemOne == l.Street && i.Door != "" {
		c.Answer = "дверь закрыта"
	}

}

// 3
func backPackTrigger(r ResultsAfterSplit, c *Commands, l Location, i *Items) {
	if r.Action == c.PutOn && l.StatusLocation[l.Room] == true {
		if r.ItemOne == i.Backpack && i.Backpack != "" {
			i.BackpackTrigger = true
			c.Answer = "вы надели: " + i.Backpack
			i.Backpack = ""
		}
	} else if r.Action == c.PutOn && l.StatusLocation[l.Room] != true {
		c.Answer = "нет такого"
	}
}

// 4
func playerTakeSomething(r ResultsAfterSplit, i *Items, c *Commands, l Location) {
	if r.Action == c.Take {
		if i.BackpackTrigger == true && l.StatusLocation[l.Room] == true {
			if r.ItemOne == i.Key || r.ItemOne == i.Abstracts {
				c.Answer = "предмет добавлен в инвентарь: " + r.ItemOne
				if r.ItemOne == i.Key {
					i.Key = ""
				}
				if r.ItemOne == i.Abstracts {
					i.Abstracts = ""
					i.onTable = "пустая комната"
				}
				return
			}
			c.Answer = "нет такого"
			return
		}
		c.Answer = "некуда класть"
	}
}

// 5
func playerApplyKey(r ResultsAfterSplit, i *Items, c *Commands, l Location) {
	if r.Action == c.Apply {
		if l.StatusLocation[l.Corridor] == true {
			if i.Key == "" {
				c.Answer = "дверь открыта"
				i.Door = ""
			} else {
				c.Answer = "нет предмета в инвентаре - " + r.ItemOne
			}
		} else {
			c.Answer = "не к чему применить"
		}
	}
}

// 6
func deleteResultsAfterSplit(r *ResultsAfterSplit) {
	r.Action = ""
	r.ItemOne = ""
	r.ItemTwo = ""
}

func initGame() {
	/*fmt.Scanln(&Cmd)*/ Cmd = []string{"осмотреться", "идти коридор",
		"идти кухня", "идти комната", "идти коридор",
		"идти кухня", "идти коридор", "идти комната", "осмотреться", "взять ключи",
		"надеть рюкзак", "взять ключи", "осмотреться",
		"взять конспекты", "идти коридор", "применить ключи дверь", "идти улица"}

	location := newLocation("коридор", "кухня", "улица", "комната")
	items := newItems(false, "ключи", "конспекты", "рюкзак", "дверь", "на столе: ")
	commands := newCommands("", "осмотреться", "идти", "надеть", "взять", "применить")
	resultsAfterSplit := newResultsAfterSplit("", "", "")
	initializeMap(location)

	for i := range Cmd {
		resultsAfterSplit.splitTheCommand(Cmd[i])
		lookAround(*location, *resultsAfterSplit, commands, *items)
		playerGoToLocation(*resultsAfterSplit, commands, location, *items)
		backPackTrigger(*resultsAfterSplit, commands, *location, items)
		playerTakeSomething(*resultsAfterSplit, items, commands, *location)
		playerApplyKey(*resultsAfterSplit, items, commands, *location)

		fmt.Println("cmd:", Cmd[i], "\n", "result:", commands.Answer)
		/*fmt.Println("corridor", location.StatusLocation[location.Corridor], "/", "Kitchen", location.StatusLocation[location.Kitchen],
		"/", "Room", location.StatusLocation[location.Room], "/", "Street", location.StatusLocation[location.Street])*/

		deleteResultsAfterSplit(resultsAfterSplit)

		command = Cmd[i]
		Answer = commands.Answer
		handleCommand(command)
		commands.Answer = ""
	}

}

var Cmd []string
var command string
var Answer string

func handleCommand(cmd string) string {
	return Answer
}

func main() {
	initGame()
}
