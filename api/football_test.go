package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
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

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	duration time.Duration,
) {
	tokenString, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, tokenString)
	request.Header.Set("Authorization", authorizationHeader)
}

func TestNameFootballclub(t *testing.T) {

	testFootballclub := FootballclubDataTest()

	testCases := []struct {
		name                  string
		inputFootballclubName string
		setupAuth             func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration)
		buildStubs            func(store *mockdb.MockTransaction)
		checkResponse         func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:                  "TestCase1 -- OK",
			inputFootballclubName: testFootballclub.ClubName,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration) {
				addAuthorization(t, request, tokenMaker, authorizationType, user, duration)
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().GetfootballclubByName(gomock.Any(), gomock.Eq(testFootballclub.ClubName)).Times(1).Return(testFootballclub, nil)
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
			tServer := NewServer(trans, testConfig)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/footballclub/nameFootballclub/%v", tc.inputFootballclubName)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, tServer.TokenMaker, testAuthorizationTypeBearer, testUser, testTime)

			tServer.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}
}

func TestCreateFootballclub(t *testing.T) {
	//testing football club data to preapare a request
	footballClubData := FootballclubDataTest()

	//test cases
	testCases := []struct {
		name          string
		inputData     createFootballclub
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration)
		buildStubs    func(store *mockdb.MockTransaction)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "TestCase1 -- OK",
			inputData: createFootballclub{
				ClubName:  footballClubData.ClubName,
				CountryFc: footballClubData.CountryFc,
				Balance:   footballClubData.Balance,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType, user string, duration time.Duration) {
				addAuthorization(t, request, tokenMaker, authorizationType, user, duration)
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				arg := DB.CreatefootballclubParams{
					ClubName:  footballClubData.ClubName,
					CountryFc: footballClubData.CountryFc,
					Balance:   footballClubData.Balance,
				}
				trans.EXPECT().GetfootballclubByName(gomock.Any(), gomock.Eq(arg.ClubName)).Times(1).Return(DB.Footballclub{}, sql.ErrNoRows)
				trans.EXPECT().Createfootballclub(gomock.Any(), gomock.Eq(arg)).Times(1).Return(footballClubData, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		// {},
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
			tServer := NewServer(trans, testConfig)
			recorder := httptest.NewRecorder()
			url := "/footballclub/createFootballclub"
			data, err := json.Marshal(tc.inputData)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, tServer.TokenMaker, testAuthorizationTypeBearer, testUser, testTime)

			tServer.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func FootballclubDataTest() DB.Footballclub {
	return DB.Footballclub{
		FcID:      int32(Util.RandomIntGenerator(10, 1000)),
		ClubName:  Util.RandomfootballclubName(),
		CountryFc: Util.RandomCountryPl(),
		Balance:   Util.Randombalance(),
	}
}
