package warhammer_aos

import "embed"

//go:embed data/*/*.json
var Files embed.FS

const (
	DataDir = "data/armies"
)
