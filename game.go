package main

import (
	"fmt"
	"math/rand"
)



func main() {
	for true {
		player := Player{}

		game := Game{}

		name := getPlayerName()

		newPlayer := player.NewPlayer(name)

		newGame := game.NewGame(newPlayer)

		newGame.StartGame()
	}
}

type Game struct {
	Player       Player
}

func (g Game) NewGame(player Player) Game {
	return Game{
		Player:       player,
	}
}
func (g Game) StartGame() {
	FNum := 0
	fmt.Print("Sevimli raqamingizni kiriting : ")
	fmt.Scan(&FNum)

	fmt.Printf("Salom %s\n", g.Player.Name)
	fmt.Println(" Bu kompyuter o'ylagan sonni topish o'yini ")
	chances := 0
	fmt.Print(" O'yinda necha marta urinib ko'rmoqchisiz: ") 
	fmt.Scan(&chances)
	son := 0
	if chances<5{
		son=50
	} else if chances >=5 && chances <10{
		son=100
	} else {
		son=500
	}
    var	RandomNumber int=rand.Intn(son)
	for i := 0; i < chances; i++ {
		var n int
		fmt.Printf("%d - urinish \n son kiriting: ", i+1)
		fmt.Scan(&n)

		if n == RandomNumber {
			fmt.Println("TABRIKLAYMIZ")
			return
		} else if n > RandomNumber {
			fmt.Println("Random son siz kiritgan sondan kichkina")
		} else if n < RandomNumber {
			fmt.Println("Random son siz kiritgan sondam katta ")
		} else if n==FNum{
			fmt.Println("Random son sizning sevimli soningizga teng")
		} else {
			fmt.Println("Xato")
		}
	}
	fmt.Println(" Yutqazdingiz\n kompyuter o'ylagan son ", RandomNumber, " edi")
}

type Player struct {
	Name            string
	Age             int
	FNum            int
	Chances         int
}

func (p Player) NewPlayer(name string) Player {
	return Player{
		Name:    name,
	}
}
func getPlayerName() string{
	var (
		name            string
		age             int
	)
	fmt.Print("Ismingizni kiriting : ")
	fmt.Scan(&name)
	fmt.Print("Yoshingizni kiriting : ")
	fmt.Scan(&age)
	
	return name
}
