package main 

import (
	"fmt"
	"math/rand"
	
)

func main() {
	//func1()
	//func2()
	//func3()
	//func7()
	//func8()
	//func9()
	//uhomework3()
	//fmt.Println(ikkisonniqoshish(2,999))
	//fmt.Println(faktorial(5))
	//uhomework6()
	uhomework11()

}
// 1. Array va butun son berilgan. Ushbu arrayning ikkinchi manfiy va
// to'rtinchi manfiy elementlari orasidagi sonlar yig'indisini qaytaradigan dastur tuzing
func func1() {
	var len int = 10
	var arr = []int{1, -2, 3, -4, 5, -6, 7, -8, 9, -10}
	var (
		sum     int = 0
		counter     = 0
		g       int = 0
		k           = 0
	)

	for i := 0; i < len; i++ {
		if arr[i] < 0 {
			counter++

			if counter == 2 {
				k = i
			}

			if counter == 4 {
				g = i
			}
		}
	}

	//	fmt.Print(k, g)
	for j := k + 1; j < g; j++ {
		sum = sum + arr[j]
		//fmt.Println(arr[j])
	}

	fmt.Println("Yig'indi ", sum)
}
// Kiruvchi ma'lumot sifatida int array va uning uzunligi beriladi.
// Sizning vazifangiz ushbu arrayda nechta musbat, manfiy, nol, juft va toq sonlar borligini 
// aniqlaydigan funksiya yasang
func func2() {
	var (
		musbat = 0
		manfiy = 0
		nol    = 0
		juft   = 0
		toq    = 0
	)
	var arr = []int{0, -1, 2, -3, 4, 0, 5, -6, 7, 0}
	//	fmt.Println("array 2 is here", arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
		if arr[i] > 0 {
			musbat++
		} else if arr[i] < 0 {
			manfiy++
		} else if arr[i] == 0 {
			nol++
		}
		if arr[i]%2 == 0 {
			juft++
		}
		if (arr[i]%2 == 1) || (arr[i]%2 == -1) {
			toq++
		}
	}
	fmt.Println("Musbat sonlar ", musbat)
	fmt.Println("Manfiy sonlar ", manfiy)
	fmt.Println("Nolga teng sonlar ", nol)
	fmt.Println("Juft sonlar ", juft)
	fmt.Println("Toq sonlar ", toq)
}
// 3. Funksiya orqali sliceni juft va toqlarga ajrating
func func3() {
	var len int = 10
	var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Juftlar")
	for i := 0; i < len; i++ {
		if slice1[i]%2 == 0 {
			fmt.Print(slice1[i], " ")
		}

	}
	fmt.Println()
	fmt.Println("Toqlar")
	for j := 0; j < len; j++ {
		if slice1[j]%2 == 1 {
			fmt.Print(slice1[j], " ")
		}

	}
}
// 4. 10ta butun sondan iborat int slice berilgan va 
// juft indeksdagi elementlar ko'paytmasini va toq elementlar yig'indisini 
// topadigan dastur tuzing
func func4() {

	var len int = 10
	var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var juft int = 1
	for i := 0; i < len; i += 2 {
		juft = juft * slice1[i]
	}
	var toq int = 0
	for j := 0; j < len; j += 2 {
		toq = toq + slice1[j]
	}
	fmt.Println(juft)
	fmt.Println(toq)
}
// 5. 10 talik uzunlikdagi int sliceni elementlari
// orasidagi ikkinchi eng kichik qiymatni aniqlovchi dastur tuzing
func func5() {
	var len int = 10
	var a = []int{1, 4, 3, 2, 5, 7, 6, 8, 9, 10}
	min := 0
	min = a[9]
	for i := 0; i < len; i++ {
         if(a[i]<min){
			min = a[i+1]
		 }
		 
		}
		fmt.Println(min)
		
	}
	// 6. 10 talik uzunlikdagi int sliceni elementlari
// orasidagi ikkinchi eng katta qiymatni aniqlovchi dastur tuzing

	func function6() {
		var len int = 10
		var a = []int{1, 4, 3, 2, 5, 7, 6, 8, 9, 10}
		max := 0
		max = a[5]
		for i := 0; i < len; i++ {
			 if(a[i]>max){
				max = a[i-1]
			 }
			 
			}
			fmt.Println(max)
			
		}
