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
INSERT INTO leases (
  lease_number, external_doc_id, lease_pdf,
  tenant_id, landlord_id, apartment_id,
  lease_start_date, lease_end_date, rent_amount,
  status, created_by, updated_by,
  previous_lease_id, tenant_signing_url
) VALUES (
  $1, $2, $3,
  $4, $5, $6,
  $7, $8, $9,
  $10, $11, $12,
  $13, $14
)
RETURNING id, lease_number, external_doc_id, lease_pdf, tenant_id, landlord_id, apartment_id, lease_start_date, lease_end_date, rent_amount, status, created_by, updated_by, created_at, updated_at, previous_lease_id, tenant_signing_url
`

type CreateLeaseParams struct {
	LeaseNumber      int64          `json:"lease_number"`
	ExternalDocID    string         `json:"external_doc_id"`
	LeasePdf         []byte         `json:"lease_pdf"`
	TenantID         int64          `json:"tenant_id"`
	LandlordID       int64          `json:"landlord_id"`
	ApartmentID      int64          `json:"apartment_id"`
	LeaseStartDate   pgtype.Date    `json:"lease_start_date"`
	LeaseEndDate     pgtype.Date    `json:"lease_end_date"`
	RentAmount       pgtype.Numeric `json:"rent_amount"`
	Status           LeaseStatus    `json:"status"`
	CreatedBy        int64          `json:"created_by"`
	UpdatedBy        int64          `json:"updated_by"`
	PreviousLeaseID  pgtype.Int8    `json:"previous_lease_id"`
	TenantSigningUrl pgtype.Text    `json:"tenant_signing_url"`
}

func (q *Queries) CreateLease(ctx context.Context, arg CreateLeaseParams) (Lease, error) {
	row := q.db.QueryRow(ctx, createLease,
		arg.LeaseNumber,
		arg.ExternalDocID,
		arg.LeasePdf,
		arg.TenantID,
		arg.LandlordID,
		arg.ApartmentID,
		arg.LeaseStartDate,
		arg.LeaseEndDate,
		arg.RentAmount,
		arg.Status,
		arg.CreatedBy,
		arg.UpdatedBy,
		arg.PreviousLeaseID,
		arg.TenantSigningUrl,
	)
	var i Lease
	err := row.Scan(
		&i.ID,
		&i.LeaseNumber,
		&i.ExternalDocID,
		&i.LeasePdf,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.Status,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PreviousLeaseID,
		&i.TenantSigningUrl,
	)
	return i, err
}

const expireLeasesEndingToday = `-- name: ExpireLeasesEndingToday :one
WITH expired_leases AS (
    UPDATE leases
    SET status = 'expired', updated_at = NOW()
    WHERE status = 'active' AND lease_end_date <= CURRENT_DATE
    RETURNING id
)
SELECT 
    COUNT(*) as expired_count,
    CASE 
        WHEN COUNT(*) = 0 THEN 'No leases expired today'
        WHEN COUNT(*) = 1 THEN '1 lease expired today'
        ELSE COUNT(*) || ' leases expired today'
    END as message
FROM expired_leases
`

type ExpireLeasesEndingTodayRow struct {
	ExpiredCount int64       `json:"expired_count"`
	Message      interface{} `json:"message"`
}

func (q *Queries) ExpireLeasesEndingToday(ctx context.Context) (ExpireLeasesEndingTodayRow, error) {
	row := q.db.QueryRow(ctx, expireLeasesEndingToday)
	var i ExpireLeasesEndingTodayRow
	err := row.Scan(&i.ExpiredCount, &i.Message)
	return i, err
}

const getConflictingActiveLease = `-- name: GetConflictingActiveLease :one
SELECT id, lease_number, external_doc_id, lease_pdf, tenant_id, landlord_id, apartment_id, lease_start_date, lease_end_date, rent_amount, status, created_by, updated_by, created_at, updated_at, previous_lease_id, tenant_signing_url FROM leases
WHERE tenant_id = $1
  AND status = 'active'
  AND lease_start_date <= $3
  AND lease_end_date >= $2
