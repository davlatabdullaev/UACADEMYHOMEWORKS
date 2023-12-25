package funcforprod

import (
	_"github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
	"marketbasa/product"
	"time"
)

type Inventory struct{
	db *sql.DB
  }
  
  func New(db *sql.DB) Inventory{
	return Inventory{
	  db: db,
	}
  }
  
  func ConnectDB() (*sql.DB, error) {
	  db, err := sql.Open("postgres", "host=localhost port=5432 user=davlat password=1 database=market sslmode=disable")
	  if err!=nil{
		  fmt.Println("Error during connecting", err)
		  return nil, err
	  }
	  return db, nil
  }

func ProductFuncs(){
	db, err := ConnectDB()
	if err!= nil{
		panic(err)
	}
	defer db.Close()
	inv := New(db)
	Num := 0
fmt.Print(`
            Welcome product table!
            Select a command:
            1 - insert 
            2 - update by id
            3 - get list products
            4 - get product by id
            5 - delete by id
`)
fmt.Scan(&Num)
switch(Num) {
case 1: 
if err := inv.InsertProduct(); err!=nil{
	log.Fatal("Error during Insert product ", err)
}
case 2:
if err:=inv.UpdateProductById(); err!=nil{
	log.Fatal("error during update product ", err)
}
case 3:
   products, err :=inv.GetProducts()
   if err!=nil{
	fmt.Println("Error get all products", err)
   } else {
	fmt.Println(products)
   }
case 4:
OneProduct, err :=	inv.GetProductById()
if err!=nil{
	fmt.Println("Error get product by id", err)
}
fmt.Println(OneProduct)
case 5: 
  if err := inv.DeleteProductById(); err!=nil{
	fmt.Println("error during delete product by id",err)
  }
default: fmt.Println("no such command exists")
}
}

func (i Inventory) InsertProduct() error{
	var ProductName string  = ""
	fmt.Println("Enter a product name: ")
	fmt.Scan(&ProductName)
    var ProductPrice int  = 0
	fmt.Println("Enter a product price: ")
	fmt.Scan(&ProductPrice)
	var CategoryId int = 0
	fmt.Println("Enter a category id: ")
	fmt.Scan(&CategoryId)
	TimeNow := time.Now()
	var TimeZero time.Time
	_, err := i.db.Exec(`insert into product (name, price, category_id, created_at, updated_at) values ($1, $2, $3, $4, $5)`, ProductName, ProductPrice, CategoryId, TimeNow, TimeZero)
      if err!=nil{
		fmt.Println("Error during product input", err)
		return err
	  }
	  return nil
}
func (i Inventory) UpdateProductById() error{
     var ProductId int = 0
	 fmt.Println("Enter product id")
	 fmt.Scan(&ProductId)
	 var ProductName string = ""
	 fmt.Println("enter a new product name")
	 fmt.Scan(&ProductName)
	 var ProductPrice int = 0
	 fmt.Println("Enter product id")
	 fmt.Scan(&ProductPrice)
	 now:=time.Now()
	if _, err := i.db.Exec(`update product set name = $1, price = $2, updated_at = $3 where id = $4`, ProductName, ProductPrice, now, ProductId); err!=nil{
		log.Fatal("Error during product update", err)
		return err
	}
	return nil
}

func (i Inventory) GetProducts() ([]product.Product, error){

	rows, err := i.db.Query(`select * from product`)
   if err!=nil{
	return nil, err
   }
   products := []product.Product{}
   for rows.Next(){
    product:=product.Product{}
	if err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.CreateAt, &product.UpdateAt); err!=nil{
		return nil, err
	}
	products=append(products, product)
   }
return products, nil
}
func (i Inventory) GetProductById() (product.Product, error) {
	num:=0
    fmt.Println("Enter product's id")
	fmt.Scan(&num)
	row := i.db.QueryRow(`select id, name, price, category_id, created_at, updated_at from product where id = $1`, num)
   
	GetProduct := product.Product{}

   if err := row.Scan(&GetProduct.ID, &GetProduct.Name, &GetProduct.Price, &GetProduct.CategoryID, &GetProduct.CreateAt, &GetProduct.UpdateAt); err!=nil{
	   panic(err)
   }
   
   return GetProduct, nil
}
func (i Inventory) DeleteProductById() error {
	ProductId := 0
	fmt.Println("Enter a product's id: ")
	fmt.Scan(&ProductId)
	if _, err := i.db.Exec(`delete from product where id = $1`, ProductId); err!=nil{
		fmt.Println("error during delete product by id", err)
	}
	return nil
}
