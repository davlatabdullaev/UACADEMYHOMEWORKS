package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// player := Player{}
	// game   := Game{}
	player := getPlayerInfo()
	game := NewGame(player)
	RandomNumber(player, game)
	StartGame(player, game)
}

type Player struct {
	name    string
	chances int
	age     int
	FN      int
}
type Game struct {
	RN     int
	PLAYER Player
}

func getPlayerInfo() Player {
	var (
		Name    string
		Chances int
		Age     int
		fn      int
		player  Player
	)
	fmt.Println("Ismingizni kiriting: ")
	fmt.Scan(&Name)
	fmt.Println("Nechta imkoniyat so'raysiz: ")
	fmt.Scan(&Chances)
	fmt.Println("Yoshingizni kiriting: ")
	fmt.Scan(&Age)
	fmt.Println("Qaysi raqamni yaxshi ko'rasiz: ")
	fmt.Scan(&fn)

	player.name = Name
	player.chances = Chances
	player.age = Age
	player.FN = fn

	return player
}

func NewGame(player Player) Game {
	game := Game{}
	if player.chances==1 {
		game.RN=rand.Intn(5)
	} else if player.chances <= 3 {
		game.RN = rand.Intn(10)
	} else if player.chances <= 5 {
		game.RN = rand.Intn(20)
	} else if player.chances <= 10 {
		game.RN = rand.Intn(50)
	} else {
		game.RN = rand.Intn(100)
	}
	if player.age <= 10 || player.age >= 60 {
		game.RN = rand.Intn(10)
	}
	return game
}
func RandomNumber(player Player, game Game) {
	sum := 0
	for j:=0; j<1; j++ {
	for i := 2; i <= game.RN/2; i++ {
		if game.RN%i == 0 {
			sum++
		}
	}
	if sum == 0 {
		fmt.Println("Tasodifiy son tub")
	} 
}
}

func StartGame(p Player, g Game) {
	fmt.Println("Assalomu alaykum ", p.name)
	fmt.Println("Bu tasodifiy sonni topish o'yini ")
	if g.RN == p.FN {
		fmt.Println("Tasodifiy son sizning sevimli soningizga teng")
	}
	for i := 0; i < p.chances; i++ {

		number := 0
		fmt.Println(i+1, " - sonni kiriting ")
		fmt.Scan(&number)
		if number > g.RN {
			fmt.Println("Tasodifiy son siz yozgan sondan kichkina")
		} else if number < g.RN {
			fmt.Println("Tasodifiy son siz yozgan sondan katta")
		}
		if g.RN == number {
			fmt.Println("TABRIKLAYMIZ. SIZ TASODIFIY SONNI TOPDINGIZ")
			return
		}
	}

	fmt.Println("Yutqazdingiz tasodifiy son ", g.RN, " edi ")

}
