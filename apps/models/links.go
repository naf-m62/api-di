package models

type Link struct {
	ID          int64  `json:"id" bson:"_id"`
	Url         string `json:"url" bson:"url"`
	LinkGroup   string `json:"link_group" bson:"link_group"`
	Description string `json:"description" bson:"description"`
}

type LinksByGroup struct {
	Name  string  `json:"name"`
	Count int64   `json:"count"`
	Links []*Link `json:"links"`
}
