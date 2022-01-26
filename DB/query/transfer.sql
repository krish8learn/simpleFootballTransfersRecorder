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


-- name: GettransferByPlayerid :one
SELECT * FROM transfer
WHERE player_id = $1 LIMIT 1; 


-- name: Updatetransfer :exec
UPDATE transfer
SET amount = $2
WHERE t_id = $1;


-- name: Deletetransfer :exec
DELETE FROM transfer
WHERE t_id = $1;