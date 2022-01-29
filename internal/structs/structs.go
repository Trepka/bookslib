package structs

type Book struct {
	ID     string `json:"book-id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
}
