package controller

import (
	"basa/structs"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (c Controller) InsertOrderProducts() {
if err := c.Store.OrderProductsStorage.InsertOrderProduct(getOrderProduct()); err!=nil{
	log.Fatal("error while insert order products ", err.Error())
    return
}
fmt.Println("inserting completed succesfully...")
}

func getOrderProduct() structs.OrderProducts{
	var oidStr string
	fmt.Println("enter order id: ")
	fmt.Scan(&oidStr)
	uoid := uuid.MustParse(oidStr)
	var pidStr string
	fmt.Println("enter product id: ")
	fmt.Scan(&pidStr)
	upid := uuid.MustParse(pidStr)
	quantity := 0
	fmt.Println("enter quantity: ")
	fmt.Scan(&quantity)
	price := 0
	fmt.Println("enter price: ")
	fmt.Scan(&price)
    return structs.OrderProducts{
		OrderID: uoid,
		ProductID: upid,
		Quantity: quantity,
		Price: price,
	}
}

func (c Controller) GetOrderProductId() {
    var idStr string
	fmt.Println("enter id: ")
	fmt.Scan(&idStr)
	uid := uuid.MustParse(idStr)
orderProduct, err :=	c.Store.OrderProductsStorage.GetByIDOrderProduct(uid)
if err!=nil{
	log.Fatal("error while get order product by id ...", err.Error())
    return
}
fmt.Println(orderProduct)

}

func (c Controller) GetOrderProductsList() {
orderProducts, err :=	c.Store.OrderProductsStorage.GetListOrderProducts()
if err!=nil{
	log.Fatal("error while get order products list...", err.Error())
	return
}
fmt.Println(orderProducts)
}

func (c Controller) UpdateOrderProductById() {
if err := c.Store.OrderProductsStorage.UpdateOrderProducts(getOrderProductForUpdate()); err!=nil{
	log.Fatal("error while update order product...", err.Error())
	return
}
fmt.Println("order product updated succesfully...")
}

func getOrderProductForUpdate() structs.OrderProducts{
	var id string
	fmt.Println("enter ")
	fmt.Scan(id)
	uid := uuid.MustParse(id)
	var oidStr string
	fmt.Println("enter order id: ")
	fmt.Scan(&oidStr)
	uoid := uuid.MustParse(oidStr)
	var pidStr string
	fmt.Println("enter product id: ")
	fmt.Scan(&pidStr)
	upid := uuid.MustParse(pidStr)
	quantity := 0
	fmt.Println("enter quantity: ")
	fmt.Scan(&quantity)
	price := 0
	fmt.Println("enter price: ")
	fmt.Scan(&price)
    return structs.OrderProducts{
		Id: uid,
		OrderID: uoid,
		ProductID: upid,
		Quantity: quantity,
		Price: price,
	}
}

func (c Controller) DeleteOrderProductsByID() {
	var id string
	fmt.Println("enter ")
	fmt.Scan(id)
	uid := uuid.MustParse(id)
	if err := c.Store.OrderProductsStorage.DeleteOrderProducts(uid); err!=nil{
		log.Fatal("error while delete order products...")
	}
}
