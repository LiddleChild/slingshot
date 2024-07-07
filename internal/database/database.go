package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/LiddleChild/slingshot/internal/core/models"
	_ "github.com/mattn/go-sqlite3"
)

func NewDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "slingshot.db")
	if err != nil {
		return nil, err
	}

	err = migrate(db, "connections", models.Connection{})
	if err != nil {
		return nil, err
	}

	return db, err
}

func migrate(db *sql.DB, table string, v interface{}) error {
	t := reflect.TypeOf(v)

	fields := []string{}
	for i := range t.NumField() {
		tag := t.Field(i).Tag.Get("db")
		if len(strings.Split(tag, ":")) < 2 {
			return fmt.Errorf("invalid db tag: %s", tag)
		}

		fields = append(fields, strings.ReplaceAll(tag, ":", " "))
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	createQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", table, strings.Join(fields, ","))
	_, err = tx.Exec(createQuery)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
