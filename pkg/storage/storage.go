package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"

	"sf3081/hw-02/models"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

func (s *Storage) GetTasks(taskID, authorID int) ([]models.Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE
			($1 = 0 OR id = $1) AND
			($2 = 0 OR author_id = $2)
		ORDER BY id;
	`,
		taskID,
		authorID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) AddTask(t models.Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, content, author_id, assigned_id, opened, closed)
		VALUES ($1, $2, $3, $4, COALESCE($5, extract(epoch from now())), $6) 
		RETURNING id;
		`,
		t.Title,
		t.Content,
		t.AuthorID,
		t.AssignedID,
		t.Opened,
		t.Closed,
	).Scan(&id)
	return id, err
}

func (s *Storage) UpdateTask(t models.Task) error {
	result, err := s.db.Exec(context.Background(), `
		UPDATE tasks 
		SET 
			title = $1,
			content = $2,
			assigned_id = $3,
			closed = $4
		WHERE id = $5
		`,
		t.Title,
		t.Content,
		t.AssignedID,
		t.Closed,
		t.ID,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (s *Storage) DeleteTaskByID(id int) error {
	result, err := s.db.Exec(context.Background(), `
		DELETE FROM tasks WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (s *Storage) GetTaskByID(id int) (*models.Task, error) {
	var t models.Task
	err := s.db.QueryRow(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE id = $1
		`,
		id,
	).Scan(
		&t.ID,
		&t.Opened,
		&t.Closed,
		&t.AuthorID,
		&t.AssignedID,
		&t.Title,
		&t.Content,
	)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Storage) GetTasksByAssigneeID(assignedID int) ([]models.Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE assigned_id = $1
		ORDER BY opened DESC
		`,
		assignedID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) AddLabelToTask(taskID, labelID int) error {
	_, err := s.db.Exec(context.Background(), `
		INSERT INTO tasks_labels (task_id, label_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
		`,
		taskID,
		labelID,
	)
	return err
}

func (s *Storage) RemoveLabelFromTask(taskID, labelID int) error {
	_, err := s.db.Exec(context.Background(), `
		DELETE FROM tasks_labels 
		WHERE task_id = $1 AND label_id = $2
		`,
		taskID,
		labelID,
	)
	return err
}

func (s *Storage) GetTaskLabelsByTaskID(taskID int) ([]string, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT l.name
		FROM labels l
		JOIN tasks_labels tl ON l.id = tl.label_id
		WHERE tl.task_id = $1
		ORDER BY l.name
		`,
		taskID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var labels []string
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		labels = append(labels, name)
	}
	return labels, rows.Err()
}

func (s *Storage) GetUsers() ([]models.User, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT id, name FROM users ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (s *Storage) AddUser(name string) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO users (name) VALUES ($1) RETURNING id
		`,
		name,
	).Scan(&id)
	return id, err
}

func (s *Storage) GetLabels() ([]models.Label, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT id, name FROM labels ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var labels []models.Label
	for rows.Next() {
		var l models.Label
		err = rows.Scan(&l.ID, &l.Name)
		if err != nil {
			return nil, err
		}
		labels = append(labels, l)
	}
	return labels, rows.Err()
}

func (s *Storage) Close() {
	if s.db != nil {
		s.db.Close()
	}
}
