package scencontroller

import (
	fr "github.com/TerraDharitri/drt-go-chain-vm-v1/scenarios/fileresolver"
)

// NewDefaultFileResolver yields a new DefaultFileResolver instance.
// Reexported here to avoid having all external packages importing the parser.
// DefaultFileResolver is in parse for local tests only.
func NewDefaultFileResolver() *fr.DefaultFileResolver {
	return fr.NewDefaultFileResolver()
}
