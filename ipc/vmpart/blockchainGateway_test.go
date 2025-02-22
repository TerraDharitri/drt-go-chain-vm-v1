package vmpart

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	vmcommon "github.com/TerraDharitri/drt-go-chain-vm-common"
	"github.com/TerraDharitri/drt-go-chain-vm-v1/ipc/common"
	"github.com/TerraDharitri/drt-go-chain-vm-v1/ipc/marshaling"
	"github.com/TerraDharitri/drt-go-chain-vm-v1/ipc/nodepart"
	"github.com/stretchr/testify/require"
)

func TestGateway_ProcessBuiltInFunction(t *testing.T) {
	callHook := func(gateway *BlockchainHookGateway) {
		vmOutput, err := gateway.ProcessBuiltInFunction(&vmcommon.ContractCallInput{
			Function:      "fooFunction",
			RecipientAddr: []byte("alice"),
		})
		require.NoError(t, err)
		require.NotNil(t, vmOutput)
		require.Equal(t, vmcommon.Ok, vmOutput.ReturnCode)
		require.Equal(t, [][]byte{{42}}, vmOutput.ReturnData)
	}

	handleHookCall := func(request common.MessageHandler) common.MessageHandler {
		require.Equal(t, "fooFunction", string(request.(*common.MessageBlockchainProcessBuiltinFunctionRequest).CallInput.Function))
		require.Equal(t, "alice", string(request.(*common.MessageBlockchainProcessBuiltinFunctionRequest).CallInput.RecipientAddr))
		vmOutput := &vmcommon.VMOutput{
			ReturnCode: vmcommon.Ok,
			ReturnData: [][]byte{{42}},
		}
		return common.NewMessageBlockchainProcessBuiltinFunctionResponse(vmOutput, nil)
	}

	runHookScenario(t, callHook, handleHookCall)
}

func TestGateway_GetAllState(t *testing.T) {
	callHook := func(gateway *BlockchainHookGateway) {
		state, err := gateway.GetAllState([]byte("alice"))
		require.NoError(t, err)
		require.NotNil(t, state)
		require.Equal(t, "bar", string(state["foo"]))
	}

	handleHookCall := func(request common.MessageHandler) common.MessageHandler {
		require.Equal(t, "alice", string(request.(*common.MessageBlockchainGetAllStateRequest).Address))
		state := make(map[string][]byte)
		state["foo"] = []byte("bar")
		return common.NewMessageBlockchainGetAllStateResponse(state, nil)
	}

	runHookScenario(t, callHook, handleHookCall)
}

func TestGateway_GetUserAccount(t *testing.T) {
	callHook := func(gateway *BlockchainHookGateway) {
		account, err := gateway.GetUserAccount([]byte("alice"))
		require.NoError(t, err)
		require.NotNil(t, account)
		require.Equal(t, 42, int(account.GetNonce()))
		require.Equal(t, big.NewInt(43), account.GetDeveloperReward())
	}

	handleHookCall := func(request common.MessageHandler) common.MessageHandler {
		require.Equal(t, "alice", string(request.(*common.MessageBlockchainGetUserAccountRequest).Address))
		return common.NewMessageBlockchainGetUserAccountResponse(&common.Account{Nonce: 42, DeveloperReward: big.NewInt(43)}, nil)
	}

	runHookScenario(t, callHook, handleHookCall)
}

func TestGateway_GetUserAccount_WithError(t *testing.T) {
	errFoo := fmt.Errorf("foo error")

	callHook := func(gateway *BlockchainHookGateway) {
		account, err := gateway.GetUserAccount([]byte("alice"))
		require.Error(t, err, errFoo)
		require.Nil(t, account)
	}

	handleHookCall := func(request common.MessageHandler) common.MessageHandler {
		require.Equal(t, "alice", string(request.(*common.MessageBlockchainGetUserAccountRequest).Address))
		return common.NewMessageBlockchainGetUserAccountResponse(nil, errFoo)
	}

	runHookScenario(t, callHook, handleHookCall)
}

