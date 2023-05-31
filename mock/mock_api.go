// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/evcc-io/evcc/api (interfaces: Charger,ChargeState,PhaseSwitcher,Identifier,Meter,MeterEnergy,Vehicle,ChargeRater,Battery,Tariff)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	api "github.com/evcc-io/evcc/api"
	gomock "github.com/golang/mock/gomock"
)

// MockCharger is a mock of Charger interface.
type MockCharger struct {
	ctrl     *gomock.Controller
	recorder *MockChargerMockRecorder
}

// MockChargerMockRecorder is the mock recorder for MockCharger.
type MockChargerMockRecorder struct {
	mock *MockCharger
}

// NewMockCharger creates a new mock instance.
func NewMockCharger(ctrl *gomock.Controller) *MockCharger {
	mock := &MockCharger{ctrl: ctrl}
	mock.recorder = &MockChargerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharger) EXPECT() *MockChargerMockRecorder {
	return m.recorder
}

// Enable mocks base method.
func (m *MockCharger) Enable(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockChargerMockRecorder) Enable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockCharger)(nil).Enable), arg0)
}

// Enabled mocks base method.
func (m *MockCharger) Enabled() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enabled indicates an expected call of Enabled.
func (mr *MockChargerMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockCharger)(nil).Enabled))
}

// MaxCurrent mocks base method.
func (m *MockCharger) MaxCurrent(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxCurrent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MaxCurrent indicates an expected call of MaxCurrent.
func (mr *MockChargerMockRecorder) MaxCurrent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxCurrent", reflect.TypeOf((*MockCharger)(nil).MaxCurrent), arg0)
}

// Status mocks base method.
func (m *MockCharger) Status() (api.ChargeStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(api.ChargeStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockChargerMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockCharger)(nil).Status))
}

// MockChargeState is a mock of ChargeState interface.
type MockChargeState struct {
	ctrl     *gomock.Controller
	recorder *MockChargeStateMockRecorder
}

// MockChargeStateMockRecorder is the mock recorder for MockChargeState.
type MockChargeStateMockRecorder struct {
	mock *MockChargeState
}

// NewMockChargeState creates a new mock instance.
func NewMockChargeState(ctrl *gomock.Controller) *MockChargeState {
	mock := &MockChargeState{ctrl: ctrl}
	mock.recorder = &MockChargeStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChargeState) EXPECT() *MockChargeStateMockRecorder {
	return m.recorder
}

// Status mocks base method.
func (m *MockChargeState) Status() (api.ChargeStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(api.ChargeStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockChargeStateMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockChargeState)(nil).Status))
}

// MockPhaseSwitcher is a mock of PhaseSwitcher interface.
type MockPhaseSwitcher struct {
	ctrl     *gomock.Controller
	recorder *MockPhaseSwitcherMockRecorder
}

// MockPhaseSwitcherMockRecorder is the mock recorder for MockPhaseSwitcher.
type MockPhaseSwitcherMockRecorder struct {
	mock *MockPhaseSwitcher
}

// NewMockPhaseSwitcher creates a new mock instance.
func NewMockPhaseSwitcher(ctrl *gomock.Controller) *MockPhaseSwitcher {
	mock := &MockPhaseSwitcher{ctrl: ctrl}
	mock.recorder = &MockPhaseSwitcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhaseSwitcher) EXPECT() *MockPhaseSwitcherMockRecorder {
	return m.recorder
}

// Phases1p3p mocks base method.
func (m *MockPhaseSwitcher) Phases1p3p(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Phases1p3p", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Phases1p3p indicates an expected call of Phases1p3p.
func (mr *MockPhaseSwitcherMockRecorder) Phases1p3p(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Phases1p3p", reflect.TypeOf((*MockPhaseSwitcher)(nil).Phases1p3p), arg0)
}

// MockIdentifier is a mock of Identifier interface.
type MockIdentifier struct {
	ctrl     *gomock.Controller
	recorder *MockIdentifierMockRecorder
}

// MockIdentifierMockRecorder is the mock recorder for MockIdentifier.
type MockIdentifierMockRecorder struct {
	mock *MockIdentifier
}

// NewMockIdentifier creates a new mock instance.
func NewMockIdentifier(ctrl *gomock.Controller) *MockIdentifier {
	mock := &MockIdentifier{ctrl: ctrl}
	mock.recorder = &MockIdentifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentifier) EXPECT() *MockIdentifierMockRecorder {
	return m.recorder
}

// Identify mocks base method.
func (m *MockIdentifier) Identify() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identify")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Identify indicates an expected call of Identify.
func (mr *MockIdentifierMockRecorder) Identify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identify", reflect.TypeOf((*MockIdentifier)(nil).Identify))
}

