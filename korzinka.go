package main

import "fmt"

type Xaridor struct{
  Id int
  ismi string 
  savati []MahsulotKartasi
}
type MahsulotKartasi struct{
	mahsulotId  int
	OlchamId int
	miqdori int
}

type Mahsulot struct{
	ID int
	nomi string
	olchami []Olcham
}
type Olcham struct{
	id int
	nomi string
	narxi int
}

func main() {
	Kassa()
	}
	


func Kassa() {
   
   Xaridorlar := []Xaridor{
	{Id: 1, ismi: "Jamshid", savati: []MahsulotKartasi{
		{mahsulotId: 11, OlchamId: 2, miqdori: 4},
		{mahsulotId: 8, OlchamId: 1, miqdori: 5},
		{mahsulotId: 3, OlchamId: 3, miqdori: 2},
	}},
	{Id: 2, ismi: "Otabek", savati: []MahsulotKartasi{
		{mahsulotId: 11, OlchamId: 2, miqdori: 4},
		{mahsulotId: 3, OlchamId: 1, miqdori: 5},
		{mahsulotId: 7, OlchamId: 3, miqdori: 2},
	}},
	{Id: 3, ismi: "Muhammad", savati: []MahsulotKartasi{
		{mahsulotId: 8, OlchamId: 2, miqdori: 4},
		{mahsulotId: 2, OlchamId: 1, miqdori: 5},
		{mahsulotId: 5, OlchamId: 3, miqdori: 2},
		{mahsulotId: 14, OlchamId: 2, miqdori: 19},
	}},
	{Id: 4, ismi: "Anvar", savati: []MahsulotKartasi{
		{mahsulotId: 7, OlchamId: 2, miqdori: 4},
		{mahsulotId: 4, OlchamId: 1, miqdori: 5},
		{mahsulotId: 2, OlchamId: 3, miqdori: 2},
	}},
	{Id: 5, ismi: "Islom", savati: []MahsulotKartasi{
		{mahsulotId: 10, OlchamId: 2, miqdori: 4},
		{mahsulotId: 9, OlchamId: 1, miqdori: 5},
		{mahsulotId: 3, OlchamId: 3, miqdori: 2},
	}},
	{Id: 6, ismi: "Jasur", savati: []MahsulotKartasi{
		{mahsulotId: 1, OlchamId: 2, miqdori: 4},
		{mahsulotId: 8, OlchamId: 1, miqdori: 5},
		{mahsulotId: 4, OlchamId: 3, miqdori: 2},
	}},
	{Id: 7, ismi: "Ali", savati: []MahsulotKartasi{
		{mahsulotId: 1, OlchamId: 2, miqdori: 4},
		{mahsulotId: 12, OlchamId: 1, miqdori: 5},
		{mahsulotId: 3, OlchamId: 3, miqdori: 22},
		{mahsulotId: 11, OlchamId: 1, miqdori: 5},
		{mahsulotId: 7, OlchamId: 3, miqdori: 21},

	}},
	{Id: 8, ismi: "Bilol", savati: []MahsulotKartasi{
		{mahsulotId: 8, OlchamId: 2, miqdori: 4},
		{mahsulotId: 9, OlchamId: 1, miqdori: 35},
		{mahsulotId: 5, OlchamId: 3, miqdori: 2},
		{mahsulotId: 8, OlchamId: 2, miqdori: 4},
		{mahsulotId: 9, OlchamId: 1, miqdori: 5},
		{mahsulotId: 5, OlchamId: 3, miqdori: 12},
	}},
	{Id: 9, ismi: "Sulaymon", savati: []MahsulotKartasi{
		{mahsulotId: 4, OlchamId: 2, miqdori: 4},
		{mahsulotId: 1, OlchamId: 1, miqdori: 5},
		{mahsulotId: 13, OlchamId: 3, miqdori: 2},
	}},
	{Id: 10, ismi: "Ibrohim", savati: []MahsulotKartasi{
		{mahsulotId: 6, OlchamId: 2, miqdori: 4},
		{mahsulotId: 13, OlchamId: 1, miqdori: 5},
		{mahsulotId: 12, OlchamId: 3, miqdori: 2},
	}},
}

   Mahsulotlar := []Mahsulot{
   {ID: 1, nomi: "Non", olchami: []Olcham{
	{id: 1, nomi: "250gram", narxi: 2000},
	{id: 2, nomi: "500gram", narxi: 3000},
	{id: 3, nomi: "800gram", narxi: 4000 },
   }},
   {ID: 2, nomi: "Olma", olchami: []Olcham{
	{id: 1, nomi: "sariq", narxi: 19990},
	{id: 2, nomi: "qizil", narxi: 24990},
	{id: 3, nomi: "yashil", narxi: 21990},
   }},
   {ID: 3, nomi: "Kolbasa", olchami: []Olcham{
	{id: 1, nomi: "kichkina", narxi: 19990},
	{id: 2, nomi: "o'rtacha", narxi: 26990},
	{id: 3, nomi: "katta", narxi: 46990},
   }},
   {ID: 4, nomi: "Snickers", olchami: []Olcham{
	{id: 1, nomi: "kichkina", narxi: 6990},
	{id: 2, nomi: "o'rtacha", narxi: 9990},
	{id: 3, nomi: "katta", narxi: 12990},
   }},
   {ID: 5, nomi: "Borjomi", olchami: []Olcham{
	{id: 1, nomi: "0,5L", narxi: 11990},
	{id: 2, nomi: "1,0L", narxi: 14990},
	{id: 3, nomi: "1,5L", narxi: 18990},
   }},
     {ID: 6, nomi: "Ahmad Tea", olchami: []Olcham{
	{id: 1, nomi: "kichkina", narxi: 13990},
	{id: 2, nomi: "o'rtacha", narxi: 17990},
	{id: 3, nomi: "katta", narxi: 19990},
   }},
   {ID: 7, nomi: "Nescafe", olchami: []Olcham{
	{id: 1, nomi: "classic", narxi: 21990},
	{id: 2, nomi: "gold", narxi: 26990},
	{id: 3, nomi: "platinum", narxi: 33990},
   }},
   {ID: 8, nomi: "Pizza", olchami: []Olcham{
	{id: 1, nomi: "25sm", narxi: 44990},
	{id: 2, nomi: "30sm", narxi: 52990},
	{id: 3, nomi: "35sm", narxi: 60990},
   }},
   {ID: 9, nomi: "Hot Dog", olchami: []Olcham{
	{id: 1, nomi: "kichkina", narxi: 15000},
	{id: 2, nomi: "o'rtacha", narxi: 19000},
	{id: 3, nomi: "katta", narxi: 23000},
   }},
   {ID: 10, nomi: "Coca cola", olchami: []Olcham{
	{id: 1, nomi: "0,5L", narxi: 6000},
	{id: 2, nomi: "1,0L", narxi: 10000},
	{id: 3, nomi: "1,5L", narxi: 13000},
   }},
   {ID: 11, nomi: "Fanta", olchami: []Olcham{
	{id: 1, nomi: "0,5L", narxi: 7000},
	{id: 2, nomi: "1,0L", narxi: 11000},
	{id: 3, nomi: "1,5L", narxi: 13000},
   }},
   {ID: 12, nomi: "Pepsi Cola", olchami: []Olcham{
	{id: 1, nomi: "0,5L", narxi: 8000},
	{id: 2, nomi: "1,0L", narxi: 12000},
	{id: 3, nomi: "1,5L", narxi: 14000},
   }},
   {ID: 13, nomi: "Shampoon Clear Men", olchami: []Olcham{
	{id: 1, nomi: "250ml", narxi: 23990},
	{id: 2, nomi: "380ml", narxi: 32990},
	{id: 3, nomi: "450ml", narxi: 39990},
   }},
   {ID: 14, nomi: "Colgate", olchami: []Olcham{
	{id: 1, nomi: "75ml", narxi: 19490},
	{id: 2, nomi: "100ml", narxi: 24990},
	{id: 3, nomi: "120ml", narxi: 28990},
   }},
}
// Har bitta xaridorning savatida necha so'mlik mahsulot borligini topish
   for _, xaridor := range Xaridorlar{
	sum:=0
	for _, mahsulotlari := range xaridor.savati{
		for _, mahsulotlar := range Mahsulotlar{
			if mahsulotlari.mahsulotId==mahsulotlar.ID{
				for _, qanchaligi := range mahsulotlar.olchami{
                   if mahsulotlari.OlchamId==qanchaligi.id{
					  sum+=mahsulotlari.miqdori*qanchaligi.narxi
				   }
				}
			}
		}
	}
	fmt.Print(xaridor.ismi, "ning savatidagi mahsulotlarning umumiy narxi: ", sum," so'm\n")
   }
   fmt.Println("\nKassaga eng ko'p urilgan element")
// Korzinka kassasiga qaysi mahsulot eng ko'p qo'shilganligini topish
// Buning uchun korzinkada bor mahsulotlarni ularning soni bo'yicha aylanamiz
// keyin ushbu mahsulotlar xaridor savatida bor bo'lsa miqdorini qo'shib boramiz

var kopurilganmahsulotnomi string

max:=0
for _, zmahsulotlar := range Mahsulotlar{

hisob:=0
   for _, zxaridor := range Xaridorlar{
	for _, zsavatdagiMahsulotlar := range zxaridor.savati{
		if zsavatdagiMahsulotlar.mahsulotId==zmahsulotlar.ID{
			hisob+=zsavatdagiMahsulotlar.miqdori
		}
	}
   }
   if max<hisob{
	max=hisob
	kopurilganmahsulotnomi=zmahsulotlar.nomi

   }
}
fmt.Print("Kassaga ",kopurilganmahsulotnomi, " dan ", max, " ta qo'shildi ")

}