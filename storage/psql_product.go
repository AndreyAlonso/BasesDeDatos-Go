package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
							id SERIAL NOT NULL,
							name VARCHAR(25) NOT NULL,
							observations VARCHAR(100),
							price INT NOT NULL,
							created_at TIMESTAMP NOT NULL DEFAULT now(),
							updated_at TIMESTAMP,
							CONSTRAINT products_id_pk PRIMARY KEY(id)
						)

	`
)

// PsqlProduct used for work with postgres - product
type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct return nuevo apuntador de PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate implementa la interfaz product.Storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migración de producto ejecutada correctamente")
	return nil
}
