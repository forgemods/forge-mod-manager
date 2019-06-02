package config

type LoggingFile struct {
	ID   string `json:"id"`
	SHA1 string `json:"sha1"`
	Size int    `json:"size"`
	Url  string `json:"url"`
}

type LoggingConfig struct {
	Argument string       `json:"argument"`
	File     *LoggingFile `json:"file"`
	Type     string       `json:"type"`
}

type ArtifactDownloadInfo struct {
	Path string `json:"path"`
	SHA1 string `json:"sha1"`
	Size int    `json:"size"`
	Url  string `json:"url"`
}

type DownloadInfo struct {
	Artifact *ArtifactDownloadInfo `json:"artifact"`
}

type Library struct {
	Name      string        `json:"id"`
	Downloads *DownloadInfo `json:"downloads,omitempty"`
	Url       string        `json:"url,omitempty"`
	Checksums []string      `json:"checksums,omitempty"`
	ServerReq bool          `json:"serverreq,omitempty"`
	ClientReq bool          `json:"clientreq,omitempty"`
}

type InheritedVersion struct {
	ID                 string                    `json:"id"`
	Time               string                    `json:"time"`
	ReleaseTime        string                    `json:"releaseTime"`
	Type               string                    `json:"type"`
	MinecraftArguments string                    `json:"minecraftArguments"`
	MainClass          string                    `json:"mainClass"`
	InheritsFrom       string                    `json:"inheritsFrom"`
	Jar                string                    `json:"jar"`
	Logging            map[string]*LoggingConfig `json:"logging"`
	Libraries          []*Library                `json:"libraries"`
}
