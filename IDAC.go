package main

type ITodoDAC interface {
	getall() ([]Todo, error)
	getbyid(id string) (Todo, error)
}