LIMIT 1
`

type GetConflictingActiveLeaseParams struct {
	TenantID       int64       `json:"tenant_id"`
	LeaseEndDate   pgtype.Date `json:"lease_end_date"`
	LeaseStartDate pgtype.Date `json:"lease_start_date"`
}

func (q *Queries) GetConflictingActiveLease(ctx context.Context, arg GetConflictingActiveLeaseParams) (Lease, error) {
	row := q.db.QueryRow(ctx, getConflictingActiveLease, arg.TenantID, arg.LeaseEndDate, arg.LeaseStartDate)
	var i Lease
	err := row.Scan(
		&i.ID,
		&i.LeaseNumber,
		&i.ExternalDocID,
		&i.LeasePdf,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.Status,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PreviousLeaseID,
		&i.TenantSigningUrl,
	)
	return i, err
}

const getDuplicateLease = `-- name: GetDuplicateLease :one
SELECT id, lease_number, external_doc_id, lease_pdf, tenant_id, landlord_id, apartment_id, lease_start_date, lease_end_date, rent_amount, status, created_by, updated_by, created_at, updated_at, previous_lease_id, tenant_signing_url FROM leases
WHERE tenant_id = $1
  AND apartment_id = $2
  AND status = $3
LIMIT 1
`

type GetDuplicateLeaseParams struct {
	TenantID    int64       `json:"tenant_id"`
	ApartmentID int64       `json:"apartment_id"`
	Status      LeaseStatus `json:"status"`
}

func (q *Queries) GetDuplicateLease(ctx context.Context, arg GetDuplicateLeaseParams) (Lease, error) {
	row := q.db.QueryRow(ctx, getDuplicateLease, arg.TenantID, arg.ApartmentID, arg.Status)
	var i Lease
	err := row.Scan(
		&i.ID,
		&i.LeaseNumber,
		&i.ExternalDocID,
		&i.LeasePdf,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.Status,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PreviousLeaseID,
		&i.TenantSigningUrl,
	)
	return i, err
}

const getLeaseByExternalDocID = `-- name: GetLeaseByExternalDocID :one
SELECT id, lease_number, external_doc_id, lease_pdf, tenant_id, landlord_id, apartment_id, lease_start_date, lease_end_date, rent_amount, status, created_by, updated_by, created_at, updated_at, previous_lease_id, tenant_signing_url FROM leases
WHERE external_doc_id = $1
LIMIT 1
`

func (q *Queries) GetLeaseByExternalDocID(ctx context.Context, externalDocID string) (Lease, error) {
	row := q.db.QueryRow(ctx, getLeaseByExternalDocID, externalDocID)
	var i Lease
	err := row.Scan(
		&i.ID,
		&i.LeaseNumber,
		&i.ExternalDocID,
		&i.LeasePdf,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.Status,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PreviousLeaseID,
		&i.TenantSigningUrl,
	)
	return i, err
}

const getLeaseByID = `-- name: GetLeaseByID :one
SELECT lease_number,
    external_doc_id,
    lease_pdf,
    tenant_id,
    landlord_id,
    apartment_id,
    lease_start_date,
    lease_end_date,
    rent_amount,
    status,
    created_by,
    updated_by,
    previous_lease_id
FROM leases
WHERE id = $1
`

type GetLeaseByIDRow struct {
	LeaseNumber     int64          `json:"lease_number"`
	ExternalDocID   string         `json:"external_doc_id"`
	LeasePdf        []byte         `json:"lease_pdf"`
	TenantID        int64          `json:"tenant_id"`
	LandlordID      int64          `json:"landlord_id"`
	ApartmentID     int64          `json:"apartment_id"`
	LeaseStartDate  pgtype.Date    `json:"lease_start_date"`
	LeaseEndDate    pgtype.Date    `json:"lease_end_date"`
	RentAmount      pgtype.Numeric `json:"rent_amount"`
	Status          LeaseStatus    `json:"status"`
	CreatedBy       int64          `json:"created_by"`
	UpdatedBy       int64          `json:"updated_by"`
	PreviousLeaseID pgtype.Int8    `json:"previous_lease_id"`
}

func (q *Queries) GetLeaseByID(ctx context.Context, id int64) (GetLeaseByIDRow, error) {
	row := q.db.QueryRow(ctx, getLeaseByID, id)
	var i GetLeaseByIDRow
	err := row.Scan(
		&i.LeaseNumber,
		&i.ExternalDocID,
		&i.LeasePdf,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.Status,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.PreviousLeaseID,
	)
	return i, err
}

const listActiveLeases = `-- name: ListActiveLeases :one
SELECT id, lease_number, external_doc_id, lease_pdf, tenant_id, landlord_id, apartment_id, lease_start_date, lease_end_date, rent_amount, status, created_by, updated_by, created_at, updated_at, previous_lease_id, tenant_signing_url FROM leases
WHERE status = 'active'
LIMIT 1
`

func (q *Queries) ListActiveLeases(ctx context.Context) (Lease, error) {
	row := q.db.QueryRow(ctx, listActiveLeases)
	var i Lease
	err := row.Scan(
		&i.ID,
		&i.LeaseNumber,
		&i.ExternalDocID,
		&i.LeasePdf,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.Status,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PreviousLeaseID,
		&i.TenantSigningUrl,
	)
	return i, err
}

const listLeases = `-- name: ListLeases :many
SELECT id, lease_number,
    external_doc_id,
    lease_pdf,
    tenant_id,
    landlord_id,
    apartment_id,
    lease_start_date,
    lease_end_date,
    rent_amount,
    status,
    created_by,
    updated_by,
    previous_lease_id
