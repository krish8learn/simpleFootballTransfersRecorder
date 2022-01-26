// Code generated by sqlc. DO NOT EDIT.
// source: player.sql

package DB

import (
	"context"
)

const createplayer = `-- name: Createplayer :one
INSERT INTO player(
    player_name,
    position,
    country_pl,
    value,
    footballclub_id
)VALUES (
    $1, $2, $3, $4, $5
)
RETURNING p_id, player_name, position, country_pl, value, footballclub_id, created_at
`

type CreateplayerParams struct {
	PlayerName     string `json:"player_name"`
	Position       string `json:"position"`
	CountryPl      string `json:"country_pl"`
	Value          int64  `json:"value"`
	FootballclubID int32  `json:"footballclub_id"`
}

func (q *Queries) Createplayer(ctx context.Context, arg CreateplayerParams) (Player, error) {
	row := q.db.QueryRowContext(ctx, createplayer,
		arg.PlayerName,
		arg.Position,
		arg.CountryPl,
		arg.Value,
		arg.FootballclubID,
	)
	var i Player
	err := row.Scan(
		&i.PID,
		&i.PlayerName,
		&i.Position,
		&i.CountryPl,
		&i.Value,
		&i.FootballclubID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteplayer = `-- name: Deleteplayer :exec
DELETE FROM player
WHERE player_name = $1
`

func (q *Queries) Deleteplayer(ctx context.Context, playerName string) error {
	_, err := q.db.ExecContext(ctx, deleteplayer, playerName)
	return err
}

const getplayerByCountry = `-- name: GetplayerByCountry :many
SELECT p_id, player_name, position, country_pl, value, footballclub_id, created_at FROM player
WHERE country_pl = $1
`

func (q *Queries) GetplayerByCountry(ctx context.Context, countryPl string) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, getplayerByCountry, countryPl)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.PID,
			&i.PlayerName,
			&i.Position,
			&i.CountryPl,
			&i.Value,
			&i.FootballclubID,
			&i.CreatedAt,
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

const getplayerByFootballclub = `-- name: GetplayerByFootballclub :many
SELECT p_id, player_name, position, country_pl, value, footballclub_id, created_at FROM player
WHERE footballclub_id = $1
`

func (q *Queries) GetplayerByFootballclub(ctx context.Context, footballclubID int32) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, getplayerByFootballclub, footballclubID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.PID,
			&i.PlayerName,
			&i.Position,
			&i.CountryPl,
			&i.Value,
			&i.FootballclubID,
			&i.CreatedAt,
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

const getplayerByName = `-- name: GetplayerByName :one
SELECT p_id, player_name, position, country_pl, value, footballclub_id, created_at FROM player 
WHERE player_name = $1
`

func (q *Queries) GetplayerByName(ctx context.Context, playerName string) (Player, error) {
	row := q.db.QueryRowContext(ctx, getplayerByName, playerName)
	var i Player
	err := row.Scan(
		&i.PID,
		&i.PlayerName,
		&i.Position,
		&i.CountryPl,
		&i.Value,
		&i.FootballclubID,
		&i.CreatedAt,
	)
	return i, err
}

const getplayerByPosition = `-- name: GetplayerByPosition :many
SELECT p_id, player_name, position, country_pl, value, footballclub_id, created_at FROM player
WHERE position = $1
`

func (q *Queries) GetplayerByPosition(ctx context.Context, position string) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, getplayerByPosition, position)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.PID,
			&i.PlayerName,
			&i.Position,
			&i.CountryPl,
			&i.Value,
			&i.FootballclubID,
			&i.CreatedAt,
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

const getplayerByValueHigherthan = `-- name: GetplayerByValueHigherthan :many
SELECT p_id, player_name, position, country_pl, value, footballclub_id, created_at FROM player
WHERE value >= $1
`

func (q *Queries) GetplayerByValueHigherthan(ctx context.Context, value int64) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, getplayerByValueHigherthan, value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.PID,
			&i.PlayerName,
			&i.Position,
			&i.CountryPl,
			&i.Value,
			&i.FootballclubID,
			&i.CreatedAt,
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

const getplayerByValueLessthan = `-- name: GetplayerByValueLessthan :many
SELECT p_id, player_name, position, country_pl, value, footballclub_id, created_at FROM player 
WHERE value <= $1
`

func (q *Queries) GetplayerByValueLessthan(ctx context.Context, value int64) ([]Player, error) {
	rows, err := q.db.QueryContext(ctx, getplayerByValueLessthan, value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.PID,
			&i.PlayerName,
			&i.Position,
			&i.CountryPl,
			&i.Value,
			&i.FootballclubID,
			&i.CreatedAt,
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

const updateplayer = `-- name: Updateplayer :exec
UPDATE player
SET value = $2, footballclub_id =$3
WHERE p_id = $1
`

type UpdateplayerParams struct {
	PID            int32 `json:"p_id"`
	Value          int64 `json:"value"`
	FootballclubID int32 `json:"footballclub_id"`
}

func (q *Queries) Updateplayer(ctx context.Context, arg UpdateplayerParams) error {
	_, err := q.db.ExecContext(ctx, updateplayer, arg.PID, arg.Value, arg.FootballclubID)
	return err
}
