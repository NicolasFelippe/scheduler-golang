package models

type Company struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"nome,omitempty"`
}
