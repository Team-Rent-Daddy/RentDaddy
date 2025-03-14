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
    status,
    last_login,
    updated_at,
    created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9,$10
) RETURNING id, clerk_id, first_name, last_name, email, phone,role, created_at
`

type CreateUserParams struct {
	ClerkID   string           `json:"clerk_id"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	Phone     pgtype.Text      `json:"phone"`
	Role      Role             `json:"role"`
	Status    AccountStatus    `json:"status"`
	LastLogin pgtype.Timestamp `json:"last_login"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
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
		arg.Status,
		arg.LastLogin,
		arg.UpdatedAt,
		arg.CreatedAt,
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

const deleteUserByClerkID = `-- name: DeleteUserByClerkID :exec
DELETE FROM users
WHERE clerk_id = $1
`

func (q *Queries) DeleteUserByClerkID(ctx context.Context, clerkID string) error {
	_, err := q.db.Exec(ctx, deleteUserByClerkID, clerkID)
	return err
}

const getAdminByClerkID = `-- name: GetAdminByClerkID :one
SELECT id, clerk_id, first_name, last_name, email, role, unit_number, status, created_at
FROM users
WHERE clerk_id = $1 AND role = 'admin'
`

type GetAdminByClerkIDRow struct {
	ID         int64            `json:"id"`
	ClerkID    string           `json:"clerk_id"`
	FirstName  string           `json:"first_name"`
	LastName   string           `json:"last_name"`
	Email      string           `json:"email"`
	Role       Role             `json:"role"`
	UnitNumber pgtype.Int2      `json:"unit_number"`
	Status     AccountStatus    `json:"status"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) GetAdminByClerkID(ctx context.Context, clerkID string) (GetAdminByClerkIDRow, error) {
	row := q.db.QueryRow(ctx, getAdminByClerkID, clerkID)
	var i GetAdminByClerkIDRow
	err := row.Scan(
		&i.ID,
		&i.ClerkID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Role,
		&i.UnitNumber,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getAllTenants = `-- name: GetAllTenants :many
SELECT id, clerk_id, first_name, last_name, email, role, unit_number, status, created_at
FROM users
WHERE clerk_id = $1 AND role = 'tenant'
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type GetAllTenantsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetAllTenantsRow struct {
	ID         int64            `json:"id"`
	ClerkID    string           `json:"clerk_id"`
	FirstName  string           `json:"first_name"`
	LastName   string           `json:"last_name"`
	Email      string           `json:"email"`
	Role       Role             `json:"role"`
	UnitNumber pgtype.Int2      `json:"unit_number"`
	Status     AccountStatus    `json:"status"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) GetAllTenants(ctx context.Context, arg GetAllTenantsParams) ([]GetAllTenantsRow, error) {
	rows, err := q.db.Query(ctx, getAllTenants, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllTenantsRow
	for rows.Next() {
		var i GetAllTenantsRow
		if err := rows.Scan(
			&i.ID,
			&i.ClerkID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Role,
			&i.UnitNumber,
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

const getTenantByClerkID = `-- name: GetTenantByClerkID :one
SELECT id, clerk_id, first_name, last_name, email, role, unit_number, status, created_at
FROM users
WHERE clerk_id = $1 AND role = 'tenant'
`

type GetTenantByClerkIDRow struct {
	ID         int64            `json:"id"`
	ClerkID    string           `json:"clerk_id"`
	FirstName  string           `json:"first_name"`
	LastName   string           `json:"last_name"`
	Email      string           `json:"email"`
	Role       Role             `json:"role"`
	UnitNumber pgtype.Int2      `json:"unit_number"`
	Status     AccountStatus    `json:"status"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) GetTenantByClerkID(ctx context.Context, clerkID string) (GetTenantByClerkIDRow, error) {
	row := q.db.QueryRow(ctx, getTenantByClerkID, clerkID)
	var i GetTenantByClerkIDRow
	err := row.Scan(
		&i.ID,
		&i.ClerkID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Role,
		&i.UnitNumber,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByClerkID = `-- name: GetUserByClerkID :one
SELECT id, clerk_id, first_name, last_name, email, role, unit_number, status, created_at
FROM users
WHERE clerk_id = $1
`

type GetUserByClerkIDRow struct {
	ID         int64            `json:"id"`
	ClerkID    string           `json:"clerk_id"`
	FirstName  string           `json:"first_name"`
	LastName   string           `json:"last_name"`
	Email      string           `json:"email"`
	Role       Role             `json:"role"`
	UnitNumber pgtype.Int2      `json:"unit_number"`
	Status     AccountStatus    `json:"status"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) GetUserByClerkID(ctx context.Context, clerkID string) (GetUserByClerkIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByClerkID, clerkID)
	var i GetUserByClerkIDRow
	err := row.Scan(
		&i.ID,
		&i.ClerkID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Role,
		&i.UnitNumber,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, clerk_id, first_name, last_name, email, role, unit_number, status, created_at
FROM users
WHERE role = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetUsersParams struct {
	Role   Role  `json:"role"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetUsersRow struct {
	ID         int64            `json:"id"`
	ClerkID    string           `json:"clerk_id"`
	FirstName  string           `json:"first_name"`
	LastName   string           `json:"last_name"`
	Email      string           `json:"email"`
	Role       Role             `json:"role"`
	UnitNumber pgtype.Int2      `json:"unit_number"`
	Status     AccountStatus    `json:"status"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUsersRow, error) {
	rows, err := q.db.Query(ctx, getUsers, arg.Role, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersRow
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.ClerkID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Role,
			&i.UnitNumber,
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

const updateUserCredentials = `-- name: UpdateUserCredentials :exec
UPDATE users
SET first_name = $2, last_name = $3, email = $4
WHERE clerk_id = $1
`

type UpdateUserCredentialsParams struct {
	ClerkID   string `json:"clerk_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (q *Queries) UpdateUserCredentials(ctx context.Context, arg UpdateUserCredentialsParams) error {
	_, err := q.db.Exec(ctx, updateUserCredentials,
		arg.ClerkID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
	)
	return err
}

const updateUserRole = `-- name: UpdateUserRole :exec
UPDATE users
SET role = $2
WHERE clerk_id = $1
`

type UpdateUserRoleParams struct {
	ClerkID string `json:"clerk_id"`
	Role    Role   `json:"role"`
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) error {
	_, err := q.db.Exec(ctx, updateUserRole, arg.ClerkID, arg.Role)
	return err
}