FROM leases ORDER BY created_at DESC
`

type ListLeasesRow struct {
	ID              int64          `json:"id"`
	LeaseNumber     int64          `json:"lease_number"`
	ExternalDocID   string         `json:"external_doc_id"`
	LeasePdf        []byte         `json:"lease_pdf"`
	TenantID        int64          `json:"tenant_id"`
	LandlordID      int64          `json:"landlord_id"`
	ApartmentID     int64          `json:"apartment_id"`
	LeaseStartDate  pgtype.Date    `json:"lease_start_date"`
	LeaseEndDate    pgtype.Date    `json:"lease_end_date"`
	RentAmount      pgtype.Numeric `json:"rent_amount"`
	Status          LeaseStatus    `json:"status"`
	CreatedBy       int64          `json:"created_by"`
	UpdatedBy       int64          `json:"updated_by"`
	PreviousLeaseID pgtype.Int8    `json:"previous_lease_id"`
}

func (q *Queries) ListLeases(ctx context.Context) ([]ListLeasesRow, error) {
	rows, err := q.db.Query(ctx, listLeases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLeasesRow
	for rows.Next() {
		var i ListLeasesRow
		if err := rows.Scan(
			&i.ID,
			&i.LeaseNumber,
			&i.ExternalDocID,
			&i.LeasePdf,
			&i.TenantID,
			&i.LandlordID,
			&i.ApartmentID,
			&i.LeaseStartDate,
			&i.LeaseEndDate,
			&i.RentAmount,
			&i.Status,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.PreviousLeaseID,
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

const markLeaseAsSignedBothParties = `-- name: MarkLeaseAsSignedBothParties :exec
UPDATE leases
SET status = 'active', updated_at = now()
WHERE id = $1
RETURNING lease_number,
    external_doc_id,
    lease_pdf,
    tenant_id,
    landlord_id,
    apartment_id,
    lease_start_date,
    lease_end_date,
    rent_amount,
    status,
    created_by,
    updated_by,
    previous_lease_id
