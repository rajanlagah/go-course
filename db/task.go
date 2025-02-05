package db

import "context"

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