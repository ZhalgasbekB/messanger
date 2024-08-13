package messanger

import "database/sql"

type MessangerSqlite struct {
	db *sql.DB
}

func NewMesssangerSqlite(db *sql.DB) *MessangerSqlite {
	return &MessangerSqlite{db: db}
}

// Some code with related chat
func CreateMessage() error {
	return nil
}

func DeleteMessage() error {
	return nil
}

func UpdateMessage() error {
	return nil
}

func Message() error {
	return nil
}
