-- name: Createtransfer :one 
INSERT INTO transfer(
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
SELECT * FROM transfer
WHERE t_id = $1 LIMIT 1;


-- name: GettransferList :many
SELECT * FROM transfer
ORDER BY t_id OFFSET $1 LIMIT $2;


-- name: GettransferByPlayerid :many
SELECT * FROM transfer
WHERE player_id = $1; 


-- name: Updatetransfer :exec
UPDATE transfer
SET amount = $2
WHERE t_id = $1;


-- name: Highesttransfer :one
SELECT * FROM transfer
WHERE amount = (
    SELECT MAX(amount)
    FROM transfer
);

-- name: Latesttransfer :one 
SELECT * FROM transfer
WHERE player_id = $1 AND destination_club = $2
ORDER BY created_at DESC 
LIMIT 1;

-- name: Deletetransfer :exec
DELETE FROM transfer
WHERE t_id = $1;

-- name: GetLasttransferByPlayerid :one
SELECT * FROM transfer
WHERE player_id = $1
ORDER BY created_at DESC 
LIMIT 1;; 