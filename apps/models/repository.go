package models

import (
	"gopkg.in/mgo.v2"
)

type Repository struct {
	Session *mgo.Session
}

// collection получаем коллекцию useful_links
func (r *Repository) collection() *mgo.Collection {
	return r.Session.DB("useful_links").C("links")
}

// GetAll получить все ссылки
func (r *Repository) GetAll() (links []*Link, err error) {
	err = r.collection().Find(nil).Sort("link_group", "id").All(&links)
	return links, err
}

// GetById получить все ссылки
func (r *Repository) GetById(id int64) (link *Link, err error) {
	err = r.collection().FindId(id).One(&link)
	return link, err
}

// CreateLink создать запись
func (r *Repository) CreateLink(l *Link) error {
	return r.collection().Insert(l)
}

// UpdateLink обновить запись
func (r *Repository) UpdateLink(l *Link) error {
	return r.collection().UpdateId(l.ID, l)
}

// DeleteLink удалить ссылку
func (r *Repository) DeleteLink(id int64) (err error) {
	err = r.collection().RemoveId(id)
	return err
}

// GetGroups получить все группы
func (r *Repository) GetGroups() (groups []string, err error) {
	err = r.collection().Find(nil).Distinct("link_group", &groups)
	return groups, err
}
