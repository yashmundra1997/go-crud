package api

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"strconv"
	"encoding/json"
)

type Task struct{
	Name string `json:"name"`
	Id string `json:"id"`
}

func (t *Task)InsertIntoDatabase(db *sql.DB) error {
	var err error
	sql:="insert into task (task) values(?)"
	_,err=db.Exec(sql,t.Name)
	if err!=nil{
		return err
	}
	return err
	
}

func (t *Task) UpdateDatabase(db *sql.DB) error{
	var err error
	id,_:=strconv.Atoi(t.Id)
	sql:="update task set task=(?) where id =(?)"
	_,err=db.Exec(sql,t.Name,id)
	if err!=nil{
		return err
	}
	return err
}


func (t *Task) DeleteFromDatabase(db *sql.DB) error{
	var err error
	id,_:=strconv.Atoi(t.Id)
	sql:="delete from task where id =(?)"
	_,err=db.Exec(sql,id)
	if err!=nil{
		return err
	}
	return err
}


func (t *Task) RetrieveFromDatabase(db *sql.DB) []byte{
	var tasks []Task
	rows,err:=db.Query("select id,task from task")
	if err!=nil{
		panic(err.Error())
	}
	for rows.Next() {
        var id string
        var task string
        rows.Scan(&id ,&task)
        tasks = append(tasks, Task{id, task})
    }

    taskBytes, _ := json.Marshal(&tasks)
	db.Close()
    return taskBytes
	
}
