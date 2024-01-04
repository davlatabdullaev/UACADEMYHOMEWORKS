package controller

import (
	"basa/structs"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (c Controller) InsertOrders() {
	var amount string
	fmt.Println("enter amount: ")
	fmt.Scan(&amount)
	var idStr string
	fmt.Println("enter user id: ")
	fmt.Scan(&idStr)
	id, err := uuid.Parse(idStr)
	if err!=nil{
		log.Fatal("this id is not uuid type ", err.Error())
	}
	if err := c.Store.OrdersStorage.InsertOrders(amount,id); err!=nil{
		log.Fatal("error while insert orders ", err.Error())
	}
	fmt.Println("order inserting succesfully...")
}

func (c Controller) GetOrderById() {
	var idStr string
	fmt.Println("enter user id: ")
	fmt.Scan(&idStr)
	id, err := uuid.Parse(idStr)
	if err!=nil{
		log.Fatal("this id is not uuid type ", err.Error())
	}
order, err :=	c.Store.OrdersStorage.GetByIDOrder(id)
if err!=nil{
	log.Fatal("error while get order by id err: ", err.Error())
}
	fmt.Println(order)	
}

func (c Controller) GetOrdersList() {
orders, err :=	c.Store.OrdersStorage.GetListOrder()
if err != nil{
	log.Fatal("error while get orders list: ", err.Error())
}
fmt.Println(orders)
}
 
func (c Controller) UpdateOrder() {
	orders := ForUpdateOrders()
	if err := c.Store.OrdersStorage.UpdateOrders(orders); err!=nil{
		log.Fatal("error while updating order ", err.Error())
	}
	fmt.Println("order updated succesfully...")
}
func ForUpdateOrders() structs.Orders{
	var amount string
	fmt.Println("enter amount: ")
	fmt.Scan(&amount)
	var idStr string
	fmt.Println("enter user id: ")
	fmt.Scan(&idStr)
	uid, err := uuid.Parse(idStr)
	if err!=nil{
		log.Fatal("this id is not uuid type ", err.Error())
	}
	return structs.Orders{
		Amount: amount,
		UserID: uid,
	}
}

func (c Controller) DeleteOrder() {
	var idStr string
	fmt.Println("enter user id: ")
	fmt.Scan(&idStr)
	uid, err := uuid.Parse(idStr)
	if err!=nil{
		log.Fatal("this id is not uuid type ", err.Error())
	}
	if err := c.Store.OrdersStorage.DeleteOrders(uid); err!=nil{
		log.Fatal("error while delete order...")
	}
	fmt.Println("order succesfully deleted...")
}