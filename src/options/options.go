package options

import "github.com/dennis1248/OpenOEM/src/types"

// GetOptions returns the internal options of this app
func GetOptions() types.Options {
	return types.Options{
		PackageName: "config.json"}
}
