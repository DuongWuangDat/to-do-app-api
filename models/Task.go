package models

var Task struct {
	ID        string `bson:"_id,omitempty" json:"_id,omitempty"`
	Title     string `bson:"title" json:"title"`
	IsDone    bool   `bson:"isdone" json:"isdone"`
	CreatedAt int    `bson:"createdat" json:"createdat"`
}
