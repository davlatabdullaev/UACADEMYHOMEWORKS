package funcsforcompany

import (
	connectdb "basa-company/ConnectDB"
	structforcompany "basa-company/StructForCompany"
	"database/sql"
	"fmt"
	"log"
	_"github.com/lib/pq"
)

type BasaCompany struct {
	db *sql.DB
}

func New(db *sql.DB) BasaCompany {
	return BasaCompany{
		db: db,
	}
}

func FuncsForCompany() {
	db, err := connectdb.ConnectDB()
	if err != nil {
		log.Fatal("error during connect database ", err)
	}
	defer db.Close()
	DataBaseCompany := New(db)
	command := 0
	fmt.Println("Please select command: ")
	fmt.Print(`
	1 - insert company
	2 - get company by id
	3 - get all companies
	4 - update company by id
	5 - delete company by id
	`)
	fmt.Scan(&command)
	switch command {
	case 1:
		// insert new company
		if err := DataBaseCompany.InsertCompany(); err != nil {
			log.Fatal("error during insert new company... ", err)
		}
	case 2:
		// get company by id
		GetCompany, err := DataBaseCompany.GetCompanyById()
		if err != nil {
			log.Fatal("error during get company by id... ", err)
		} else {
			fmt.Println(GetCompany)
		}
	case 3:
		// get all companies
		AllCompanies, err := DataBaseCompany.GetAllCompanies()
		if err!= nil{
			log.Fatal("error during get all companies ... ", err)
		} else {
			fmt.Println(AllCompanies)
		}
	case 4:
		// update company
	if err := DataBaseCompany.UpdateCompany(); err!=nil{
		log.Fatal("error during update company...", err)
	}
	case 5:
		//delete company
	if err :=	DataBaseCompany.DeleteCompany(); err != nil{
		log.Fatal("error during delete company...", err)
	}
	default:
		fmt.Println("no such command exists")
	}
}

func (b BasaCompany) InsertCompany() error {
	data := structforcompany.ForInsertCompany()
	if _, err := b.db.Exec(`insert into company (company_name, employees_count) values ($1, $2)`, data.CompanyName, data.EmployeesCount); err != nil {
		return err
	}
	log.Fatal("insert company completed succesfully...")
	return nil
}

func (b BasaCompany) GetCompanyById() (structforcompany.Company, error) {
	get := structforcompany.ForGetCompanyByID()
	row := b.db.QueryRow(`select id, company_name, employees_count from company where id = $1`, get.ID)
	GetCompany := structforcompany.Company{}
	if err := row.Scan(&GetCompany.ID, &GetCompany.CompanyName, &GetCompany.EmployeesCount); err != nil {
		return structforcompany.Company{}, err
	}
	return GetCompany, nil
}

func (b BasaCompany) GetAllCompanies() ([]structforcompany.Company, error) {
	rows, err := b.db.Query(`select id, company_name, employees_count from company`)
	if err!=nil{
		log.Fatal("error during get all companies ", err)
	}
	AllCompanies := []structforcompany.Company{}

	for rows.Next() {
       OneCompany := structforcompany.Company{}
	if err :=   rows.Scan(&OneCompany.ID, &OneCompany.CompanyName, &OneCompany.EmployeesCount); err!=nil{
		return nil, err
	}
	AllCompanies = append(AllCompanies, OneCompany)
	}
	return AllCompanies, nil
}

func (b BasaCompany) UpdateCompany() error{
	updating:=structforcompany.ForUpdateCompanyById()
 if _, err :=   b.db.Exec(`update company set company_name = $1, employees_count = $2 where id = $3`, updating.CompanyName, updating.EmployeesCount, updating.ID); err!=nil{
	return err
 }
fmt.Println("update company completed succesfully...")
return nil
}

func (b BasaCompany) DeleteCompany() error{
	deleting := structforcompany.ForGetCompanyByID()
	if _, err := b.db.Exec(`delete from company where id = $1`, deleting.ID); err!=nil{
		return err
	}
	fmt.Println("deleting completed succesfully...")
	return nil
}

