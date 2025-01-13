package todo

import (
	"reflect"
	"testing"
)

func TestService_Search(t *testing.T) {
	type fields struct {
		todos []Item
	}
	type args struct {
		query string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &Service{
				todos: tt.fields.todos,
			}
			if got := svc.Search(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
