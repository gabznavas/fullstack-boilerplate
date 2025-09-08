package dtos

type CreateTodoRequest struct {
	Title string `json:"title"`
}

type CreateTodoResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
