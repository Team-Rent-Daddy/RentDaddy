// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: parking_permits.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createParkingPermit = `-- name: CreateParkingPermit :one
INSERT INTO parking_permits (
    permit_number,
    created_by,
    expires_at,
    updated_at
)
VALUES (
    $1,
    $2,
    $3,
    now()
)
RETURNING id, permit_number, created_by, updated_at, expires_at
`

type CreateParkingPermitParams struct {
	PermitNumber int64            `json:"permit_number"`
	CreatedBy    int64            `json:"created_by"`
	ExpiresAt    pgtype.Timestamp `json:"expires_at"`
}

func (q *Queries) CreateParkingPermit(ctx context.Context, arg CreateParkingPermitParams) (ParkingPermit, error) {
	row := q.db.QueryRow(ctx, createParkingPermit, arg.PermitNumber, arg.CreatedBy, arg.ExpiresAt)
	var i ParkingPermit
	err := row.Scan(
		&i.ID,
		&i.PermitNumber,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const deleteParkingPermit = `-- name: DeleteParkingPermit :exec
DELETE FROM parking_permits
WHERE id = $1
`

func (q *Queries) DeleteParkingPermit(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteParkingPermit, id)
	return err
}

const getNumOfUserParkingPermits = `-- name: GetNumOfUserParkingPermits :one
SELECT COUNT(*)
FROM parking_permits
WHERE created_by = $1
`

func (q *Queries) GetNumOfUserParkingPermits(ctx context.Context, createdBy int64) (int64, error) {
	row := q.db.QueryRow(ctx, getNumOfUserParkingPermits, createdBy)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getParkingPermit = `-- name: GetParkingPermit :one
SELECT permit_number, created_by, updated_at, expires_at
FROM parking_permits
WHERE id = $1
LIMIT 1
`

type GetParkingPermitRow struct {
	PermitNumber int64            `json:"permit_number"`
	CreatedBy    int64            `json:"created_by"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	ExpiresAt    pgtype.Timestamp `json:"expires_at"`
}

func (q *Queries) GetParkingPermit(ctx context.Context, id int64) (GetParkingPermitRow, error) {
	row := q.db.QueryRow(ctx, getParkingPermit, id)
	var i GetParkingPermitRow
	err := row.Scan(
		&i.PermitNumber,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const getTenantParkingPermits = `-- name: GetTenantParkingPermits :many
SELECT id, permit_number, created_by, updated_at, expires_at
FROM parking_permits
WHERE created_by = $1
`

func (q *Queries) GetTenantParkingPermits(ctx context.Context, createdBy int64) ([]ParkingPermit, error) {
	rows, err := q.db.Query(ctx, getTenantParkingPermits, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ParkingPermit
	for rows.Next() {
		var i ParkingPermit
		if err := rows.Scan(
			&i.ID,
			&i.PermitNumber,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.ExpiresAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listParkingPermits = `-- name: ListParkingPermits :many
SELECT id, permit_number, created_by, updated_at, expires_at
FROM parking_permits
ORDER BY created_by DESC
`

func (q *Queries) ListParkingPermits(ctx context.Context) ([]ParkingPermit, error) {
	rows, err := q.db.Query(ctx, listParkingPermits)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ParkingPermit
	for rows.Next() {
		var i ParkingPermit
		if err := rows.Scan(
			&i.ID,
			&i.PermitNumber,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.ExpiresAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
