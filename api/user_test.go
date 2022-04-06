package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/krish8learn/simpleFootballTransfersRecorder/DB/mock"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
	"github.com/stretchr/testify/require"
)

func Test_newUserResp(t *testing.T) {
	type args struct {
		user DB.User
	}
	tests := []struct {
		name string
		args args
		want userResp
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				user: DB.User{},
			},
			want: userResp{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newUserResp(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newUserResp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testCreateUser(t *testing.T) (DB.User, string) {

	var passwordTest = Util.RandomStringGenerator(10)
	// password hashing
	hashedPassword, err := Util.HashPassword(passwordTest)
	require.NoError(t, err)
	return DB.User{
		Username:       Util.Randomusername(),
		HashedPassword: hashedPassword,
		FullName:       Util.Randomfullname(),
		Email:          Util.Randomemail(),
	}, passwordTest
}

type eqCreateUserParamsMatcher struct {
	arg      DB.CreateusersParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(DB.CreateusersParams)
	if !ok {
		return false
	}

	err := Util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg DB.CreateusersParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func requireBodyMatcher(t *testing.T, body *bytes.Buffer, user DB.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	result := fmt.Sprintf("User Created %v, Relogin", user.Username)
	require.Equal(t, result, strings.Trim(string(data), `"\"`))
}

func TestServer_Createusers(t *testing.T) {

	testUser, testPassword := testCreateUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(trans *mockdb.MockTransaction)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":  testUser.Username,
				"password":  testPassword,
				"full_name": testUser.FullName,
				"email":     testUser.Email,
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				arg := DB.CreateusersParams{
					Username: testUser.Username,
					FullName: testUser.FullName,
					Email:    testUser.Email,
				}
				trans.EXPECT().
					Createusers(gomock.Any(), EqCreateUserParams(arg, testPassword)).
					Times(1).
					Return(testUser, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatcher(t, recorder.Body, testUser)
			},
		},
		{
			name: "InvalidEmail",
			body: gin.H{
				"username":  testUser.Username,
				"password":  testPassword,
				"full_name": testUser.FullName,
				"email":     "invalid-email",
			},
			buildStubs: func(store *mockdb.MockTransaction) {
				store.EXPECT().
					Createusers(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TooShortPassword",
			body: gin.H{
				"username":  testUser.Username,
				"password":  "123",
				"full_name": testUser.FullName,
				"email":     testUser.Email,
			},
			buildStubs: func(store *mockdb.MockTransaction) {
				store.EXPECT().
					Createusers(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
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

			server := NewServer(trans, testConfig)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/user/createUser"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}

func TestServer_Loginuser(t *testing.T) {
	testUser, testPassword := testCreateUser(t)

	testCases := []struct {
		name          string
		body          loginUserRequest
		buildStubs    func(trans *mockdb.MockTransaction)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: loginUserRequest{
				Username: testUser.Username,
				Password: testPassword,
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().GetUsers(gomock.Any(), gomock.Eq(testUser.Username)).Times(1).Return(testUser, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "User Not found",
			body: loginUserRequest{
				Username: testUser.Username,
				Password: testPassword,
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().GetUsers(gomock.Any(), gomock.Eq(testUser.Username)).Times(1).Return(DB.User{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "Invalid Password",
			body: loginUserRequest{
				Username: testUser.Username,
				Password: "incorrect",
			},
			buildStubs: func(trans *mockdb.MockTransaction) {
				trans.EXPECT().GetUsers(gomock.Any(), gomock.Eq(testUser.Username)).Times(1).Return(testUser, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
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

			server := NewServer(trans, testConfig)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/user/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
