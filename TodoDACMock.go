package main

import "errors"

type TodoDACMock struct{}

var mockData = []Todo{
	Todo{ID: "1", Name: "jokerstudio"},
	Todo{ID: "2", Name: "ploy"},
}

func (TodoDACMock) getall() ([]Todo, error) {
	return mockData, nil
}

func (TodoDACMock) getbyid(id string) (Todo, error) {
	result := Todo{}
	for _, todo := range mockData {
		if todo.ID == id {
			result = todo
			return result, nil
		}
	}
	err := errors.New("not found")
	return result, err
}
