package CommonStorageInterface

import "sf3081/hw-02/models"

type CommonStorageInterface interface {
	// Работа с задачами
	GetTasks(taskID, authorID int) ([]models.Task, error)
	AddTask(models.Task) (int, error)
	UpdateTask(models.Task) error
	DeleteTaskByID(id int) error
	GetTaskByID(id int) (*models.Task, error)
	GetTasksByAssigneeID(assignedID int) ([]models.Task, error)

	AddLabelToTask(taskID, labelID int) error
	RemoveLabelFromTask(taskID, labelID int) error
	GetTaskLabelsByTaskID(taskID int) ([]string, error)
	GetLabels() ([]models.Label, error)

	GetUsers() ([]models.User, error)
	AddUser(name string) (int, error)

	Close()
}
