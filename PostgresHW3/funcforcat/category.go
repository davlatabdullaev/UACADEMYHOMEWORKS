package funcforcat

import (
	_"github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
	"marketbasa/category"
	"time"
)

type Inventory struct {
	db *sql.DB
}

func New(db *sql.DB) Inventory {
	return Inventory{
		db: db,
	}
}
func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=davlat password=1 database=market sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CategoryFuncs() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	inv := New(db)
	fmt.Print(`
	Select category command :
	1 - insert category
	2 - update category by id
	3 - category lists
	4 - get category by id
	5 - delete category by id
	`)
	commands := 0
	fmt.Scan(&commands)
	switch commands {
	case 1:
		if err := inv.InsertCategory(); err != nil {
			log.Fatal("Error insert category: ")
		}
	case 2:
		if err = inv.UpdateCategoryById(); err != nil {
			log.Fatal("error during update category command", err)
		}
	case 3:
		All, err := inv.GetAllCategory()
		if err != nil {
			log.Fatal("error during get all category ", err)
		} else {
			fmt.Println(All)
		}
	case 4:
     Category, err := inv.GetCategoryById()
	 if err!=nil{
		fmt.Println("error during get category by id")
	 } else {
		fmt.Println(Category)
	 }
	case 5: 
	if err = inv.DeleteCategoryById(); err!=nil{
		log.Fatal("error during delete category", err)
	}
default: 
fmt.Println("no such command exists")
	}
}

func (i Inventory) InsertCategory() error {
	var CategoryName string = ""
	fmt.Println("Enter a category name: ")
	fmt.Scan(&CategoryName)
	TimeNow := time.Now()
	var TimeZero time.Time
	if _, err := i.db.Exec(`insert into category (name, created_at, updated_at) values ($1, $2, $3)`, CategoryName, TimeNow, TimeZero); err != nil {
		log.Fatal("error during insert category ", err)
		return err
	}
	return nil
}
func (i Inventory) UpdateCategoryById() error {
	CategoryId := 0
	fmt.Println("Enter a category id: ")
	fmt.Scan(&CategoryId)
	var CategoryName string = ""
	fmt.Println("Enter a new category name: ")
	fmt.Scan(&CategoryName)
	TimeNow := time.Now()
	_, err := i.db.Exec(`update category set name = $1, updated_at = $2 where id = $3`, CategoryName, TimeNow, CategoryId)
	if err != nil {
		log.Fatal("error during update category")
	}
	return nil
}

func (i Inventory) GetAllCategory() ([]category.Category, error) {
	rows, err := i.db.Query(`select * from category`)
	if err != nil {
		return nil, err
	}
	AllCategory := []category.Category{}

	for rows.Next() {
		Category := category.Category{}
		if err = rows.Scan(&Category.ID, &Category.Name, &Category.CreatedAt, &Category.UpdateAt); err != nil {
			return nil, err
		}
		AllCategory = append(AllCategory, Category)
	}
	return AllCategory, nil
}
func (i Inventory) GetCategoryById() (category.Category, error) {
	row := i.db.QueryRow(`select id, name, created_at, updated_at from category where id = $1`,)

   GetCategory := category.Category{}

   if err := row.Scan(&GetCategory.ID, &GetCategory.Name, &GetCategory.CreatedAt, &GetCategory.UpdateAt); err!=nil{
	fmt.Println("error during get category by id", err)
   }
   return category.Category{}, nil
}
func (i Inventory) DeleteCategoryById() error{
   num:=0
   fmt.Println("Enter a category id for delete ")
   fmt.Scan(&num)
 _, err:=  i.db.Exec(`delete from category where id = $1`, num)
   if err!=nil{
log.Fatal("error during delete category")
   }
   return nil
}