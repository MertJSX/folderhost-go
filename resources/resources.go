package resources

import (
	"embed"
)

//go:embed default_config.yml
var DefaultConfig embed.FS
