package mockstorage

import tsk "sf3081/hw-02/task"

type DB []tsk.Task

func (db DB) Tasks(int, int) ([]tsk.Task, error) {
	return db, nil
}

func (db DB) NewTask(tsk.Task) (int, error) {
	return 0, nil
}
