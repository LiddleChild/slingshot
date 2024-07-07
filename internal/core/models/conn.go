package models

type Connection struct {
	Name string `db:"name:VARCHAR(255) PRIMARY KEY"`
	Url  string `db:"url:TEXT"`
}
