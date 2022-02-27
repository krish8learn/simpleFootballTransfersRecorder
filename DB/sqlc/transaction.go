package DB

import (
	"context"
	"database/sql"
	"fmt"
)

//we we will execute transaction using this struct
type Transaction struct {
	*Queries
	db *sql.DB
}

//creates new object of Transaction struct
func NewTransaction(db *sql.DB) *Transaction {
	return &Transaction{
		db:      db,
		Queries: New(db),
	}
}

//execute multiple DB queries in one function to complete one transaction
func (transaction *Transaction) execTx(ctx context.Context, fn func(*Queries) error) error {

	//begin
	tx, txErr := transaction.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}

	queryObject := New(tx)

	queryErr := fn(queryObject)
	if queryErr != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return fmt.Errorf("tx error: %v, rollBack error: %v ", queryErr, rollBackErr)
		}
		return fmt.Errorf("tx error, while creating query object: %v", rollBackErr)
	}

	//commit
	commitErr := tx.Commit()
	if commitErr != nil {
		return fmt.Errorf("tx error, while commiting: %v", commitErr)
	}

	return nil

}

//this contain  the parameters of creating a transfer record
type TransferTxParams struct {
	Season            int64 `json:"season"`
	PlayerID          int32 `json:"player_id"`
	SourceClubID      int32 `json:"source_club"`
	DestinationClubID int32 `json:"destination_club"`
	Amount            int64 `json:"amount"`
}

//this contain results after all the transaction
type TransferTxResult struct {
	Transfer        Transfer     `json:"transfer"`
	SourceClub      Footballclub `json:"source_footballclub"`
	DestinationClub Footballclub `json:"destination_footballclub"`
	Player          Player       `json:"player"`
}

func (t Transaction) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	execTxErr := t.execTx(ctx, func(q *Queries) error {

		var err error

		//performing transaction
		//creating transfer Record
		result.Transfer, err = q.Createtransfer(ctx, CreatetransferParams{
			Season:          arg.Season,
			PlayerID:        arg.PlayerID,
			SourceClub:      arg.SourceClubID,
			DestinationClub: arg.DestinationClubID,
			Amount:          arg.Amount,
		})

		if err != nil {
			return err
		}

		//updating player data(his club, value) based on transfer data
		err = q.Updateplayer(ctx, UpdateplayerParams{
			PID:            arg.PlayerID,
			Value:          arg.Amount,
			FootballclubID: arg.DestinationClubID,
		})

		if err != nil {
			return nil
		}

		//getting data of source club
		getSourceClubData, err := q.GetfootballclubByID(ctx, arg.SourceClubID)

		if err != nil {
			return nil
		}

		//updating balance of selling club
		err = q.UpdatefootballclubBalance(ctx, UpdatefootballclubBalanceParams{
			FcID:    arg.SourceClubID,
			Balance: getSourceClubData.Balance + arg.Amount,
		})

		//getting data of destination club
		getDestinationClubData, err := q.GetfootballclubByID(ctx, arg.DestinationClubID)

		if err != nil {
			return nil
		}

		//updating balance of buying  club
		err = q.UpdatefootballclubBalance(ctx, UpdatefootballclubBalanceParams{
			FcID:    arg.DestinationClubID,
			Balance: getDestinationClubData.Balance - arg.Amount,
		})

		//get player data inside the result
		result.Player, _ = q.GetplayerByID(ctx, arg.PlayerID)
		//get sourceClub in result
		result.SourceClub, _ = q.GetfootballclubByID(ctx, arg.SourceClubID)
		//get destinationClub into result
		result.DestinationClub, _ = q.GetfootballclubByID(ctx, arg.DestinationClubID)

		return nil
	})

	return result, execTxErr
}
