package token

import (
	"testing"
	"time"

	"github.com/krish8learn/simpleFootballTransfersRecorder/Util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	symmetricKey := Util.RandomStringGenerator(32)
	maker, err := NewPasetoMaker(symmetricKey)

	require.NoError(t, err)
	username := Util.RandomStringGenerator(32)
	duration := time.Minute
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	tokenString, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	payload, err := maker.VerifyToken(tokenString)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, username, payload.UserName)
	require.NotZero(t, payload.ID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Minute)
	require.WithinDuration(t, expiresAt, payload.ExpiredAt, time.Minute)
}
