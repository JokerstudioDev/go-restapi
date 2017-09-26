package main

import (
	"gopkg.in/mgo.v2"
)

func connectdb() (*mgo.Database, *mgo.Session) {
	session, err := mgo.Dial("mongodb://jokerstudio:Jk065807588@ds149124.mlab.com:49124/godb")
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("godb")
	return db, session
}
