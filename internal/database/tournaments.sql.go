// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: tournaments.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createTournament = `-- name: CreateTournament :one
INSERT INTO tournaments (
  name,
  created_at
) VALUES (?, ?) RETURNING id, name, created_at, updated_at
`

type CreateTournamentParams struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func (q *Queries) CreateTournament(ctx context.Context, arg CreateTournamentParams) (Tournament, error) {
	row := q.db.QueryRowContext(ctx, createTournament, arg.Name, arg.CreatedAt)
	var i Tournament
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTournamentByID = `-- name: DeleteTournamentByID :exec
DELETE FROM tournaments WHERE id = ?
`

func (q *Queries) DeleteTournamentByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTournamentByID, id)
	return err
}

const getAllTournaments = `-- name: GetAllTournaments :many
SELECT id, name, created_at, updated_at FROM tournaments
ORDER BY id DESC
`

func (q *Queries) GetAllTournaments(ctx context.Context) ([]Tournament, error) {
	rows, err := q.db.QueryContext(ctx, getAllTournaments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tournament
	for rows.Next() {
		var i Tournament
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getTournamentByID = `-- name: GetTournamentByID :one
SELECT id, name, created_at, updated_at FROM tournaments
WHERE id = ?
`

func (q *Queries) GetTournamentByID(ctx context.Context, id int64) (Tournament, error) {
	row := q.db.QueryRowContext(ctx, getTournamentByID, id)
	var i Tournament
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTournamentByID = `-- name: UpdateTournamentByID :one
UPDATE tournaments
SET name = ?, updated_at = ?
WHERE id = ?
RETURNING id, name, created_at, updated_at
`

type UpdateTournamentByIDParams struct {
	Name      string       `json:"name"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
	ID        int64        `json:"id"`
}

func (q *Queries) UpdateTournamentByID(ctx context.Context, arg UpdateTournamentByIDParams) (Tournament, error) {
	row := q.db.QueryRowContext(ctx, updateTournamentByID, arg.Name, arg.UpdatedAt, arg.ID)
	var i Tournament
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}