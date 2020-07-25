package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	// Se accede a:
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// NewPostgresDB realiza la conexión con la base de datos
func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "user=postgres dbname=curso-go sslmode=disable")
		if err != nil {
			log.Fatalf("No se puede abrir la base de datos: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("No se puede hacer ping con la base de datos: %v", err)
		}

	})
	fmt.Println("Conectado a Postgres")

}

// Pool retorna una unica instancia de db
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}
