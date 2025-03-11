-- name: CreateApartment :one
INSERT INTO apartments (
    unit_number,
    price,
    size,
    management_id,
    availability,
    lease_id,
    lease_start_date,
    lease_end_date,
    updated_at,
    created_at
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetApartment :one
SELECT id,
  unit_number,
  price,
  size,
  management_id,
  availability,
  lease_id,
  lease_start_date,
  lease_end_date
FROM apartments
WHERE id = $1
LIMIT 1;

-- name: ListApartments :many
SELECT id,
  unit_number,
  price,
  size,
  management_id,
  availability,
  lease_id,
  lease_start_date,
  lease_end_date
FROM apartments
ORDER BY unit_number DESC
LIMIT $1 OFFSET $2;

-- name: UpdateApartment :exec
UPDATE apartments
SET price = $2,
  management_id = $3,
  availability = $4,
  lease_id = $5,
  lease_start_date = $6,
  lease_end_date = $7,
  updated_at = $8,
  WHERE id = $1;

-- name: DeleteApartment :exec
DELETE FROM apartments
WHERE id = $1;