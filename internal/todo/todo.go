package todo

import "errors"

type TodoItem struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

type Service struct {
	todos []TodoItem
}

func NewService() *Service {
	return &Service{
		todos: make([]TodoItem, 0),
	}
}

/* Adding logics that we will need */
func (svc *Service) Add(todo TodoItem) error {
	for _, t := range svc.GetAll() {
		if t.Item == todo.Item {
			return errors.New("todo is not unique")
		}
	}
	svc.todos = append(svc.todos, todo)
	return nil
}

func (svc *Service) GetAll() []TodoItem {
	return svc.todos
}

func (svc *Service) Remove(index int) {
	allTodos := svc.todos
	allTodos = append(allTodos[:index], allTodos[index+1:]...)
}
