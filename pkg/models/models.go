package models

// Task - модель задачи
type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

// User - модель пользователя
type User struct {
	ID   int
	Name string
}

// Label - модель метки
type Label struct {
	ID   int
	Name string
}
