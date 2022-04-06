package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func TestNamePlayer(t *testing.T) {
	playerData := PlayerDataTest()

	testCases := []struct {
		name            string
		inputPlayerName string
		setupAuth       func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration)
		buildStubs      func(trans *mockdb.MockTransaction)
		checkResponse   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:            "TestCase1 -- OK",
			inputPlayerName: playerData.PlayerName,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration) {
				addAuthorization(t, request, tokenMaker, authorizationType, user, duration)
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().GetplayerByName(gomock.Any(), gomock.Eq(playerData.PlayerName)).Times(1).Return(playerData, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:            "TestCase2 -- Not Found",
			inputPlayerName: playerData.PlayerName,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration) {
				addAuthorization(t, request, tokenMaker, authorizationType, user, duration)
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().GetplayerByName(gomock.Any(), gomock.Eq(playerData.PlayerName)).Times(1).Return(DB.Player{}, sql.ErrNoRows)
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
			url := fmt.Sprintf("/player/namePlayer/%v", tc.inputPlayerName)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, tserver.TokenMaker, testAuthorizationTypeBearer, testUser, testTime)

			tserver.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListPlayer(t *testing.T) {
	var listTestPlayers []DB.Player

	for i := 0; i < 10; i++ {
		listTestPlayers = append(listTestPlayers, PlayerDataTest())
	}

	testCases := []struct {
		name          string
		input         listPlayers
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration)
		buildStubs    func(trans *mockdb.MockTransaction)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "TestCase1 -- OK",
			input: listPlayers{
				PageID:   2,
				PageSize: 5,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration) {
				addAuthorization(t, request, tokenMaker, authorizationType, user, duration)
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				arg := DB.GetPlayersListParams{
					Offset: (2 - 1) * 5,
					Limit:  5,
				}
				trans.EXPECT().GetPlayersList(gomock.Any(), gomock.Eq(arg)).Times(1).Return(listTestPlayers, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "TestCase2 -- Bad Request -- Invalid PageID",
			input: listPlayers{
				PageID:   0,
				PageSize: 4,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration) {
				addAuthorization(t, request, tokenMaker, authorizationType, user, duration)
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().GetPlayersList(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TestCase2 -- Bad Request -- Invalid PageSize",
			input: listPlayers{
				PageID:   2,
				PageSize: 10000,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration) {
				addAuthorization(t, request, tokenMaker, authorizationType, user, duration)
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().GetPlayersList(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
			url := fmt.Sprintf("/player/listPlayers")
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.input.PageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.input.PageSize))
			request.URL.RawQuery = q.Encode()

			tc.setupAuth(t, request, tserver.TokenMaker, testAuthorizationTypeBearer, testUser, testTime)

			tserver.Router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestUpdatePlayer(t *testing.T) {
	amount := Util.RandomAmount()
	testPlayer, testClub := updatePlayerData()

	testCases := []struct {
		name          string
		input         updateValuePlayer
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration)
		buildStubs    func(trans *mockdb.MockTransaction)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "TestCase1 -- OK",
			input: updateValuePlayer{
				PlayerName:       testPlayer.PlayerName,
				Value:            amount,
				FootballclubName: testClub.ClubName,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration) {
				addAuthorization(t, request, tokenMaker, authorizationType, user, duration)
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().GetplayerByName(gomock.Any(), gomock.Eq(testPlayer.PlayerName)).Times(1).Return(testPlayer, nil)
				trans.EXPECT().GetfootballclubByName(gomock.Any(), gomock.Eq(testClub.ClubName)).Times(1).Return(testClub, nil)
				arg := DB.UpdateplayerParams{
					PID:            testPlayer.PID,
					Value:          amount,
					FootballclubID: testClub.FcID,
				}
				trans.EXPECT().Updateplayer(gomock.Any(), gomock.Eq(arg)).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
			url := fmt.Sprintf("/player/updatevaluePlayer")
			data, err := json.Marshal(tc.input)
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, tserver.TokenMaker, testAuthorizationTypeBearer, testUser, testTime)

			tserver.Router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func PlayerDataTest() DB.Player {
	footballClub := FootballclubDataTest()
	return DB.Player{
		PID:            int32(Util.RandomIntGenerator(1, 1000)),
		PlayerName:     Util.RandomPlayername(),
		Position:       Util.RandomPosition(),
		CountryPl:      Util.RandomCountryPl(),
		Value:          Util.RandomPlayerValue(),
		FootballclubID: footballClub.FcID,
	}
}

func updatePlayerData() (DB.Player, DB.Footballclub) {
	footballClub := FootballclubDataTest()
	player := DB.Player{
		PID:            int32(Util.RandomIntGenerator(1, 1000)),
		PlayerName:     Util.RandomPlayername(),
		Position:       Util.RandomPosition(),
		CountryPl:      Util.RandomCountryPl(),
		Value:          Util.RandomPlayerValue(),
		FootballclubID: footballClub.FcID,
	}
	return player, footballClub
}

func requireBodyMatchAccounts(t *testing.T, body *bytes.Buffer, accounts []DB.Player) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccounts []DB.Player
	err = json.Unmarshal(data, &gotAccounts)
	require.NoError(t, err)
	require.Equal(t, accounts, gotAccounts)
}
