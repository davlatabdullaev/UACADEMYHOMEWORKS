package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("Hello go function")
	//TubSon(17)
	//TeskariRaqam(224234)
	//fmt.Println(RaqamlarYigindisi(34516543))
	//fmt.Println(NgachaSonlarYigindisi(5))
	//ABoraliq(10, 3)
	//func6(13)
	//func7(2, 7, 5)
	//func8(433)
	//fmt.Println(func9(2, 3))
	//func11(12345)
	//func12(28)
	//fmt.Println(func13(6))
	//fmt.Println(func14(1, 2, 4))
	//fmt.Println(func15(1.1, 2))
	// func16(7)
	//	func17(12)
	//fmt.Println(xonasoni(282939))
	//fmt.Println(ikkisonniqoshish(123, 456))
	//faktorial(5)
	//RevNum(67243)
	TrilionGacha(999999999999)
}

// 1. Tub sonni topadigan funksiya yarating
func TubSon(n int) {
	var sum int = 0
	for i := 2; i <= n/2; i++ {
		if n%2 == 0 {
			sum++
		}
	}
	if sum == 0 {
		fmt.Println("Tub son")
	} else if sum > 0 {
		fmt.Println("Tub son emas")
	}

}

// 2. Funksiya orqali n sonining raqamlarini teskari tartibda chiqaring
func TeskariRaqam(n int) {
	var sum int = 0
	var m int = n
	for n != 0 {
		n = n / 10
		sum++
	}
	var k int = 0
	for i := 0; i < sum; i++ {
		k = m % 10
		fmt.Print(k, " ")
		m /= 10
	}
}

// 3. Funksiya orqali n sonining raqamlar yig'indisini chiqaring
func RaqamlarYigindisi(n int) int {
	var sum int = 0
	var m int = n
	for n != 0 {
		n = n / 10
		sum++
	}
	var k int = 0
	var yigindi int = 0
	for i := 0; i < sum; i++ {
		k = m % 10
		yigindi += k
		m /= 10
	}
	return yigindi
}

