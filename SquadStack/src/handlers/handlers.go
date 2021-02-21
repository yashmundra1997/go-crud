package handlers

import(
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"api"
)
type App struct {
	Router *mux.Router
	DB *sql.DB
}

type Task struct{
	Name string `json:"name"`
	Id string `json:"id"`
}



func  (a *App) CreateTask(rw http.ResponseWriter,r *http.Request){
	t:=api.Task{}
	var err error
	err = json.NewDecoder(r.Body).Decode(&t)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusBadRequest)
        return
    }
	err=t.InsertIntoDatabase(a.DB)
	if err!=nil{
		http.Error(rw,err.Error(),http.StatusBadRequest)
	}
	fmt.Fprintf(rw,"Inserted",http.StatusOK)
}


func (a *App) UpdateTask(rw http.ResponseWriter, r *http.Request){
	t:=api.Task{}
	var err error
	err=json.NewDecoder(r.Body).Decode(&t)
	if err!=nil{
		http.Error(rw,err.Error(),http.StatusBadRequest)
		return 
	}
	err=t.UpdateDatabase(a.DB)
	if err!=nil{
		http.Error(rw,err.Error(),http.StatusBadRequest)
		return 
	}
	fmt.Fprintf(rw,"Updated",http.StatusOK)
}


func (a *App) DeleteTask(rw http.ResponseWriter,r *http.Request){
	t:= api.Task{}
	var err error
	err=json.NewDecoder(r.Body).Decode(&t)
	if err!=nil{
		http.Error(rw,err.Error(),http.StatusBadRequest)
		return 
	}
	err=t.DeleteFromDatabase(a.DB)
	if err!=nil{
		http.Error(rw,err.Error(),http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw,"Deleted",http.StatusOK)
}


func (a *App) RetrieveData(rw http.ResponseWriter,r *http.Request){
	t:=api.Task{}
	tasks:=t.RetrieveFromDatabase(a.DB)

    //taskBytes, _ := json.Marshal(&tasks)

    rw.Write(tasks)
    a.DB.Close()
	
}