package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"handlers"
)


func main(){
	a:=handlers.App{}
	var err error
	a.Router =mux.NewRouter()
	a.DB, err = sql.Open("mysql", "root:mundra@tcp(127.0.0.1:3306)/squadstack")
	if err!=nil{
		panic(err.Error())
	}
	a.Router.HandleFunc("/insert",a.CreateTask).Methods("POST")
	a.Router.HandleFunc("/update",a.UpdateTask).Methods("POST")
	a.Router.HandleFunc("/delete",a.DeleteTask).Methods("POST")
	a.Router.HandleFunc("/retrieve",a.RetrieveData).Methods("GET")
	http.ListenAndServe(":8080",a.Router)
}