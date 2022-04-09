package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockdb "github.com/krish8learn/simpleFootballTransfersRecorder/DB/mock"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
	"github.com/krish8learn/simpleFootballTransfersRecorder/token"
	"github.com/stretchr/testify/require"
)

func testTransferData(t *testing.T) (*[]DB.Footballclub, *DB.Player, *[]DB.Transfer) {

	sendClub := DB.Footballclub{
		FcID:      int32(Util.RandomIntGenerator(10, 1000)),
		ClubName:  Util.RandomfootballclubName(),
		CountryFc: Util.RandomCountryPl(),
		Balance:   Util.Randombalance(),
	}

	receiveClub := DB.Footballclub{
		FcID:      int32(Util.RandomIntGenerator(10, 1000)),
		ClubName:  Util.RandomfootballclubName(),
		CountryFc: Util.RandomCountryPl(),
		Balance:   Util.Randombalance(),
	}
	clubs := []DB.Footballclub{sendClub, receiveClub}

	player := DB.Player{
		PID:            int32(Util.RandomIntGenerator(1, 1000)),
		PlayerName:     Util.RandomPlayername(),
		Position:       Util.RandomPosition(),
		CountryPl:      Util.RandomCountryPl(),
		Value:          Util.RandomPlayerValue(),
		FootballclubID: sendClub.FcID,
	}

	transfers := []DB.Transfer{
		{
			TID:             int32(Util.RandomIntGenerator(10, 1000)),
			Season:          Util.RandomSeason(),
			PlayerID:        player.PID,
			SourceClub:      sendClub.FcID,
			DestinationClub: receiveClub.FcID,
			Amount:          Util.RandomAmount(),
		},
	}

	return &clubs, &player, &transfers
}

func TestPlayerNameTransfer(t *testing.T) {
	_, testPlayer, testTransfer := testTransferData(t)

	testCases := []struct {
		name          string
		input         string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration)
		buildStubs    func(trans *mockdb.MockTransaction)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:  "OK",
			input: testPlayer.PlayerName,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration) {
				addAuthorization(t, request, tokenMaker, authorizationType, user, duration)
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().GetplayerByName(gomock.Any(), gomock.Eq(testPlayer.PlayerName)).Times(1).Return(*testPlayer, nil)
				trans.EXPECT().GettransferByPlayerid(gomock.Any(), gomock.Eq(testPlayer.PID)).Times(1).Return(*testTransfer, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:  "Not found",
			input: "No Player",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration) {
				addAuthorization(t, request, tokenMaker, authorizationType, user, duration)
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().GetplayerByName(gomock.Any(), gomock.Eq("No Player")).Times(1).Return(DB.Player{}, sql.ErrNoRows)
				// trans.EXPECT().GettransferByPlayerid(gomock.Any(), gomock.Eq(testPlayer.PID)).Times(1).Return(*testTransfer, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//create a controller for mocking
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			trans := mockdb.NewMockTransaction(ctrl)
			tc.buildStubs(trans)

			testConfig := Util.Config{
				SecurityKey: Util.RandomStringGenerator(32),
				AccessTime:  testTime,
			}

			tserver := NewServer(trans, testConfig)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/transfer/playerNameTransfer/%v", tc.input)
			request, err := http.NewRequest(http.MethodGet, url, nil)

			require.NoError(t, err)

			tc.setupAuth(t, request, tserver.TokenMaker, testAuthorizationTypeBearer, testUser, testTime)

			tserver.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestMaxTransfer(t *testing.T) {
	_, _, testTransfer := testTransferData(t)
	testCases := []struct {
		name string
		// input         string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration)
		buildStubs    func(trans *mockdb.MockTransaction)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration) {
				addAuthorization(t, request, tokenMaker, authorizationType, user, duration)
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().Highesttransfer(gomock.Any()).Times(1).Return((*testTransfer)[0], nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//create a controller for mocking
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			trans := mockdb.NewMockTransaction(ctrl)
			tc.buildStubs(trans)

			testConfig := Util.Config{
				SecurityKey: Util.RandomStringGenerator(32),
				AccessTime:  testTime,
			}

			tserver := NewServer(trans, testConfig)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/transfer/maxTransfer")
			request, err := http.NewRequest(http.MethodGet, url, nil)

			require.NoError(t, err)

			tc.setupAuth(t, request, tserver.TokenMaker, testAuthorizationTypeBearer, testUser, testTime)

			tserver.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
