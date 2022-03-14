-- name: Createtransfer :one 
INSERT INTO transfers(
   season,
   player_id,
   source_club,
   destination_club,
   amount
)VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;


-- name: GettransferByTransferid :one
SELECT * FROM transfers
WHERE t_id = $1 LIMIT 1;


-- name: GettransferList :many
SELECT * FROM transfers
ORDER BY t_id OFFSET $1 LIMIT $2;


-- name: GettransferByPlayerid :many
SELECT * FROM transfers
WHERE player_id = $1; 


-- name: Updatetransfer :exec
UPDATE transfers
SET amount = $2
WHERE t_id = $1;


-- name: Highesttransfer :one
SELECT * FROM transfers
WHERE amount = (
    SELECT MAX(amount)
    FROM transfers
);

-- name: Latesttransfer :one 
SELECT * FROM transfers
WHERE player_id = $1 AND destination_club = $2
ORDER BY created_at DESC 
LIMIT 1;

-- name: Deletetransfer :exec
DELETE FROM transfers
WHERE t_id = $1;

-- name: GetLasttransferByPlayerid :one
SELECT * FROM transfers
WHERE player_id = $1
ORDER BY created_at DESC 
LIMIT 1;; 