func TestGateway_GetShardOfAddress(t *testing.T) {
	callHook := func(gateway *BlockchainHookGateway) {
		shard := gateway.GetShardOfAddress([]byte("alice"))
		require.Equal(t, 3, int(shard))
	}

	handleHookCall := func(request common.MessageHandler) common.MessageHandler {
		require.Equal(t, "alice", string(request.(*common.MessageBlockchainGetShardOfAddressRequest).Address))
		return common.NewMessageBlockchainGetShardOfAddressResponse(3)
	}

	runHookScenario(t, callHook, handleHookCall)
}

func TestGateway_IsSmartContract(t *testing.T) {
	callHook := func(gateway *BlockchainHookGateway) {
		result := gateway.IsSmartContract([]byte("contract"))
		require.True(t, result)
	}

	handleHookCall := func(request common.MessageHandler) common.MessageHandler {
		require.Equal(t, "contract", string(request.(*common.MessageBlockchainIsSmartContractRequest).Address))
		return common.NewMessageBlockchainIsSmartContractResponse(true)
	}

	runHookScenario(t, callHook, handleHookCall)
}

func TestGateway_IsPayable(t *testing.T) {
	callHook := func(gateway *BlockchainHookGateway) {
		result, err := gateway.IsPayable(nil, []byte("contract"))
		require.True(t, result)
		require.Nil(t, err)
	}

	handleHookCall := func(request common.MessageHandler) common.MessageHandler {
		require.Equal(t, "contract", string(request.(*common.MessageBlockchainIsPayableRequest).Address))
		return common.NewMessageBlockchainIsPayableResponse(true, nil)
	}

	runHookScenario(t, callHook, handleHookCall)
}

func TestBlockchainHookGateway_GetCompiledCode(t *testing.T) {
	callHook := func(gateway *BlockchainHookGateway) {
		result, code := gateway.GetCompiledCode([]byte("contract"))
		require.True(t, result)
		require.Equal(t, code, []byte("contract"))
	}

	handleHookCall := func(request common.MessageHandler) common.MessageHandler {
		require.Equal(t, "contract", string(request.(*common.MessageBlockchainGetCompiledCodeRequest).CodeHash))
		return common.NewMessageBlockchainGetCompiledCodeResponse(true, []byte("contract"))
	}

	runHookScenario(t, callHook, handleHookCall)
}

func TestBlockchainHookGateway_SaveCompiledCode(t *testing.T) {
	callHook := func(gateway *BlockchainHookGateway) {
		gateway.SaveCompiledCode([]byte("contract"), []byte("contract"))
	}

	handleHookCall := func(request common.MessageHandler) common.MessageHandler {
		require.Equal(t, "contract", string(request.(*common.MessageBlockchainSaveCompiledCodeRequest).CodeHash))
		return common.NewMessageBlockchainSaveCompiledCodeResponse()
	}

	runHookScenario(t, callHook, handleHookCall)
}

func runHookScenario(t *testing.T, callHook func(*BlockchainHookGateway), handleHookCall func(common.MessageHandler) common.MessageHandler) {
	testFiles := createTestFiles(t)
	marshalizer := marshaling.CreateMarshalizer(marshaling.JSON)
	nodeMessenger := nodepart.NewNodeMessenger(testFiles.inputOfNode, testFiles.outputOfNode, marshalizer)
	vmMessenger := NewVMMessenger(testFiles.inputOfVM, testFiles.outputOfVM, marshalizer)
	gateway := NewBlockchainHookGateway(vmMessenger)

	go func() {
		request, err := nodeMessenger.Receive(0)
		require.NoError(t, err)
		response := handleHookCall(request)
		err = nodeMessenger.SendHookCallResponse(response)
		require.NoError(t, err)
	}()

	callHook(gateway)
}

type testFiles struct {
	outputOfNode *os.File
	inputOfVM    *os.File
	outputOfVM   *os.File
	inputOfNode  *os.File
}

func createTestFiles(t *testing.T) testFiles {
	files := testFiles{}

	var err error
	files.inputOfVM, files.outputOfNode, err = os.Pipe()
	require.NoError(t, err)
	files.inputOfNode, files.outputOfVM, err = os.Pipe()
	require.NoError(t, err)

	return files
}
