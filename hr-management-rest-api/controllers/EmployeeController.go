package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/anuragmahadik004/hr_api/datalayer"
	"github.com/anuragmahadik004/hr_api/interfaces"
	"github.com/anuragmahadik004/hr_api/models"
	"github.com/gorilla/mux"
)

//Dependency Injection
var _EmployeeRepository interfaces.EmployeeRepository

// var _IEmployee interfaces.IEmployee

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	jwtToken := r.Header.Get("jwttoken")

	jwtTokenIsValid := false

	if len(jwtToken) == 0 {
		json.NewEncoder(w).Encode("User not logged in")
	} else {

		jwtTokenIsValid = VerifyJWTToken(jwtToken)

		fmt.Println(jwtTokenIsValid)

		if jwtTokenIsValid {

			lstEmployees := _EmployeeRepository.GetAllEmployees()

			json.NewEncoder(w).Encode(lstEmployees)

		} else {

			json.NewEncoder(w).Encode("User not logged in")

		}
	}

}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	jwtToken := r.Header.Get("jwttoken")

	jwtTokenIsValid := false

	if len(jwtToken) == 0 {
		json.NewEncoder(w).Encode("User not logged in")
	} else {
		jwtTokenIsValid = VerifyJWTToken(jwtToken)

		if jwtTokenIsValid {

			params := mux.Vars(r)

			Employee := _EmployeeRepository.GetEmployee(params["empId"])

			json.NewEncoder(w).Encode(Employee)

		} else {

			json.NewEncoder(w).Encode("User not logged in")

		}
	}

}

func SaveEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	jwtToken := r.Header.Get("jwttoken")

	jwtTokenIsValid := false

	if len(jwtToken) == 0 {
		json.NewEncoder(w).Encode("User not logged in")
	} else {
		jwtTokenIsValid = VerifyJWTToken(jwtToken)

		if jwtTokenIsValid {

			var Employee models.Employee

			_ = json.NewDecoder(r.Body).Decode(&Employee)

			fmt.Println(Employee)

			EmployeeSaved := _EmployeeRepository.SaveEmployee(Employee)

			json.NewEncoder(w).Encode(EmployeeSaved)

		} else {

			json.NewEncoder(w).Encode("User not logged in")

		}
	}

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	jwtToken := r.Header.Get("jwttoken")

	jwtTokenIsValid := false

	if len(jwtToken) == 0 {
		json.NewEncoder(w).Encode("User not logged in")
	} else {
		jwtTokenIsValid = VerifyJWTToken(jwtToken)

		if jwtTokenIsValid {

			params := mux.Vars(r)

			EmployeeId := params["empId"]

			EmployeeDeleted := _EmployeeRepository.DeleteEmployee(EmployeeId)

			json.NewEncoder(w).Encode(EmployeeDeleted)

		} else {

			json.NewEncoder(w).Encode("User not logged in")

		}
	}

}
