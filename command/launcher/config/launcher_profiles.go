package config

// Root model for launcher_profiles.json file
type ProfileSettings struct {
	Settings               *Settings           `json:"settings"`               // Contains all the launcher settings
	LauncherVersion        *Version            `json:"launcherVersion"`        // Contains the current launcher build name, format and profiles format.
	Profiles               map[string]*Profile `json:"profiles"`               // All the launcher profiles and their configurations.
	AuthenticationDatabase map[string]*AuthDB  `json:"authenticationDatabase"` // All the logged in accounts. Every account in this key contains a UUID-hashed map of account info
	SelectedUser           *SelectedUser       `json:"selectedUser"`           // Contains the UUID-hashed account and the UUID of the currently selected user
	AnalyticsToken         string              `json:"analyticsToken"`         // The latest token for tracking analysts. Those are used locally by the launcher for managing and tracking sessions
	AnalyticsFailCount     int64               `json:"analyticsFailcount"`     // The fail count for analytics (???)
	ClientToken            string              `json:"clientToken"`            // The currently logged in client token.
}

// Minecraft Java Launcher Settings
type Settings struct {
	Channel        string `json:"channel"`        // what type of version channel to use, e.g. release
	Locale         string `json:"locale"`         // language to use in the game
	ShowMenu       bool   `json:"showMenu"`       // flag indication of whether to show the menu or hide it
	EnableAdvanced bool   `json:"enableAdvanced"` // flag indication of whether advanced settings are enabled
}

// Minecraft Java Launcher Version
type Version struct {
	Name           string `json:"name"` // version name
	Format         int64  `json:"format"`
	ProfilesFormat int64  `json:"profilesFormat"`
}

type ProfileType string // Type of game version profile configuration

const (
	ProfileTypeCustom         ProfileType = "custom" // manually created by the user
	ProfileTypeLatestRelease  ProfileType = "latest-release" // uses the latest stable release
	ProfileTypeLatestSnapshot ProfileType = "latest-snapshot" // uses the latest build of Minecraft
)

// Game version specific configuration for launcher
type Profile struct {
	Name     string      `json:"name"`     // profile name
	Type     ProfileType `json:"type"`     // type specifying what way to download and setup files and dependencies before launching
	Icon     string      `json:"icon"`     // base64 encoded string of the icon image to use. e.g. data:image/png;base64,iVBORw0KG...
	GameDir  string      `json:"gameDir"`  // working directory for Minecraft Java process to use
	JavaArgs string      `json:"javaArgs"` // JVM arguments for controlling e.g. allowed memory usage

	// Below: Date variables with format `2019-03-28T07:00:38.537Z`
	Created       string `json:"created"`       // creation date of this profile
	LastUsed      string `json:"lastUsed"`      // date when this profile were last launched
	LastVersionID string `json:"lastVersionId"` // id of what game version configuration to use
}

// Authentication details configured by the Minecraft Launcher user
type AuthDB struct {
	AccessToken string                    `json:"accessToken"` // access token from the Mojang authentication servers
	Username    string                    `json:"username"`    // username of the user logged in with the launcher
	Profiles    map[string]*AuthDBProfile `json:"profiles"`    // list of user profiles configured in the launcher
}

// Profile entry in AuthDB
type AuthDBProfile struct {
	DisplayName string `json:"displayName"` // Display name of the user profile
}

// Currently selected user in the Launcher
type SelectedUser struct {
	Account string `json:"account"` // Account hash
	Profile string `json:"profile"` // Account profile hash
}
