package server

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"AddEmployee", http.MethodPost, "/employees", AddEmployee},
	Route{"ListEmployees", http.MethodGet, "/employees", ListEmployees},
	Route{"GetEmployee", http.MethodGet, "/employees/{eid}", GetEmployee},
	Route{"UpdateEmployee", http.MethodPut, "/employees/{eid}", UpdateEmployee},
	Route{"RemoveEmployee", http.MethodDelete, "/employees/{eid}", RemoveEmployee},
}
