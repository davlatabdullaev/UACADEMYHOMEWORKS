package main

import (
	funcsforcompany "basa-company/FuncsForCompany"
	funcsforemployees "basa-company/FuncsForEmployees"
	//trigger "basa-company/Trigger"
	"fmt"
)

func main() {
	cmd := 0
	fmt.Print(`
        	Select the table below: 
	         1 - Employees
                 2 - Companies
	`)
	fmt.Scan(&cmd)
	switch cmd {
	case 1:
		funcsforemployees.FuncsForEmployees()
	case 2:
		funcsforcompany.FuncsForCompany()
	default:
		fmt.Println("no such command exists ... ")
	}

}