// 7. n ta butun sondan iborat int slice berilgan va undagi 1-chi 
// va oxirgi elementlarini almashtiring
	func function7(){
		var n int =5
		var a = []int{1,2,3,4,5}
		b:= a[4]
		a[4]=a[0]
		a[0]=b
		for i:=0; i<n; i++{
			fmt.Print(a[i], " ")
		}
	}	
//8. n ta butun sondan iborat slice berilgan va maxdan(maxni o'zi inobatga olinmasin)
// keyin nechta son borligini aniqlaydigan dastur tuzing 	
func function8(){		
	var a = []int{2,9,3,-4,5,11, 5,4}
	var n int = len(a)
	max:=0
	m:=0
	son:=0
	max=a[0]
	for i:=0; i<n;i++{
		if a[i]>max{
			max=a[i]
             m=i
		}
	}
	for j:=m+1; j<n; j++{
		son++
	}
	fmt.Println(son, " ta element bor ")
}
//9. n ta butun sondan iborat slice berilgan va mindan(minni o'zi inobatga olinmasin)
// oldin nechta son borligini aniqlaydigan dastur tuzing 	

func function9(){		
	var a = []int{2,9,3,-4,5,-9}
	var n int = len(a)
	min:=0
	m:=0
	son:=0
	min=a[0]
	for i:=0; i<n;i++{
		if a[i]<min{
			min=a[i]
             m=i
		}
	}
	for j:=0; j<m; j++{
		son++
	}
	fmt.Println(son, " ta element bor ")
}
//10. nta butun sondan iborat int slice berilgan va ulardagi max va min qiymatlar(max va min inobatga olinmasin)
// orasida nechta son borligini aniqlaydigan dastur yarating
func function10() {
	var a = []int{2,-9,3,-4,5,5,6,3,4,6,11}
	var len = len(a)
	max, min := 0, 0
	m, n := 0, 0
	max = a[0]
	min = a[0]
	for i:=0; i<len;i++{
		if a[i]<min{
			min=a[i]
             m=i
		}
	}
	for j:=0; j<len;j++{
		if a[j]>max{
			max=a[j]
             n=j
		}
	}
	if n<m{
		hisob:=0
	 for i:=n+1; i<m; i++{
       hisob++
	 }
	 fmt.Println(hisob, " ta son bor ")
	}
	if n>m{
		hisob:=0
	 for i:=m+1; i<n; i++{
       hisob++
	 }
	 fmt.Println(hisob, " ta son bor ")
	}
	

}


