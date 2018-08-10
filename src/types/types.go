package types

// types
// a type pre-defines what an object / array contains

// Config struct for the config files
// Started this stuct using: https://mholt.github.io/json-to-go/
type Config struct {
	Info                 string   `json:"INFO,omitempty"`
	Info1                string   `json:"INFO1,omitempty"`
	Info2                string   `json:"INFO2,omitempty"`
	ProgramsSlash        string   `json:"// programs,omitempty"`
	Programs             []string `json:"programs"`
	WallpaperSlash       string   `json:"// wallpaper,omitempty"`
	Wallpaper            string   `json:"wallpaper"`
	ThemeColorSlash      string   `json:"// themeColor,omitempty"`
	ThemeColor           string   `json:"themeColor"`
	SearchSlash          string   `json:"// search,omitempty"`
	Search               string   `json:"search"`
	TaskViewSlash        string   `json:"// taskView,omitempty"`
	TaskView             bool     `json:"taskView"`
	RemoveJunkSlash      string   `json:"// removeJunkApps,omitempty"`
	RemoveJunk           bool     `json:"removeJunkApps"`
	RemoveEdigeIconSlash string   `json:"// removeEdgeIcon,omitempty"`
	RemoveEdigeIcon      bool     `json:"removeEdgeIcon"`
	RemovePeopleSlash    string   `json:"// removePeople,omitempty"`
	RemovePeople         bool     `json:"removePeople"`
}

// Options bind to the app internal options.
type Options struct {
	PackageName string
}
