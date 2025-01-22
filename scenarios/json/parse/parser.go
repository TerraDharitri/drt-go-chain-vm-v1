package scenjsonparse

import (
	ei "github.com/TerraDharitri/drt-go-chain-vm-v1/scenarios/expression/interpreter"
	fr "github.com/TerraDharitri/drt-go-chain-vm-v1/scenarios/fileresolver"
)

// Parser performs parsing of both json tests (older) and scenarios (new).
type Parser struct {
	ExprInterpreter ei.ExprInterpreter
}

// NewParser provides a new Parser instance.
func NewParser(fileResolver fr.FileResolver) Parser {
	return Parser{
		ExprInterpreter: ei.ExprInterpreter{
			FileResolver: fileResolver,
		},
	}
}
