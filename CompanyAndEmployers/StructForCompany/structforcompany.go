package structforcompany

import (
	"fmt"
	_"github.com/lib/pq"
)

type Company struct {
	ID             int
	CompanyName    string
	EmployeesCount int
}
func ForInsertCompany() Company{
	fmt.Println("Enter company name: ")
    var NewCompanyName string = ""
	fmt.Scan(&NewCompanyName)
	fmt.Println("Enter company employees count: ")
    var NewEmployeesCount int = 0
	fmt.Scan(&NewEmployeesCount)
	NewCompany := Company{
		CompanyName: NewCompanyName,
		EmployeesCount: NewEmployeesCount,
	}
	return NewCompany
}

func ForGetCompanyByID() Company{
	fmt.Println("Enter company id: ")
	var GetCompById int = 0
	fmt.Scan(&GetCompById)
	CompanyID := Company {
		ID: GetCompById,
	}
	return CompanyID
}
func ForUpdateCompanyById() Company{
	fmt.Println("Enter company id: ")
	var CompanyId int = 0
	fmt.Scan(&CompanyId)
	fmt.Println("Enter New Company Name: ")
	var NewCompanyNameForUpdate string
	fmt.Scan(&NewCompanyNameForUpdate)
	fmt.Println("Enter New Company Employees Count: ")
	var NewCompanyEmloyeesCount = 0
	fmt.Scan(&NewCompanyEmloyeesCount)
	NewCompany:=Company{
		ID: CompanyId,
		CompanyName: NewCompanyNameForUpdate,
		EmployeesCount: NewCompanyEmloyeesCount,
	}
	return NewCompany
}
