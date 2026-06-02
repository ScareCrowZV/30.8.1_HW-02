package main

import (
	"fmt"

	dbInterface "sf3081/hw-02/commonStorageInterface"
	"sf3081/hw-02/models"
	"sf3081/hw-02/storage"
)

var db dbInterface.CommonStorageInterface

func main() {

	storage, err := storage.New("postgresql://postgres:Qwerty123@localhost:5432/tasks")
	if err != nil {
		fmt.Println("Ошибка подключения к tasks:", err)
		return
	}
	defer storage.Close()

	db = storage

	userID, err := db.AddUser("Голубев Иван")
	if err != nil {
		fmt.Println("Ошибка создания пользователя:", err)
	} else {
		fmt.Println("Создан пользователь ID:", userID)
	}

	newTask := models.Task{
		Title:      "Нарезать хлебушек и колбасу",
		Content:    "Хлеб в холодильнике, а колбаса в магазине",
		AuthorID:   userID,
		AssignedID: userID,
		Opened:     0,
		Closed:     0,
	}

	taskID, err := db.AddTask(newTask)
	if err != nil {
		fmt.Println("Ошибка создания задачи:", err)
	} else {
		fmt.Println("Создана задача ID:", taskID)
	}

	users, err := db.GetUsers()
	if err != nil {
		fmt.Println("Ошибка получения пользователей:", err)
	} else {
		fmt.Println("\nПользователи:")
		for _, u := range users {
			fmt.Printf("  %d. %s\n", u.ID, u.Name)
		}
	}

	labels, err := db.GetLabels()
	if err != nil {
		fmt.Println("Ошибка получения меток:", err)
	} else {
		fmt.Println("\nМетки:")
		for _, l := range labels {
			fmt.Printf("  %d. %s\n", l.ID, l.Name)
		}
	}

	tasks, err := db.GetTasks(0, 0)
	if err != nil {
		fmt.Println("Ошибка получения задач:", err)
	} else {
		fmt.Println("\nЗадачи:")
		for _, t := range tasks {
			fmt.Printf("  %d. %s (автор: %d, исполнитель: %d)\n", t.ID, t.Title, t.AuthorID, t.AssignedID)
		}
	}
}
