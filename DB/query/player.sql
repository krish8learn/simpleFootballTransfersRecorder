-- name: Createplayer :one 
INSERT INTO player(
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
SELECT * FROM player 
WHERE player_name = $1;

-- name: GetplayerByValueHigherthan :many
SELECT * FROM player
WHERE value >= $1;

-- name: GetplayerByValueLessthan :many
SELECT * FROM player 
WHERE value <= $1;

-- name: GetplayerByFootballclub :many
SELECT * FROM player
WHERE footballclub_id = $1;

-- name: GetplayerByCountry :many
SELECT * FROM player
WHERE country_pl = $1;

-- name: GetplayerByPosition :many
SELECT * FROM player
WHERE position = $1; 

-- name: Deleteplayer :exec
DELETE FROM player
WHERE player_name = $1;

-- name: Updateplayer :exec
UPDATE player
SET value = $2, footballclub_id =$3
WHERE p_id = $1;