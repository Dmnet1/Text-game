package main

import "fmt"

type iCmd interface {
	getCommand(command string) (cmd []string)
	clearCommand()
}

type iDescription interface {
	descriptionAfterTakingItem()
}

type iCheckerLocation interface {
	checker(location map[string]location) (checkStatus bool)
}

type iLocation interface {
	giveDescription(keyLocation string) string
	giveMoreDescription(keyLocation string) string
	showItemsInLocation(keyLocation string) []Item
	deleteItemInLocation(keyLocation, key string)
}

type bagCreator interface {
	createBag(keyLocation string) iBag
}

// Для взаимодействия с предметами
type iPlayerComplexAction interface {
	takeSomethingFromLocation(keyLocation string, bag iBag) string
	applySomething(bag iBag) string
	addItemInBag(item string, bag iBag)
}

// Для перемещения
type iPlayersSimpleActions interface {
	goToLocation(location map[string]location) string
	lookAround(keyLocation string) string
}

// Создает любой тип инвентаря
type iInventory interface {
	newItem(item string) interface{}
}

// Создает любой тип сумки с возможностью проверить содержимое
type iBag interface {
	showItemsInBag() interface{}
}

var answer string

// var cmd *string
var Action string
var ItemOne string
var ItemTwo string
var PreviousLocation string

func main() {
	/*var creature string = "shark"
	var pointer *string = &creature

	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)

	fmt.Println("*pointer =", *pointer)

	*pointer = "jellyfish"
	fmt.Println("*pointer =", *pointer)*/
	incomingCmd := []string{"осмотреться", "идти коридор"}
	GameInit()
	playerOne := NewPlayer()
	command := newResultsAfterSplit()
	PreviousLocation = "кухня"
	answer = "not nil"
	ItemOne = "кухня"
	for i := range incomingCmd {
		cmd := command.getCommand(incomingCmd[i])
		if cmd[0] == "идти" && (cmd[1] == "кухня" || cmd[1] == "коридор" || cmd[1] == "комната" || cmd[1] == "улица") {
			playerOne.simpleActions.goToLocation(Locations)
		}
		if cmd[0] == "осмотреться" {
			//playerOne.lookAround(cmd[0])
			playerOne.lookAround(cmd[0])

		}
		if cmd[0] == "взять" {

		}

		if cmd[0] == "применить" {

		}

		fmt.Println(answer, cmd[0], Locations[PreviousLocation])

	}
	fmt.Println(answer, Locations[PreviousLocation])
}
