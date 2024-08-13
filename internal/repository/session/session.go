package session

import (
	"database/sql"

	"forum/internal/models"
)

type SessionSqlite struct {
	db *sql.DB
}

func NewSessionSqlite(db *sql.DB) *SessionSqlite {
	return &SessionSqlite{db: db}
}

func (r *SessionSqlite) Create(session *models.Session) error {
	query := "INSERT INTO sessions (user_id, uuid, expire_at) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, session.User_id, session.UUID, session.ExpireAt)
	return err
}

func (r *SessionSqlite) DeleteByUUID(uuid string) error {
	query := "DELETE FROM sessions WHERE uuid = $1"
	_, err := r.db.Exec(query, uuid)
	return err
}

func (r *SessionSqlite) GetByUserId(userId int) (*models.Session, error) {
	session := models.Session{}
	query := "SELECT * FROM sessions WHERE user_id = $1"
	err := r.db.QueryRow(query, userId).Scan(&session.User_id, &session.UUID, &session.ExpireAt)
	return &session, err
}

func (r *SessionSqlite) GetByUUID(uuid string) (*models.Session, error) {
	session := models.Session{}
	query := "SELECT * FROM sessions WHERE uuid = $1"
	err := r.db.QueryRow(query, uuid).Scan(&session.User_id, &session.UUID, &session.ExpireAt)
	return &session, err
}