//    HOMEWORKS
// 1. n ta butun sondan iborat int slice berilgan  va undagi max va min 
// elementlarini almashtiring
func uhomework1() {
var a =[]int{2,9,3,-4,5,-11}
var len int =len(a)
min, max := 0, 0
min = a[0]
max = a[0]
maxindex, minindex := 0, 0
for i:=0; i<len; i++{
	if a[i]>max{
		max = a[i]
		maxindex=i
	}
	if a[i]<min{
		min= a[i]
		minindex=i
	}
}
c:=0
c=a[maxindex]
a[maxindex]=a[minindex]
a[minindex]=c 
 for j:=0; j<len; j++{
	fmt.Print(a[j]," ")
 }
 fmt.Println()
}
// 2. Foydalanuvchi tomonidan int tipidagi 10 ta butun sonlar massivini
// kiritiladi. Shu massiv ichida tub sonlar sonini aniqlovchi dastur tuzing.
// Dastur natijada  massiv ichidagi tub sonlar ro'yhati va ularning 
// sonini chop etadi. 
func uhomework2() {
var massiv = [10]int{}	

var tubson int =0
for i:=0; i<10; i++{
fmt.Print(i, " - indeksdagi elementni kiriting ")
fmt.Scan(&massiv[i])
}
for i:=0; i<10; i++{
	var tubhisob int =0
	for j:=1; j<=(massiv[i]); j++{
       if massiv[i]%j==0{
		tubhisob=tubhisob+1
	   } 
	}
	
	if tubhisob==2{
fmt.Print(massiv[i], " ")
		tubson=tubson+1
	   }
}

	
fmt.Println("\n Ushbu massivda ", tubson, " ta tub son mavjud.")

}
//3. Foydalanuvchi tomonidan int tipida istalgan miqdorda butun sonlar massivini 
// kiritiladi. Agar indexlar yig'indisi massiv qiymatlari yig'indisidan katta bo'lsa,
// INDEX yozuvi chop etilsin. Agar qiymatlar yig'indisi indexlar yig'indisidan katta bo'lsa, VALUES, teng bo'lsa WOW degan yozuv 
// chop etilsin
func uhomework3() {
	var massivhisobi int = 0
	var indexhisobi int = 0

	
	fmt.Print("Massivga nechta element kiritmoqchisiz ")
	var son int = 0 
	fmt.Scan(&son)
	var massiv = make([]int, son)	
	for i:=0; i<son; i++ {
	fmt.Println(i, " - indeksdagi elementni kiriting ")
    fmt.Scan(&massiv[i])	
}
	for i:=0; i<son; i++{
     massivhisobi=massivhisobi+massiv[i]
	 indexhisobi=indexhisobi+i
	}
	if indexhisobi>massivhisobi{
		fmt.Println("INDEX")
	} else if indexhisobi<massivhisobi{
		fmt.Println("VALUES")
	} else if indexhisobi==massivhisobi{
		fmt.Println("WOW")
	}
}
func uhomework3version2() {
	
	fmt.Print("Massivga nechta element kiritmoqchisiz ")
	var son int = 0 
	fmt.Scan(&son)
	var massiv = []int{}	
	for i:=0; i<son; i++ {
	fmt.Println(i, " - indeksdagi elementni kiriting ")
    element := 0
	fmt.Scan(&element)
	massiv = append(massiv, element)	
}
	var massivhisobi int = 0
	var indexhisobi int = 0

	for i:=0; i<son; i++{
     massivhisobi=massivhisobi+massiv[i]
	 indexhisobi=indexhisobi+i
	}
	if indexhisobi>massivhisobi{
		fmt.Println("INDEX")
	} else if indexhisobi<massivhisobi{
		fmt.Println("VALUES")
	} else if indexhisobi==massivhisobi{
		fmt.Println("WOW")
	}
}

