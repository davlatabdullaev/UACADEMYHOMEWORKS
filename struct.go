package main

import "fmt"

type Student struct {
	name        string
	scholarship int
	course      int
}

func main() {
	classwork3()
}

// HOMEWORK1
//
//	Student nomli struktura yarating(uning elementlari: Name, Scholarship, Course). Ushbu strukturani 10ta ma'lumot bilan to'ldiring. Barcha 2-kurs studentlariga
//
// umumiy hisobda qancha scholarship to'langanligini aniqlang va chiqaring.
func struct1() {

	students := []Student{
		{name: "Ruslan", scholarship: 550000, course: 4},
		{name: "Abdulaziz", scholarship: 350000, course: 2},
		{name: "Tolibjon", scholarship: 400000, course: 4},
		{name: "Anvar", scholarship: 550000, course: 2},
		{name: "Gofur", scholarship: 550000, course: 4},
		{name: "Jamshid", scholarship: 650000, course: 2},
		{name: "Davlat", scholarship: 980000, course: 4},
		{name: "Siroj", scholarship: 550000, course: 4},
		{name: "Muhammad", scholarship: 500000, course: 2},
		{name: "Timur", scholarship: 550000, course: 4},
	}
	l := len(students)
	sum := 0
	for i := 0; i < l; i++ {

		if students[i].course == 2 {
			sum = sum + students[i].scholarship
		}
	}
	fmt.Println(sum)
}

// HOMEWORK2
// Student nomli struktura yarating(uning elementlari: Name, Scholarship, Course). Ushbu strukturani 10ta ma'lumot bilan to'ldiring.
// Namening uzunligi 5tadan kam bo'lgan studentlarni chiqaring
func struct2() {
	students := []Student{
		{name: "Ruslan", scholarship: 550000, course: 4},
		{name: "Abdulaziz", scholarship: 350000, course: 2},
		{name: "Tolibjon", scholarship: 400000, course: 4},
		{name: "Anvar", scholarship: 550000, course: 2},
		{name: "Gofur", scholarship: 550000, course: 4},
		{name: "Jamshid", scholarship: 650000, course: 2},
		{name: "Davlat", scholarship: 980000, course: 4},
		{name: "Siroj", scholarship: 550000, course: 4},
		{name: "Muhammad", scholarship: 500000, course: 2},
		{name: "Timur", scholarship: 550000, course: 4},
	}
	l := len(students)
	for i := 0; i < l; i++ {
		if len(students[i].name) <= 5 {
			fmt.Println(students[i])
		}
	}
}

// HOMEWORK3
// AEROPORT nomli struktura yarating(uning elementlari: PlaneType, FlightDate, FromCity, ToCity, FlightTime).
// Ushbu strukturani 10 ta ma'lumot bilan to'ldiring. Yoz oylarida uchadigan
// samolyotlar haqida ma'lumotni chiqaring.
func struct3() {
	type Aeroport struct {
		PlaneType  string
		FlightDate []int
		FromCity   string
		ToCity     string
		FlightTime int
	}
	FlightsList := []Aeroport{
		{PlaneType: "Boeing 747", FlightDate: []int{01, 02, 2024}, FromCity: "Tashkent", ToCity: "Moscow", FlightTime: 4},
		{PlaneType: "AirBus A380", FlightDate: []int{01, 6, 2024}, FromCity: "Tashkent", ToCity: "New York", FlightTime: 18},
		{PlaneType: "AirBus A380", FlightDate: []int{02, 6, 2024}, FromCity: "Tashkent", ToCity: "Moscow", FlightTime: 3},
		{PlaneType: "Boeing 737", FlightDate: []int{20, 2, 2024}, FromCity: "London", ToCity: "Tashkent", FlightTime: 10},
		{PlaneType: "Boeing 747", FlightDate: []int{19, 5, 2024}, FromCity: "Tashkent", ToCity: "Astana", FlightTime: 2},
		{PlaneType: "AirBus A380", FlightDate: []int{21, 7, 2024}, FromCity: "Stanbul", ToCity: "Tashkent", FlightTime: 7},
		{PlaneType: "Boeing 747", FlightDate: []int{27, 6, 2024}, FromCity: "Tashkent", ToCity: "Bishkek", FlightTime: 3},
		{PlaneType: "Boeing 737", FlightDate: []int{19, 12, 2024}, FromCity: "Tashkent", ToCity: "Urgench", FlightTime: 2},
		{PlaneType: "Boeing 747", FlightDate: []int{32, 10, 2024}, FromCity: "Tashkent", ToCity: "Moskva", FlightTime: 5},
		{PlaneType: "AirBus A380", FlightDate: []int{31, 8, 2024}, FromCity: "Tashkent", ToCity: "Dushanbe", FlightTime: 3},
	}
	l := len(FlightsList)
	for i := 0; i < l; i++ {
		if FlightsList[i].FlightDate[1] == 6 || FlightsList[i].FlightDate[1] == 7 || FlightsList[i].FlightDate[1] == 8 {
			fmt.Println(FlightsList[i], " ")
		}
	}

}

// HOMEWORK4
// AEROPORT nomli struktura yarating(uning elementlari: PlaneType, FlightDate, FromCity, ToCity, FlightTime).
// Ushbu strukturani 10 ta ma'lumot bilan to'ldiring. Uchish soati 2 va 3 oralig'idagi va
// qo'nish shaxri 'Toshkent' bo'lgan samolyotlarni chiqaring

