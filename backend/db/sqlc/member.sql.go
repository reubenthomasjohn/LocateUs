// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: member.sql

package db

import (
	"context"
	"database/sql"
)

const createMember = `-- name: CreateMember :one
INSERT INTO members (
  full_name,
  latitude,
  longitude,
  phone_number,
  status
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, full_name, phone_number, latitude, longitude, address, is_family, created_at, status
`

type CreateMemberParams struct {
	FullName    sql.NullString `json:"full_name"`
	Latitude    float64        `json:"latitude"`
	Longitude   float64        `json:"longitude"`
	PhoneNumber sql.NullString `json:"phone_number"`
	Status      NullUserStatus `json:"status"`
}

func (q *Queries) CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, createMember,
		arg.FullName,
		arg.Latitude,
		arg.Longitude,
		arg.PhoneNumber,
		arg.Status,
	)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.PhoneNumber,
		&i.Latitude,
		&i.Longitude,
		&i.Address,
		&i.IsFamily,
		&i.CreatedAt,
		&i.Status,
	)
	return i, err
}

const deleteMember = `-- name: DeleteMember :exec
DELETE FROM members 
WHERE id = $1
`

func (q *Queries) DeleteMember(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMember, id)
	return err
}

const getMember = `-- name: GetMember :one
SELECT id, full_name, phone_number, latitude, longitude, address, is_family, created_at, status FROM members 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMember(ctx context.Context, id int64) (Member, error) {
	row := q.db.QueryRowContext(ctx, getMember, id)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.PhoneNumber,
		&i.Latitude,
		&i.Longitude,
		&i.Address,
		&i.IsFamily,
		&i.CreatedAt,
		&i.Status,
	)
	return i, err
}

const getMemberByNumber = `-- name: GetMemberByNumber :one
SELECT id, full_name, phone_number, latitude, longitude, address, is_family, created_at, status FROM members
WHERE phone_number = $1 LIMIT 1
`

func (q *Queries) GetMemberByNumber(ctx context.Context, phoneNumber sql.NullString) (Member, error) {
	row := q.db.QueryRowContext(ctx, getMemberByNumber, phoneNumber)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.PhoneNumber,
		&i.Latitude,
		&i.Longitude,
		&i.Address,
		&i.IsFamily,
		&i.CreatedAt,
		&i.Status,
	)
	return i, err
}

const listMembers = `-- name: ListMembers :many
SELECT id, full_name, phone_number, latitude, longitude, address, is_family, created_at, status FROM members
ORDER BY id
`

func (q *Queries) ListMembers(ctx context.Context) ([]Member, error) {
	rows, err := q.db.QueryContext(ctx, listMembers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Member{}
	for rows.Next() {
		var i Member
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.PhoneNumber,
			&i.Latitude,
			&i.Longitude,
			&i.Address,
			&i.IsFamily,
			&i.CreatedAt,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMember = `-- name: UpdateMember :one
UPDATE members 
    SET latitude = $2, longitude = $3, full_name = $4, status = $5
WHERE id = $1
RETURNING id, full_name, phone_number, latitude, longitude, address, is_family, created_at, status
`

type UpdateMemberParams struct {
	ID        int64          `json:"id"`
	Latitude  float64        `json:"latitude"`
	Longitude float64        `json:"longitude"`
	FullName  sql.NullString `json:"full_name"`
	Status    NullUserStatus `json:"status"`
}

func (q *Queries) UpdateMember(ctx context.Context, arg UpdateMemberParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, updateMember,
		arg.ID,
		arg.Latitude,
		arg.Longitude,
		arg.FullName,
		arg.Status,
	)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.PhoneNumber,
		&i.Latitude,
		&i.Longitude,
		&i.Address,
		&i.IsFamily,
		&i.CreatedAt,
		&i.Status,
	)
	return i, err
}

const updateMemberName = `-- name: UpdateMemberName :one
UPDATE members 
    SET full_name = $2, status = $3
WHERE id = $1
RETURNING id, full_name, phone_number, latitude, longitude, address, is_family, created_at, status
`

type UpdateMemberNameParams struct {
	ID       int64          `json:"id"`
	FullName sql.NullString `json:"full_name"`
	Status   NullUserStatus `json:"status"`
}

func (q *Queries) UpdateMemberName(ctx context.Context, arg UpdateMemberNameParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, updateMemberName, arg.ID, arg.FullName, arg.Status)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.PhoneNumber,
		&i.Latitude,
		&i.Longitude,
		&i.Address,
		&i.IsFamily,
		&i.CreatedAt,
		&i.Status,
	)
	return i, err
}
