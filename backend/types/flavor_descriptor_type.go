package types

type FlavorDescriptor struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
