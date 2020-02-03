package models

import (
	"gopkg.in/mgo.v2"
)

type Repository struct {
	Session *mgo.Session
}

func (r *Repository) collection() *mgo.Collection {
	return r.Session.DB("useful_links").C("links")
}

func (r *Repository) GetAll() (links []*Links, err error) {
	err = r.collection().Find(nil).All(&links)
	return links, err
}
