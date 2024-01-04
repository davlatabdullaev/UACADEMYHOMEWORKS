package controller

import (
	"basa/structs"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (c Controller) InsertProduct() {
	price :=  0
	fmt.Println("Enter product price: ")
	fmt.Scan(&price)
	var ProductName string
	fmt.Println("enter product name: ")
	fmt.Scan(&ProductName)
	if err := c.Store.ProductsStorage.InsertProducts(price,ProductName); err!=nil{
		log.Fatal("error while insert product...", err.Error())
		return
	}
	fmt.Println("product succesfully added...")
}

func (c Controller) GetProductByID() {
	var idStr string
	fmt.Println("enter products id: ")
	fmt.Scan(&idStr)
	uid, err := uuid.Parse(idStr)
	if err!=nil{
		log.Fatalln("this id is not uuid type ", err.Error())
		return
	}
	product, err:=c.Store.ProductsStorage.GetByIDProduct(uid)
	if err!=nil{
		log.Fatal("error while get product by id...", err.Error())
		return
	}
	fmt.Println(product)
}

func (c Controller) GetProductList() {
products, err :=	c.Store.ProductsStorage.GetListProduct()
if err!=nil{
	log.Fatal("error while get lists product...", err.Error())
	return
}
fmt.Println(products)
}

func (c Controller) UpdateProductByID() {
	product := GetProductInfo()
if err := c.Store.ProductsStorage.UpdateProducts(product); err!=nil{
	log.Fatal("error while update product by id...", err.Error())
}
fmt.Println("product updated succesfully...")
}

func GetProductInfo() structs.Products{
	var idStr string
	fmt.Println("enter products id: ")
	fmt.Scan(&idStr)
	var price int = 0
	fmt.Println("enter products price: ")
	fmt.Scan(&price)
	var Name string
	fmt.Println("enter products name: ")
	fmt.Scan(&Name)
	uid := uuid.MustParse(idStr)
	return structs.Products{
		ID: uid,
		Price: price,
		ProductName: Name,
	}
}

func (c Controller) DeleteProductByID() {
	var idStr string
	fmt.Println("enter products id: ")
	fmt.Scan(&idStr)
	uid := uuid.MustParse(idStr)
	if err := c.Store.ProductsStorage.DeleteProducts(uid); err!=nil{
		log.Fatal("error while product deleting ...", err.Error())
	}
	fmt.Println("product succesfully deleted")
}
