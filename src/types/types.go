package types

// types
// a type pre defines what a object / array contains

// Generated json types using: https://mholt.github.io/json-to-go/
type Config struct {
	ProgramsSlash string   `json:"// programs"`
	Programs      []string `json:"programs"`
}

// Type structure to bind to options.
type Options struct {
	PackageName string
}
