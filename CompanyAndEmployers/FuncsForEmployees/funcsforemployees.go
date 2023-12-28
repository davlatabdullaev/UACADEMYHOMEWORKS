package funcsforemployees

import (
	connectdb "basa-company/ConnectDB"
	structforemployees "basa-company/StructForEmployees"
	"database/sql"
	"fmt"
	"log"
)


type BasaEmployees struct{
	db *sql.DB
}

func New(db *sql.DB) BasaEmployees{
	return BasaEmployees{
      db: db,
	}
}

func FuncsForEmployees() {
	db, err := connectdb.ConnectDB()
	if err!=nil{
		log.Fatal("error during connecting database...")
	}
	DataBaseEmployee := New(db)
	defer db.Close()
	command := 0
	fmt.Println("Please select command")
	fmt.Print(`
	1 - insert employee
	2 - get employee by id
	3 - get all employees
	4 - update employee by id
	5 - delete employee by id
	`)
	fmt.Scan(&command)
	switch(command){
	case 1:
		// insert employee
	err :=	DataBaseEmployee.InsertEmployee()
	if err!=nil{
		fmt.Println("error during insert employee...  ", err)
	}
    case 2: 
	// get employee by id
	employee, err := DataBaseEmployee.GetEmployeeById()
	if err!= nil{
		log.Fatal("error during get employee by id ...  ")
	}
	fmt.Println(employee)
	case 3: 
	// get all employee
	Employees, err := DataBaseEmployee.GetAllEmployee()
	if err!=nil{
		log.Fatal("error during get all employee ... ")
	}
	fmt.Println(Employees)
	case 4:
		// update employee by id 
	if err = DataBaseEmployee.UpdateEmployeeByID(); err!=nil{
		log.Fatal("error during update employee by id ... ")
	}
	case 5:
		// delete employee by id
		err := DataBaseEmployee.DeleteEmployeeById()
		if err!=nil{
			log.Fatal("error during delete employee by id ... ")
		}
	default: 
	fmt.Println("no such command exists...")
	}
}

func (b BasaEmployees) InsertEmployee() error{
	insertEmployee:=structforemployees.ForInsertEmployee()
	if _, err := b.db.Exec(`insert into employees (id, employee_name, phone, company_id, salary, emloyee_type) values ($1, $2, $3, $4, $5, $6)`, insertEmployee.ID, insertEmployee.EmployeeName, insertEmployee.Phone, insertEmployee.CompanyID, insertEmployee.Salary, insertEmployee.EmployeeType); err!=nil{
		return err
	}
	var companyIDExists bool
	err := b.db.QueryRow("SELECT EXISTS(SELECT 1 FROM company WHERE id = $1)", insertEmployee.CompanyID).Scan(&companyIDExists)
	if err != nil {
		log.Fatal("Error checking if company_id exists:", err)
		return err
	}
	if !companyIDExists {
	
		log.Println("Warning: The company_id does not exist in the company table.")
	}

	fmt.Println("insert employee completed succesfully...")
	return nil
}
 
func (b BasaEmployees) GetEmployeeById() (structforemployees.Employees, error) {
	id :=structforemployees.ForGetEmployeeById()
row :=	b.db.QueryRow(`select id, employee_name, phone, company_id, salary, emloyee_type from employees where id = $1`, id.ID)
employee := structforemployees.Employees{}
 err := row.Scan(&employee.ID, &employee.EmployeeName, &employee.Phone, &employee.CompanyID, &employee.Salary, &employee.EmployeeType)
if err!=nil{
	return structforemployees.Employees{}, err
}
return employee, nil
}

func (b BasaEmployees) GetAllEmployee() ([]structforemployees.Employees, error) {
	rows, err  := b.db.Query(`select id, employee_name, phone, company_id, salary, emloyee_type from employees`)
	if err!=nil{
		return nil, err
	}
	Employees := []structforemployees.Employees{}
	for rows.Next() {
		Employee := structforemployees.Employees{}
    rows.Scan(&Employee.ID, &Employee.EmployeeName, &Employee.Phone, &Employee.CompanyID, &Employee.Salary, &Employee.EmployeeType)  	
    Employees = append(Employees, Employee)	
}
return Employees, nil
}

func (b BasaEmployees) UpdateEmployeeByID() error{
	updating := structforemployees.ForUpdateEmployeeById()
	_, err := b.db.Exec(`update employees set employee_name = $1, phone = $2, company_id = $3, salary = $4, emloyee_type = $5 where id = $6`, updating.EmployeeName, updating.Phone, updating.CompanyID, updating.Salary, updating.EmployeeType, updating.ID)
       if err!=nil{
		log.Fatal("Error during update employee by id:", err)
		return err
	   }
	   fmt.Println(" updating completed succesfully ... ")
  return nil
	}

func (b BasaEmployees)  DeleteEmployeeById() error{
	deleting := structforemployees.ForGetEmployeeById()
	_, err := b.db.Exec(`delete from employees where id = $1`, deleting.ID)
	if err!=nil {
		return err
	}
	fmt.Println("deleting completed succesfully...")
	return nil
}	
