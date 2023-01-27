package main

import "fmt"

type room struct {
	backpackInTheRoom, abstractsInTneRoom, keyInTneRoom bool
}

// 1.начальное состояние предметов в комнате.
func (r *room) addThingsInTheRoom() {
	(*r).abstractsInTneRoom = true
	(*r).keyInTneRoom = true
	(*r).backpackInTheRoom = true
	//fmt.Println("Лекции находятся в комнате - ", (*r).abstractsInTneRoom)
	//fmt.Println("Ключи находятся в комнате - ", (*r).keyInTneRoom)
	//fmt.Println("Рюкзак находятся в комнате - ", (*r).backpackInTheRoom)
}

// Состояние меняется, когда вещь добавляется в рюкзак. Возможно нужно сделать обычным присваиванием в методе putTheItemInTheBackpack,
// чтобы не использовать лишние методы.
func (r *room) deleteAbstractsInTneRoom() {
	(*r).abstractsInTneRoom = false
	//fmt.Println("Состояние конспектов в комнате", (*r).abstractsInTneRoom)
}
func (r *room) deleteKeyInTneRoom() {
	(*r).keyInTneRoom = false
	//fmt.Println("Состояние ключей в комнате", (*r).keyInTneRoom)
}
func (r *room) deleteBackpackInTneRoom() {
	(*r).backpackInTheRoom = false
	//fmt.Println("Состояние рюкзака в комнате", (*r).backpackInTheRoom)
}

type Player struct {
	Command, Answer                                        string
	backpackIsOn, abstractsInTheBackpack, keyInTheBackpack bool
}

// Игровой мир
type GameWorld struct {
	corridorStatus, kitchenStatus, roomStatus bool
	*Player
	*room

	//street
	//Player
}

// Вещи, которые налутал игрок
func (g *GameWorld) getItemsFromTheRoom(s string) string {
	if (*g).abstractsInTneRoom == true || (*g).keyInTneRoom == true {
		switch s {
		case "взять конспекты":
			(*g).abstractsInTheBackpack = true
			(*g).Answer = "предмет добавлен в инвентарь: конспекты"

		case "взять ключи":
			(*g).keyInTheBackpack = true
			(*g).Answer = "предмет добавлен в инвентарь: ключи"
		}
	} else {
		(*g).Answer = "нет такого"
	}

	//fmt.Println((*g).Answer)
	return (*g).Answer
}

// 2. Стартовая позиция игрока относительно мира.
func (g *GameWorld) addPlayerPosition() {
	(*g).kitchenStatus = true
	(*g).corridorStatus = false
	(*g).roomStatus = false
}

