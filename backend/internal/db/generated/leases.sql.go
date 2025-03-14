// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: leases.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createLease = `-- name: CreateLease :one
INSERT INTO leases (external_doc_id, tenant_id, landlord_id, lease_start_date, lease_end_date, rent_amount, lease_status)
VALUES ($1, $2, $3, $4, $5, $6, 'DRAFT')
RETURNING document_id
`

type CreateLeaseParams struct {
	ExternalDocID  string         `json:"external_doc_id"`
	TenantID       int64          `json:"tenant_id"`
	LandlordID     int64          `json:"landlord_id"`
	LeaseStartDate pgtype.Date    `json:"lease_start_date"`
	LeaseEndDate   pgtype.Date    `json:"lease_end_date"`
	RentAmount     pgtype.Numeric `json:"rent_amount"`
}

// PLEASE SHRINK THE LEASES TABLE AND MAKE SURE THESE QUERIES ACTUALLY WORK VIA
// sqlc vet at the backend folder
// ALSO MAKE SURE baseTables.sql match init.up.sql
func (q *Queries) CreateLease(ctx context.Context, arg CreateLeaseParams) (int64, error) {
	row := q.db.QueryRow(ctx, createLease,
		arg.ExternalDocID,
		arg.TenantID,
		arg.LandlordID,
		arg.LeaseStartDate,
		arg.LeaseEndDate,
		arg.RentAmount,
	)
	var document_id int64
	err := row.Scan(&document_id)
	return document_id, err
}

const getLeaseByID = `-- name: GetLeaseByID :one
SELECT document_id, external_doc_id, tenant_id, landlord_id, lease_start_date, lease_end_date, rent_amount, payment_status, lease_status, created_at, updated_at FROM leases WHERE document_id = $1 LIMIT 1
`

func (q *Queries) GetLeaseByID(ctx context.Context, documentID int64) (Lease, error) {
	row := q.db.QueryRow(ctx, getLeaseByID, documentID)
	var i Lease
	err := row.Scan(
		&i.DocumentID,
		&i.ExternalDocID,
		&i.TenantID,
		&i.LandlordID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.PaymentStatus,
		&i.LeaseStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listLeases = `-- name: ListLeases :many
SELECT document_id, external_doc_id, tenant_id, landlord_id, lease_start_date, lease_end_date, rent_amount, payment_status, lease_status, created_at, updated_at FROM leases ORDER BY created_at DESC
`

func (q *Queries) ListLeases(ctx context.Context) ([]Lease, error) {
	rows, err := q.db.Query(ctx, listLeases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Lease
	for rows.Next() {
		var i Lease
		if err := rows.Scan(
			&i.DocumentID,
			&i.ExternalDocID,
			&i.TenantID,
			&i.LandlordID,
			&i.LeaseStartDate,
			&i.LeaseEndDate,
			&i.RentAmount,
			&i.PaymentStatus,
			&i.LeaseStatus,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const renewLease = `-- name: RenewLease :one
UPDATE leases
SET lease_end_date = $1, updated_at = now()
WHERE document_id  = $2 AND lease_status = 'active'
RETURNING document_id, external_doc_id, tenant_id, landlord_id, lease_start_date, lease_end_date, rent_amount, payment_status, lease_status, created_at, updated_at
`

type RenewLeaseParams struct {
	LeaseEndDate pgtype.Date `json:"lease_end_date"`
	DocumentID   int64       `json:"document_id"`
}

func (q *Queries) RenewLease(ctx context.Context, arg RenewLeaseParams) (Lease, error) {
	row := q.db.QueryRow(ctx, renewLease, arg.LeaseEndDate, arg.DocumentID)
	var i Lease
	err := row.Scan(
		&i.DocumentID,
		&i.ExternalDocID,
		&i.TenantID,
		&i.LandlordID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.PaymentStatus,
		&i.LeaseStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const terminateLease = `-- name: TerminateLease :one
UPDATE leases
SET lease_status = 'terminated', updated_at = now()
WHERE document_id  = $1
RETURNING document_id, external_doc_id, tenant_id, landlord_id, lease_start_date, lease_end_date, rent_amount, payment_status, lease_status, created_at, updated_at
`

func (q *Queries) TerminateLease(ctx context.Context, documentID int64) (Lease, error) {
	row := q.db.QueryRow(ctx, terminateLease, documentID)
	var i Lease
	err := row.Scan(
		&i.DocumentID,
		&i.ExternalDocID,
		&i.TenantID,
		&i.LandlordID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.PaymentStatus,
		&i.LeaseStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateLease = `-- name: UpdateLease :one
UPDATE leases
SET document_id = $1,
    tenant_id = $2,
    lease_status = $3,
    lease_start_date = $4,
    lease_end_date = $5,
    updated_at = now()
WHERE document_id = $6
RETURNING document_id, external_doc_id, tenant_id, landlord_id, lease_start_date, lease_end_date, rent_amount, payment_status, lease_status, created_at, updated_at
`

type UpdateLeaseParams struct {
	DocumentID     int64       `json:"document_id"`
	TenantID       int64       `json:"tenant_id"`
	LeaseStatus    string      `json:"lease_status"`
	LeaseStartDate pgtype.Date `json:"lease_start_date"`
	LeaseEndDate   pgtype.Date `json:"lease_end_date"`
	DocumentID_2   int64       `json:"document_id_2"`
}

func (q *Queries) UpdateLease(ctx context.Context, arg UpdateLeaseParams) (Lease, error) {
	row := q.db.QueryRow(ctx, updateLease,
		arg.DocumentID,
		arg.TenantID,
		arg.LeaseStatus,
		arg.LeaseStartDate,
		arg.LeaseEndDate,
		arg.DocumentID_2,
	)
	var i Lease
	err := row.Scan(
		&i.DocumentID,
		&i.ExternalDocID,
		&i.TenantID,
		&i.LandlordID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.PaymentStatus,
		&i.LeaseStatus,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
