package model

type ID struct {
	ID uint32 `gorm:"primary_key" json:"id"`
}

type CreatedOn struct {
	CreatedOn uint32 `json:"created_on"`
}

type ModifiedOn struct {
	ModifiedOn uint32 `json:"modified_on"`
}

type Model struct {
	ID
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn
	ModifiedOn
	DeletedOn uint32 `json:"deleted_on"`
}