// 3. Метод должен работать в главном цикле, после него должна находиться логика добавления предметов.
// В методе описано позиционирование игрока.
func (g *GameWorld) getPlayerPosition(command string) (c string) {
	if (*g).kitchenStatus == true {
		switch command {
		case "осмотреться":
			(*g).Answer = "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
			(*g).corridorStatus = false
			(*g).roomStatus = false
			(*g).kitchenStatus = true
			c = command
			//fmt.Println("1", Command, "-", g.Answer)
			//fmt.Println("kitchenStatus-", g.kitchenStatus, "corridorStatus-", g.corridorStatus, "roomStatus-", g.roomStatus)
		case "идти коридор":
			(*g).Answer = "ничего интересного. можно пройти - кухня, комната, улица"
			(*g).corridorStatus = true
			(*g).kitchenStatus = false
			(*g).roomStatus = false
			c = command
			//fmt.Println("2", Command, "-", g.Answer)
			//fmt.Println("kitchenStatus-", g.kitchenStatus, "corridorStatus-", g.corridorStatus, "roomStatus-", g.roomStatus)
		case "идти комната":
			(*g).Answer = "нет пути в комната"
			(*g).kitchenStatus = true
			(*g).corridorStatus = false
			(*g).roomStatus = false
			c = command
			//fmt.Println("3", Command, "-", g.Answer)
			//fmt.Println("kitchenStatus-", g.kitchenStatus, "corridorStatus-", g.corridorStatus, "roomStatus-", g.roomStatus)
		}
		//fmt.Println("kitchenStatus-", g.kitchenStatus, "corridorStatus-", g.corridorStatus, "roomStatus-", g.roomStatus)
		c = command
	}

	if (*g).corridorStatus == true {
		if command == "идти комната" {
			(*g).Answer = "ты в своей комнате. можно пройти - коридор"
			(*g).roomStatus = true
			(*g).kitchenStatus = false
			(*g).corridorStatus = false
			c = command
			//fmt.Println("4", Command, "-", g.Answer)
			//fmt.Println("kitchenStatus-", g.kitchenStatus, "corridorStatus-", g.corridorStatus, "roomStatus-", g.roomStatus)
		}
		c = command
	}

	//fmt.Println("kitchenStatus-", g.kitchenStatus, "corridorStatus-", g.corridorStatus, "roomStatus-", g.roomStatus)
	if (*g).roomStatus == true {
		if g.abstractsInTneRoom == true && g.keyInTneRoom == true && g.backpackInTheRoom == true {
			if command == "осмотреться" {
				(*g).Answer = "на столе: ключи, конспекты, на стуле - рюкзак. можно пройти - коридор"
				(*g).roomStatus = true
				(*g).kitchenStatus = false
				(*g).corridorStatus = false
				c = command
				//fmt.Println("5", Command, "-", g.Answer)
				//fmt.Println("kitchenStatus-", g.kitchenStatus, "corridorStatus-", g.corridorStatus, "roomStatus-", g.roomStatus)
			}
			if command == "идти коридор" {
				(*g).Answer = "ничего интересного. можно пройти - кухня, комната, улица"
				(*g).corridorStatus = true
				(*g).kitchenStatus = false
				(*g).roomStatus = false
				c = command
			}
		}
		if g.abstractsInTneRoom == true && g.keyInTneRoom == true && g.backpackInTheRoom == false {
			if command == "осмотреться" {
				(*g).Answer = "на столе: ключи, конспекты. можно пройти - коридор"
				(*g).roomStatus = true
				(*g).kitchenStatus = false
				(*g).corridorStatus = false
				c = command
				//fmt.Println("6", Command, "-", g.Answer)
				//fmt.Println("kitchenStatus-", g.kitchenStatus, "corridorStatus-", g.corridorStatus, "roomStatus-", g.roomStatus)
			}
			if command == "идти коридор" {
				(*g).Answer = "ничего интересного. можно пройти - кухня, комната, улица"
				(*g).corridorStatus = true
				(*g).kitchenStatus = false
				(*g).roomStatus = false
				c = command
			}
		}
		if g.abstractsInTneRoom == true && g.keyInTneRoom == false && g.backpackInTheRoom == false {
			if command == "осмотреться" {
				(*g).Answer = "на столе: конспекты. можно пройти - коридор"
				(*g).roomStatus = true
				(*g).kitchenStatus = false
				(*g).corridorStatus = false
				c = command
				//fmt.Println("7", Command, "-", g.Answer)
				//fmt.Println("kitchenStatus-", g.kitchenStatus, "corridorStatus-", g.corridorStatus, "roomStatus-", g.roomStatus)
			}
			if command == "идти коридор" {
				(*g).Answer = "ничего интересного. можно пройти - кухня, комната, улица"
				(*g).corridorStatus = true
				(*g).kitchenStatus = false
				(*g).roomStatus = false
				c = command
			}
		}
		if g.abstractsInTneRoom == false && g.keyInTneRoom == false && g.backpackInTheRoom == false {
			if command == "осмотреться" {
				(*g).Answer = "пустая комната. можно пройти - коридор"
				(*g).roomStatus = true
				(*g).kitchenStatus = false
				(*g).corridorStatus = false
				c = command
				//fmt.Println("8", Command, "-", g.Answer)
				//fmt.Println("kitchenStatus-", g.kitchenStatus, "corridorStatus-", g.corridorStatus, "roomStatus-", g.roomStatus)
			}
			if command == "идти коридор" {
				(*g).Answer = "ничего интересного. можно пройти - кухня, комната, улица"
				(*g).corridorStatus = true
				(*g).kitchenStatus = false
				(*g).roomStatus = false
				c = command
			}
		}
		c = command
	}
	//fmt.Println("kitchenStatus-", g.kitchenStatus, "corridorStatus-", g.corridorStatus, "roomStatus-", g.roomStatus)
	return c
}

