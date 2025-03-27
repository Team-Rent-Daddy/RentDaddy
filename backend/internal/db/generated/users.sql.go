// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    clerk_id,
    first_name,
    last_name,
    email,
    phone,
    role,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, now()
) RETURNING id, clerk_id, first_name, last_name, email, phone, role, created_at
`

type CreateUserParams struct {
	ClerkID   string      `json:"clerk_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Phone     pgtype.Text `json:"phone"`
	Role      Role        `json:"role"`
}

type CreateUserRow struct {
	ID        int64            `json:"id"`
	ClerkID   string           `json:"clerk_id"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	Phone     pgtype.Text      `json:"phone"`
	Role      Role             `json:"role"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ClerkID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.Role,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.ClerkID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE
FROM users
WHERE clerk_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, clerkID string) error {
	_, err := q.db.Exec(ctx, deleteUser, clerkID)
	return err
}

const getTenantsWithNoLease = `-- name: GetTenantsWithNoLease :many
SELECT id, clerk_id, first_name, last_name, email, phone, role, status
FROM users
WHERE role = 'tenant' 
AND id NOT IN (SELECT tenant_id FROM leases)
`

type GetTenantsWithNoLeaseRow struct {
	ID        int64         `json:"id"`
	ClerkID   string        `json:"clerk_id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Phone     pgtype.Text   `json:"phone"`
	Role      Role          `json:"role"`
	Status    AccountStatus `json:"status"`
}

func (q *Queries) GetTenantsWithNoLease(ctx context.Context) ([]GetTenantsWithNoLeaseRow, error) {
	rows, err := q.db.Query(ctx, getTenantsWithNoLease)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTenantsWithNoLeaseRow
	for rows.Next() {
		var i GetTenantsWithNoLeaseRow
		if err := rows.Scan(
			&i.ID,
			&i.ClerkID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.Role,
			&i.Status,
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

const getUser = `-- name: GetUser :one
SELECT id, clerk_id, first_name, last_name, email, phone, role, status, created_at
FROM users
WHERE clerk_id = $1
LIMIT 1
`

type GetUserRow struct {
	ID        int64            `json:"id"`
	ClerkID   string           `json:"clerk_id"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	Phone     pgtype.Text      `json:"phone"`
	Role      Role             `json:"role"`
	Status    AccountStatus    `json:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) GetUser(ctx context.Context, clerkID string) (GetUserRow, error) {
	row := q.db.QueryRow(ctx, getUser, clerkID)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.ClerkID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByClerkId = `-- name: GetUserByClerkId :one
SELECT id, clerk_id, first_name, last_name, email, phone, role, status, created_at
FROM users
WHERE clerk_id = $1
LIMIT 1
`

type GetUserByClerkIdRow struct {
	ID        int64            `json:"id"`
	ClerkID   string           `json:"clerk_id"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	Phone     pgtype.Text      `json:"phone"`
	Role      Role             `json:"role"`
	Status    AccountStatus    `json:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) GetUserByClerkId(ctx context.Context, clerkID string) (GetUserByClerkIdRow, error) {
	row := q.db.QueryRow(ctx, getUserByClerkId, clerkID)
	var i GetUserByClerkIdRow
	err := row.Scan(
		&i.ID,
		&i.ClerkID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, clerk_id, first_name, last_name, email, phone,role, status
FROM users
WHERE id = $1
LIMIT 1
`

type GetUserByIDRow struct {
	ID        int64         `json:"id"`
	ClerkID   string        `json:"clerk_id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Phone     pgtype.Text   `json:"phone"`
	Role      Role          `json:"role"`
	Status    AccountStatus `json:"status"`
}

func (q *Queries) GetUserByID(ctx context.Context, id int64) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.ClerkID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.Role,
		&i.Status,
	)
	return i, err
}

const listTenantsWithLeases = `-- name: ListTenantsWithLeases :many
SELECT 
    users.id,
    users.clerk_id,
    users.first_name,
    users.last_name,
    users.email,
    users.phone,
    users.role,
    users.status,
    users.created_at,
    leases.status,
    leases.lease_start_date,
    leases.lease_end_date
FROM users
LEFT JOIN leases
ON users.id = leases.tenant_id
WHERE users.role = 'tenant'
ORDER BY users.created_at DESC
`

type ListTenantsWithLeasesRow struct {
	ID             int64            `json:"id"`
	ClerkID        string           `json:"clerk_id"`
	FirstName      string           `json:"first_name"`
	LastName       string           `json:"last_name"`
	Email          string           `json:"email"`
	Phone          pgtype.Text      `json:"phone"`
	Role           Role             `json:"role"`
	Status         AccountStatus    `json:"status"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	Status_2       NullLeaseStatus  `json:"status_2"`
	LeaseStartDate pgtype.Date      `json:"lease_start_date"`
	LeaseEndDate   pgtype.Date      `json:"lease_end_date"`
}

func (q *Queries) ListTenantsWithLeases(ctx context.Context) ([]ListTenantsWithLeasesRow, error) {
	rows, err := q.db.Query(ctx, listTenantsWithLeases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListTenantsWithLeasesRow
	for rows.Next() {
		var i ListTenantsWithLeasesRow
		if err := rows.Scan(
			&i.ID,
			&i.ClerkID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.Role,
			&i.Status,
			&i.CreatedAt,
			&i.Status_2,
			&i.LeaseStartDate,
			&i.LeaseEndDate,
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

const listUsersByRole = `-- name: ListUsersByRole :many
SELECT id,
       clerk_id,
       first_name,
       last_name,
       email,
       phone,
       role,
       status,
       created_at
FROM users
WHERE role = $1
ORDER BY created_at DESC
`

type ListUsersByRoleRow struct {
	ID        int64            `json:"id"`
	ClerkID   string           `json:"clerk_id"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	Phone     pgtype.Text      `json:"phone"`
	Role      Role             `json:"role"`
	Status    AccountStatus    `json:"status"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) ListUsersByRole(ctx context.Context, role Role) ([]ListUsersByRoleRow, error) {
	rows, err := q.db.Query(ctx, listUsersByRole, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersByRoleRow
	for rows.Next() {
		var i ListUsersByRoleRow
		if err := rows.Scan(
			&i.ID,
			&i.ClerkID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.Role,
			&i.Status,
			&i.CreatedAt,
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

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET first_name = $2,
    last_name  = $3,
    email      = $4,
    phone      = $5,
    updated_at = now()
WHERE clerk_id = $1
`

type UpdateUserParams struct {
	ClerkID   string      `json:"clerk_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Phone     pgtype.Text `json:"phone"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ClerkID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
	)
	return err
}
