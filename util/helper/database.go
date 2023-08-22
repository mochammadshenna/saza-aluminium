package helper

import (
	"context"
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover() // overlaping
	if err != nil {
		errorRollback := tx.Rollback()
		PanicError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicError(errorCommit)
	}
}

func BeginDB(ctx context.Context, db *sql.DB) (*sql.Tx, error) {
	return db.BeginTx(ctx, nil)
}