// 4. Метод описывает взаимодействие игрока и рюкзака.
func (g *GameWorld) takeTheBackpack(command string) (cmd string) { //метод вызывается после того, как в определенной комнате будет команда "надеть рюкзак"
	if (*g).roomStatus == true {
		if command == "надеть рюкзак" {
			//fmt.Println("зашел")
			(*g).backpackIsOn = true
			(*g).backpackInTheRoom = false
			(*g).Answer = "вы надели: рюкзак"
			//fmt.Println(Command, (*g).backpackIsOn)
			//fmt.Println("состояние рюкзака в комнате - ", (*g).backpackInTheRoom)
			//awr = "вы надели: рюкзак"
		}
	}
	cmd = command
	return cmd
}

// 5. Метод описывает добавление предметов в рюкзак.
func (g *GameWorld) putTheItemInTheBackpack(c string) /*(Answer string)*/ {
	if (*g).backpackIsOn == true {
		switch c {
		case "взять ключи":
			g.getItemsFromTheRoom(c)
			g.deleteKeyInTneRoom()
		case "взять конспекты":
			g.getItemsFromTheRoom(c)
			g.deleteAbstractsInTneRoom()
		case "взять телефон":
			(*g).Answer = "нет такого"
			//fmt.Println("попытка взять несуществующий телефон -", (*g).Answer)
		}
	} else if (*g).backpackIsOn == false && (c == "взять ключи" || c == "взять конспекты") {
		(*g).Answer = "некуда класть"
	}

	//return Answer
}

/*type street struct {
}*/

func (g *GameWorld) GetAnswer() string {
	fmt.Println((*g).Answer)
	return (*g).Answer
}

func handleCommand(cmd string) string {
	/*c := GameWorld.GetAnswer()
	fmt.Println(cmd, "-", c)*/

	return
}

func initGame() {
	cmd := []string{"осмотреться", "идти коридор",
		"идти комната", "осмотреться", "взять ключи",
		"надеть рюкзак", "взять ключи", "осмотреться",
		"взять конспекты", "идти коридор"}

	g := newGameWorld(false, false,
		false,
		*newPlayer("", "", false, false, false),
		*newRoom(false, false, false))

	g.addThingsInTheRoom()
	g.addPlayerPosition()

	for i := range cmd {
		c := cmd[i]
		g.putTheItemInTheBackpack(g.takeTheBackpack(g.getPlayerPosition(c)))
		//g.GetAnswer()
		handleCommand(g.Player.Command)

		fmt.Println("ИТЕРАЦИЯ-", i, "/", cmd[i], "/", g.Answer)

	}
}
func newPlayer(Command, Answer string, backpackIsOn, abstractsInTheBackpack, keyInTheBackpack bool) *Player {
	return &Player{
		Command:                Command,
		Answer:                 Answer,
		backpackIsOn:           backpackIsOn,
		abstractsInTheBackpack: abstractsInTheBackpack,
		keyInTheBackpack:       keyInTheBackpack,
	}
}

func newRoom(backpackInTheRoom, abstractsInTneRoom, keyInTneRoom bool) *room {
	return &room{
		backpackInTheRoom:  backpackInTheRoom,
		abstractsInTneRoom: abstractsInTneRoom,
		keyInTneRoom:       keyInTneRoom,
	}
}

func newGameWorld(corridorStatus, kitchenStatus, roomStatus bool, Player Player, room room) *GameWorld {
	return &GameWorld{
		corridorStatus: corridorStatus,
		kitchenStatus:  kitchenStatus,
		roomStatus:     roomStatus,
		Player:         &Player,
		room:           &room,
	}
}

func main() {
	initGame()

}
