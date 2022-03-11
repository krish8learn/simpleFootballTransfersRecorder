package token

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
	"github.com/stretchr/testify/require"
)

func TestNewJWTMaker(t *testing.T) {
	secretKey := Util.RandomStringGenerator(32)
	type args struct {
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		want    Maker
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestCase1",
			args: args{
				secretKey: secretKey,
			},
			want: &JWTMaker{
				secretKey: secretKey,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewJWTMaker(tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJWTMaker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJWTMaker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTMaker_CreateToken(t *testing.T) {
	maker, err := NewJWTMaker(Util.RandomStringGenerator(32))
	require.NoError(t, err)

	tokenString, err := maker.CreateToken(Util.Randomusername(), time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, tokenString)
}

func TestJWTMaker_VerifyToken(t *testing.T) {
	type fields struct {
		secretKey string
	}
	type args struct {
		token string
	}
	secretKey := &fields{
		secretKey: Util.RandomStringGenerator(32),
	}
	id, _ := uuid.NewRandom()
	username := Util.Randomusername()
	issueTime := time.Now()
	expiredTime := time.Now().Add(time.Minute)
	maker, _ := NewJWTMaker(secretKey.secretKey)
	token, _ := maker.CreateToken(username, time.Minute)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Payload
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestCase1",
			fields: *secretKey,
			args: args{
				token: token,
			},
			want: &Payload{
				ID:        id,
				UserName:  username,
				IssuedAt:  issueTime,
				ExpiredAt: expiredTime,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maker := &JWTMaker{
				secretKey: tt.fields.secretKey,
			}
			got, err := maker.VerifyToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTMaker.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.UserName, tt.want.UserName) {
				t.Errorf("JWTMaker.VerifyToken() = %v, want %v", got, tt.want)
			}
			require.NotZero(t, got.ID)
			
		})
	}
}
