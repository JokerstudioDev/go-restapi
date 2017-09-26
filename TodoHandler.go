package main

// TodoHandler : todo handler object
type TodoHandler struct {
	tododac ITodoDAC
}

// NewTodoHandler : create todo handler
func NewTodoHandler(tododac ITodoDAC) *TodoHandler {
	return &TodoHandler{tododac}
}

// GetAllTodos :  get all todos
func (todohandler *TodoHandler) GetAllTodos() ([]Todo, error) {
	return todohandler.tododac.getall()
}

// GetTodo :  get all todos
func (todohandler *TodoHandler) GetTodo(id string) (Todo, error) {
	return todohandler.tododac.getbyid(id)
}