// MockMeter is a mock of Meter interface.
type MockMeter struct {
	ctrl     *gomock.Controller
	recorder *MockMeterMockRecorder
}

// MockMeterMockRecorder is the mock recorder for MockMeter.
type MockMeterMockRecorder struct {
	mock *MockMeter
}

// NewMockMeter creates a new mock instance.
func NewMockMeter(ctrl *gomock.Controller) *MockMeter {
	mock := &MockMeter{ctrl: ctrl}
	mock.recorder = &MockMeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeter) EXPECT() *MockMeterMockRecorder {
	return m.recorder
}

// CurrentPower mocks base method.
func (m *MockMeter) CurrentPower() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentPower")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentPower indicates an expected call of CurrentPower.
func (mr *MockMeterMockRecorder) CurrentPower() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentPower", reflect.TypeOf((*MockMeter)(nil).CurrentPower))
}

// MockMeterEnergy is a mock of MeterEnergy interface.
type MockMeterEnergy struct {
	ctrl     *gomock.Controller
	recorder *MockMeterEnergyMockRecorder
}

// MockMeterEnergyMockRecorder is the mock recorder for MockMeterEnergy.
type MockMeterEnergyMockRecorder struct {
	mock *MockMeterEnergy
}

// NewMockMeterEnergy creates a new mock instance.
func NewMockMeterEnergy(ctrl *gomock.Controller) *MockMeterEnergy {
	mock := &MockMeterEnergy{ctrl: ctrl}
	mock.recorder = &MockMeterEnergyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeterEnergy) EXPECT() *MockMeterEnergyMockRecorder {
	return m.recorder
}

// TotalEnergy mocks base method.
func (m *MockMeterEnergy) TotalEnergy() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalEnergy")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalEnergy indicates an expected call of TotalEnergy.
func (mr *MockMeterEnergyMockRecorder) TotalEnergy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalEnergy", reflect.TypeOf((*MockMeterEnergy)(nil).TotalEnergy))
}

// MockVehicle is a mock of Vehicle interface.
type MockVehicle struct {
	ctrl     *gomock.Controller
	recorder *MockVehicleMockRecorder
}

// MockVehicleMockRecorder is the mock recorder for MockVehicle.
type MockVehicleMockRecorder struct {
	mock *MockVehicle
}

// NewMockVehicle creates a new mock instance.
func NewMockVehicle(ctrl *gomock.Controller) *MockVehicle {
	mock := &MockVehicle{ctrl: ctrl}
	mock.recorder = &MockVehicleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVehicle) EXPECT() *MockVehicleMockRecorder {
	return m.recorder
}

// Capacity mocks base method.
func (m *MockVehicle) Capacity() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capacity")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Capacity indicates an expected call of Capacity.
func (mr *MockVehicleMockRecorder) Capacity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capacity", reflect.TypeOf((*MockVehicle)(nil).Capacity))
}

// Icon mocks base method.
func (m *MockVehicle) Icon() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Icon")
	ret0, _ := ret[0].(string)
	return ret0
}

// Icon indicates an expected call of Icon.
func (mr *MockVehicleMockRecorder) Icon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Icon", reflect.TypeOf((*MockVehicle)(nil).Icon))
}

// Identifiers mocks base method.
func (m *MockVehicle) Identifiers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifiers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Identifiers indicates an expected call of Identifiers.
func (mr *MockVehicleMockRecorder) Identifiers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifiers", reflect.TypeOf((*MockVehicle)(nil).Identifiers))
}

// OnIdentified mocks base method.
func (m *MockVehicle) OnIdentified() api.ActionConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnIdentified")
	ret0, _ := ret[0].(api.ActionConfig)
	return ret0
}

// OnIdentified indicates an expected call of OnIdentified.
func (mr *MockVehicleMockRecorder) OnIdentified() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnIdentified", reflect.TypeOf((*MockVehicle)(nil).OnIdentified))
}

