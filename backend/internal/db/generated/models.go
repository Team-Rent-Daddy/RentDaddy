// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package generated

import (
	"database/sql"
)

type Tenant struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt sql.NullTime
}
