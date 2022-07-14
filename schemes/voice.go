package schemes

// VoiceServerConfig External altv-voice-server configuration
type VoiceServerConfig struct {
	// Host IP address of the external voice server used to receive the alt:V servers voice connection
	// Can be a private ip or 0.0.0.0 to accept all
	Host string `alt:"host" default:"0.0.0.0"`
	// Port used in combination with the ip above
	Port int `alt:"port" default:"7798"`
	// PlayerHost IP address which is used to receive the clients connections
	// Should be a public ip or 0.0.0.0 to accept all
	PlayerHost string `alt:"playerHost" default:"0.0.0.0"`
	// PlayerPort used for the clients connections
	PlayerPort int `alt:"playerPort" default:"7799"`
	// Secret shared between the alt:V server and the external voice server
	Secret string `alt:"secret"`
}
