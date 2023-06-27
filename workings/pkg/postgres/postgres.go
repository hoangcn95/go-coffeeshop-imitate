package postgres

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
)

const (
	_defaultConnAttempts = 3
	_defaultConnTimeout  = time.Second
)

type DBConnString string

type postgres struct {
	connAttempts int
	connTimeout  time.Duration
	db           *sql.DB
}

func NewPostgresDB(url DBConnString) (DBEngine, error) {
	slog.Info("CONN", "connect string", url)

	pg := &postgres{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	var err error
	for pg.connAttempts > 0 {
		pg.db, err = sql.Open("postgres", string(url))
		if err == nil {
			slog.Info("📰 connected to postgresdb 🎉")
			return pg, nil
		}

		log.Printf("Postgres is trying to connect, attempts left: %d, err: %v", pg.connAttempts, err)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	return nil, err
}

func (p *postgres) Configure(opts ...Option) DBEngine {
	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *postgres) GetDB() *sql.DB {
	return p.db
}

func (p *postgres) Close() {
	if p.db != nil {
		p.db.Close()
	}
}