func struct4() {
	type Aeroport struct {
		PlaneType  string
		FlightDate []int
		FromCity   string
		ToCity     string
		FlightTime int
	}
	FlightsList := []Aeroport{
		{PlaneType: "Boeing 747", FlightDate: []int{01, 02, 2024}, FromCity: "Tashkent", ToCity: "Moscow", FlightTime: 4},
		{PlaneType: "AirBus A380", FlightDate: []int{01, 6, 2024}, FromCity: "Tashkent", ToCity: "New York", FlightTime: 18},
		{PlaneType: "AirBus A380", FlightDate: []int{02, 6, 2024}, FromCity: "Tashkent", ToCity: "Moscow", FlightTime: 3},
		{PlaneType: "Boeing 737", FlightDate: []int{20, 2, 2024}, FromCity: "London", ToCity: "Tashkent", FlightTime: 10},
		{PlaneType: "Boeing 747", FlightDate: []int{19, 5, 2024}, FromCity: "Tashkent", ToCity: "Astana", FlightTime: 2},
		{PlaneType: "AirBus A380", FlightDate: []int{21, 7, 2024}, FromCity: "Stanbul", ToCity: "Tashkent", FlightTime: 7},
		{PlaneType: "Boeing 747", FlightDate: []int{27, 6, 2024}, FromCity: "Tashkent", ToCity: "Bishkek", FlightTime: 3},
		{PlaneType: "Boeing 737", FlightDate: []int{19, 12, 2024}, FromCity: "Urgench", ToCity: "Tashkent", FlightTime: 2},
		{PlaneType: "Boeing 747", FlightDate: []int{32, 10, 2024}, FromCity: "Tashkent", ToCity: "Moskva", FlightTime: 5},
		{PlaneType: "AirBus A380", FlightDate: []int{31, 8, 2024}, FromCity: "Bishkek", ToCity: "Tashkent", FlightTime: 3},
	}
	l := len(FlightsList)
	for i := 0; i < l; i++ {
		if (FlightsList[i].FlightTime == 2 || FlightsList[i].FlightTime == 3) && FlightsList[i].ToCity == "Tashkent" {
			fmt.Println(FlightsList[i], " ")
		}
	}

}

// HOMEWORK5
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

// CLASSWORK1
// Car nomli strukturani yarating va uning elementlarini Name, Price, Model va HorsePower bo'yicha
// ma'lumotlarni kiriting. Siznig vazifangiz strukturaning hamma ma'lumotlarini
// Modeli bo'yicha alfavit tartibida chiqaring
func classwork1() {

	type Car struct {
		name       string
		price      int
		model      string
		horsepower int
	}

	cars := []Car{
     {name: "Chevrolet", price: 14000, model: "Cobalt", horsepower: 106},
	 {name: "Chevrolet", price: 4000, model: "Matiz", horsepower: 100},
	 {name: "Chevrolet", price: 13000, model: "Gentra", horsepower: 105},
	 {name: "Chevrolet", price: 11000, model: "Spark", horsepower: 103},
	 {name: "Chevrolet", price: 8000, model: "Damas", horsepower: 101},
	 {name: "Chevrolet", price: 11000, model: "Neksiya", horsepower: 105},
	 {name: "Chevrolet", price: 80000, model: "Taho", horsepower: 325},
	 {name: "Chevrolet", price: 14000, model: "Traverse", horsepower: 318},

	}
	l:=len(cars)
	for i:=0; i<l; i++{
		for j:=i+1; j<l; j++{
			if cars[i].model>cars[j].model{
              cars[i].model, cars[j].model=cars[j].model, cars[i].model
			}
		}
	}
	for i:=0; i<l; i++{
		fmt.Println(cars[i])
	}

}
// CLASSWORK2 VA HOMEWORK5 BIR XIL

// CLASSWORK3
// 3. Student nomli struktura e'lon qiling va uning 
// elementlarini (1-Name, 2-Age, 3-Schoalrship va 4-studentning 5 ta fandan bahosi(5talik slice))
// Nta element bilan to'ldiring. O'rtacha bahosi 4 bo'lgan Studentlar
// haqida ma'lumot bering
func classwork3(){
  
	type Student struct{
		name string
		age int
        scholarship int 
		score []int
	}

	students := []Student{
		{name: "Ruslan", age: 25, scholarship: 550000, score: []int{4,5,4,4,4}},
		{name: "Abdulaziz", age: 19, scholarship: 350000, score: []int{5,4,5,4,5}},
		{name: "Tolibjon", age: 21, scholarship: 400000, score: []int{4,4,5,4,5}},
		{name: "Anvar", age: 20, scholarship: 550000, score: []int{5,5,1,1,5}},
		{name: "Gofur", age: 20, scholarship: 550000, score: []int{1,4,5,1,5}},
		{name: "Jamshid",age: 21, scholarship: 650000, score: []int{4,4,4,4,5}},
		{name: "Davlat", age: 21, scholarship: 980000, score: []int{5,5,5,5,4}},
		{name: "Siroj", age: 21, scholarship: 550000, score: []int{5,3,4,5,3}},
		{name: "Muhammad",age: 20, scholarship: 500000, score: []int{4,5,5,4,3}},
		{name: "Timur", age: 22, scholarship: 550000, score: []int{5,4,3,1,1}},
	}
l:=len(students)
for i:=0; i<l; i++{
	sum:=0
	l1:=len(students[i].score)
	for j:=0; j<l1; j++{
		sum=sum+students[i].score[j]
	}
	if sum/5==4{
  fmt.Println(students[i])
	}
}

}