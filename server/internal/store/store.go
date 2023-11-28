package store

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type InterfaceStore interface {
	Open()
	Close()
	Users() UserService
	Chats() ChatService
	Messages() MessageService
}

type store struct {
	db *sqlx.DB
}

func (s *store) Open() {
	database, err := sqlx.Open("pgx", "postgres://notes_db_sebp_user:mGA7RJ0scHFknI60vZ6RGkQmaEovnTNU@dpg-clgdq9ef27hc739jfplg-a.frankfurt-postgres.render.com/notes_db_sebp")
	if err != nil {
		panic(err)
	}
	s.db = database
}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) Users() UserService {
	return &user{db: s.db}
}

func (s *store) Chats() ChatService {
	return &chat{db: s.db}
}

func (s *store) Messages() MessageService {
	return &message{db: s.db}
}

func NewStore() InterfaceStore {
	return &store{}
}
