package main

import "fmt"

func main() {

	struct5()

}

// Team nomli strukturani e'lon qiling va uning elementlarini Name, Coach, PlayersCount soni Captain bo'yicha
// ma'lumotlarni kiriting. Sizning vazifangiz name nomli string seriyasini
// kiritamiz va ushbu seriyali Teamlar haqida ma'lumotni chiqaring. Barcha Teamlarni ishtirokchilar soni
// bo'yicha kamayish tartibida chiqaring
func struct5() {

	type Team struct {
		name         string
		coach        string
		PlayersCount int
		Captain      string
	}
	Teams := []Team{
		{name: "Villareal", coach: "Enuai Emeri", PlayersCount: 23, Captain: "Mario Gaspar"},
		{name: "RealMadrid", coach: "Karlo Anchelotti", PlayersCount: 21, Captain: "Luca Modric"},
		{name: "Barcelona", coach: "Xavi Ernandes", PlayersCount: 27, Captain: "Sergio Buskets"},
		{name: "AtletikoMadrid", coach: "Diego Simyone", PlayersCount: 31, Captain: "Gabi"},
		{name: "Jirona", coach: "Xuan Karlos", PlayersCount: 28, Captain: "Ivan Balliu"},
		{name: "Hetafe", coach: "Hose Bordalas", PlayersCount: 29, Captain: "Jene Dakonam"},
		{name: "Valensiya", coach: "Xavi Grasiya", PlayersCount: 17, Captain: "Pontus Jansson"},
		{name: "RealBetis", coach: "Ektor Kuper", PlayersCount: 15, Captain: "Nabil Fekir"},
		{name: "Granada", coach: "Xoes Luis Oltra", PlayersCount: 18, Captain: "Toni Adams"},
		{name: "Osasuna", coach: "Yagoba Arrasate", PlayersCount: 23, Captain: "Oier Sanjurjo"},
	}
	b := false
	l := len(Teams)
	NewTeam := []Team{}
	var a string

	for i := 0; i < l; i++ {
		fmt.Print("Enter a club name: ")
		fmt.Scan(&a)
   for j:=0; j<l; j++ {
		if a == Teams[j].name {
			b = true
			NewTeam = append(NewTeam, Teams[j])
		}
	}
}

if b {
	l1:=len(NewTeam)
	for i := 0; i < l1; i++ {
		for j := i + 1; j < l1; j++ {
			if NewTeam[i].PlayersCount < NewTeam[j].PlayersCount {
				NewTeam[i], NewTeam[j] = NewTeam[j], NewTeam[i]
			}
		}
	}
}
	if b {
		l1:=len(NewTeam)
		for i := 0; i < l1; i++ {
			fmt.Printf("Name: %s, Coach: %s, PlayersCount: %d, Captain: %s\n",
				NewTeam[i].name, NewTeam[i].coach, NewTeam[i].PlayersCount, NewTeam[i].Captain)
		}
	}
	

}
