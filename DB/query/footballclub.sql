-- name: Createfootballclub :one
INSERT INTO footballclubs (
  club_name,
  country_fc,
  balance
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetfootballclubByID :one
SELECT * FROM footballclubs
WHERE fc_id = $1 LIMIT 1;

-- name: GetfootballclubByName :one 
SELECT * FROM footballclubs
WHERE club_name = $1;

-- name: GetfootballclubByCountry :many
SELECT * FROM footballclubs
WHERE country_fc = $1;

-- name: Listfootballclub :many
SELECT * FROM footballclubs
ORDER BY fc_id OFFSET $1 LIMIT $2;

-- name: UpdatefootballclubBalance :exec
UPDATE footballclubs
SET balance = $2
WHERE fc_id = $1;

-- name: Deletefootballclub :exec
DELETE FROM footballclubs
WHERE club_name = $1;
