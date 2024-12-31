// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: pre_go_crm_user_c.sql

package database

import (
	"context"
)

const getUserByEmailSQLC = `-- name: GetUserByEmailSQLC :one
SELECT usr_email, usr_id FROM ` + "`" + `pre_go_crm_user_c` + "`" + ` WHERE usr_email = ? LIMIT 1
`

type GetUserByEmailSQLCRow struct {
	UsrEmail string
	UsrID    uint32
}

func (q *Queries) GetUserByEmailSQLC(ctx context.Context, usrEmail string) (GetUserByEmailSQLCRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmailSQLC, usrEmail)
	var i GetUserByEmailSQLCRow
	err := row.Scan(&i.UsrEmail, &i.UsrID)
	return i, err
}

const updateUserStatusByUserId = `-- name: UpdateUserStatusByUserId :exec
UPDATE ` + "`" + `pre_go_crm_user_c` + "`" + `
SET usr_status = $2, usr_updated_at = $3
WHERE usr_id = $1
`

func (q *Queries) UpdateUserStatusByUserId(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateUserStatusByUserId)
	return err
}
