//   HOME WORKS

package main

import "fmt"

func main() {

	// IF
	// work 1

	/*	for true {
			fmt.Println("A sonni kiriting ")
			var a float32
			fmt.Scan(&a)
			fmt.Println("B sonni kiriting ")
			var b float32
			fmt.Scan(&b)
			if a > b {
				a = a + b
				b = a - b
				a = a - b
				fmt.Println(" a = ", a, "\n b = ", b)
			} else {
				fmt.Println(" a = ", a, "\n b = ", b)
			}
		}
	*/

	// 2-masala
	/*
		for true {
			fmt.Println("A sonni kiriting ")
			var a float32
			fmt.Scan(&a)
			fmt.Println("B sonni kiriting ")
			var b float32
			fmt.Scan(&b)
			if a > b || a < b {
				fmt.Println("Yig'indi  ", (a + b))
			} else if a == b {
				fmt.Println("0")
			}
		}
	*/

	// 3-masala
	/*
		for true {
			fmt.Println("A sonni kiriting ")
			var a float32
			fmt.Scan(&a)
			fmt.Println("B sonni kiriting ")
			var b float32
			fmt.Scan(&b)
			if a > b {
				fmt.Println(a)
			} else if a == b {
				fmt.Println("0")
			} else {
				fmt.Println(b)
			}
		}
	*/

	// 4-masala
	/*
		for true {
			fmt.Println("A sonni kiriting ")
			var a int
			fmt.Scan(&a)
			fmt.Println("B sonni kiriting ")
			var b int
			fmt.Scan(&b)
			fmt.Println("C sonni kiriting ")
			var c int
			fmt.Scan(&c)
			if (b > c) && (a > c) {
				fmt.Println("Eng kichik son ", c)
			} else if (c > b) && (a > b) {
				fmt.Println("Eng kichik son ", b)
			} else if (b > a) && (c > a) {
				fmt.Println("Eng kichik son ", a)
			}
		}
	*/
	// 5-masala
	/*
		for true {
			fmt.Println("A sonni kiriting ")
			var a int
			fmt.Scan(&a)
			fmt.Println("B sonni kiriting ")
			var b int
			fmt.Scan(&b)
			fmt.Println("C sonni kiriting ")
			var c int
			fmt.Scan(&c)
			if (b > c) && (a < c) {
				fmt.Println("O'rtacha son ", c)
			} else if (a > c) && (b < c) {
				fmt.Println("O'rtacha son ", c)
			} else if (c > b) && (a < b) {
				fmt.Println("O'rtacha son ", b)
			} else if (a > b) && (c < b) {
				fmt.Println("O'rtacha son ", b)
			} else if (b > a) && (c < a) {
				fmt.Println("O'rtacha son ", a)
			} else if (c > a) && (b < a) {
				fmt.Println("O'rtacha son ", a)
			}
		} */

	// 6-masala
	/*
		for true {
			fmt.Println("A sonni kiriting ")
			var a int
			fmt.Scan(&a)
			fmt.Println("B sonni kiriting ")
			var b int
			fmt.Scan(&b)
			fmt.Println("C sonni kiriting ")
			var c int
			fmt.Scan(&c)
			if (b > c) && (a > c) {
				fmt.Println("Eng kichik son ", c)
			} else if (c > b) && (a > b) {
				fmt.Println("Eng kichik son ", b)
			} else if (b > a) && (c > a) {
				fmt.Println("Eng kichik son ", a)
			}
			if (b < c) && (a < c) {
				fmt.Println("Eng katta son ", c)
			} else if (c < b) && (a < b) {
				fmt.Println("Eng katta son ", b)
			} else if (b < a) && (c < a) {
				fmt.Println("Eng katta son ", a)
			}
		}
	*/

	// 7-masala
	/*
		for true {
			fmt.Println("A sonni kiriting ")
			var a int
			fmt.Scan(&a)
			fmt.Println("B sonni kiriting ")
			var b int
			fmt.Scan(&b)
			fmt.Println("C sonni kiriting ")
			var c int
			fmt.Scan(&c)
			if (b > c) && (a < c) {
				fmt.Println(b, " va ", c)
			} else if (a > c) && (b < c) {
				fmt.Println(a, " va ", c)
			} else if (c > b) && (a < b) {
				fmt.Println(c, " va ", b)
			} else if (a > b) && (c < b) {
				fmt.Println(a, " va ", b)
			} else if (b > a) && (c < a) {
				fmt.Println(b, " va ", a)
			} else if (c > a) && (b < a) {
				fmt.Println(c, " va ", a)
			}
		}
	*/

	// 8-masala
	/*
		for true {
			fmt.Println(" A = ")
			var a int
			fmt.Scan(&a)
			fmt.Println(" B = ")
			var b int
			fmt.Scan(&b)
			fmt.Println(" C = ")
			var c int
			fmt.Scan(&c)
			if (a < b) && (b < c) {
				fmt.Println(2*a, 2*b, 2*c)
			} else {
				fmt.Println(-a, -b, -c)
			}
		}
	*/

	// 9-masala
	/*
		for true {
			fmt.Println(" A = ")
			var a int
			fmt.Scan(&a)
			fmt.Println(" B = ")
			var b int
			fmt.Scan(&b)
			fmt.Println(" C = ")
			var c int
			fmt.Scan(&c)
			if (a < b) && (b < c) {
				fmt.Println(2*a, 2*b, 2*c)
			} else if (a > b) && (b > c) {
				fmt.Println(2*c, 2*b, 2*a)
			} else {
				fmt.Println(-a, -b, -c)
			}
		}
	*/

	// 10-masala
	/*
		for true {
			fmt.Println("Birinchi sonni kiriting ")
			var a int
			fmt.Scan(&a)
			fmt.Println("Ikkinchi sonni kiriting ")
			var b int
			fmt.Scan(&b)
			fmt.Println("Uchinchi sonni kiriting ")
			var c int
			fmt.Scan(&c)
			if a == b {
				fmt.Println("Uchinchi")
			} else if a == c {
				fmt.Println("Ikkinchi")
			} else if b == c {
				fmt.Println("Birinchi")
			} else {
				fmt.Println("Masala sharti qanoatlantirilmadi")
			}
		}
	*/

	// 11-masala
	/*
		for true {
			fmt.Println("Birinchi sonni kiriting ")
			var a int
			fmt.Scan(&a)
			fmt.Println("Ikkinchi sonni kiriting ")
			var b int
			fmt.Scan(&b)
			fmt.Println("Uchinchi sonni kiriting ")
			var c int
			fmt.Scan(&c)
			fmt.Println("To'rtinchi sonni kiriting ")
			var d int
			fmt.Scan(&d)
			if (a == b) && (b == d) {
				fmt.Println("Uchinchi")
			} else if (a == c) && (c == d) {
				fmt.Println("Ikkinchi")
			} else if (b == c) && (c == d) {
				fmt.Println("Birinchi")
			} else if (b == c) && (c == a) {
				fmt.Println("To'rtinchi")
			} else {
				fmt.Println("Masala sharti qanoatlantirilmadi")
			}
		}
	*/

	// FOR
	// 1-masala
	/*
		for true {
			fmt.Print("\n n  = ")
			var n int
			fmt.Scan(&n)
			for i := n; i > 1; i-- {
				if i%2 == 0 {
					fmt.Print(" ", i, " ")
				}
			}

		}
	*/

	// 2-masala
	/*
		for true {
			fmt.Print("\n n  = ")
			var n int
			fmt.Scan(&n)
			for i := n; i > 1; i-- {
				if i%2 == 1 {
					fmt.Print(" ", i, " ")
				}
			}
	*/
	// 3-masala
	/*
		for i := 10; i < 100; i++ {
			var sum int
			for j := 2; j < i/2; j++ {
				if i%j == 0 {
					sum++
				}
			}
			if sum == 0 {
				fmt.Print(i, " ")
			}
		}
	*/

	// 4-masala
	/*
		for i := 100; i < 1000; i++ {
			var birlik int = i % 10
			var onlik int = (i % 100) / 10
			var yuzlik int = i / 100
			if (birlik == yuzlik) || (yuzlik == onlik) || (onlik == birlik) {
				fmt.Println(i, " ")
			}
		}
	*/

	// 5-masala
	/*
		for true {
			fmt.Print("\n n = ")
			var n int
			fmt.Scan(&n)
			var m int = n
			var sum int = 0
			for n > 0 {
				n = n / 10
				sum++
			}
			fmt.Println(sum, "xonali son kiritildi")
			var nol int = 1
			for i := 1; i < sum; i++ {
				nol = nol * 10
			}
			fmt.Println(nol)
			var k int
			for i := 0; i < sum; i++ {
				k = (m / nol)
				fmt.Print(k, " ")
				m = m % nol
				nol = nol / 10
			}

		} */

	// 6-masala

	// SWITCH

	// 1-masala
	/*
		for true {
			fmt.Println("\n Robot boshlang'ich qaysi tomonga qarab turibdi \n s - shimol \n j - janub \n q - sharq \n g - garb")
			var holat string
			fmt.Scan(&holat)
			fmt.Println("Robotga komanda bering \n 0 - harakatni davom ettir \n 1 - chapga buril \n 2 - o'nga buril ")
			var kom int
			fmt.Scan(&kom)
			if holat == "s" {
				switch kom {
				case 0:
					fmt.Println("Komandadan keyingi holati: Shimol")
				case 1:
					fmt.Println("Komandadan keyingi holati: G'arb")
				case 2:
					fmt.Println("Komandadan keyingi holati: Sharq")
				}

			} else if holat == "j" {
				switch kom {
				case 0:
					fmt.Println("Komandadan keyingi holati: Janub")
				case 1:
					fmt.Println("Komandadan keyingi holati: Sharq")
				case 2:
					fmt.Println("Komandadan keyingi holati: G'arb")
				}

			} else if holat == "q" {
				switch kom {
				case 0:
					fmt.Println("Komandadan keyingi holati: Sharq")
				case 1:
					fmt.Println("Komandadan keyingi holati: Shimol")
				case 2:
					fmt.Println("Komandadan keyingi holati: Janub")
				}

			} else if holat == "g" {
				switch kom {
				case 0:
					fmt.Println("Komandadan keyingi holati: G'arb")
				case 1:
					fmt.Println("Komandadan keyingi holati: Janub")
				case 2:
					fmt.Println("Komandadan keyingi holati: Shimol")
				}

			}

		}
	*/

	// 2-masala

	for true {
		fmt.Println("\n Lokatr boshlang'ich qaysi tomonga qaratilgan \n s - shimol \n j - janub \n q - sharq \n g - garb")
		var holat string
		fmt.Scan(&holat)
		fmt.Println("Lokatrga birinchi komandani bering  \n 0 - o'nga buril \n 1 - chapga buril \n 2 - burilish 180 ")
		var kom1 int
		fmt.Scan(&kom1)
		var holatA string
		fmt.Println("Lokatrga ikkinchi komandani bering  \n 0 - o'nga buril \n 1 - chapga buril \n 2 - burilish 180 ")
		var kom2 int
		fmt.Scan(&kom2)

		if holat == "s" {
			switch kom1 {
			case 0:
				holatA = "q"
			case 1:
				holatA = "g"
			case 2:
				holatA = "j"
			}

		} else if holat == "j" {
			switch kom1 {
			case 0:
				holatA = "g"
			case 1:
				holatA = "q"
			case 2:
				holatA = "s"
			}

		} else if holat == "q" {
			switch kom1 {
			case 0:
				holatA = "j"
			case 1:
				holatA = "s"
			case 2:
				holatA = "g"
			}
		} else if holat == "g" {
			switch kom1 {
			case 0:
				holatA = "s"
			case 1:
				holatA = "j"
			case 2:
				holatA = "q"
			}

		}

		if holatA == "s" {
			switch kom2 {
			case 0:
				fmt.Println("Komandadan keyingi holati: Sharq")
			case 1:
				fmt.Println("Komandadan keyingi holati: G'arb")
			case 2:
				fmt.Println("Komandadan keyingi holati: Janub")
			}

		} else if holatA == "j" {
			switch kom2 {
			case 0:
				fmt.Println("Komandadan keyingi holati: G'arb")
			case 1:
				fmt.Println("Komandadan keyingi holati: Sharq")
			case 2:
				fmt.Println("Komandadan keyingi holati: Shimol")
			}

		} else if holatA == "q" {
			switch kom2 {
			case 0:
				fmt.Println("Komandadan keyingi holati: Janub")
			case 1:
				fmt.Println("Komandadan keyingi holati: Shimol")
			case 2:
				fmt.Println("Komandadan keyingi holati: G'arb")
			}

		} else if holatA == "g" {
			switch kom2 {
			case 0:
				fmt.Println("Komandadan keyingi holati: Shimol")
			case 1:
				fmt.Println("Komandadan keyingi holati: Janub")
			case 2:
				fmt.Println("Komandadan keyingi holati: Sharq")
			}

		}

	}
}
