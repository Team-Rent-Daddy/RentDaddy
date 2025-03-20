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
    unit_number,
    image_url,
    role,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, now()
) RETURNING id, clerk_id, first_name, last_name, email, phone, unit_number,role, created_at
`

type CreateUserParams struct {
	ClerkID    string      `json:"clerk_id"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Email      string      `json:"email"`
	Phone      pgtype.Text `json:"phone"`
	UnitNumber pgtype.Int2 `json:"unit_number"`
	ImageUrl   pgtype.Text `json:"image_url"`
	Role       Role        `json:"role"`
}

type CreateUserRow struct {
	ID         int64            `json:"id"`
	ClerkID    string           `json:"clerk_id"`
	FirstName  string           `json:"first_name"`
	LastName   string           `json:"last_name"`
	Email      string           `json:"email"`
	Phone      pgtype.Text      `json:"phone"`
	UnitNumber pgtype.Int2      `json:"unit_number"`
	Role       Role             `json:"role"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ClerkID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.UnitNumber,
		arg.ImageUrl,
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
		&i.UnitNumber,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE clerk_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, clerkID string) error {
	_, err := q.db.Exec(ctx, deleteUser, clerkID)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, clerk_id, first_name, last_name, email, phone, role, unit_number, status, created_at
FROM users
WHERE clerk_id = $1
LIMIT 1
`

type GetUserRow struct {
	ID         int64            `json:"id"`
	ClerkID    string           `json:"clerk_id"`
	FirstName  string           `json:"first_name"`
	LastName   string           `json:"last_name"`
	Email      string           `json:"email"`
	Phone      pgtype.Text      `json:"phone"`
	Role       Role             `json:"role"`
	UnitNumber pgtype.Int2      `json:"unit_number"`
	Status     AccountStatus    `json:"status"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
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
		&i.UnitNumber,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listUsersByRole = `-- name: ListUsersByRole :many
SELECT id, clerk_id, first_name, last_name, email, phone, role, unit_number, status, created_at
FROM users
WHERE role = $1
ORDER BY created_at DESC
`

type ListUsersByRoleRow struct {
	ID         int64            `json:"id"`
	ClerkID    string           `json:"clerk_id"`
	FirstName  string           `json:"first_name"`
	LastName   string           `json:"last_name"`
	Email      string           `json:"email"`
	Phone      pgtype.Text      `json:"phone"`
	Role       Role             `json:"role"`
	UnitNumber pgtype.Int2      `json:"unit_number"`
	Status     AccountStatus    `json:"status"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
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

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET first_name = $2, last_name = $3, email = $4, phone = $5, image_url = $6, updated_at = now()
WHERE clerk_id = $1
`

type UpdateUserParams struct {
	ClerkID   string      `json:"clerk_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Phone     pgtype.Text `json:"phone"`
	ImageUrl  pgtype.Text `json:"image_url"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ClerkID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.ImageUrl,
	)
	return err
}
