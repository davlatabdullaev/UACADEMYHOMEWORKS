package structforemployees

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Employees struct {
	ID           uuid.UUID
	EmployeeName string
	Phone        string
	CompanyID    int
	Salary       int
	EmployeeType string
}

func ForInsertEmployee() Employees{
	var InsertEmployeeName string = ""
	fmt.Println("Enter employee name: ")
	fmt.Scan(&InsertEmployeeName)
	var InsertEmployeePhone string = ""
	fmt.Println("Enter employee phone: ")
	fmt.Scan(&InsertEmployeePhone)
	var InsertEmployeeCompanyId int = 0
	fmt.Println("Enter employee company id: ")
	fmt.Scan(&InsertEmployeeCompanyId)
    var InsertEmployeeSalary int = 0
	fmt.Println("Enter employee salary: ")
	fmt.Scan(&InsertEmployeeSalary)
	var InsertEmployeeType string = ""
	fmt.Println("Enter employee type: ")
	fmt.Scan(&InsertEmployeeType)
	return Employees{
		ID: uuid.New(),
		EmployeeName: InsertEmployeeName,
		Phone: InsertEmployeePhone,
		CompanyID: InsertEmployeeCompanyId,
		Salary: InsertEmployeeSalary,
		EmployeeType: InsertEmployeeType,
	}
}
 
func ForGetEmployeeById() Employees{
	var EmployeeId string = ""
	fmt.Println("Enter employee Id: ")
	fmt.Scan(&EmployeeId)
	convUuid, err := uuid.Parse(EmployeeId)
	if err!= nil {
    log.Fatal("error... ", err)
	}
	return Employees{
		ID: convUuid,
	}
}

func ForUpdateEmployeeById() Employees{
	var UpdateEmployeeId string = ""
	fmt.Println("Enter employee ID: ")
	fmt.Scan(&UpdateEmployeeId)
	var UpdateEmployeeName string = ""
	fmt.Println("Enter new employee name: ")
	fmt.Scan(&UpdateEmployeeName)
	var UpdateEmployeePhone string = ""
	fmt.Println("Enter new employee phone: ")
	fmt.Scan(&UpdateEmployeePhone)
	var UpdateEmployeeCompanyId int = 0
	fmt.Println("Enter  new employee company id: ")
	fmt.Scan(&UpdateEmployeeCompanyId)
    var UpdateEmployeeSalary int = 0
	fmt.Println("Enter new employee salary: ")
	fmt.Scan(&UpdateEmployeeSalary)
	var UpdateEmployeeType string = ""
	fmt.Println("Enter new employee type: ")
	fmt.Scan(&UpdateEmployeeType)
	uuidtype, err := uuid.Parse(UpdateEmployeeId)
	if err!=nil{
		log.Fatal("error during convertation string in uuid type... ")
	}
	return Employees{
		ID: uuidtype,
		EmployeeName: UpdateEmployeeName,
		Phone: UpdateEmployeePhone,
		CompanyID: UpdateEmployeeCompanyId,
		Salary: UpdateEmployeeSalary,
		EmployeeType: UpdateEmployeeType,
	}
}
