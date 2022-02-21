package DB

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
	"github.com/stretchr/testify/require"
)

func TestTransaction_TransferTx(t *testing.T) {
	testTransact := NewTransaction(testDbConnect)

	sellingClubData := CreateTestfootballclubdataSource()
	buyingClubData := CreateTestfootballclubdataDestionation()
	playerData, testErrData := testQueries.Createplayer(context.Background(), CreateplayerParams{
		PlayerName:     Util.RandomPlayername(),
		Position:       Util.RandomPosition(),
		CountryPl:      Util.RandomCountryPl(),
		Value:          Util.RandomPlayerValue(),
		FootballclubID: sellingClubData.FcID,
	})

	if testErrData != nil {
		log.Fatalln("Database Error", testErrData)
	}

	//run concurent transaction
	transactionCount := 2
	amount := int64(100000)

	errsChan := make(chan error)
	resultChan := make(chan TransferTxResult)
	for i := 0; i < transactionCount; i++ {
		go func() {
			result, err := testTransact.TransferTx(context.Background(), TransferTxParams{
				Season:            2021,
				PlayerID:          playerData.PID,
				SourceClubID:      sellingClubData.FcID,
				DestinationClubID: buyingClubData.FcID,
				Amount:            amount,
			})

			errsChan <- err
			resultChan <- result
		}()
	}

	for i := 0; i < transactionCount; i++ {
		err := <-errsChan
		fmt.Println(err)
		require.NoError(t, err)

		results := <-resultChan
		require.NotEmpty(t, results)
		require.Equal(t, buyingClubData.FcID, results.Player.FootballclubID)
		require.Equal(t, sellingClubData.FcID, results.SourceClub.FcID)
		require.Equal(t, buyingClubData.FcID, results.DestinationClub.FcID)
		require.Equal(t, amount, results.Transfer.Amount)

		require.NotZero(t, results.Transfer.TID)
		//checking data
	}
}
