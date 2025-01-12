package todo

import "errors"

type TodoItem struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

type Item struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}

type Service struct {
	todos []Item
}

func NewService() *Service {
	return &Service{
		todos: make([]Item, 0),
	}
}

/* Adding logics that we will need */
func (svc *Service) Add(todo string) error {
	for _, t := range svc.todos {
		if t.Task == todo {
			return errors.New("todo is not unique")
		}
	}
	svc.todos = append(svc.todos, Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	})
	return nil
}

func (svc *Service) GetAll() []Item {
	return svc.todos
}

func (svc *Service) Remove(index int) {
	allTodos := svc.todos
	allTodos = append(allTodos[:index], allTodos[index+1:]...)
}
