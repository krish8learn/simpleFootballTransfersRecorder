// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc (interfaces: Transaction)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	DB "github.com/krish8learn/simpleFootballTransfersRecorder/DB/sqlc"
)

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// Createfootballclub mocks base method.
func (m *MockTransaction) Createfootballclub(arg0 context.Context, arg1 DB.CreatefootballclubParams) (DB.Footballclub, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Createfootballclub", arg0, arg1)
	ret0, _ := ret[0].(DB.Footballclub)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Createfootballclub indicates an expected call of Createfootballclub.
func (mr *MockTransactionMockRecorder) Createfootballclub(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Createfootballclub", reflect.TypeOf((*MockTransaction)(nil).Createfootballclub), arg0, arg1)
}

// Createplayer mocks base method.
func (m *MockTransaction) Createplayer(arg0 context.Context, arg1 DB.CreateplayerParams) (DB.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Createplayer", arg0, arg1)
	ret0, _ := ret[0].(DB.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Createplayer indicates an expected call of Createplayer.
func (mr *MockTransactionMockRecorder) Createplayer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Createplayer", reflect.TypeOf((*MockTransaction)(nil).Createplayer), arg0, arg1)
}

// Createtransfer mocks base method.
func (m *MockTransaction) Createtransfer(arg0 context.Context, arg1 DB.CreatetransferParams) (DB.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Createtransfer", arg0, arg1)
	ret0, _ := ret[0].(DB.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Createtransfer indicates an expected call of Createtransfer.
func (mr *MockTransactionMockRecorder) Createtransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Createtransfer", reflect.TypeOf((*MockTransaction)(nil).Createtransfer), arg0, arg1)
}

// Createusers mocks base method.
func (m *MockTransaction) Createusers(arg0 context.Context, arg1 DB.CreateusersParams) (DB.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Createusers", arg0, arg1)
	ret0, _ := ret[0].(DB.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Createusers indicates an expected call of Createusers.
func (mr *MockTransactionMockRecorder) Createusers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Createusers", reflect.TypeOf((*MockTransaction)(nil).Createusers), arg0, arg1)
}

// DeletePlayerByClubID mocks base method.
func (m *MockTransaction) DeletePlayerByClubID(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlayerByClubID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlayerByClubID indicates an expected call of DeletePlayerByClubID.
func (mr *MockTransactionMockRecorder) DeletePlayerByClubID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlayerByClubID", reflect.TypeOf((*MockTransaction)(nil).DeletePlayerByClubID), arg0, arg1)
}

// Deletefootballclub mocks base method.
func (m *MockTransaction) Deletefootballclub(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deletefootballclub", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deletefootballclub indicates an expected call of Deletefootballclub.
func (mr *MockTransactionMockRecorder) Deletefootballclub(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deletefootballclub", reflect.TypeOf((*MockTransaction)(nil).Deletefootballclub), arg0, arg1)
}

// Deleteplayer mocks base method.
func (m *MockTransaction) Deleteplayer(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deleteplayer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deleteplayer indicates an expected call of Deleteplayer.
func (mr *MockTransactionMockRecorder) Deleteplayer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deleteplayer", reflect.TypeOf((*MockTransaction)(nil).Deleteplayer), arg0, arg1)
}

// Deletetransfer mocks base method.
func (m *MockTransaction) Deletetransfer(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deletetransfer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deletetransfer indicates an expected call of Deletetransfer.
func (mr *MockTransactionMockRecorder) Deletetransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deletetransfer", reflect.TypeOf((*MockTransaction)(nil).Deletetransfer), arg0, arg1)
}

// GetLasttransferByPlayerid mocks base method.
func (m *MockTransaction) GetLasttransferByPlayerid(arg0 context.Context, arg1 int32) (DB.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLasttransferByPlayerid", arg0, arg1)
	ret0, _ := ret[0].(DB.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLasttransferByPlayerid indicates an expected call of GetLasttransferByPlayerid.
func (mr *MockTransactionMockRecorder) GetLasttransferByPlayerid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLasttransferByPlayerid", reflect.TypeOf((*MockTransaction)(nil).GetLasttransferByPlayerid), arg0, arg1)
}

// GetPlayersList mocks base method.
func (m *MockTransaction) GetPlayersList(arg0 context.Context, arg1 DB.GetPlayersListParams) ([]DB.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlayersList", arg0, arg1)
	ret0, _ := ret[0].([]DB.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlayersList indicates an expected call of GetPlayersList.
func (mr *MockTransactionMockRecorder) GetPlayersList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlayersList", reflect.TypeOf((*MockTransaction)(nil).GetPlayersList), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockTransaction) GetUsers(arg0 context.Context, arg1 string) (DB.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1)
	ret0, _ := ret[0].(DB.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockTransactionMockRecorder) GetUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockTransaction)(nil).GetUsers), arg0, arg1)
}

// GetfootballclubByCountry mocks base method.
func (m *MockTransaction) GetfootballclubByCountry(arg0 context.Context, arg1 string) ([]DB.Footballclub, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetfootballclubByCountry", arg0, arg1)
	ret0, _ := ret[0].([]DB.Footballclub)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetfootballclubByCountry indicates an expected call of GetfootballclubByCountry.
func (mr *MockTransactionMockRecorder) GetfootballclubByCountry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetfootballclubByCountry", reflect.TypeOf((*MockTransaction)(nil).GetfootballclubByCountry), arg0, arg1)
}

// GetfootballclubByID mocks base method.
func (m *MockTransaction) GetfootballclubByID(arg0 context.Context, arg1 int32) (DB.Footballclub, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetfootballclubByID", arg0, arg1)
	ret0, _ := ret[0].(DB.Footballclub)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetfootballclubByID indicates an expected call of GetfootballclubByID.
func (mr *MockTransactionMockRecorder) GetfootballclubByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetfootballclubByID", reflect.TypeOf((*MockTransaction)(nil).GetfootballclubByID), arg0, arg1)
}

// GetfootballclubByName mocks base method.
func (m *MockTransaction) GetfootballclubByName(arg0 context.Context, arg1 string) (DB.Footballclub, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetfootballclubByName", arg0, arg1)
	ret0, _ := ret[0].(DB.Footballclub)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetfootballclubByName indicates an expected call of GetfootballclubByName.
func (mr *MockTransactionMockRecorder) GetfootballclubByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetfootballclubByName", reflect.TypeOf((*MockTransaction)(nil).GetfootballclubByName), arg0, arg1)
}

// GetplayerByCountry mocks base method.
func (m *MockTransaction) GetplayerByCountry(arg0 context.Context, arg1 string) ([]DB.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetplayerByCountry", arg0, arg1)
	ret0, _ := ret[0].([]DB.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetplayerByCountry indicates an expected call of GetplayerByCountry.
func (mr *MockTransactionMockRecorder) GetplayerByCountry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetplayerByCountry", reflect.TypeOf((*MockTransaction)(nil).GetplayerByCountry), arg0, arg1)
}

// GetplayerByFootballclub mocks base method.
func (m *MockTransaction) GetplayerByFootballclub(arg0 context.Context, arg1 int32) ([]DB.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetplayerByFootballclub", arg0, arg1)
	ret0, _ := ret[0].([]DB.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetplayerByFootballclub indicates an expected call of GetplayerByFootballclub.
func (mr *MockTransactionMockRecorder) GetplayerByFootballclub(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetplayerByFootballclub", reflect.TypeOf((*MockTransaction)(nil).GetplayerByFootballclub), arg0, arg1)
}

// GetplayerByID mocks base method.
func (m *MockTransaction) GetplayerByID(arg0 context.Context, arg1 int32) (DB.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetplayerByID", arg0, arg1)
	ret0, _ := ret[0].(DB.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetplayerByID indicates an expected call of GetplayerByID.
func (mr *MockTransactionMockRecorder) GetplayerByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetplayerByID", reflect.TypeOf((*MockTransaction)(nil).GetplayerByID), arg0, arg1)
}

// GetplayerByName mocks base method.
func (m *MockTransaction) GetplayerByName(arg0 context.Context, arg1 string) (DB.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetplayerByName", arg0, arg1)
	ret0, _ := ret[0].(DB.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetplayerByName indicates an expected call of GetplayerByName.
func (mr *MockTransactionMockRecorder) GetplayerByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetplayerByName", reflect.TypeOf((*MockTransaction)(nil).GetplayerByName), arg0, arg1)
}

// GetplayerByPosition mocks base method.
func (m *MockTransaction) GetplayerByPosition(arg0 context.Context, arg1 string) ([]DB.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetplayerByPosition", arg0, arg1)
	ret0, _ := ret[0].([]DB.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetplayerByPosition indicates an expected call of GetplayerByPosition.
func (mr *MockTransactionMockRecorder) GetplayerByPosition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetplayerByPosition", reflect.TypeOf((*MockTransaction)(nil).GetplayerByPosition), arg0, arg1)
}

// GetplayerByValueHigherthan mocks base method.
func (m *MockTransaction) GetplayerByValueHigherthan(arg0 context.Context, arg1 int64) ([]DB.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetplayerByValueHigherthan", arg0, arg1)
	ret0, _ := ret[0].([]DB.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetplayerByValueHigherthan indicates an expected call of GetplayerByValueHigherthan.
func (mr *MockTransactionMockRecorder) GetplayerByValueHigherthan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetplayerByValueHigherthan", reflect.TypeOf((*MockTransaction)(nil).GetplayerByValueHigherthan), arg0, arg1)
}

// GetplayerByValueLessthan mocks base method.
func (m *MockTransaction) GetplayerByValueLessthan(arg0 context.Context, arg1 int64) ([]DB.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetplayerByValueLessthan", arg0, arg1)
	ret0, _ := ret[0].([]DB.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetplayerByValueLessthan indicates an expected call of GetplayerByValueLessthan.
func (mr *MockTransactionMockRecorder) GetplayerByValueLessthan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetplayerByValueLessthan", reflect.TypeOf((*MockTransaction)(nil).GetplayerByValueLessthan), arg0, arg1)
}

// GettransferByPlayerid mocks base method.
func (m *MockTransaction) GettransferByPlayerid(arg0 context.Context, arg1 int32) ([]DB.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GettransferByPlayerid", arg0, arg1)
	ret0, _ := ret[0].([]DB.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GettransferByPlayerid indicates an expected call of GettransferByPlayerid.
func (mr *MockTransactionMockRecorder) GettransferByPlayerid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GettransferByPlayerid", reflect.TypeOf((*MockTransaction)(nil).GettransferByPlayerid), arg0, arg1)
}

// GettransferByTransferid mocks base method.
func (m *MockTransaction) GettransferByTransferid(arg0 context.Context, arg1 int32) (DB.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GettransferByTransferid", arg0, arg1)
	ret0, _ := ret[0].(DB.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GettransferByTransferid indicates an expected call of GettransferByTransferid.
func (mr *MockTransactionMockRecorder) GettransferByTransferid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GettransferByTransferid", reflect.TypeOf((*MockTransaction)(nil).GettransferByTransferid), arg0, arg1)
}

// GettransferList mocks base method.
func (m *MockTransaction) GettransferList(arg0 context.Context, arg1 DB.GettransferListParams) ([]DB.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GettransferList", arg0, arg1)
	ret0, _ := ret[0].([]DB.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GettransferList indicates an expected call of GettransferList.
func (mr *MockTransactionMockRecorder) GettransferList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GettransferList", reflect.TypeOf((*MockTransaction)(nil).GettransferList), arg0, arg1)
}

// Highesttransfer mocks base method.
func (m *MockTransaction) Highesttransfer(arg0 context.Context) (DB.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Highesttransfer", arg0)
	ret0, _ := ret[0].(DB.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Highesttransfer indicates an expected call of Highesttransfer.
func (mr *MockTransactionMockRecorder) Highesttransfer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Highesttransfer", reflect.TypeOf((*MockTransaction)(nil).Highesttransfer), arg0)
}

// Latesttransfer mocks base method.
func (m *MockTransaction) Latesttransfer(arg0 context.Context, arg1 DB.LatesttransferParams) (DB.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Latesttransfer", arg0, arg1)
	ret0, _ := ret[0].(DB.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Latesttransfer indicates an expected call of Latesttransfer.
func (mr *MockTransactionMockRecorder) Latesttransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Latesttransfer", reflect.TypeOf((*MockTransaction)(nil).Latesttransfer), arg0, arg1)
}

// Listfootballclub mocks base method.
func (m *MockTransaction) Listfootballclub(arg0 context.Context, arg1 DB.ListfootballclubParams) ([]DB.Footballclub, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listfootballclub", arg0, arg1)
	ret0, _ := ret[0].([]DB.Footballclub)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Listfootballclub indicates an expected call of Listfootballclub.
func (mr *MockTransactionMockRecorder) Listfootballclub(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listfootballclub", reflect.TypeOf((*MockTransaction)(nil).Listfootballclub), arg0, arg1)
}

// TransferTx mocks base method.
func (m *MockTransaction) TransferTx(arg0 context.Context, arg1 DB.TransferTxParams) (DB.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(DB.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockTransactionMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockTransaction)(nil).TransferTx), arg0, arg1)
}

// UpdatefootballclubBalance mocks base method.
func (m *MockTransaction) UpdatefootballclubBalance(arg0 context.Context, arg1 DB.UpdatefootballclubBalanceParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatefootballclubBalance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatefootballclubBalance indicates an expected call of UpdatefootballclubBalance.
func (mr *MockTransactionMockRecorder) UpdatefootballclubBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatefootballclubBalance", reflect.TypeOf((*MockTransaction)(nil).UpdatefootballclubBalance), arg0, arg1)
}

// Updateplayer mocks base method.
func (m *MockTransaction) Updateplayer(arg0 context.Context, arg1 DB.UpdateplayerParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updateplayer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Updateplayer indicates an expected call of Updateplayer.
func (mr *MockTransactionMockRecorder) Updateplayer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updateplayer", reflect.TypeOf((*MockTransaction)(nil).Updateplayer), arg0, arg1)
}

// Updatetransfer mocks base method.
func (m *MockTransaction) Updatetransfer(arg0 context.Context, arg1 DB.UpdatetransferParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updatetransfer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Updatetransfer indicates an expected call of Updatetransfer.
func (mr *MockTransactionMockRecorder) Updatetransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updatetransfer", reflect.TypeOf((*MockTransaction)(nil).Updatetransfer), arg0, arg1)
}
