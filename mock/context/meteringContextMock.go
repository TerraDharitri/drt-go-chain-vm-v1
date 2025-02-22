package mock

import (
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain-vm-v1/config"
	"github.com/TerraDharitri/drt-go-chain-vm-v1/vmhost"
)

var _ vmhost.MeteringContext = (*MeteringContextMock)(nil)

// MeteringContextMock is used in tests to check the MeteringContext interface method calls
type MeteringContextMock struct {
	GasCost           *config.GasCost
	GasLeftMock       uint64
	GasLockedMock     uint64
	GasComputedToLock uint64
	BlockGasLimitMock uint64
	Err               error
}

// InitState mocked method
func (m *MeteringContextMock) InitState() {
}

// PushState mocked method
func (m *MeteringContextMock) PushState() {
}

// PopSetActiveState mocked method
func (m *MeteringContextMock) PopSetActiveState() {
}

// PopDiscard mocked method
func (m *MeteringContextMock) PopDiscard() {
}

// ClearStateStack mocked method
func (m *MeteringContextMock) ClearStateStack() {
}

// SetGasSchedule mocked method
func (m *MeteringContextMock) SetGasSchedule(gasSchedule config.GasScheduleMap) {
	gasCostConfig, _ := config.CreateGasConfig(gasSchedule)
	m.GasCost = gasCostConfig
}

// GasSchedule mocked method
func (m *MeteringContextMock) GasSchedule() *config.GasCost {
	return m.GasCost
}

// UseGas mocked method
func (m *MeteringContextMock) UseGas(_ uint64) {
}

// FreeGas mocked method
func (m *MeteringContextMock) FreeGas(_ uint64) {
}

// RestoreGas mocked method
func (m *MeteringContextMock) RestoreGas(_ uint64) {
}

// GasLeft mocked method
func (m *MeteringContextMock) GasLeft() uint64 {
	return m.GasLeftMock
}

// AddToUsedGas -
func (m *MeteringContextMock) AddToUsedGas(_ []byte, _ uint64) {
}

// ForwardGas mocked method
func (m *MeteringContextMock) ForwardGas(_ []byte, _ []byte, _ uint64) {
}

// GetForwardedGas -
func (m *MeteringContextMock) GetForwardedGas(_ []byte) uint64 {
	return 0
}

// InitStateFromContractCallInput mocked method
func (m *MeteringContextMock) InitStateFromContractCallInput(_ *vmcommon.VMInput) {
}

// GasUsedByContract mocked method
func (m *MeteringContextMock) GasUsedByContract() (uint64, uint64) {
	return 0, 0
}

// GasUsedForExecution mocked method
func (m *MeteringContextMock) GasUsedForExecution() uint64 {
	return 0
}

// GasSpentByContract mocked method
func (m *MeteringContextMock) GasSpentByContract() uint64 {
	return 0
}

// GetGasForExecution mocked method
func (m *MeteringContextMock) GetGasForExecution() uint64 {
	return 0
}

// GetGasProvided mocked method
func (m *MeteringContextMock) GetGasProvided() uint64 {
	return 0
}

// GetSCPrepareInitialCost mocked method
func (m *MeteringContextMock) GetSCPrepareInitialCost() uint64 {
	return 0
}

// BoundGasLimit mocked method
func (m *MeteringContextMock) BoundGasLimit(value int64) uint64 {
	gasLeft := m.GasLeft()
	limit := uint64(value)

	if gasLeft < limit {
		return gasLeft
	}
	return limit
}

// ComputeGasLockedForAsync mocked method
func (m *MeteringContextMock) ComputeGasLockedForAsync() uint64 {
	return m.GasComputedToLock
}

// DeductGasIfAsyncStep mocked method
func (m *MeteringContextMock) DeductGasIfAsyncStep() error {
	return m.Err
}

// UseGasBounded mocked method
func (m *MeteringContextMock) UseGasBounded(_ uint64) error {
	return m.Err
}

// UnlockGasIfAsyncCallback mocked method
func (m *MeteringContextMock) UnlockGasIfAsyncCallback() {
}

// UseGasForAsyncStep mocked method
func (m *MeteringContextMock) UseGasForAsyncStep() error {
	return m.Err
}

// UnlockGasIfAsyncStep mocked method
func (m *MeteringContextMock) UnlockGasIfAsyncStep() {
}

// GetGasLocked mocked method
func (m *MeteringContextMock) GetGasLocked() uint64 {
	return m.GasLockedMock
}

// BlockGasLimit mocked method
func (m *MeteringContextMock) BlockGasLimit() uint64 {
	return m.BlockGasLimitMock
}

// DeductInitialGasForExecution mocked method
func (m *MeteringContextMock) DeductInitialGasForExecution(_ []byte) error {
	return m.Err
}

// DeductInitialGasForDirectDeployment mocked method
func (m *MeteringContextMock) DeductInitialGasForDirectDeployment(_ vmhost.CodeDeployInput) error {
	return m.Err
}

// DeductInitialGasForIndirectDeployment mocked method
func (m *MeteringContextMock) DeductInitialGasForIndirectDeployment(_ vmhost.CodeDeployInput) error {
	return m.Err
}

// SetTotalUsedGas mocked method
func (m *MeteringContextMock) SetTotalUsedGas(_ uint64) {
}

// GetPreviousTotalUsedGas mocked method
func (m *MeteringContextMock) GetPreviousTotalUsedGas() uint64 {
	return 0
}
