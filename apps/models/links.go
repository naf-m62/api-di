package models

type Links struct {
	ID int64 `json:"id" bson:"_id"`
	Url string `json:"url" bson:"url"`
	Description string `json:"description" bson:"description"`
}
