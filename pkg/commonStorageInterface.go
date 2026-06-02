package CommonStorageInterface

import tsk "sf3081/hw-02/task"

type CommonStorageInterface interface {
	Tasks(int, int) ([]tsk.Task, error)
	NewTask(tsk.Task) (int, error)
}
