package DB

import (
	"context"
	"log"

	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
)

func CreateTestfootballclubdataSource() Footballclub {
	testFootballClubSource, testFootballClubSouceErr := testQueries.Createfootballclub(context.Background(), CreatefootballclubParams{
		ClubName:  Util.RandomfootballclubName(),
		CountryFc: Util.Randomcountryfc(),
		Balance:   Util.Randombalance(),
	})

	if testFootballClubSouceErr != nil {
		log.Fatalln("Cannot create footballclub Data in DB for testing", testFootballClubSouceErr)
	}
	return testFootballClubSource
}

func CreateTestfootballclubdataDestionation() Footballclub {
	testFootballClubDestionation, testFootballClubErrDestionation := testQueries.Createfootballclub(context.Background(), CreatefootballclubParams{
		ClubName:  Util.RandomfootballclubName(),
		CountryFc: Util.Randomcountryfc(),
		Balance:   Util.Randombalance(),
	})

	if testFootballClubErrDestionation != nil {
		log.Fatalln("Cannot create footballclub Data in DB for testing", testFootballClubErrDestionation)
	}
	return testFootballClubDestionation
}

func CreateTestplayerdata() Player {
	testFootballClubSource := CreateTestfootballclubdataSource()

	args := CreateplayerParams{
		PlayerName:     Util.RandomPlayername(),
		Position:       Util.RandomPosition(),
		CountryPl:      Util.RandomCountryPl(),
		Value:          Util.RandomPlayerValue(),
		FootballclubID: testFootballClubSource.FcID,
	}

	testPlayerData, testErrData := testQueries.Createplayer(context.Background(), args)

	if testErrData != nil {
		log.Fatalln("Cannot create player Data in DB for testing", testErrData)
	}
	return testPlayerData

}

func CreateTestfootballclubdata() Footballclub {
	testFootballClub, testFootballClubErr := testQueries.Createfootballclub(context.Background(), CreatefootballclubParams{
		ClubName:  Util.RandomfootballclubName(),
		CountryFc: Util.Randomcountryfc(),
		Balance:   Util.Randombalance(),
	})

	if testFootballClubErr != nil {
		log.Fatalln("Cannot create footballclub Data in DB for testing", testFootballClubErr)
	}
	return testFootballClub
}

func CreateTestTransferdata() Transfer {
	testPlayerData = CreateTestplayerdata()
	testFootballClubDestionation = CreateTestfootballclubdataDestionation()

	arg := CreatetransferParams{
		Season:          Util.RandomSeason(),
		PlayerID:        testPlayerData.PID,
		SourceClub:      testPlayerData.FootballclubID,
		DestinationClub: testFootballClubDestionation.FcID,
		Amount:          Util.RandomAmount(),
	}

	testTransfer, testTransferErr := testQueries.Createtransfer(context.Background(), arg)

	if testTransferErr != nil {
		log.Fatalln("Cannot create transfer Data in DB for testing", testTransferErr)
	}

	return testTransfer
}

func CreateTestUser() User {
	arg := CreateusersParams{
		Username:       Util.Randomusername(),
		HashedPassword: "",
		FullName:       Util.Randomfullname(),
		Email:          Util.Randomemail(),
	}

	user, testErr := testQueries.Createusers(context.Background(), arg)

	if testErr != nil {
		log.Fatalln("Cannot create user Data in DB for testing")
	}

	return user
}
