package routers

import (
	"encoding/json"
	"net/http"

	"github.com/anuragmahadik004/hr_api/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/employee/getemployees", controllers.GetAllEmployees).Methods("GET")
	router.HandleFunc("/api/employee/getemployee/{empId}", controllers.GetEmployee).Methods("GET")
	router.HandleFunc("/api/employee/saveemployee", controllers.SaveEmployee).Methods("POST")
	router.HandleFunc("/api/employee/deleteemployee/{empId}", controllers.DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/api/auth/loginuser", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/api/auth/saveuser", controllers.SaveUser).Methods("POST")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode("HRMS System")

	}).Methods("GET")

	return router

}
