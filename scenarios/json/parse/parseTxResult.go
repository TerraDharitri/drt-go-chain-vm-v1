package scenjsonparse

import (
	"errors"
	"fmt"

	mj "github.com/TerraDharitri/drt-go-chain-vm-v1/scenarios/json/model"
	oj "github.com/TerraDharitri/drt-go-chain-vm-v1/scenarios/orderedjson"
)

func (p *Parser) processTxExpectedResult(blrRaw oj.OJsonObject) (*mj.TransactionResult, error) {
	blrMap, isMap := blrRaw.(*oj.OJsonMap)
	if !isMap {
		return nil, errors.New("unmarshalled block result is not a map")
	}

	blr := mj.TransactionResult{
		Status:          mj.JSONCheckBigIntUnspecified(),
		Message:         mj.JSONCheckBytesUnspecified(),
		Gas:             mj.JSONCheckUint64Unspecified(),
		Refund:          mj.JSONCheckBigIntUnspecified(),
		LogsStar:        true,
		LogsUnspecified: true,
	}
	var err error
	for _, kvp := range blrMap.OrderedKV {
		switch kvp.Key {
		case "out":
			blr.Out, err = p.parseCheckBytesList(kvp.Value)
			if err != nil {
				return nil, fmt.Errorf("invalid block result out: %w", err)
			}
		case "status":
			blr.Status, err = p.processCheckBigInt(kvp.Value, bigIntSignedBytes)
			if err != nil {
				return nil, fmt.Errorf("invalid block result status: %w", err)
			}
		case "message":
			blr.Message, err = p.parseCheckBytes(kvp.Value)
			if err != nil {
				return nil, fmt.Errorf("invalid block result message: %w", err)
			}
		case "logs":
			blr.LogsUnspecified = false
			if IsStar(kvp.Value) {
				blr.LogsStar = true
			} else {
				blr.LogsStar = false
				blr.LogHash, err = p.parseString(kvp.Value)
				if err != nil {
					var logListErr error
					blr.Logs, logListErr = p.processLogList(kvp.Value)
					if logListErr != nil {
						return nil, logListErr
					}
				}
			}
		case "gas":
			blr.Gas, err = p.processCheckUint64(kvp.Value)
			if err != nil {
				return nil, fmt.Errorf("invalid block result gas: %w", err)
			}
		case "refund":
			blr.Refund, err = p.processCheckBigInt(kvp.Value, bigIntUnsignedBytes)
			if err != nil {
				return nil, fmt.Errorf("invalid block result refund: %w", err)
			}
		default:
			return nil, fmt.Errorf("unknown tx result field: %s", kvp.Key)
		}
	}

	return &blr, nil
}
