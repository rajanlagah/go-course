package db

import (
	"context"
	"time"
)

type Task struct {}

var TaskRepository = Task{}

type PostTaskPayload struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status string `json:"status"`
}


func (t Task) SaveTaskQuery(payload PostTaskPayload) (int, error){
	var id int;

	query := `Insert into tasks (title, content, status) VALUES ($1, $2, $3) RETURNING id;`

	err := DB.QueryRow(context.Background(), query, payload.Title, payload.Content, payload.Status).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

type TaskType struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (t Task) ReadTaskQuery() ([]TaskType, error){
	var tasks []TaskType;

	query := `Select id, title, content, status, created_at FROM tasks ORDER BY created_at DESC LIMIT 10;`

	rows, err := DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for  rows.Next() {
		var item TaskType;
		err := rows.Scan(&item.ID, &item.Title, &item.Content, &item.Status, &item.CreatedAt)

		if err != nil {
			return nil, err
		}
		tasks = append(tasks, item)
	}

	return tasks, nil
}

type UpdateTaskPayload struct {
	ID int `json:"id" binding:"required"`
	Title string `json:"title" binding:"max=100"`
	Content string `json:"content" binding:"max=1000"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (t Task) UpdateTask(payload UpdateTaskPayload) error {
	query := `
		UPDATE tasks
		SET title = $1, content = $2, status = $3 
		WHERE id = $4
	`

  _, err := DB.Exec(context.Background(), query, payload.Title,payload.Content, payload.Status, payload.ID)
  return err
}


func (t Task) GetTaskById(id int) (TaskType, error) {

	var task TaskType

	query := `Select id, title, content, status, created_at from tasks where id=$1`

	err := DB.QueryRow(context.Background(), query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Content,
		&task.Status,
		&task.CreatedAt,
	)
	if err != nil {
		return TaskType{}, err
	}

	return task, nil
}

func (t Task) DeleteTaskQuery(id int) error {

	query := ` Delete from tasks WHERE id = $1`

	_, err := DB.Exec(context.Background(), query, id)

	return err
}