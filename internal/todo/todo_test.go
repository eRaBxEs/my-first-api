package todo_test

import (
	"my-first-api/internal/todo"
	"reflect"
	"testing"
)

func TestService_Search(t *testing.T) {

	tests := []struct {
		name           string
		todosToAdd     []string
		query          string
		expectedResult []string
	}{
		{
			name:           "given a todo of shop and a search of sh, I should get shop back",
			todosToAdd:     []string{"shop"},
			query:          "sh",
			expectedResult: []string{"shop"},
		},
		{
			name:           "still returns shop even if the case doesn't match",
			todosToAdd:     []string{"Shopping"},
			query:          "sh",
			expectedResult: []string{"Shopping"},
		},
		{
			name:           "spaces",
			todosToAdd:     []string{"go Shopping"},
			query:          "go",
			expectedResult: []string{"go Shopping"},
		},
		{
			name:           "space at the start of a word",
			todosToAdd:     []string{" Space at the beginning"},
			query:          "space",
			expectedResult: []string{" Space at the beginning"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := todo.NewService()
			for _, toAdd := range tt.todosToAdd {
				err := svc.Add(toAdd)
				if err != nil {
					t.Error(err)
				}
			}
			if got := svc.Search(tt.query); !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}
