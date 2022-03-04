package book

type CreateBookDTO struct {
	Name   string `json:"name,omitempty"`
	Year   int    `json:"year,omitempty"`
	Author string `json:"author,omitempty"`
}

type UpdateBookDTO struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Author string `json:"author"`
	Busy   bool   `json:"busy"`
	Owner  string `json:"owner"`
}
