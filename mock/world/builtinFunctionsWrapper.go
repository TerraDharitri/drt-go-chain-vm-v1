package worldmock

import (
	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain-vm-common/builtInFunctions"
	"github.com/TerraDharitri/drt-go-chain-vm-common/mock"
	"github.com/TerraDharitri/drt-go-chain-vm-v1/config"
)

// WorldMarshalizer is the global marshalizer to be used by the components of
// the BuiltinFunctionsWrapper.
var WorldMarshalizer = &GogoProtoMarshalizer{}

// BuiltinFunctionsWrapper manages and initializes a BuiltInFunctionContainer
// along with its dependencies
type BuiltinFunctionsWrapper struct {
	Container       vmcommon.BuiltInFunctionContainer
	MapDNSAddresses map[string]struct{}
	World           *MockWorld
	Marshalizer     vmcommon.Marshalizer
}

// NewBuiltinFunctionsWrapper creates a new BuiltinFunctionsWrapper with
// default dependencies.
func NewBuiltinFunctionsWrapper(
	world *MockWorld,
	gasMap config.GasScheduleMap,
) (*BuiltinFunctionsWrapper, error) {

	dnsMap := makeDNSAddresses(numDNSAddresses)

	argsBuiltIn := builtInFunctions.ArgsCreateBuiltInFunctionContainer{
		GasMap:                           gasMap,
		MapDNSAddresses:                  dnsMap,
		MapDNSV2Addresses:                makeDNSAddresses(numDNSAddresses),
		Marshalizer:                      WorldMarshalizer,
		Accounts:                         world.AccountsAdapter,
		ShardCoordinator:                 world,
		EnableEpochsHandler:              &mock.EnableEpochsHandlerStub{},
		GuardedAccountHandler:            world.GuardedAccountHandler,
		MaxNumOfAddressesForTransferRole: 100,
	}

	builtinFuncFactory, err := builtInFunctions.NewBuiltInFunctionsCreator(argsBuiltIn)
	if err != nil {
		return nil, err
	}

	err = builtinFuncFactory.CreateBuiltInFunctionContainer()
	if err != nil {
		return nil, err
	}

	err = builtinFuncFactory.SetPayableHandler(world)
	if err != nil {
		return nil, err
	}

	builtinFuncsWrapper := &BuiltinFunctionsWrapper{
		Container:       builtinFuncFactory.BuiltInFunctionContainer(),
		MapDNSAddresses: argsBuiltIn.MapDNSAddresses,
		World:           world,
	}

	return builtinFuncsWrapper, nil
}

// ProcessBuiltInFunction delegates the execution of a real builtin function to
// the inner BuiltInFunctionContainer.
func (bf *BuiltinFunctionsWrapper) ProcessBuiltInFunction(input *vmcommon.ContractCallInput) (*vmcommon.VMOutput, error) {
	caller := bf.getAccountSharded(input.CallerAddr)
	recipient := bf.getAccountSharded(input.RecipientAddr)

	function, err := bf.Container.Get(input.Function)
	if err != nil {
		return nil, err
	}

	return function.ProcessBuiltinFunction(caller, recipient, input)
}

// GetBuiltinFunctionNames returns the list of defined builtin-in functions.
func (bf *BuiltinFunctionsWrapper) GetBuiltinFunctionNames() vmcommon.FunctionNames {
	return bf.Container.Keys()
}

// TODO change AccountMap to support this instead
func (bf *BuiltinFunctionsWrapper) getAccountSharded(address []byte) vmcommon.UserAccountHandler {
	accountShard := bf.World.ComputeId(address)
	if accountShard != bf.World.SelfId() {
		return nil
	}
	return bf.World.AcctMap.GetAccount(address)
}
