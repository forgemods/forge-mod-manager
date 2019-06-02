package config

type ProfileSettings struct {
	Settings               *Settings           `json:"settings"`
	LauncherVersion        *Version            `json:"launcherVersion"`
	Profiles               map[string]*Profile `json:"profiles"`
	AuthenticationDatabase map[string]*AuthDB  `json:"authenticationDatabase"`
	SelectedUser           *SelectedUser       `json:"selectedUser"`
	AnalyticsToken         string              `json:"analyticsToken"`
	AnalyticsFailCount     int64               `json:"analyticsFailcount"`
	ClientToken            string              `json:"clientToken"`
}

type Settings struct {
	Channel        string `json:"channel"`
	Locale         string `json:"locale"`
	ShowMenu       bool   `json:"showMenu"`
	EnableAdvanced bool   `json:"enableAdvanced"`
}

type Version struct {
	Name           string `json:"name"`
	Format         int64  `json:"format"`
	ProfilesFormat int64  `json:"profilesFormat"`
}

type Profile struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Icon     string `json:"icon"`
	GameDir  string `json:"gameDir"`
	JavaArgs string `json:"javaArgs"`

	// Date variables with format `2019-03-28T07:00:38.537Z`
	Created       string `json:"created"`
	LastUsed      string `json:"lastUsed"`
	LastVersionID string `json:"lastVersionId"`
}

type AuthDB struct {
	AccessToken string                    `json:"accessToken"`
	Username    string                    `json:"username"`
	Profiles    map[string]*AuthDBProfile `json:"profiles"`
}

type AuthDBProfile struct {
	DisplayName string `json:"displayName"`
}

type SelectedUser struct {
	Account string `json:"account"`
	Profile string `json:"profile"`
}
