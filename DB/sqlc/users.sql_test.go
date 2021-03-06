// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package DB

import (
	"context"
	"reflect"
	"testing"

	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
	"github.com/stretchr/testify/require"
)

func TestQueries_Createusers(t *testing.T) {
	// type fields struct {
	// 	db DBTX
	// }
	hashedPassword, err := Util.HashPassword(Util.RandomStringGenerator(6))
	require.NoError(t, err)
	type args struct {
		ctx context.Context
		arg CreateusersParams
	}
	argument :=args{
		ctx: context.Background(),
		arg: CreateusersParams{
			Username:Util.Randomusername(),
			HashedPassword: hashedPassword,
			FullName: Util.Randomfullname(),
			Email: Util.Randomemail(),
		},
	}
	tests := []struct {
		name    string
		// fields  fields
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestCase1",
			args: argument,
			want: User{
				Username: argument.arg.Username,
				HashedPassword: "",
				FullName: argument.arg.FullName,
				Email: argument.arg.Email,
			},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// q := &Queries{
			// 	db: tt.fields.db,
			// }
			got, err := testQueries.Createusers(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.Createusers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Username, tt.want.Username) {
				t.Errorf("Queries.Createusers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_GetUsers(t *testing.T) {
	// type fields struct {
	// 	db DBTX
	// }
	//create an user
	user := CreateTestUser()

	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		// fields  fields
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestCase1",
			args: args{
				ctx: context.Background(),
				username: user.Username,
			},
			want: user, 
			wantErr : false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// q := &Queries{
			// 	db: tt.fields.db,
			// }
			got, err := testQueries.GetUsers(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
