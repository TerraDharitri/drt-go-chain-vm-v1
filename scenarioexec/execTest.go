package scenarioexec

import (
	"fmt"

	mj "github.com/TerraDharitri/drt-go-chain-vm-v1/scenarios/json/model"
)

// ExecuteTest executes an individual test.
func (ae *VMTestExecutor) ExecuteTest(test *mj.Test) error {
	// reset world
	ae.World.Clear()
	ae.World.Blockhashes = mj.JSONBytesFromStringValues(test.BlockHashes)

	for _, acct := range test.Pre {
		account, err := convertAccount(acct)
		if err != nil {
			return err
		}

		ae.World.AcctMap.PutAccount(account)
	}

	for _, block := range test.Blocks {
		for txIndex, tx := range block.Transactions {
			txName := fmt.Sprintf("%d", txIndex)

			// execute
			output, err := ae.executeTx(txName, tx)
			if err != nil {
				return err
			}

			blResult := block.Results[txIndex]

			// check results
			err = checkTxResults(txName, blResult, test.CheckGas, output)
			if err != nil {
				return err
			}
		}
	}

	return ae.checkAccounts(test.PostState)
}
