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
    hisob:=0
	for m > 0 {
		m /= 10
		sum++
	}
	kxona := sum

	for i := 1; i < sum; i++ {
		xona = xona * 10
	}
	qiymat := 0
	aylanishlarsoni := sum
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
			case 0:
				fmt.Print("")
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
			case 0:
				fmt.Print("")
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
			case 0:
				fmt.Print("")
			}
		}
	   if sum == 12 &&  kxona<12 && qiymat!=0{
			fmt.Print(" milliard ")
		} else if sum == 9 && kxona<9 && qiymat!=0 {
			fmt.Print(" million ")
		} else if sum == 6 && kxona<6 && qiymat!=0{
			fmt.Print(" ming ")
		} else if sum == 11 && kxona<11 && qiymat!=0{
			fmt.Print(" milliard ")
		} else if sum == 8 && kxona<8 &&  qiymat!=0 {
			fmt.Print(" million ")
		} else if sum == 5 && kxona<5 && qiymat!=0{
			fmt.Print(" ming ")
		} else if (sum == 10) && qiymat !=0{
			fmt.Print(" milliard ")
		} else if (sum == 7 ) && qiymat !=0 {
			fmt.Print(" million ")
		} else if (sum == 4) && qiymat!=0 {
			fmt.Print(" ming ")
		}

		kson = kson % xona
		xona = xona / 10
		sum = sum - 1
		if qiymat==0{
          hisob++
		}

	}
	if (hisob==kxona-1) && (kxona!=4) && (kxona!=7) && (kxona!=10){
	if kxona>9{
		fmt.Print(" milliard ")
	} else if kxona>6{
		fmt.Print(" million")
	} else if kxona>3{
		fmt.Print(" ming ")
	}
}
	fmt.Println()
}
