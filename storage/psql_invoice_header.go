package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
							id SERIAL NOT NULL,
							client VARCHAR(100) NOT NULL,
							created_at TIMESTAMP NOT NULL DEFAULT now(),
							updated_at TIMESTAMP,
							CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
						)

	`
)

// PsqlInvoiceHeader used for work with postgres - InvoiceHeader
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// NewPsqlInvoiceHeader return nuevo apuntador de PsqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Migrate implementa la interfaz InvoiceHeader.Storage
func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migración de InvoiceHeadero ejecutada correctamente")
	return nil
}
