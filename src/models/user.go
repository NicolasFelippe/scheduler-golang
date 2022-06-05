package models

type User struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"nome,omitempty"`
	CompanyId uint64 `json:"empresa_id,omitempty"`
}
