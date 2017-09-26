package main

import "gopkg.in/mgo.v2/bson"

type TodoDAC struct{}

func (TodoDAC) getall() ([]Todo, error) {
	db, session := connectdb()
	defer session.Close()
	todos := db.C("todos")
	result := []Todo{}
	err := todos.Find(nil).All(&result)

	return result, err
}

func (TodoDAC) getbyid(id string) (Todo, error) {
	db, session := connectdb()
	defer session.Close()
	todos := db.C("todos")
	result := Todo{}
	err := todos.Find(bson.M{"_id": id}).One(&result)

	return result, err
}
