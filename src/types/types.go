package types

// types
// a type pre defines what a object / array contains

// Generated json types using: https://mholt.github.io/json-to-go/
type Config struct {
	ProgramsSlash   string   `json:"// programs,omitempty"`
	Programs        []string `json:"programs"`
	WallpaperSlash  string   `json:"// wallpaper,omitempty"`
	Wallpaper       string   `json:"wallpaper"`
	ThemeColorSlash string   `json:"// themeColor,omitempty"`
	ThemeColor      string   `json:"themeColor"`
	SearchSlash     string   `json:"// search,omitempty"`
	Search          string   `json:"search"`
	TaskViewSlash   string   `json:"// taskView,omitempty"`
	TaskView        bool     `json:"taskView"`
	RemoveJunkSlash string   `json:"// removeJunkApps,omitempty"`
	RemoveJunk      bool     `json:"removeJunkApps"`
}

// Type structure to bind to options.
type Options struct {
	PackageName string
}
