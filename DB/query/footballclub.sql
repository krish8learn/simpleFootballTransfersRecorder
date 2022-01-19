-- name: Createfootballclub :one
INSERT INTO footballclub (
  club_name,
  country_fc,
  balance
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: Getfootballclub :one
SELECT * FROM footballclub
WHERE fc_id = $1 LIMIT 1;

-- name: Listfootballclub :many
SELECT * FROM footballclub
ORDER BY fc_id LIMIT $1 OFFSET $2;

-- name: UpdatefootballclubBalance :exec
UPDATE footballclub
SET balance = $2
WHERE fc_id = $1;

-- name: Deletefootballclub :exec
DELETE FROM footballclub
WHERE club_name = $1;
