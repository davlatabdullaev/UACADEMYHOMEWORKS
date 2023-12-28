package trigger

import (
	"database/sql"
	"fmt"
	"log"
)

type TriggerBasa struct {
	db *sql.DB
}

func New(db *sql.DB) TriggerBasa {
	return TriggerBasa{
		db: db,
	}
}

func (t TriggerBasa) FuncsForTrigger() error {
	_, err := t.db.Exec(`
	CREATE OR REPLACE FUNCTION update_company_employees_count()
	RETURNS TRIGGER AS $$
	BEGIN
		UPDATE company
		SET employees_count = employees_count + 1
		WHERE id = NEW.company_id;
		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

	CREATE TRIGGER update_employees_count_trigger
	AFTER INSERT ON employees
	FOR EACH ROW
	EXECUTE FUNCTION update_company_employees_count();
	`)
	return err
}

func (t TriggerBasa) Trigger() {
	rows, err := t.db.Query("SELECT * FROM company")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id, employeesCount int
		var companyName string
		err := rows.Scan(&id, &companyName, &employeesCount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Company Name: %s, Employees Count: %d\n", id, companyName, employeesCount)
	}
}
