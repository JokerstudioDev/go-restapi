package main

import (
	"testing"
)

var tododacMock = TodoDACMock{}
var todoHandlertest = NewTodoHandler(tododacMock)

func TestGetAllTodos(t *testing.T) {
	_, err := todoHandlertest.GetAllTodos()
	if err != nil {
		t.Error(err)
	}
}

func TestGetTodo(t *testing.T) {
	_, err := todoHandlertest.GetTodo("3")
	if err != nil {
		t.Error(err)
	}
}
