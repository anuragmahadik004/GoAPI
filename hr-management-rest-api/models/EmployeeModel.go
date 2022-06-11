package models

type Employee struct {
	EmployeeId         string        `json:"EmployeeId"`
	EmployeeFirstName  string        `json:"EmployeeFirstName"`
	EmployeeMiddleName string        `json:"EmployeeMiddleName"`
	EmployeeLastName   string        `json:"EmployeeLastName"`
	Email              string        `json:"Email"`
	Address            string        `json:"Address"`
	BirthDate          string        `json:"BirthDate"`
	PhoneNumber        string        `json:"PhoneNumber"`
	PinCode            string        `json:"PinCode"`
	City               string        `json:"City"`
	State              string        `json:"State"`
	Country            string        `json:"Country"`
	PreviousEmployer   string        `json:"PreviousEmployer"`
	ExperienceInYears  string        `json:"ExperienceInYears"`
	MaritalStatus      MaritalStatus `json:"MaritalStatus"`
	Gender             Gender        `json:"Gender"`
	CreatedBy          string        `json:"CreatedBy"`
	CreatedDate        string        `json:"CreatedDate"`
	ModifiedBy         string        `json:"ModifiedBy"`
	ModifiedDate       string        `json:"ModifiedDate"`
}

//defining enums
type MaritalStatus int

const (
	Single   MaritalStatus = 1
	Married  MaritalStatus = 2
	Divorced MaritalStatus = 3
)

type Gender int

const (
	Male   Gender = 1
	Female Gender = 2
	Other  Gender = 3
)
