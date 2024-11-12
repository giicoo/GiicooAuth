package sqlite

import (
	"database/sql"
	"os"
	"sync"

	"github.com/giicoo/GiicooAuth/internal/config"
	"github.com/sirupsen/logrus"
)

type Sqlite struct {
	cfg *config.Config
	log *logrus.Logger
	db  *sql.DB
	m   *sync.RWMutex
}

func NewRepo(cfg *config.Config, log *logrus.Logger, db *sql.DB) *Sqlite {
	return &Sqlite{
		cfg: cfg,
		log: log,
		db:  db,
	}
}

func (sq *Sqlite) InitDB() error {
	file, err := os.ReadFile(sq.cfg.DB.PathToSQL + "create_db.sql")
	if err != nil {
		return err
	}
	stmt := string(file)
	_, err = sq.db.Exec(stmt)
	if err != nil {
		return err
	}

	return sq.db.Ping()
}
