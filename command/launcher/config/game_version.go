package config

// Information of artifact to download and use as dependency
type ArtifactDownloadInfo struct {
	Path string `json:"path"` // where to put downloaded file
	SHA1 string `json:"sha1"` // sha1 to verify the downloaded file
	Size int    `json:"size"` // size of file to download
	Url  string `json:"url"`  // download url
}

// Download information for dependency
type DownloadInfo struct {
	Artifact *ArtifactDownloadInfo `json:"artifact"`
}

// Dependency of the game version
type Dependency struct {
	Name      string        `json:"id"`                  // dependency ref. as used in Gradle
	Downloads *DownloadInfo `json:"downloads,omitempty"` // Files to download for the dependency
	Url       string        `json:"url,omitempty"`       // Download url
	Checksums []string      `json:"checksums,omitempty"` // Checksums for verifying downloaded files
	ServerReq bool          `json:"serverreq,omitempty"` // Flag to specify whether, or not, dependency is required by server
	ClientReq bool          `json:"clientreq,omitempty"` // Flag to specify whether, or not, dependency is required by client
}

// Version inheriting configuration from another version file
// This is useful e.g. when using a ModPack version on top of a Forge Version
type InheritedVersion struct {
	ID                 string                    `json:"id"` // name or id of the game version
	Time               string                    `json:"time"` // creation time of version
	ReleaseTime        string                    `json:"releaseTime"` // release time of game version
	Type               string                    `json:"type"` // version type, e.g. release
	MinecraftArguments string                    `json:"minecraftArguments"` // arguments to send to minecraft application
	MainClass          string                    `json:"mainClass"` // main class to use when launching game
	InheritsFrom       string                    `json:"inheritsFrom"` // id of game version to inherit from
	Jar                string                    `json:"jar"` // main jar file to use for launching game
	Logging            map[string]*LoggingConfig `json:"logging"` // configuration of logging
	Libraries          []*Dependency             `json:"libraries"` // libraries and dependencies
}

// File to log to
type LoggingFile struct {
	ID   string `json:"id"`
	SHA1 string `json:"sha1"`
	Size int    `json:"size"`
	Url  string `json:"url"`
}

// Configuration of Minecraft client logging
type LoggingConfig struct {
	Argument string       `json:"argument"`
	File     *LoggingFile `json:"file"` // file to log to
	Type     string       `json:"type"`
}
