// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package graph

import (
	time "time"
)

type CreateTodo struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}
type Event struct {
	ID           string    `json:"id"`
	Timestamp    time.Time `json:"timestamp"`
	Aggregate_id string    `json:"aggregate_id"`
	Data         EventData `json:"data"`
}
type EventData interface{}
type Filter struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
type TodoCreatedV1 struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	Author_name string `json:"author_name"`
}
type TodoTextUpdatedV1 struct {
	Text string `json:"text"`
}
type UpdateTodo struct {
	ID   string  `json:"id"`
	Text *string `json:"text"`
}
type User struct {
	Name string `json:"name"`
}