// Phases mocks base method.
func (m *MockVehicle) Phases() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Phases")
	ret0, _ := ret[0].(int)
	return ret0
}

// Phases indicates an expected call of Phases.
func (mr *MockVehicleMockRecorder) Phases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Phases", reflect.TypeOf((*MockVehicle)(nil).Phases))
}

// SetTitle mocks base method.
func (m *MockVehicle) SetTitle(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTitle", arg0)
}

// SetTitle indicates an expected call of SetTitle.
func (mr *MockVehicleMockRecorder) SetTitle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTitle", reflect.TypeOf((*MockVehicle)(nil).SetTitle), arg0)
}

// Soc mocks base method.
func (m *MockVehicle) Soc() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Soc")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Soc indicates an expected call of Soc.
func (mr *MockVehicleMockRecorder) Soc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Soc", reflect.TypeOf((*MockVehicle)(nil).Soc))
}

// Title mocks base method.
func (m *MockVehicle) Title() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Title")
	ret0, _ := ret[0].(string)
	return ret0
}

// Title indicates an expected call of Title.
func (mr *MockVehicleMockRecorder) Title() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Title", reflect.TypeOf((*MockVehicle)(nil).Title))
}

// MockChargeRater is a mock of ChargeRater interface.
type MockChargeRater struct {
	ctrl     *gomock.Controller
	recorder *MockChargeRaterMockRecorder
}

// MockChargeRaterMockRecorder is the mock recorder for MockChargeRater.
type MockChargeRaterMockRecorder struct {
	mock *MockChargeRater
}

// NewMockChargeRater creates a new mock instance.
func NewMockChargeRater(ctrl *gomock.Controller) *MockChargeRater {
	mock := &MockChargeRater{ctrl: ctrl}
	mock.recorder = &MockChargeRaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChargeRater) EXPECT() *MockChargeRaterMockRecorder {
	return m.recorder
}

// ChargedEnergy mocks base method.
func (m *MockChargeRater) ChargedEnergy() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChargedEnergy")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChargedEnergy indicates an expected call of ChargedEnergy.
func (mr *MockChargeRaterMockRecorder) ChargedEnergy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChargedEnergy", reflect.TypeOf((*MockChargeRater)(nil).ChargedEnergy))
}

// MockBattery is a mock of Battery interface.
type MockBattery struct {
	ctrl     *gomock.Controller
	recorder *MockBatteryMockRecorder
}

// MockBatteryMockRecorder is the mock recorder for MockBattery.
type MockBatteryMockRecorder struct {
	mock *MockBattery
}

// NewMockBattery creates a new mock instance.
func NewMockBattery(ctrl *gomock.Controller) *MockBattery {
	mock := &MockBattery{ctrl: ctrl}
	mock.recorder = &MockBatteryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBattery) EXPECT() *MockBatteryMockRecorder {
	return m.recorder
}

// Soc mocks base method.
func (m *MockBattery) Soc() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Soc")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Soc indicates an expected call of Soc.
func (mr *MockBatteryMockRecorder) Soc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Soc", reflect.TypeOf((*MockBattery)(nil).Soc))
}

// MockTariff is a mock of Tariff interface.
type MockTariff struct {
	ctrl     *gomock.Controller
	recorder *MockTariffMockRecorder
}

// MockTariffMockRecorder is the mock recorder for MockTariff.
type MockTariffMockRecorder struct {
	mock *MockTariff
}

// NewMockTariff creates a new mock instance.
func NewMockTariff(ctrl *gomock.Controller) *MockTariff {
	mock := &MockTariff{ctrl: ctrl}
	mock.recorder = &MockTariffMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTariff) EXPECT() *MockTariffMockRecorder {
	return m.recorder
}

// Rates mocks base method.
func (m *MockTariff) Rates() (api.Rates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rates")
	ret0, _ := ret[0].(api.Rates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rates indicates an expected call of Rates.
func (mr *MockTariffMockRecorder) Rates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rates", reflect.TypeOf((*MockTariff)(nil).Rates))
}

// Type mocks base method.
func (m *MockTariff) Type() api.TariffType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(api.TariffType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockTariffMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockTariff)(nil).Type))
}
