package main

import (
	"basa/config"
	"basa/controller"
	"basa/storage/postgres"
	"fmt"
	"log"
)

func main() {
cfg := config.Load()

store, err := postgres.New(cfg)
if err!=nil{
	log.Fatalln("error while connecting to db err: ", err.Error())
	return
}
defer store.DB.Close()

 con := controller.New(store)

cmd := 0
fmt.Print(`
ENTER COMMAND:
1 - users
2 - orders
3 - products
4 - order products
`)
fmt.Scan(&cmd)
switch(cmd){
case 1: 
userscmd := 0
fmt.Print(`
ENTER USERS COMMAND:
1 - insert user
2 - get user by id
3 - get users list
4 - update user
5 - delete user
`)
fmt.Scan(&userscmd)
switch(userscmd) {
case 1:
	con.InsertUser()
case 2:
	con.GetUserByID()
case 3:
	con.GetUsersList()
case 4:
	con.UpdateUser()
case 5:
	con.DeleteUser()
default:
	fmt.Println("command not found...")
}
case 2:
	orderscmd := 0
fmt.Print(`
ENTER ORDERS COMMAND:
1 - insert order
2 - get order by id
3 - get orders list
4 - update order
5 - delete order
`)
fmt.Scan(&orderscmd)
switch(orderscmd) {
case 1:
	con.InsertOrders()
case 2:
	con.GetOrderById()
case 3:
	con.GetOrdersList()
case 4:
	con.UpdateOrder()
case 5:
	con.DeleteOrder()
default:
	fmt.Println("command not found...")
}

case 3:
	productscmd := 0
fmt.Print(`
ENTER PRODUCTS COMMAND:
1 - insert product
2 - get product by id
3 - get products list
4 - update product
5 - delete product
`)
fmt.Scan(&productscmd)
switch(productscmd) {
case 1:
	con.InsertProduct()
case 2:
	con.GetProductByID()
case 3:
	con.GetProductList()
case 4:
	con.UpdateProductByID()
case 5:
	con.DeleteProductByID()
default:
	fmt.Println("command not found...")
}
case 4: 
	orderproductcmd := 0
fmt.Print(`
ENTER ORDER PRODUCT COMMAND:
1 - insert orderproduct
2 - get order product by id
3 - get order products list
4 - update order product
5 - delete order product
`)
fmt.Scan(&orderproductcmd)
switch(orderproductcmd) {
case 1:
	con.InsertOrderProducts()
case 2:
	con.GetOrderProductId()
case 3:
	con.GetOrderProductsList()
case 4:
	con.UpdateOrderProductById()
case 5:
	con.DeleteOrderProductsByID()
default:
	fmt.Println("command not found...")
}
default:
	fmt.Println("command not found...")

}
}