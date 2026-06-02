package main

import (
	"fmt"

	dbInterface "sf3081/hw-02/commonStorageInterface"
	// ms "sf3081/hw-02/mockStorage"
	str "sf3081/hw-02/storage"
	tsk "sf3081/hw-02/task"
)

var db dbInterface.CommonStorageInterface

func main() {
	t := tsk.Task{ID: 1, Opened: 0, Closed: 0, AuthorID: 1, AssignedID: 2, Title: "ttl", Content: "text"}

	db, err := str.New("postgresql://postgres:Qwerty123@localhost:5432/tasks")
	// db := ms.DB{}

	tr, err := db.NewTask(t)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tr)
}
