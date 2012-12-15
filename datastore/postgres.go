package datastore


import (
    "database/sql"
    _ "github.com/bmizerany/pq"
)

type PostgresDataHandler struct {
    db_name string
    db *sql.DB
}

func (p PostgresDataHandler) CreateNewDB(db_name string) error {
    temp_db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
    if err != nil {
        return err
    }
    p.db = temp_db
    return nil
}