//4. -12 dan 26 gacha random sonlar bilan n ta elementdan iborat 
// sliceni to'ldiring va ushbu seriyadagi manfiy qiymatli sonlarni nolga 
// va musbat sonlarni 1 ga almashtiring va natijani chiqaring
func uhomework4() {
	min:=-12
	max:=26
	fmt.Print("Massiv nechta elementdan iborat bo'lsin ")
	eson:=0
	fmt.Scan(&eson)
	slice := []int{}
	for i:=0; i<eson; i++ {
		num := rand.Intn(max-min)+min
		slice = append(slice, num)
		fmt.Print(num, " ")
	}
	fmt.Println()
	for i:=0; i<eson; i++{
        if slice[i]>0{
			fmt.Print(" 1 ")
		} else {
			fmt.Print(" 0 ")
		}
	}
}
// 5. 14 dan 35 gacha random sonlar bilan n ta elementdan iborat 
// int sliceni to'ldiring va ushbu slicedagi faqat juft qiymatli sonlarni 
// + belgisini chiqaring
func uhomework5() {
	min :=14
    max :=35
	fmt.Print("Slice nechta elementdan iborat bo'lsin ")
	esoni:=0
	slice := []int{}
	fmt.Scan(&esoni)
	 var i int = 0;
	 for i=0; i<esoni; i++{
	  num := rand.Intn(max-min)+min
	  slice = append(slice, num)
	  fmt.Print(num, " ")
	 }
	 fmt.Println()
     for i:=0; i<esoni; i++{
		if slice[i]%2==0{
			fmt.Print(" + ")
		} else {
			fmt.Print(" ",slice[i]," ")
		} 
	 }
	 fmt.Println()
}
//6. -46 dan 5 gacha random sonlar biln n ta elementdan iborat int 
// sliceni to'ldiring va ushbu slicedagi faqat toq qiymatli sonlarni
// tag chiziq (_) chiqaring 
func uhomework6() {
	min := -46
	max := 5
	fmt.Print("Slice nechta elementdan iborat bo'lsin ")
	num:=0
	fmt.Scan(&num)
	slice := []int{}
    for i:=0; i<num; i++{
		number:=rand.Intn(max-min)+min
		slice = append(slice, number)
		fmt.Print(number, " ")
	}
	fmt.Println()
	for i:=0; i<num; i++ {
		if slice[i]%2==1 || slice[i]%2==-1{
			fmt.Print(" _ ")
		} else{
			fmt.Print(" ", slice[i], " ")
		}
	} 
	fmt.Println()
} 
/* 7 - 8 - masala bir xil
n ta sondan iborat int sliceni -5 dan 25 gacha random sonlar bilan 
to'ldiring va ushbu slice elementlarini chapga 1 ta elementga suring 
va saqlang.
input: n=5 array=10 7 0 -5 -2
output: 7 0 -5 -2 10
*/
func uhomework7and8() {
    min := -5
	max := 25
	fmt.Print("Slice nechta elementdan iborat bo'lsin ")
	son := 0
	fmt.Scan(&son)
    slice := []int{}
	for i:=0; i<son; i++{
		num:=rand.Intn(max-min)+min
		slice = append(slice, num)
		fmt.Print(num," ")
	}
	fmt.Println()
	b:=slice[0]
	for i:=0; i<son; i++{
		if i==son-1{
			fmt.Print(b," ")
		} else {
		slice[i]=slice[i+1]
		fmt.Print(slice[i]," ")
		}
	}
	fmt.Println()
}
/* 9. n ta sondan iborat int sliceni -15 dan 8 gacha random sonlar 
bilan to'ldiring va ushbu slice elementlarini o'nga 1 ta elementga suring va saqlang.
input: n=5 array=-10 5 0 1 -2
output: -2 -10 5 0 1
*/
func uhomework9(){
	min:=-15
	max:=8
	fmt.Print("Slice uzunligini kiriting ")
	uzunlik:=0
	fmt.Scan(&uzunlik)
	slice := []int{}
    for i:=0; i<uzunlik; i++{
		num:=rand.Intn(max-min)+min
		slice = append(slice, num)
		fmt.Print(num," ")
	}
	fmt.Println()
	b:=slice[uzunlik-1]
	for i:=0; i<uzunlik; i++{
		if i==0{
			fmt.Print(b," ")
		} else {
			x:=i-1
			fmt.Print(slice[x]," ")
		}
	}
	fmt.Println()
}
//10. n ta sondan iborat int sliceni 8 dan 21 gacha random sonlar bilan to'ldiring
// ushbu slice elementlarini chapga k ta elementga suring va saqlang.
// Input: n=10 array 1 2 3 4 5 6 7 8 9 10 k=3
// Output: 4 5 6 7 8 9 10 1 2 3
func uhomework10() {
	min:=8
	max:=21
	fmt.Print(" Slice uzunligini kiriting ")
	uzunlik:=0
	slice := []int{}
	fmt.Scan(&uzunlik)
	for i:=0; i<uzunlik; i++{
		num := rand.Intn(max-min)+min
		slice = append(slice, num)
        fmt.Print(num," ")
	}
	fmt.Println()
	fmt.Print(" slice elementlarini chapga nechta elementga surish kerak ")
	k:=0
	fmt.Scan(&k)
	for i:=k; i<uzunlik; i++{
        fmt.Print(slice[i], " ")
	}
	for i:=0; i<k; i++{
		fmt.Print(slice[i], " ")
	}
	fmt.Println()
}
/* 11. n ta sondan iborat int seriyani -11 dan 70 gacha random sonnlar 
bilan to'ldiring  va ushbu seriya elementlarini o'nga k ta elementga suring va saqlang
Input: n=10 array = 1 2 3 4 5 6 7 8 9 10 k=3
Output: 8 9 10 1 2 3 4 5 6 7 
*/
func uhomework11() {
	min:=-11
	max:=70
	fmt.Print(" Enter a slice's length ")
	len:=0
	fmt.Scan(&len)
	slice:=make([]int, len)
	for i:=0; i<len; i++{
      slice[i]=rand.Intn(max-min)+min
	  fmt.Print(slice[i]," ")
	}
	fmt.Println()
	fmt.Print("Enter a k(int) ")
	k:=0
	fmt.Scan(&k)
	for i:=len-k; i<len; i++{
		fmt.Print(slice[i]," ")
	}
	for i:=0; i<len-k; i++{
		fmt.Print(slice[i]," ")
	}
	fmt.Println()
}