// 4. Funksiya orqali 1 dan n gacha bo'lan sonlar yigi'indisi chiqarilsin
func NgachaSonlarYigindisi(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// 5. Funksiya orqali A va B oraliqdagi (A va B ning qiymatlarini hisobga olgan holda)
// sonlarni agar A<B bo'lsa o'sish tartibida, aks holda esa kamayish tartibida chiqaruvchi
// funksiya tuzing
func ABoraliq(a, b int) {
	if a < b {
		for i := a; i <= b; i++ {
			fmt.Print(i, " ")
		}
	} else if a > b {
		for i := a; i >= b; i-- {
			fmt.Print(i, " ")
		}
	}

}

// 6. Funksiya orqali 1 dan N gacha (N soni funksiya prametri sifatida ishlatiladi)
// sonlarni juftlarini teskari tartibda chiqaring.
func func6(n int) {
	if n%2 == 0 {
		for i := n; i > 1; i -= 2 {
			fmt.Print(i, " ")
		}
	}
	if n%2 == 1 {
		for i := n - 1; i > 1; i -= 2 {
			fmt.Print(i, " ")
		}
	}
}

// 7. Funksiya parametri sifatida uchta butun son berilgan va
// ushbu sonlarni eng kichik va eng katta qiymatlar orasidagi sonni
// chiqaruvchi funksiya tuzing
func func7(a, b, c int) {
	if (a > b) && (b > c) {
		fmt.Println(b)
	} else if (c > b) && (b > a) {
		fmt.Println(b)
	} else if (c > a) && (a > b) {
		fmt.Println(a)
	} else if (b > a) && (a > c) {
		fmt.Println(a)
	} else if (b > c) && (c > a) {
		fmt.Println(c)
	} else if (a > c) && (c > b) {
		fmt.Println(c)
	}
}

// 8. Funksiya parametri sifatida butun son kiritiladi va ushbu sonning
// raqamlar yig'indisi juft bo'lsa true toq bo'lsa false qaytarilsin
func func8(n int) {
	var sum int = 0
	var m int = n
	for n != 0 {
		n = n / 10
		sum++
	}
	var k int = 0
	var yigindi int = 0
	for i := 0; i < sum; i++ {
		k = m % 10
		yigindi += k
		m /= 10
	}
	if yigindi%2 == 0 {
		fmt.Println("True")
	} else if yigindi%2 == 1 {
		fmt.Println("False")
	}
}

// 9. Funksiya parametri sifatida ikkita butun son berilgan va ushbu
// sonlarning bo'linmasini qaytaradigan funksiya yozing
func func9(x, y int) float64 {
	var z float64 = float64(x) / float64(y)
	return z
}

// 10. Funksiya parametri sifatida N soni kiritiladi va 1 dan Ngacha
// sonlarning yig'indisini qaytaradigan funksiya tuzing
func func10(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// 11. Funksiya orqali N sonining (N soni funksiya parametri sifatida ishlatiladi)
// raqamlarini to'g'ri tartibda chiqaring
func func11(n int) {
	var sum int = 0
	var m int = n
	for n > 0 {
		n /= 10
		sum++
	}
	var nol int = 1
	for i := 1; i < sum; i++ {
		nol = nol * 10
	}
	for i := 0; i < sum; i++ {
		fmt.Print(m/nol, " ")
		m = m % nol
		nol = nol / 10

	}
}

// 12. Funksiya orqali kiritilgan sonning tub bo'luvchilarini chiqaruvchi
// funksiya tuzing
func func12(n int) {
	for j := 2; j <= n; j++ {
		var sum int = 0
		for i := 2; i <= j/2; i++ {
			if j%i == 0 {
				sum++
			}
		}
		if sum == 0 {
			if n%j == 0 {
				fmt.Print(j, " ")
			}
		}
	}
}

// 13. Funksiya parametri sifatida N soni kiritiladi va 1 dan Ngacha sonlarning
// ko'paytmasini qaytaradigan funksiya tuzing
func func13(n int) int {
	var num int = 1
	for i := 1; i <= n; i++ {
		num *= i
	}
	return num
}

// 14. Funksiya parametri sifatida a, b, va c butun sonlar kiritiladi
// va a+b ning qiymati c qiymatiga teng bo'lsa funksiya 1 ni aks holda 0 ni qaytarsin.
func func14(a, b, c int) int {
	if a+b == c {
		return 1
	} else {
		return 0
	}
}

// 15. Funksiya parametri sifatida haqiqiy son va butun son berilgan. Ushbu
// haqiqiy sonni butun son darajasini topadigan funksiya tuzing
func func15(x float64, y int) float64 {
	var z float64 = math.Pow(x, float64(y))
	return z
}

// 16. Funksiya orqali  1 dan N gacha sonlarni
// toqlarini to'g'ri tartibda chiqaring
func func16(n int) {
	for i := 1; i < n; i += 2 {
		fmt.Print(i, " ")
	}
}

// 17. 1 dan N gacha karra jadvalni chiqaruvchi funksiya tuzing
func func17(n int) {
	var a int = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= 10; j++ {
			a = i * j
			fmt.Println(j, " * ", i, " = ", a)
			if j == 10 {
				fmt.Println("_____________")
			}
		}
	}
}

// 22-noyabr kuni berilgan uyga vazifalar
// Kiritilgan sonning necha xonali ekanligini toping
func xonasoni(n float64) int {
	var sum int = int(math.Log10(n)) + 1
	return sum
}

// 2 ta son berilgan ularni bir-biriga qo'shing
// Masalan: a=123; b=456; result=123456
func ikkisonniqoshish(a, b float64) float64 {
	var bxona int = int(math.Log10(b)) + 1
	var sum int = 1
	for i := 1; i <= bxona; i++ {
		sum *= 10
	}
	var result = a*float64(sum) + b
	return result
}

// Factorial
func faktorial(n int) {
	var fakt int = 1
	for i := 1; i <= n; i++ {
		fakt = fakt * i
	}
	fmt.Println(fakt)
}

// Reverse number
func RevNum(n int) {
	var m int = 0
	for n > 0 {
		m = n % 10
		fmt.Print(m)
		n = n / 10
	}
}

// 999 999 999 999 gacha sonlarni yozma tartibda chiqarish
func TrilionGacha(n int) {
	if n >= 1 && n < 1000000000000 {

		if n >= 1000000000 {
			var yuzmilliardlar = (n % 1000000000000) / 100000000000
			switch yuzmilliardlar {
			case 1:
				fmt.Print(" bir yuz ")
			case 2:
				fmt.Print(" ikki yuz ")
			case 3:
				fmt.Print(" uch yuz ")
			case 4:
				fmt.Print(" to'rt yuz  ")
			case 5:
				fmt.Print(" besh yuz")
			case 6:
				fmt.Print(" olti yuz ")
			case 7:
				fmt.Print(" yetti yuz ")
			case 8:
				fmt.Print(" sakkiz yuz")
			case 9:
				fmt.Print(" to'qqiz yuz ")

			}

			var onmilliardlar = (n % 100000000000) / 10000000000
			switch onmilliardlar {
			case 1:
				fmt.Print(" on ")
			case 2:
				fmt.Print(" yigirma ")
			case 3:
				fmt.Print(" o'ttiz ")
			case 4:
				fmt.Print(" qirq ")
			case 5:
				fmt.Print(" ellik ")
			case 6:
				fmt.Print(" oltmish ")
			case 7:
				fmt.Print(" yetmish ")
			case 8:
				fmt.Print(" sakson ")
			case 9:
				fmt.Print(" to'qson ")

			}

			var birmilliardlar = (n % 10000000000) / 1000000000
			switch birmilliardlar {
			case 1:
				fmt.Print(" bir ")
			case 2:
				fmt.Print(" ikki ")
			case 3:
				fmt.Print(" uch ")
			case 4:
				fmt.Print(" to'rt ")
			case 5:
				fmt.Print(" besh ")
			case 6:
				fmt.Print(" olti ")
			case 7:
				fmt.Print(" yetti ")
			case 8:
				fmt.Print(" sakkiz ")
			case 9:
				fmt.Print(" to'qqiz ")

			}
			var milliard string = " milliard "
			fmt.Print(milliard)
		}

		if n >= 1000000 {
			var yuzmillionlar = (n % 1000000000) / 100000000
			switch yuzmillionlar {
			case 1:
				fmt.Print(" bir yuz ")
			case 2:
				fmt.Print(" ikki yuz ")
			case 3:
				fmt.Print(" uch yuz ")
			case 4:
				fmt.Print(" to'rt yuz  ")
			case 5:
				fmt.Print(" besh yuz")
			case 6:
				fmt.Print(" olti yuz ")
			case 7:
				fmt.Print(" yetti yuz ")
			case 8:
				fmt.Print(" sakkiz yuz")
			case 9:
				fmt.Print(" to'qqiz yuz ")

			}

			var onmillionlar = (n % 100000000) / 10000000
			switch onmillionlar {
			case 1:
				fmt.Print(" on ")
			case 2:
				fmt.Print(" yigirma ")
			case 3:
				fmt.Print(" o'ttiz ")
			case 4:
				fmt.Print(" qirq ")
			case 5:
				fmt.Print(" ellik ")
			case 6:
				fmt.Print(" oltmish ")
			case 7:
				fmt.Print(" yetmish ")
			case 8:
				fmt.Print(" sakson ")
			case 9:
				fmt.Print(" to'qson ")

			}

			var birmillionlar = (n % 10000000) / 1000000
			switch birmillionlar {
			case 1:
				fmt.Print(" bir ")
			case 2:
				fmt.Print(" ikki ")
			case 3:
				fmt.Print(" uch ")
			case 4:
				fmt.Print(" to'rt ")
			case 5:
				fmt.Print(" besh ")
			case 6:
				fmt.Print(" olti ")
			case 7:
				fmt.Print(" yetti ")
			case 8:
				fmt.Print(" sakkiz ")
			case 9:
				fmt.Print(" to'qqiz ")

			}
			var million string = " million "
			fmt.Print(million)
		}
		if n >= 1000 {
			var yuzminglar = (n % 1000000) / 100000
			switch yuzminglar {
			case 1:
				fmt.Print(" bir yuz ")
			case 2:
				fmt.Print(" ikki yuz ")
			case 3:
				fmt.Print(" uch yuz ")
			case 4:
				fmt.Print(" to'rt yuz ")
			case 5:
				fmt.Print(" besh yuz ")
			case 6:
				fmt.Print(" olti yuz ")
			case 7:
				fmt.Print(" yetti yuz ")
			case 8:
				fmt.Print(" sakkiz yuz ")
			case 9:
				fmt.Print(" to'qqiz yuz ")

			}
			var onminglar = (n % 100000) / 10000
			switch onminglar {
			case 1:
				fmt.Print(" on  ")
			case 2:
				fmt.Print(" yigirma ")
			case 3:
				fmt.Print(" o'ttiz ")
			case 4:
				fmt.Print(" qirq ")
			case 5:
				fmt.Print(" ellik ")
			case 6:
				fmt.Print(" oltmish ")
			case 7:
				fmt.Print(" yetmish ")
			case 8:
				fmt.Print(" sakson ")
			case 9:
				fmt.Print(" to'qson ")

			}

			var minglar = (n % 10000) / 1000
			switch minglar {
			case 1:
				fmt.Print(" bir ")
			case 2:
				fmt.Print(" ikki ")
			case 3:
				fmt.Print(" uch ")
			case 4:
				fmt.Print(" to'rt ")
			case 5:
				fmt.Print(" besh ")
			case 6:
				fmt.Print(" olti ")
			case 7:
				fmt.Print(" yetti ")
			case 8:
				fmt.Print(" sakkiz ")
			case 9:
				fmt.Print(" to'qqiz ")

			}
			var ming string = " ming "
			fmt.Print(ming)
		}

		var yuzlar = (n % 1000) / 100
		switch yuzlar {
		case 1:
			fmt.Print(" bir yuz ")
		case 2:
			fmt.Print(" ikki yuz ")
		case 3:
			fmt.Print(" uch yuz ")
		case 4:
			fmt.Print(" to'rt yuz ")
		case 5:
			fmt.Print(" besh yuz ")
		case 6:
			fmt.Print(" olti yuz ")
		case 7:
			fmt.Print(" yetti yuz ")
		case 8:
			fmt.Print(" sakkiz yuz ")
		case 9:
			fmt.Print(" to'qqiz yuz ")
		}
		var onlar = (n % 100) / 10
		switch onlar {
		case 1:
			fmt.Print(" on ")
		case 2:
			fmt.Print(" yigirma ")
		case 3:
			fmt.Print(" o'ttiz ")
		case 4:
			fmt.Print(" qirq ")
		case 5:
			fmt.Print(" ellik ")
		case 6:
			fmt.Print(" oltmish ")
		case 7:
			fmt.Print(" yetmish ")
		case 8:
			fmt.Print(" sakson ")
		case 9:
			fmt.Print(" to'qson ")
		}
		var birlar int = n % 10
		switch birlar {
		case 1:
			fmt.Print(" bir ")
		case 2:
			fmt.Print(" ikki ")
		case 3:
			fmt.Print(" uch ")
		case 4:
			fmt.Print(" to'rt ")
		case 5:
			fmt.Print(" besh ")
		case 6:
			fmt.Print(" olti ")
		case 7:
			fmt.Print(" yetti ")
		case 8:
			fmt.Print(" sakkiz ")
		case 9:
			fmt.Print(" to'qqiz ")
		}

	} else {
		fmt.Println("Error")
	}

}