`

func (q *Queries) MarkLeaseAsSignedBothParties(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, markLeaseAsSignedBothParties, id)
	return err
}

const renewLease = `-- name: RenewLease :one
INSERT INTO leases (
  lease_number, external_doc_id, tenant_id, landlord_id, apartment_id,
  lease_start_date, lease_end_date, rent_amount, status, lease_pdf,
  created_by, updated_by, previous_lease_id, tenant_signing_url
)
VALUES (
  $1, $2, $3, $4, $5,
  $6, $7, $8, $9, $10,
  $11, $12, $13, $14
)
RETURNING id, lease_number
`

type RenewLeaseParams struct {
	LeaseNumber      int64          `json:"lease_number"`
	ExternalDocID    string         `json:"external_doc_id"`
	TenantID         int64          `json:"tenant_id"`
	LandlordID       int64          `json:"landlord_id"`
	ApartmentID      int64          `json:"apartment_id"`
	LeaseStartDate   pgtype.Date    `json:"lease_start_date"`
	LeaseEndDate     pgtype.Date    `json:"lease_end_date"`
	RentAmount       pgtype.Numeric `json:"rent_amount"`
	Status           LeaseStatus    `json:"status"`
	LeasePdf         []byte         `json:"lease_pdf"`
	CreatedBy        int64          `json:"created_by"`
	UpdatedBy        int64          `json:"updated_by"`
	PreviousLeaseID  pgtype.Int8    `json:"previous_lease_id"`
	TenantSigningUrl pgtype.Text    `json:"tenant_signing_url"`
}

type RenewLeaseRow struct {
	ID          int64 `json:"id"`
	LeaseNumber int64 `json:"lease_number"`
}

func (q *Queries) RenewLease(ctx context.Context, arg RenewLeaseParams) (RenewLeaseRow, error) {
	row := q.db.QueryRow(ctx, renewLease,
		arg.LeaseNumber,
		arg.ExternalDocID,
		arg.TenantID,
		arg.LandlordID,
		arg.ApartmentID,
		arg.LeaseStartDate,
		arg.LeaseEndDate,
		arg.RentAmount,
		arg.Status,
		arg.LeasePdf,
		arg.CreatedBy,
		arg.UpdatedBy,
		arg.PreviousLeaseID,
		arg.TenantSigningUrl,
	)
	var i RenewLeaseRow
	err := row.Scan(&i.ID, &i.LeaseNumber)
	return i, err
}

const storeGeneratedLeasePDF = `-- name: StoreGeneratedLeasePDF :exec
UPDATE leases
SET lease_pdf = $1, external_doc_id = $2, updated_at = now()
WHERE id = $3
RETURNING lease_pdf
`

type StoreGeneratedLeasePDFParams struct {
	LeasePdf      []byte `json:"lease_pdf"`
	ExternalDocID string `json:"external_doc_id"`
	ID            int64  `json:"id"`
}

func (q *Queries) StoreGeneratedLeasePDF(ctx context.Context, arg StoreGeneratedLeasePDFParams) error {
	_, err := q.db.Exec(ctx, storeGeneratedLeasePDF, arg.LeasePdf, arg.ExternalDocID, arg.ID)
	return err
}

const terminateLease = `-- name: TerminateLease :one
UPDATE leases
SET 
    
    status = 'terminated', 
    updated_by = $1, 
    updated_at = now()
WHERE id = $2
RETURNING id, lease_number, external_doc_id, tenant_id, landlord_id, apartment_id, 
    lease_start_date, lease_end_date, rent_amount, status, 
    updated_by, updated_at, previous_lease_id
`

type TerminateLeaseParams struct {
	UpdatedBy int64 `json:"updated_by"`
	ID        int64 `json:"id"`
}

type TerminateLeaseRow struct {
	ID              int64            `json:"id"`
	LeaseNumber     int64            `json:"lease_number"`
	ExternalDocID   string           `json:"external_doc_id"`
	TenantID        int64            `json:"tenant_id"`
	LandlordID      int64            `json:"landlord_id"`
	ApartmentID     int64            `json:"apartment_id"`
	LeaseStartDate  pgtype.Date      `json:"lease_start_date"`
	LeaseEndDate    pgtype.Date      `json:"lease_end_date"`
	RentAmount      pgtype.Numeric   `json:"rent_amount"`
	Status          LeaseStatus      `json:"status"`
	UpdatedBy       int64            `json:"updated_by"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	PreviousLeaseID pgtype.Int8      `json:"previous_lease_id"`
}

func (q *Queries) TerminateLease(ctx context.Context, arg TerminateLeaseParams) (TerminateLeaseRow, error) {
	row := q.db.QueryRow(ctx, terminateLease, arg.UpdatedBy, arg.ID)
	var i TerminateLeaseRow
	err := row.Scan(
		&i.ID,
		&i.LeaseNumber,
		&i.ExternalDocID,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.Status,
		&i.UpdatedBy,
		&i.UpdatedAt,
		&i.PreviousLeaseID,
	)
	return i, err
}

const updateLease = `-- name: UpdateLease :exec
UPDATE leases
SET 
    tenant_id = $1,
    status = $2,
    status = $2,
    lease_start_date = $3,
    lease_end_date = $4,
    rent_amount = $5,
    updated_by = $6,
    updated_at = now()
WHERE id = $7
RETURNING lease_number,
    external_doc_id,
    lease_pdf,
    tenant_id,
    landlord_id,
    apartment_id,
    lease_start_date,
    lease_end_date,
    rent_amount,
    status,
    created_by,
    updated_by,
    previous_lease_id
`

