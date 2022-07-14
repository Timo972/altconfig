package schemes

// JSModuleConfig Settings related to js-module - part of server.cfg
type JSModuleConfig struct {
	// Inspector Enable node.js inspector
	Inspector struct {
		// Host Define inspector ip
		Host string `alt:"host"`
		// Port Define inspector port
		Port uint16 `alt:"port"`
	}
	// SourceMaps https://nodejs.org/api/cli.html#--enable-source-maps
	SourceMaps bool `alt:"source-maps"`
	// HeapProfiler https://nodejs.org/api/cli.html#--heap-prof
	HeapProfiler bool `alt:"heap-profiler"`
	// Profiler Enable profiler
	Profiler bool `alt:"profiler"`
	// GlobalFetch https://nodejs.org/api/cli.html#--experimental-fetch
	GlobalFetch bool `alt:"global-fetch"`
	// GlobalWebCrypto https://nodejs.org/api/cli.html#--experimental-global-webcrypto
	GlobalWebCrypto bool `alt:"global-webcrypto"`
	// NetworkImports https://nodejs.org/api/cli.html#--experimental-network-imports
	NetworkImports bool `alt:"network-imports"`
	// ExtraCliArgs Add extra cli arguments to the node environment https://nodejs.org/api/cli.html
	ExtraCliArgs []string `alt:"extra-cli-args"`
}

// CSharpModuleConfig Settings related to c#-module - part of server.cfg
type CSharpModuleConfig struct {
	// DisableDependencyDownload Disable dependency (NuGet) check and download at server startup, this is recommended if you have a bad connection to the NuGet server (e.g china)
	DisableDependencyDownload bool `alt:"disableDependencyDownload"`
}

// VoiceConfig integrated voice configuration (needs to be set to enable voice chat) - part of server.cfg
type VoiceConfig struct {
	// Bitrate The bitrate of the voice server
	Bitrate uint64 `alt:"bitrate"`
	// ExternalSecret The secret for the external server (only needed when using externalHost)
	ExternalSecret string `alt:"externalSecret"`
	// ExternalHost The external host address (leave 127.0.0.1, if voice-server is on same machine)
	ExternalHost string `alt:"externalHost"`
	// ExternalPort The external host port
	ExternalPort uint16 `alt:"externalPort"`
	// ExternalPublicHost The external host public address (should be the ip address of your server, not localhost!)
	ExternalPublicHost string `alt:"externalPublicHost"`
	// ExternalPublicPort The external host public port
	ExternalPublicPort uint16 `alt:"externalPublicPort"`
}

// ServerConfig altv-server configuration - server.cfg
type ServerConfig struct {
	// Name The display name of your server
	Name string `alt:"name"`
	// Host The binding address of your server (usually 0.0.0.0)
	Host string `alt:"host" default:"0.0.0.0"`
	// Port The port of your server (default 7788)
	Port int `alt:"port" default:"7788"`
	// Slots The amount of players that can play on your server concurrently
	Slots uint16 `alt:"players"`
	// Description The description of your server
	Description string `alt:"description"`
	// Password The password required to join your server
	Password string `alt:"password"`
	// Announce If the server should be visible on the masterlist in the alt:V client
	Announce bool `alt:"announce" default:"false"`
	// Token Your private masterlist token
	Token string `alt:"token"`
	// GameMode The gamemode your server is running
	GameMode string `alt:"gamemode"`
	// Website The website of your server
	Website string `alt:"website"`
	// Language The language of your server
	Language string `alt:"language" default:"English"`
	// Debug If the debug mode should be allowed (Debug mode allows debugging functions like reconnect or the CEF debugger)
	Debug bool `alt:"debug" default:"false"`
	// StreamingDistance The stream in distance for entities
	StreamingDistance uint16 `alt:"streamingDistance"`
	// MigrationDistance The migration distance for entities
	MigrationDistance uint16 `alt:"migrationDistance"`
	// Timeout The timeout multiplier (must be >= 1)
	Timeout uint64 `alt:"timeout"`
	// AnnounceRetryErrorDelay The delay that is used when the announceRetryErrorAttempts are reached (in ms)
	AnnounceRetryErrorDelay uint64 `alt:"announceRetryErrorDelay" default:"10000"`
	// AnnounceRetryErrorAttempts Max retries until announceRetryErrorDelay will be used
	AnnounceRetryErrorAttempts uint64 `alt:"announceRetryErrorAttempts" default:"50"`
	// DuplicatePlayers Max players which can connect with the same ip address
	DuplicatePlayers uint16 `alt:"duplicatePlayers"`
	// Modules An array of all modules that should be loaded
	Modules []string `alt:"modules"`
	// Resources An array of all resources that should be loaded
	// Since alt:V 10.0 you can also use resources in subdirectories
	Resources []string `alt:"resources"`
	// JSModule See JSModuleConfig comments
	JSModule JSModuleConfig `alt:"js-module"`
	// CSharpModule See CSharpModuleConfig comments
	CSharpModule CSharpModuleConfig `alt:"csharp-module"`
	// Tags The tags for your server (max. 4)
	Tags []string `alt:"tags"`
	// UseEarlyAuth Should early auth be used for your server
	UseEarlyAuth bool `alt:"useEarlyAuth"`
	// EarlyAuthUrl The url for the early auth login page (only used when useEarlyAuth is true)
	EarlyAuthUrl string `alt:"earlyAuthUrl"`
	// UseCDN Should a CDN be used for your server
	UseCDN bool `alt:"useCdn"`
	// CDNUrl The url for the CDN page
	CDNUrl string `alt:"cdnUrl"`
	// Voice See VoiceConfig comments
	Voice VoiceConfig `alt:"voice"`
}
