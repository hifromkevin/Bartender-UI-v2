package types

type Tool struct {
	ID       string `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	ImageURL string `db:"image_url" json:"image_url"`
}
