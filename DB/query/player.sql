-- name: Createplayer :one 
INSERT INTO players(
    player_name,
    position,
    country_pl,
    value,
    footballclub_id
)VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetplayerByName :one
SELECT * FROM players 
WHERE player_name = $1;

-- name: GetplayerByID :one
SELECT * FROM players 
WHERE p_id = $1;

-- name: GetPlayersList :many
SELECT * FROM players
ORDER BY p_id OFFSET $1 LIMIT $2;

-- name: GetplayerByValueHigherthan :many
SELECT * FROM players
WHERE value >= $1;

-- name: GetplayerByValueLessthan :many
SELECT * FROM players 
WHERE value <= $1;

-- name: GetplayerByFootballclub :many
SELECT * FROM players
WHERE footballclub_id = $1;

-- name: GetplayerByCountry :many
SELECT * FROM players
WHERE country_pl = $1;

-- name: GetplayerByPosition :many
SELECT * FROM players
WHERE position = $1; 

-- name: DeletePlayerByClubID :exec
DELETE FROM players
WHERE footballclub_id = $1;

-- name: Deleteplayer :exec
DELETE FROM players
WHERE player_name = $1;

-- name: Updateplayer :exec
UPDATE players
SET value = $2, footballclub_id =$3
WHERE p_id = $1;