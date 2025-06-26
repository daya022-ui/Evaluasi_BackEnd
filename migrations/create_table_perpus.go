package migrations

import "database/sql"

type createPerpusTable struct{}

func (m *createPerpusTable) SkipProd() bool {
	return false
}

func getCreatePerpusTable() migration {
	return &createPerpusTable{}
}

func (m *createPerpusTable) Name() string {
	return "create-perpus"
}

func (m *createPerpusTable) Up(conn *sql.Tx) error {
	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS perpus (
			id SERIAL PRIMARY KEY,
			judul VARCHAR(255) NOT NULL UNIQUE,
			penulis VARCHAR(255) NOT NULL,
			status VARCHAR(10) NOT NULL DEFAULT 'available',
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`)

	if err != nil {
		return err
	}

	return err
}

func (m *createPerpusTable) Down(conn *sql.Tx) error {
	_, err := conn.Exec("DROP TABLE perpus")

	return err
}