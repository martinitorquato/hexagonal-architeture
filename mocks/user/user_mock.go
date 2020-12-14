package mock_user

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/martinitorquato/rest-api/internal/constants/model"
	routers "github.com/martinitorquato/rest-api/platform/routers"
	v10 "github.com/savsgio/atreugo/v10"
)

// MockPersistence is a mock of Persistence interface
type MockPersistence struct {
	ctrl     *gomock.Controller
	recorder *MockPersistenceMockRecorder
}

// MockPersistenceMockRecorder is the mock recorder for MockPersistence
type MockPersistenceMockRecorder struct {
	mock *MockPersistence
}

// NewMockPersistence creates a new mock instance
func NewMockPersistence(ctrl *gomock.Controller) *MockPersistence {
	mock := &MockPersistence{ctrl: ctrl}
	mock.recorder = &MockPersistenceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPersistence) EXPECT() *MockPersistenceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockPersistence) Create(ctx context.Context, user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockPersistenceMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPersistence)(nil).Create), ctx, user)
}

// FindByID mocks base method
func (m *MockPersistence) FindByID(ctx context.Context, userID int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockPersistenceMockRecorder) FindByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockPersistence)(nil).FindByID), ctx, userID)
}

// Find mocks base method
func (m *MockPersistence) Find(ctx context.Context, email, password string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, email, password)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockPersistenceMockRecorder) Find(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPersistence)(nil).Find), ctx, email, password)
}

// ChangePassword mocks base method
func (m *MockPersistence) ChangePassword(ctx context.Context, newPassword string, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, newPassword, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword
func (mr *MockPersistenceMockRecorder) ChangePassword(ctx, newPassword, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockPersistence)(nil).ChangePassword), ctx, newPassword, user)
}

// Delete mocks base method
func (m *MockPersistence) Delete(ctx context.Context, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPersistenceMockRecorder) Delete(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPersistence)(nil).Delete), ctx, user)
}

// MockCaching is a mock of Caching interface
type MockCaching struct {
	ctrl     *gomock.Controller
	recorder *MockCachingMockRecorder
}

// MockCachingMockRecorder is the mock recorder for MockCaching
type MockCachingMockRecorder struct {
	mock *MockCaching
}

// NewMockCaching creates a new mock instance
func NewMockCaching(ctrl *gomock.Controller) *MockCaching {
	mock := &MockCaching{ctrl: ctrl}
	mock.recorder = &MockCachingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCaching) EXPECT() *MockCachingMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockCaching) Save(ctx context.Context, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockCachingMockRecorder) Save(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCaching)(nil).Save), ctx, user)
}

// Get mocks base method
func (m *MockCaching) Get(ctx context.Context, userID int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCachingMockRecorder) Get(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCaching)(nil).Get), ctx, userID)
}

// Delete mocks base method
func (m *MockCaching) Delete(ctx context.Context, userID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCachingMockRecorder) Delete(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCaching)(nil).Delete), ctx, userID)
}

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// DataProfile mocks base method
func (m *MockRepository) DataProfile(ctx context.Context, userID int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataProfile", ctx, userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataProfile indicates an expected call of DataProfile
func (mr *MockRepositoryMockRecorder) DataProfile(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataProfile", reflect.TypeOf((*MockRepository)(nil).DataProfile), ctx, userID)
}

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockUsecase) Login(ctx context.Context, email, password string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, email, password)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockUsecaseMockRecorder) Login(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUsecase)(nil).Login), ctx, email, password)
}

// Profile mocks base method
func (m *MockUsecase) Profile(ctx context.Context, userID int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Profile", ctx, userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Profile indicates an expected call of Profile
func (mr *MockUsecaseMockRecorder) Profile(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Profile", reflect.TypeOf((*MockUsecase)(nil).Profile), ctx, userID)
}

// Register mocks base method
func (m *MockUsecase) Register(ctx context.Context, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockUsecaseMockRecorder) Register(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUsecase)(nil).Register), ctx, user)
}

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Test mocks base method
func (m *MockHandler) Test(ctx *v10.RequestCtx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Test", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Test indicates an expected call of Test
func (mr *MockHandlerMockRecorder) Test(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Test", reflect.TypeOf((*MockHandler)(nil).Test), ctx)
}

// CreateNewAccount mocks base method
func (m *MockHandler) CreateNewAccount(ctx *v10.RequestCtx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewAccount", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewAccount indicates an expected call of CreateNewAccount
func (mr *MockHandlerMockRecorder) CreateNewAccount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewAccount", reflect.TypeOf((*MockHandler)(nil).CreateNewAccount), ctx)
}

// SignIn mocks base method
func (m *MockHandler) SignIn(ctx *v10.RequestCtx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignIn indicates an expected call of SignIn
func (mr *MockHandlerMockRecorder) SignIn(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockHandler)(nil).SignIn), ctx)
}

// ShowProfile mocks base method
func (m *MockHandler) ShowProfile(ctx *v10.RequestCtx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowProfile", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShowProfile indicates an expected call of ShowProfile
func (mr *MockHandlerMockRecorder) ShowProfile(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowProfile", reflect.TypeOf((*MockHandler)(nil).ShowProfile), ctx)
}

// MockRoute is a mock of Route interface
type MockRoute struct {
	ctrl     *gomock.Controller
	recorder *MockRouteMockRecorder
}

// MockRouteMockRecorder is the mock recorder for MockRoute
type MockRouteMockRecorder struct {
	mock *MockRoute
}

// NewMockRoute creates a new mock instance
func NewMockRoute(ctrl *gomock.Controller) *MockRoute {
	mock := &MockRoute{ctrl: ctrl}
	mock.recorder = &MockRouteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoute) EXPECT() *MockRouteMockRecorder {
	return m.recorder
}

// Routers mocks base method
func (m *MockRoute) Routers() []*routers.Router {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Routers")
	ret0, _ := ret[0].([]*routers.Router)
	return ret0
}

// Routers indicates an expected call of Routers
func (mr *MockRouteMockRecorder) Routers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Routers", reflect.TypeOf((*MockRoute)(nil).Routers))
}
