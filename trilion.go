package main

import "fmt"

func main() {
	fmt.Print("son kiriting ")
	n := 0
	fmt.Scan(&n)

	Triliongacha(n)

}
func Triliongacha(m int) {
	sum := 0
	kson := m
	xona := 1

	for m > 0 {
		m /= 10
		sum++
	}
	for i := 1; i < sum; i++ {
		xona = xona * 10
	}
	qiymat:=0
	aylanishlarsoni:=sum
	for i := 0; i < aylanishlarsoni; i++ {
		qiymat = kson / xona
		if sum%3 == 0 {
			switch qiymat {
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
				fmt.Print(" to'qqiz yuz")
			}
		} else if sum%3 == 2 {
			switch qiymat {
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
		} else if sum%3 == 1 {
			switch qiymat {
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
		}
        if sum==10{
			fmt.Print(" milliard ")
		} else if sum==7{
			fmt.Print(" million ")
		}	else if sum==4{
			fmt.Print(" ming ")
		}
		kson=kson%xona
		xona=xona/10
        sum=sum-1
		
	}
	fmt.Println()
}