type UpdateLeaseParams struct {
	TenantID       int64          `json:"tenant_id"`
	Status         LeaseStatus    `json:"status"`
	LeaseStartDate pgtype.Date    `json:"lease_start_date"`
	LeaseEndDate   pgtype.Date    `json:"lease_end_date"`
	RentAmount     pgtype.Numeric `json:"rent_amount"`
	UpdatedBy      int64          `json:"updated_by"`
	ID             int64          `json:"id"`
}

func (q *Queries) UpdateLease(ctx context.Context, arg UpdateLeaseParams) error {
	_, err := q.db.Exec(ctx, updateLease,
		arg.TenantID,
		arg.Status,
		arg.LeaseStartDate,
		arg.LeaseEndDate,
		arg.RentAmount,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const updateLeasePDF = `-- name: UpdateLeasePDF :exec
UPDATE leases
SET 
    lease_pdf = $2, 
    updated_by = $3,
    updated_at = NOW()
WHERE id = $1
`

type UpdateLeasePDFParams struct {
	ID        int64  `json:"id"`
	LeasePdf  []byte `json:"lease_pdf"`
	UpdatedBy int64  `json:"updated_by"`
}

func (q *Queries) UpdateLeasePDF(ctx context.Context, arg UpdateLeasePDFParams) error {
	_, err := q.db.Exec(ctx, updateLeasePDF, arg.ID, arg.LeasePdf, arg.UpdatedBy)
	return err
}

const updateLeaseStatus = `-- name: UpdateLeaseStatus :one
UPDATE leases
SET status = $2, updated_by = $3, updated_at = NOW()
WHERE id = $1
RETURNING id, lease_number, external_doc_id, tenant_id, landlord_id, apartment_id, 
    lease_start_date, lease_end_date, rent_amount, status, created_by, 
    updated_by, updated_at, previous_lease_id
`

type UpdateLeaseStatusParams struct {
	ID        int64       `json:"id"`
	Status    LeaseStatus `json:"status"`
	UpdatedBy int64       `json:"updated_by"`
}

type UpdateLeaseStatusRow struct {
	ID              int64            `json:"id"`
	LeaseNumber     int64            `json:"lease_number"`
	ExternalDocID   string           `json:"external_doc_id"`
	TenantID        int64            `json:"tenant_id"`
	LandlordID      int64            `json:"landlord_id"`
	ApartmentID     int64            `json:"apartment_id"`
	LeaseStartDate  pgtype.Date      `json:"lease_start_date"`
	LeaseEndDate    pgtype.Date      `json:"lease_end_date"`
	RentAmount      pgtype.Numeric   `json:"rent_amount"`
	Status          LeaseStatus      `json:"status"`
	CreatedBy       int64            `json:"created_by"`
	UpdatedBy       int64            `json:"updated_by"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	PreviousLeaseID pgtype.Int8      `json:"previous_lease_id"`
}

func (q *Queries) UpdateLeaseStatus(ctx context.Context, arg UpdateLeaseStatusParams) (UpdateLeaseStatusRow, error) {
	row := q.db.QueryRow(ctx, updateLeaseStatus, arg.ID, arg.Status, arg.UpdatedBy)
	var i UpdateLeaseStatusRow
	err := row.Scan(
		&i.ID,
		&i.LeaseNumber,
		&i.ExternalDocID,
		&i.TenantID,
		&i.LandlordID,
		&i.ApartmentID,
		&i.LeaseStartDate,
		&i.LeaseEndDate,
		&i.RentAmount,
		&i.Status,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.UpdatedAt,
		&i.PreviousLeaseID,
	)
	return i, err
}

const updateTenantSigningURL = `-- name: UpdateTenantSigningURL :exec
UPDATE leases
SET tenant_signing_url = $2,
    updated_at = now()
WHERE id = $1
`

type UpdateTenantSigningURLParams struct {
	ID               int64       `json:"id"`
	TenantSigningUrl pgtype.Text `json:"tenant_signing_url"`
}

func (q *Queries) UpdateTenantSigningURL(ctx context.Context, arg UpdateTenantSigningURLParams) error {
	_, err := q.db.Exec(ctx, updateTenantSigningURL, arg.ID, arg.TenantSigningUrl)
	return err
}
