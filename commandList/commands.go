package commandlist

import (
	"fmt"
	"database/sql"
	"basa/user"
	_"github.com/lib/pq"
)
type Inventory struct{
	db *sql.DB
}
func New(db *sql.DB) Inventory{
	return Inventory{
		db: db,
	}
}

func CommandLists() {
	db, err := connectDB()
	if err != nil{
		panic(err)
	}
	defer db.Close()
	inv := New(db)
     var num int = 0
	 fmt.Print(`
	 Enter command:
	 1 - add user
	 2 - get user by id
	 3 - list of users
	 4 - delete user by id
	 `)
	 fmt.Scan(&num)
	 switch(num) {
	 case 1: 
	 if err := inv.AddUser(); err!=nil{
		fmt.Println("Error adding user: ",err)
		return
	 }
	 case 2:
	 user, err := inv.GetUserByID()
	 if err!=nil{
		fmt.Println("Error get user by id ", err)
	 }
	 fmt.Println(user)
	 case 3:
	 users, err := inv.GetAllUsers()
	 if err!=nil{
		fmt.Println("Error get all users ", err)
	 } else{
		fmt.Println(users)
	 }
	 case 4: 
    if err := inv.DeleteUser(); err!=nil{
		fmt.Println("Error delete user ",err)
	}
default: fmt.Println("no such command exists")
	 }

}
func (i Inventory) AddUser() error{
	fmt.Println("Enter users first name")
	var a string = ""
	fmt.Scan(&a)
	fmt.Println("Enter users last name")
	var b string = ""
	fmt.Scan(&b)
	fmt.Println("Enter users email")
	var c string = ""
	fmt.Scan(&c)
	if _, err := i.db.Exec(
		`insert into users (first_name, last_name, email) values ($1, $2, $3)`, a,b,c,
	); err!=nil{
     return err
	}
	return nil
}
func (i Inventory) GetUserByID() (user.User, error) {
	q:=0
    fmt.Println("Enter user's id")
	fmt.Scan(&q)
	row := i.db.QueryRow(`select id, first_name, last_name, email from users where id = $1`, q)
   
	m := user.User{}

   if err := row.Scan(&m.ID, &m.FirstName, &m.LastName, &m.Email); err!=nil{
	   panic(err)
   }
   
   return m, nil
}
func (i Inventory) GetAllUsers() ([]user.User, error) {
	rows, err := i.db.Query(
		`select * from users`,
	)
	if err!=nil{
		return nil, err
	}
	cs := []user.User{}

	for rows.Next(){
		c:=user.User{}
		if err = rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email); err!=nil{
			return nil, err
		}
		cs=append(cs, c)
	}
	return cs, nil
}
func (i Inventory) DeleteUser() error {
	id :=0
	fmt.Println("Enter user's id ")
	fmt.Scan(&id)
   if _, err := i.db.Exec(`delete from users where id = $1`, id); err!=nil{
	panic(err)
   }
	return nil
}
func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=davlat password=1 database=uzerz sslmode=disable")
	if err!=nil{
		return nil, err
	}
	return db, nil
}