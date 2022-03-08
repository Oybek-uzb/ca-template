package author

type CreateAuthorDTO struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type UpdateAuthorDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
