package datalayer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Knetic/go-namedParameterQuery"
	"github.com/anuragmahadik004/hr_api/models"
	"github.com/google/uuid"
)

type DLEmployee struct {
}

func (DLEmployee) GetAllEmployees() []models.Employee {

	conn, err := GetDbCOnnection()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	query := "SELECT * FROM Employee"

	rows, err := conn.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	count := 0

	var lstEmployees []models.Employee

	for rows.Next() {

		var Employee models.Employee

		err := rows.Scan(&Employee.EmployeeId,
			&Employee.EmployeeFirstName,
			&Employee.EmployeeMiddleName,
			&Employee.EmployeeLastName,
			&Employee.Email,
			&Employee.Address,
			&Employee.BirthDate,
			&Employee.PhoneNumber,
			&Employee.City,
			&Employee.State,
			&Employee.Country,
			&Employee.PreviousEmployer,
			&Employee.ExperienceInYears,
			&Employee.MaritalStatus,
			&Employee.Gender,
			&Employee.PinCode,
			&Employee.ModifiedBy,
			&Employee.CreatedBy,
			&Employee.ModifiedDate,
			&Employee.CreatedDate,
		)

		if err != nil {
			log.Fatal(err)
		}

		lstEmployees = append(lstEmployees, Employee)

		count++
	}

	if err != nil {
		log.Fatal(err)
	}

	strjson, _ := json.Marshal(lstEmployees)

	fmt.Println(string(strjson))

	return lstEmployees
}

func (DLEmployee) GetEmployee(EmployeeId string) models.Employee {

	conn, err := GetDbCOnnection()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	query := namedParameterQuery.NewNamedParameterQuery("SELECT * FROM Employee WHERE EmployeeId = :EmployeeId")

	query.SetValue("EmployeeId", EmployeeId)

	var Employee models.Employee

	err = conn.QueryRow(query.GetParsedQuery(), query.GetParsedParameters()...).Scan(&Employee.EmployeeId,
		&Employee.EmployeeFirstName,
		&Employee.EmployeeMiddleName,
		&Employee.EmployeeLastName,
		&Employee.Email,
		&Employee.Address,
		&Employee.BirthDate,
		&Employee.PhoneNumber,
		&Employee.City,
		&Employee.State,
		&Employee.Country,
		&Employee.PreviousEmployer,
		&Employee.ExperienceInYears,
		&Employee.MaritalStatus,
		&Employee.Gender,
		&Employee.PinCode,
		&Employee.ModifiedBy,
		&Employee.CreatedBy,
		&Employee.ModifiedDate,
		&Employee.CreatedDate)

	if err != nil {
		log.Fatal(err)
	}

	return Employee
}

func (DLEmployee) SaveEmployee(Employee models.Employee) bool {

	if len(Employee.EmployeeId) == 0 {
		Employee.EmployeeId = uuid.New().String()
	}

	//defining query parameters
	insQuery := `INSERT INTO [Employee]
	([EmployeeId],[EmployeeFirstName],[EmployeeMiddleName],[EmployeeLastName],[Email],[Address],[BirthDate],[PhoneNumber]
		,[PinCode],[City],[State],[Country],[PreviousEmployer],[ExperienceInYears],[MaritalStatus],[Gender],[CreatedDate],[CreatedBy]
		,[ModifiedDate],[ModifiedBy])
	VALUES(:EmployeeId, :EmployeeFirstName, :EmployeeMiddleName, :EmployeeLastName, :Email, :Address, :BirthDate, :PhoneNumber,
	:PinCode, :City, :State, :Country, :PreviousEmployer, :ExperienceInYears, :MaritalStatus, :Gender, Datediff(s, '1970-01-01', Getutcdate()), :CreatedBy, Datediff(s, '1970-01-01', Getutcdate()), :ModifiedBy)`

	updQuery := `UPDATE [Employee]
		SET   [EmployeeFirstName] = :EmployeeFirstName 
		,[EmployeeMiddleName] = :EmployeeMiddleName 
		,[EmployeeLastName] = :EmployeeLastName 
		,[Email] = :Email 
		,[Address] = :Address 
		,[BirthDate] = :BirthDate 
		,[PhoneNumber] = :PhoneNumber 
		,[PinCode] = :PinCode 
		,[City] = :City 
		,[State] = :State 
		,[Country] = :Country 
		,[PreviousEmployer] = :PreviousEmployer 
		,[ExperienceInYears] = :ExperienceInYears 
		,[MaritalStatus] = :MaritalStatus 
		,[Gender] = :Gender 
		,[ModifiedBy] = :ModifiedBy 
		,[ModifiedDate] = Datediff(s, '1970-01-01', Getutcdate())
		WHERE [EmployeeId] = :EmployeeId `

	sb := namedParameterQuery.NewNamedParameterQuery(
		`IF NOT EXISTS (SELECT * FROM Employee WHERE [EmployeeId] = :EmployeeId)` +
			` BEGIN ` +
			insQuery +
			` END ` +
			` ELSE ` +
			` BEGIN ` +
			updQuery +
			` END `,
	)

	fmt.Println(sb)

	//query params
	sb.SetValue("EmployeeId", Employee.EmployeeId)
	sb.SetValue("EmployeeFirstName", Employee.EmployeeFirstName)
	sb.SetValue("EmployeeMiddleName", Employee.EmployeeMiddleName)
	sb.SetValue("EmployeeLastName", Employee.EmployeeLastName)
	sb.SetValue("Email", Employee.Email)
	sb.SetValue("Address", Employee.Address)
	sb.SetValue("BirthDate", Employee.BirthDate)
	sb.SetValue("PhoneNumber", Employee.PhoneNumber)
	sb.SetValue("PinCode", Employee.PinCode)
	sb.SetValue("City", Employee.City)
	sb.SetValue("State", Employee.State)
	sb.SetValue("Country", Employee.Country)
	sb.SetValue("PreviousEmployer", Employee.PreviousEmployer)
	sb.SetValue("ExperienceInYears", Employee.ExperienceInYears)
	sb.SetValue("MaritalStatus", Employee.MaritalStatus)
	sb.SetValue("Gender", Employee.Gender)
	// sb.SetValue("CreatedDate", time.Now().Unix())
	sb.SetValue("CreatedBy", Employee.CreatedBy)
	// sb.SetValue("ModifiedDate", time.Now().Unix())
	sb.SetValue("ModifiedBy", Employee.ModifiedBy)

	conn, err := GetDbCOnnection()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	result, err := conn.Exec(sb.GetParsedQuery(), sb.GetParsedParameters()...)

	if err != nil {
		fmt.Println(err)
	}

	count, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		return true
	}

	return false
}

func (DLEmployee) DeleteEmployee(EmployeeId string) bool {

	conn, err := GetDbCOnnection()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	query := namedParameterQuery.NewNamedParameterQuery("DELETE Employee WHERE EmployeeId = :EmployeeId")

	query.SetValue("EmployeeId", EmployeeId)

	result, err := conn.Exec(query.GetParsedQuery(), query.GetParsedParameters()...)

	if err != nil {
		log.Fatal(err)
	}

	count, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		return true
	}

	return false
}
