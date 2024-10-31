package models

type Todo struct {
    ID          string  `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
}

var Todos = []Todo{
    {ID: "1", Title: "Foo", Description: "Bar"},
    {ID: "2", Title: "Just testing", Description: "Another test"},
    {ID: "3", Title: "A new todo", Description: "testing"},
}
